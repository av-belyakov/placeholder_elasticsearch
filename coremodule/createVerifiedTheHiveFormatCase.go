package coremodule

import (
	"fmt"
	"reflect"
	"strings"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"
	"placeholder_elasticsearch/listhandlerthehivejson"
	"placeholder_elasticsearch/mongodbinteractions"
)

func NewVerifiedTheHiveFormatCase(
	input <-chan datamodels.ChanOutputDecodeJSON,
	done <-chan bool,
	esm *elasticsearchinteractions.ElasticSearchModule,
	mongodbm *mongodbinteractions.MongoDBModule,
	logging chan<- datamodels.MessageLogging,
) {
	var (
		//Финальный объект
		verifiedCase *datamodels.VerifiedTheHiveCase = datamodels.NewVerifiedTheHiveCase()

		event        datamodels.EventMessageTheHiveCase = datamodels.EventMessageTheHiveCase{}
		eventObject  datamodels.EventCaseObject         = datamodels.EventCaseObject{}
		eventDetails datamodels.EventCaseDetails        = datamodels.EventCaseDetails{}

		eventObjectCustomFields  map[string]datamodels.CustomerFields = make(map[string]datamodels.CustomerFields)
		eventDetailsCustomFields map[string]datamodels.CustomerFields = make(map[string]datamodels.CustomerFields)
	)

	//******************* Основные обработчики для Event **********************
	// ------ EVENT ------
	listHandlerEvent := listhandlerthehivejson.NewListHandlerEventCaseElement(&event)
	// ------ EVENT OBJECT ------
	listHandlerEventObject := listhandlerthehivejson.NewListHandlerEventCaseObjectElement(&eventObject)
	// ------ EVENT OBJECT CUSTOMFIELDS ------
	listHandlerEventObjectCustomFields := listhandlerthehivejson.NewListHandlerEventObjectCustomFieldsElement(eventObjectCustomFields)
	// ------ EVENT DETAILS ------
	listHandlerEventDetails := listhandlerthehivejson.NewListHandlerEventCaseDetailsElement(&eventDetails)
	// ------ EVENT DETAILS CUSTOMFIELDS ------
	listHandlerEventDetailsCustomFields := listhandlerthehivejson.NewListHandlerEventDetailsCustomFieldsElement(eventDetailsCustomFields)

	//******************* Вспомогательный объект для Observables **********************
	so := listhandlerthehivejson.NewSupportiveObservables()
	listHandlerObservables := listhandlerthehivejson.NewListHandlerObservablesElement(so)

	//******************* Вспомогательный объект для Ttp **********************
	sttp := listhandlerthehivejson.NewSupportiveTtp()
	listHandlerTtp := listhandlerthehivejson.NewListHandlerTtpElement(sttp)

	for {
		select {
		case data := <-input:
			var (
				rootId         string
				handlerIsExist bool
			)

			verifiedCase.SetID(data.UUID)

			if source, ok := searchEventSource(data.FieldBranch, data.Value); ok {
				verifiedCase.SetSource(source)
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
			}

			// event.object element
			if lf, ok := listHandlerEventObject[data.FieldBranch]; ok {
				handlerIsExist = true

				for _, f := range lf {
					f(data.Value)
				}
			}

			// event.object.customFields element
			if lf, ok := listHandlerEventObjectCustomFields[data.FieldBranch]; ok {
				handlerIsExist = true

				for _, f := range lf {
					f(data.Value)
				}
			}

			// event.details element
			if lf, ok := listHandlerEventDetails[data.FieldBranch]; ok {
				handlerIsExist = true

				for _, f := range lf {
					f(data.Value)
				}
			}

			// event.details.customFields element
			if lf, ok := listHandlerEventDetailsCustomFields[data.FieldBranch]; ok {
				handlerIsExist = true

				for _, f := range lf {
					f(data.Value)
				}
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
			}

			//для всех полей входящих в состав observables.reports
			if strings.Contains(data.FieldBranch, "observables.reports.") {
				handlerIsExist = true
				so.HandlerReportValue(data.FieldBranch, data.Value)
			}

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
			}

			// записываем в лог-файл поля, которые не были обработаны
			if handlerIsExist {
				logging <- datamodels.MessageLogging{
					MsgData: fmt.Sprintf("event rootId: '%s', field: '%s', value: '%v'", rootId, data.FieldBranch, data.Value),
					MsgType: "case_raw_fields",
				}
			}

		case <-done:
			//Собираем объект Event
			eventObject.SetValueCustomFields(eventObjectCustomFields)
			eventDetails.SetValueCustomFields(eventDetailsCustomFields)
			event.SetValueObject(eventObject)
			event.SetValueDetails(eventDetails)

			//проверяем объек на наличие пустых полей которые должны
			//содержать дату и время
			checkDatetimeFieldsEventObject(&event)

			//собираем объект observables
			observables := datamodels.NewObservablesMessageTheHive()
			observables.SetObservables(so.GetObservables())

			//собираем объект ttp
			ttps := datamodels.NewTtpsMessageTheHive()
			ttps.SetTtps(sttp.GetTtps())

			verifiedCase.SetEvent(event)
			verifiedCase.SetObservables(*observables)
			verifiedCase.SetTtps(*ttps)

			mongodbm.ChanInputModule <- mongodbinteractions.SettingsInputChan{
				Section: "handling case",
				Command: "add new case",
				Data:    verifiedCase.Get(),
			}

			esm.ChanInputModule <- elasticsearchinteractions.SettingsInputChan{
				Section: "handling case",
				Command: "add new case",
				Data:    verifiedCase.Get(),
			}

			// ТОЛЬКО ДЛЯ ТЕСТОВ, что бы завершить гроутину вывода информации и логирования
			//при выполнения тестирования
			logging <- datamodels.MessageLogging{
				MsgData: "",
				MsgType: "STOP TEST",
			}

			return
		}
	}
}

// searchEventSource выполняет поиск источника события
func searchEventSource(fieldBranch string, value interface{}) (string, bool) {
	var (
		source string
		ok     bool
	)

	if fieldBranch == "source" {
		source, ok = value.(string)
	}

	return source, ok
}

func checkDatetimeFieldsEventObject(e *datamodels.EventMessageTheHiveCase) {
	if e.GetStartDate() == "" {
		e.SetValueStartDate("1970-01-01T03:00:00+03:00")
	}

	if e.Details.GetEndDate() == "" {
		e.Details.SetValueEndDate("1970-01-01T03:00:00+03:00")
	}

	if e.Object.GetStartDate() == "" {
		e.Object.SetValueStartDate("1970-01-01T03:00:00+03:00")
	}

	if e.Object.GetEndDate() == "" {
		e.Object.SetValueEndDate("1970-01-01T03:00:00+03:00")
	}

	if e.Object.GetCreatedAt() == "" {
		e.Object.SetValueCreatedAt("1970-01-01T03:00:00+03:00")
	}

	if e.Object.GetUpdatedAt() == "" {
		e.Object.SetValueUpdatedAt("1970-01-01T03:00:00+03:00")
	}
}
