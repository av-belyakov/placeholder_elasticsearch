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

// SettingsChanInputEEM параметры для ПРИЕМА данных в модуль
// RootId - основной идентификатор
// Source - источник
// SensorsId - искомые идентификаторы сенсоров
type SettingsChanInputEEM struct {
	RootId    string
	Source    string
	SensorsId []string
}

// SettingsChanInputEEM параметры для ОТПРАВКИ данных из модуля
// RootId - основной идентификатор
// Source - источник
// Sensors - найденная информация, где ключ карты это идентификатор сенсора
type SettingsChanOutputEEM struct {
	RootId  string
	Source  string
	Sensors map[string]FoundSensorInformation
}

// FoundSensorInformation содержит найденную о сенсоре информацию
// SensorId - идентификатор сенсора
// HostId - идентификатор сенсора, специальный, для поиска информации в НКЦКИ
// GeoCode - географический код страны
// ObjectArea - сфера деятельности объекта
// SubjectRF - субъект Российской Федерации
// INN - налоговый идентификатор
// HomeNet - перечень домашних сетей
// OrgName - наименование организации
// FullOrgName - полное наименование организации
type FoundSensorInformation struct {
	SensorId    string
	HostId      string
	GeoCode     string
	ObjectArea  string
	SubjectRF   string
	INN         string
	HomeNet     string
	OrgName     string
	FullOrgName string
}

// EventEnrichmentModule инициализированный модуль обогащения кейсов
// ChanInputModule - канал для отправки данных В модуль
// ChanOutputModule - канал для принятия данных ИЗ модуля
type EventEnrichmentModule struct {
	ChanInputModule  chan SettingsChanInputEEM
	ChanOutputModule chan SettingsChanOutputEEM
}

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

	ctxTimeout, ctxTimeoutCancel := context.WithTimeout(context.Background(), 5*time.Second)
	settingsFullOrgNameByINN, err := NewSettingsFuncFullNameOrganizationByINN(ctxTimeout, ncirccConf.URL, ncirccConf.Token, (5 * time.Second))
	if err != nil {
		ctxTimeoutCancel()
		return &module, err
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
				settingsResponse := SettingsChanOutputEEM{
					RootId:  data.RootId,
					Source:  data.Source,
					Sensors: make(map[string]FoundSensorInformation),
				}

				for _, sensorId := range data.SensorsId {
					fullInfo, err := zabbixinteractions.GetFullSensorInformationFromZabbixAPI(sensorId, zabbixConnHandler)
					if err != nil {
						_, f, l, _ := runtime.Caller(0)
						logging <- datamodels.MessageLogging{
							MsgData: fmt.Sprintf("'%v' %s:%d", err, f, l-1),
							MsgType: "error",
						}

						continue
					}

					settingsResponse.Sensors[sensorId] = FoundSensorInformation{
						SensorId:   sensorId,
						HostId:     fullInfo.HostId,
						GeoCode:    fullInfo.GeoCode,
						ObjectArea: fullInfo.ObjectArea,
						SubjectRF:  fullInfo.SubjectRF,
						INN:        fullInfo.INN,
						HomeNet:    fullInfo.HomeNet,
					}

					rd, err := settingsFullOrgNameByINN.GetFullNameOrganizationByINN(fullInfo.INN)
					if err != nil {
						_, f, l, _ := runtime.Caller(0)
						logging <- datamodels.MessageLogging{
							MsgData: fmt.Sprintf("'%v' %s:%d", err, f, l-1),
							MsgType: "error",
						}
					}

					if len(rd.Data) == 0 {
						_, f, l, _ := runtime.Caller(0)
						logging <- datamodels.MessageLogging{
							MsgData: fmt.Sprintf("'nothing was found by INN '%s'' %s:%d", fullInfo.INN, f, l-1),
							MsgType: "error",
						}

						return
					}

					if sensor, ok := settingsResponse.Sensors[sensorId]; ok {
						sensor.OrgName = rd.Data[0].Name
						sensor.FullOrgName = rd.Data[0].Sname

						settingsResponse.Sensors[sensorId] = sensor
					}
				}

				module.ChanOutputModule <- settingsResponse
			}
		}
	}()

	return &module, nil
}
