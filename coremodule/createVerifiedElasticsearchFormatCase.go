package coremodule

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"
	"placeholder_elasticsearch/eventenrichmentmodule"
	"placeholder_elasticsearch/listhandlerforesjson"
	"placeholder_elasticsearch/listhandlerthehivejson"
)

type listSensorId struct {
	sensors []string
}

// Get возвращает список идентификаторов сенсоров
func (e *listSensorId) Get() []string {
	return e.sensors
}

// AddElem добавляет только уникальные элементы
func (e *listSensorId) AddElem(sensorId string) {
	for _, v := range e.sensors {
		if v == sensorId {
			return
		}
	}

	e.sensors = append(e.sensors, sensorId)
}

func NewVerifiedElasticsearchFormatCase(
	input <-chan datamodels.ChanOutputDecodeJSON,
	done <-chan bool,
	esm *elasticsearchinteractions.ElasticSearchModule,
	cs *coreStorage,
	eemChan chan<- eventenrichmentmodule.SettingsChanInputEEM,
	logging chan<- datamodels.MessageLogging) {
	var (
		rootId string
		// список не обработанных полей
		listRawFields map[string]string = make(map[string]string)

		//Финальный объект
		verifiedCase *datamodels.VerifiedEsCase = datamodels.NewVerifiedEsCase()

		event *datamodels.EventMessageForEsCase = datamodels.NewEventMessageForEsCase()

		eventObject  *datamodels.EventForEsCaseObject = datamodels.NewEventForEsCaseObject()
		eventDetails *datamodels.EventCaseDetails     = datamodels.NewEventCaseDetails()

		eventObjectCustomFields  datamodels.CustomFields = datamodels.CustomFields{}
		eventDetailsCustomFields datamodels.CustomFields = datamodels.CustomFields{}
	)

	//******************* Основные обработчики для Event **********************
	// ------ EVENT ------
	listHandlerEvent := listhandlerforesjson.NewListHandlerEventCaseElement(event)
	// ------ EVENT OBJECT ------
	listHandlerEventObject := listhandlerforesjson.NewListHandlerEventCaseObjectElement(eventObject)
	// ------ EVENT OBJECT CUSTOMFIELDS ------
	listHandlerEventObjectCustomFields := listhandlerthehivejson.NewListHandlerEventObjectCustomFieldsElement(eventObjectCustomFields)
	// ------ EVENT DETAILS ------
	listHandlerEventDetails := listhandlerforesjson.NewListHandlerEventCaseDetailsElement(eventDetails)
	// ------ EVENT DETAILS CUSTOMFIELDS ------
	listHandlerEventDetailsCustomFields := listhandlerthehivejson.NewListHandlerEventDetailsCustomFieldsElement(eventDetailsCustomFields)

	//******************* Вспомогательный объект для Observables **********************
	so := listhandlerforesjson.NewSupportiveObservables()
	listHandlerObservables := listhandlerforesjson.NewListHandlerObservablesElement(so)

	//******************* Вспомогательный объект для Ttp **********************
	sttp := listhandlerthehivejson.NewSupportiveTtp()
	listHandlerTtp := listhandlerthehivejson.NewListHandlerTtpElement(sttp)

	for data := range input {
		var handlerIsExist bool
		verifiedCase.SetID(data.UUID)

		if source, ok := searchEventSource(data.FieldBranch, data.Value); ok {
			verifiedCase.SetSource(source)

			continue
		}

		if data.FieldBranch == "event.rootId" {
			rootId = fmt.Sprint(data.Value)
		}

		//******************************************************************
		//********** Сбор всех объектов относящихся к полю Event  **********
		// event element
		if lf, ok := listHandlerEvent[data.FieldBranch]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(data.Value)
			}

			continue
		}

		// event.object element
		if lf, ok := listHandlerEventObject[data.FieldBranch]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(data.Value)
			}

			continue
		}

		// event.object.customFields element
		if lf, ok := listHandlerEventObjectCustomFields[data.FieldBranch]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(data.Value)
			}

			continue
		}

		// event.details element
		if lf, ok := listHandlerEventDetails[data.FieldBranch]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(data.Value)
			}

			continue
		}

		// event.details.customFields element
		if lf, ok := listHandlerEventDetailsCustomFields[data.FieldBranch]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(data.Value)
			}

			continue
		}

		//************************************************************************
		//********** Сбор всех объектов относящихся к полю Observables  **********
		// для всех полей входящих в observables, кроме содержимого
		//поля reports
		if lf, ok := listHandlerObservables[data.FieldBranch]; ok {
			handlerIsExist = true

			for _, f := range lf {
				r := reflect.TypeOf(data.Value)
				switch r.Kind() {
				case reflect.Slice:
					if s, ok := data.Value.([]interface{}); ok {
						for _, value := range s {
							f(value)
						}
					}
				default:
					f(data.Value)

				}
			}

			continue
		}

		//убрал обработку observables.reports так как тип TtpsMessageEs
		//способствует росту черезмерно большого количества полей которое
		//влечет за собой превышения лимита маппинга в Elsticsearch), что
		//выражается в ошибке от СУБД типа "Limit of total fields [2000]
		//has been exceeded while adding new fields"
		//
		//для всех полей входящих в состав observables.reports
		//if strings.Contains(data.FieldBranch, "observables.reports.") {
		//		handlerIsExist = true
		//		so.HandlerReportValue(data.FieldBranch, data.Value)
		//}

		//*********************************************************************
		//********** Сбор всех объектов относящихся к полю Ttp  ***************
		if lf, ok := listHandlerTtp[data.FieldBranch]; ok {
			handlerIsExist = true

			for _, f := range lf {
				r := reflect.TypeOf(data.Value)
				switch r.Kind() {
				case reflect.Slice:
					if s, ok := data.Value.([]interface{}); ok {
						for _, value := range s {
							f(value)
						}
					}

				default:
					f(data.Value)

				}
			}

			continue
		}

		if !handlerIsExist {
			// записываем в лог-файл поля, которые не были обработаны
			listRawFields[data.FieldBranch] = fmt.Sprint(data.Value)
		}
	}

	// отправляем список полей которые не смогли обработать
	if len(listRawFields) > 0 {
		logging <- datamodels.MessageLogging{
			MsgData: joinRawFieldsToString(listRawFields, "rootId", rootId),
			MsgType: "alert_raw_fields",
		}
	}

	//проверяем значения объектов на соответствие правилам
	isAllowed := <-done
	if !isAllowed {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'the message with aler rootId %s was not sent to ES because it does not comply with the rules' %s:%d", event.GetRootId(), f, l-1),
			MsgType: "warning",
		}

		return
	}

	// Собираем объект Event
	eventObject.SetValueCustomFields(eventObjectCustomFields)
	eventDetails.SetValueCustomFields(eventDetailsCustomFields)
	event.SetValueObject(*eventObject)
	event.SetValueDetails(*eventDetails)

	// собираем объект observables
	observables := datamodels.NewObservablesMessageEs()
	observables.SetValueObservables(so.GetObservables())

	// собираем объект ttp
	ttps := datamodels.NewTtpsMessageTheHive()
	ttps.SetTtps(sttp.GetTtps())

	verifiedCase.SetEvent(*event)
	verifiedCase.SetObservables(*observables)
	verifiedCase.SetTtps(*ttps)

	sensorsId := listSensorId{
		sensors: []string(nil),
	}

	eventCase := verifiedCase.GetEvent()
	objectElem := eventCase.GetObject()
	if listSensorId, ok := objectElem.GetTags()["sensor:id"]; ok {
		for _, v := range listSensorId {
			sensorsId.AddElem(v)
		}
	}

	//отправляем кейс в Elasticsearch
	esm.ChanInputModule <- elasticsearchinteractions.SettingsInputChan{
		Section: "handling case",
		Command: "add new case",
		Data:    verifiedCase.Get(),
	}

	//делаем запрос на получение дополнительной информации о сенсорах
	if len(sensorsId.Get()) > 0 {
		//делаем запрос к модулю обогащения доп. информацией из Zabbix
		eemChan <- eventenrichmentmodule.SettingsChanInputEEM{
			RootId:    eventCase.GetRootId(),
			Source:    verifiedCase.GetSource(),
			SensorsId: sensorsId.Get(),
		}

		/*
			550001
			550119
			570048

						Некоторые сенсора отвечают вот так:
					SensorId 530134
						--->>> func 'CoreHandler' RESIVED sensors info: '{~512335920 rcmmsk [{530134 10504 Wrong Json: <!DOCTYPE html>
				<html lang="en">
				<head>
				  <meta charset="utf-8" />
				  <title>Ведомственная система обнаружения компьютерных атак</title>

				  <script src="/config/"></script>

				  <script type="module" crossorigin src="/assets/index.313374bb.js"></script>
				  <link rel="stylesheet" href="/assets/index.190ded15.css">
				  <script type="module">try{import.meta.url;import("_").catch(()=>1);}catch(e){}window.__vite_is_modern_browser=true;</script>
				  <script type="module">!function(){if(window.__vite_is_modern_browser)return;console.warn("vite: loading legacy build because dynamic import or import.meta.url is unsupported, syntax error above should be ignored");var e=document.getElementById("vite-legacy-polyfill"),n=document.createElement("script");n.src=e.src,n.onload=function(){System.import(document.getElementById('vite-legacy-entry').getAttribute('data-src'))},document.body.appendChild(n)}();</script>
				</head>...

				По этому это не добавляется в Эластик ни как

		*/

		/*
			//получаем результат выполнения запроса к модулю обогащения
			resultEventenrichment := <-eem.ChanOutputModule

			if len(resultEventenrichment.Sensors) > 0 {
				//если какая либо информация была найдена
				sai := datamodels.NewSensorAdditionalInformation()
				for _, v := range resultEventenrichment.Sensors {
					si := datamodels.NewSensorInformation()
					si.SetSensorId(v.SensorId)
					si.SetHostId(v.HostId)
					si.SetGeoCode(v.GeoCode)
					si.SetObjectArea(v.ObjectArea)
					si.SetSubjectRF(v.SubjectRF)
					si.SetINN(v.INN)
					si.SetHomeNet(v.HomeNet)
					si.SetOrgName(v.OrgName)
					si.SetFullOrgName(v.FullOrgName)

					sai.Add(*si)
				}
				verifiedCase.SetSensorAdditionalInformation(*sai)

				//отправляем, найденную о сенсорах информацию, в MongoDB
				mdbm.ChanInputModule <- mongodbinteractions.SettingsInputChan{
					Section: "handling eventenrichment",
					Command: "add sensor eventenrichment",
					Data:    resultEventenrichment,
				}
			} /*else {
				//если ничего не было найдено, то есть фактически остутствует доступ
				//к модулю взаимодействия с Zabbix
				//отправляем в MongoDB запрос со списком идентификаторов сенсоров
				//при этом прием результата выполняем в CoreHandler и отправляем
				//полученные данные в СУБД Elasticsearch для ДОПОЛНЕНИЯ информации
				//о кейсах
				mdbm.ChanInputModule <- mongodbinteractions.SettingsInputChan{
					Section: "handling eventenrichment",
					Command: "get sensor eventenrichment",
					Data:    resultEventenrichment,
				}
			}*/
	}

	//******** TEST ********
	//только в рамках тестирования, отправка обновленного объекта
	//в специальный файл
	infoUpdate, err := json.MarshalIndent(verifiedCase, "", "  ")
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}
	}
	logging <- datamodels.MessageLogging{
		MsgData: string(infoUpdate),
		MsgType: "test_object_update",
	}
	//***********************
}
