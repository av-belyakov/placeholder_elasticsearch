package eventenrichmentmodule

import (
	"context"
	"fmt"
	"placeholder_elasticsearch/confighandler"
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/zabbixinteractions"
	"runtime"
	"time"
)

// NewEventEnrichmentModule инициализирует новый модуль обогащения данными
// ctx - должен быть context.WithCancel()
// ncirccConf - параметры для доступа к НКЦКИ
// zabbixApi - параметры для доступа к API Zabbix
func NewEventEnrichmentModule(
	ctx context.Context,
	ncirccConf confighandler.NCIRCCOptions,
	zabbixApi confighandler.ZabbixJsonRPCOptions,
	logging chan<- datamodels.MessageLogging) (*EventEnrichmentModule, error) {
	module := EventEnrichmentModule{
		ChanInputModule:  make(chan SettingsChanInputEEM),
		ChanOutputModule: make(chan SettingsChanOutputEEM),
	}

	ctxTimeout, ctxTimeoutCancel := context.WithTimeout(context.Background(), 7*time.Second)
	settingsFullOrgNameByINN, err := NewSettingsFuncFullNameOrganizationByINN(ctxTimeout, ncirccConf.URL, ncirccConf.Token, (5 * time.Second))
	if err != nil {
		ctxTimeoutCancel()
		return &module, err
	}

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

	if err := checkConfZabbixAPI(zabbixApi); err != nil {
		ctxTimeoutCancel()
		return &module, err
	}

	connTimeout, err := time.ParseDuration(fmt.Sprintf("%ds", zabbixApi.ConnectionTimeout))
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		ctxTimeoutCancel()
		return &module, fmt.Errorf("'%v' %s:%d", err, f, l-1)
	}

	zabbixConnHandler, err := zabbixinteractions.NewZabbixConnectionJsonRPC(
		zabbixinteractions.SettingsZabbixConnectionJsonRPC{
			Host:              zabbixApi.NetworkHost,
			Login:             zabbixApi.Login,
			Passwd:            zabbixApi.Passwd,
			ConnectionTimeout: &connTimeout,
		})
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		ctxTimeoutCancel()
		return &module, fmt.Errorf("'%v' %s:%d", err, f, l-1)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				ctxTimeoutCancel()

				return

			case data := <-module.ChanInputModule:
				go func(
					settings SettingsChanInputEEM,
					chanOutput chan<- SettingsChanOutputEEM,
					logging chan<- datamodels.MessageLogging) {

					settingsResponse := SettingsChanOutputEEM{
						RootId:  data.RootId,
						Source:  data.Source,
						Sensors: []FoundSensorInformation(nil),
					}

					for _, sensorId := range data.SensorsId {
						fullInfo, err := zabbixinteractions.GetFullSensorInformationFromZabbixAPI(sensorId, zabbixConnHandler)
						if err != nil {
							_, f, l, _ := runtime.Caller(0)
							logging <- datamodels.MessageLogging{
								MsgData: fmt.Sprintf("'sensorId: '%s', %v' %s:%d", sensorId, err, f, l-1),
								MsgType: "error",
							}

							continue
						}

						foundInfo := FoundSensorInformation{
							SensorId:   sensorId,
							HostId:     fullInfo.HostId,
							GeoCode:    fullInfo.GeoCode,
							ObjectArea: fullInfo.ObjectArea,
							SubjectRF:  fullInfo.SubjectRF,
							INN:        fullInfo.INN,
							HomeNet:    fullInfo.HomeNet,
						}

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

						settingsResponse.Sensors = append(settingsResponse.Sensors, foundInfo)
					}

					module.ChanOutputModule <- settingsResponse
				}(data, module.ChanOutputModule, logging)
			}
		}
	}()

	return &module, nil
}
