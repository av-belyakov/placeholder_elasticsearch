package eventenrichmentmodule

import (
	"context"
	"fmt"
	"regexp"
	"runtime"
	"sync"
	"time"

	"placeholder_elasticsearch/confighandler"
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/zabbixinteractions"
)

// NewEventEnrichmentModule инициализирует новый модуль обогащения данными
func NewEventEnrichmentModule(
	ctx context.Context,
	options EventEnrichmentModuleOptions,
	logging chan<- datamodels.MessageLogging) (*EventEnrichmentModule, error) {

	module := EventEnrichmentModule{
		ChanInputModule:  make(chan SettingsChanInputEEM),
		ChanOutputModule: make(chan SettingsChanOutputEEM),
	}

	patternNumeric := regexp.MustCompile(`^[0-9]+$`)

	connTimeout, err := time.ParseDuration(fmt.Sprintf("%ds", options.ConfZabbixApi.ConnectionTimeout))
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		return &module, fmt.Errorf("'%v' %s:%d", err, f, l-1)
	}

	//поиск информации в Zabbix
	checkConfZabbixAPI := func(zabbixApi confighandler.ZabbixJsonRPCOptions) error {
		if zabbixApi.NetworkHost == "" {
			_, f, l, _ := runtime.Caller(0)
			return fmt.Errorf("'the 'NetworkHost' parameter required to access the ZabbixAPI cannot be empty' %s:%d", f, l-1)
		}

		if zabbixApi.Login == "" {
			_, f, l, _ := runtime.Caller(0)
			return fmt.Errorf("'the 'Login' parameter required to access the ZabbixAPI cannot be empty' %s:%d", f, l-1)
		}

		if zabbixApi.Passwd == "" {
			_, f, l, _ := runtime.Caller(0)
			return fmt.Errorf("'the 'Passwd' parameter required to access the ZabbixAPI cannot be empty' %s:%d", f, l-1)
		}

		return nil
	}

	if err := checkConfZabbixAPI(options.ConfZabbixApi); err != nil {
		return &module, err
	}

	ctxDone, ctxDoneCancel := context.WithCancel(context.Background())
	settingsFullOrgNameByINN, err := NewSettingsFuncFullNameOrganizationByINN(ctxDone, options.ConfNCIRCC.URL, options.ConfNCIRCC.Token, (5 * time.Second))
	if err != nil {
		ctxDoneCancel()
		return &module, err
	}

	//поиск информации по ИНН в НКЦКИ
	searchNCIRCCInfo := func(inn, sensorId string) (orgName, fullOrgName string, err error) {
		if inn == "" {
			_, f, l, _ := runtime.Caller(0)
			return "", "", fmt.Errorf("'sensorId: '%s', it is impossible to search for additional information because the INN is empty' %s:%d", sensorId, f, l-1)
		}

		rd, err := settingsFullOrgNameByINN.GetFullNameOrganizationByINN(inn)
		if err != nil {
			_, f, l, _ := runtime.Caller(0)
			return "", "", fmt.Errorf("'sensorId: '%s', %v' %s:%d", sensorId, err, f, l-1)
		}

		if len(rd.Data) == 0 {
			_, f, l, _ := runtime.Caller(0)
			return "", "", fmt.Errorf("'sensorId: '%s', nothing was found by INN '%s'' %s:%d", sensorId, inn, f, l-1)
		}

		return rd.Data[0].Name, rd.Data[0].Sname, nil
	}

	zabbixConnHandler, err := zabbixinteractions.NewZabbixConnectionJsonRPC(
		zabbixinteractions.SettingsZabbixConnectionJsonRPC{
			Host:              options.ConfZabbixApi.NetworkHost,
			Login:             options.ConfZabbixApi.Login,
			Passwd:            options.ConfZabbixApi.Passwd,
			ConnectionTimeout: &connTimeout,
		})
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		ctxDoneCancel()
		return &module, fmt.Errorf("'%v' %s:%d", err, f, l-1)
	}

	//поиск информации по георасположению в API баз данных geoip
	geoIpClient, err := NewGeoIpClient(
		context.Background(),
		WithHost(options.ConfGeoIP.Host),
		WithPort(options.ConfGeoIP.Port),
		WithPath(options.ConfGeoIP.Path),
		WithConnectionTimeout(10*time.Second))
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		ctxDoneCancel()
		return &module, fmt.Errorf("'%v' %s:%d", err, f, l-1)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				ctxDoneCancel()

				return

			case data := <-module.ChanInputModule:
				go func(
					settings SettingsChanInputEEM,
					chanOutput chan<- SettingsChanOutputEEM,
					logging chan<- datamodels.MessageLogging) {

					settingsResponse := SettingsChanOutputEEM{
						RootId:            data.RootId,
						Source:            data.Source,
						SensorsIdNotFound: []string(nil),
						Sensors:           []FoundSensorInformation(nil),
						IpAddresses:       []string(nil),
						IpAddressesInfo:   []GeoIpInformation(nil),
					}

					var wg sync.WaitGroup
					wg.Add(2)

					//поиск информации по сенсорам
					go func(sensors []string) {
						for _, sensorId := range sensors {
							//получение полной информации о сенсоре
							fullInfo, err := zabbixinteractions.GetFullSensorInformationFromZabbixAPI(sensorId, zabbixConnHandler)
							if err != nil {
								_, f, l, _ := runtime.Caller(0)
								logging <- datamodels.MessageLogging{
									MsgData: fmt.Sprintf("'sensorId: '%s', %v' %s:%d", sensorId, err, f, l-1),
									MsgType: "error",
								}

								settingsResponse.SensorsIdNotFound = append(settingsResponse.SensorsIdNotFound, sensorId)

								continue
							}

							if fullInfo.HostId == "" {
								settingsResponse.SensorsIdNotFound = append(settingsResponse.SensorsIdNotFound, sensorId)

								continue
							}

							foundInfo := FoundSensorInformation{
								SensorId:   sensorId,
								HostId:     fullInfo.HostId,
								GeoCode:    fullInfo.GeoCode,
								ObjectArea: MappingObjectArea(fullInfo.ObjectArea, options.ConfMapping.AreaActivity),
								SubjectRF:  fullInfo.SubjectRF,
								INN:        fullInfo.INN,
								HomeNet:    fullInfo.HomeNet,
							}

							if patternNumeric.MatchString(fullInfo.INN) {
								orgName, fullOrgName, err := searchNCIRCCInfo(fullInfo.INN, sensorId)
								if err != nil {
									logging <- datamodels.MessageLogging{
										MsgData: err.Error(),
										MsgType: "error",
									}
								} else {
									foundInfo.OrgName = orgName
									foundInfo.FullOrgName = fullOrgName
								}
							}

							settingsResponse.Sensors = append(settingsResponse.Sensors, foundInfo)
						}

						wg.Done()
					}(data.SensorsId)

					//поиск информации о геопозиционировании ip адресов
					go func(ipAddresses []string) {
						for _, ip := range ipAddresses {
							geoIpInfo, err := geoIpClient.GetGeoInformation(ip)
							if err != nil {
								_, f, l, _ := runtime.Caller(0)
								logging <- datamodels.MessageLogging{
									MsgData: fmt.Sprintf("'ip address: '%s', %v' %s:%d", ip, err, f, l-1),
									MsgType: "error",
								}

								continue
							}

							settingsResponse.IpAddresses = append(settingsResponse.IpAddresses, ip)
							settingsResponse.IpAddressesInfo = append(settingsResponse.IpAddressesInfo, geoIpInfo)
						}

						wg.Done()
					}(data.IpAddresses)

					wg.Wait()

					module.ChanOutputModule <- settingsResponse
				}(data, module.ChanOutputModule, logging)
			}
		}
	}()

	return &module, nil
}
