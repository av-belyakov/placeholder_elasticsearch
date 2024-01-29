package coremodule

import (
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"
	"placeholder_elasticsearch/listhandlerthehivejson"
	"reflect"
	"strings"
)

func NewVerifiedTheHiveFormat(
	input <-chan datamodels.ChanOutputDecodeJSON,
	done <-chan bool,
	esm *elasticsearchinteractions.ElasticSearchModule,
	logging chan<- datamodels.MessageLogging,
) {
	var (
		//Финальный объект
		verifiedCase *datamodels.VerifiedTheHiveCase = datamodels.NewVerifiedTheHiveCase()

		event        datamodels.EventMessageTheHive = datamodels.EventMessageTheHive{}
		eventObject  datamodels.EventObject         = datamodels.EventObject{}
		eventDetails datamodels.EventDetails        = datamodels.EventDetails{}

		eventObjectCustomFields  map[string]datamodels.CustomerFields = make(map[string]datamodels.CustomerFields)
		eventDetailsCustomFields map[string]datamodels.CustomerFields = make(map[string]datamodels.CustomerFields)
	)

	//******************* Основные обработчики для Event **********************
	// ------ EVENT ------
	listHandlerEvent := listhandlerthehivejson.NewListHandlerEventElement(&event)
	// ------ EVENT OBJECT ------
	listHandlerEventObject := listhandlerthehivejson.NewListHandlerEventObjectElement(&eventObject)
	// ------ EVENT OBJECT CUSTOMFIELDS ------
	listHandlerEventObjectCustomFields := listhandlerthehivejson.NewListHandlerEventObjectCustomFieldsElement(eventObjectCustomFields)
	// ------ EVENT DETAILS ------
	listHandlerEventDetails := listhandlerthehivejson.NewListHandlerEventDetailsElement(&eventDetails)
	// ------ EVENT DETAILS CUSTOMFIELDS ------
	listHandlerEventDetailsCustomFields := listhandlerthehivejson.NewListHandlerEventDetailsCustomFieldsElement(eventDetailsCustomFields)

	//******************* Вспомогательный объект для Observables **********************
	so := listhandlerthehivejson.NewSupportiveObservables()
	listHandlerObservables := listhandlerthehivejson.NewListHandlerObservablesElement(so)

	for {
		select {
		case data := <-input:
			if source, ok := searchEventSource(data.FieldBranch, data.Value); ok {
				verifiedCase.SetSource(source)
			}

			//******************************************************************
			//********** Сбор всех объектов относящихся к полю Event  **********
			// event element
			if lf, ok := listHandlerEvent[data.FieldBranch]; ok {
				for _, f := range lf {
					f(data.Value)
				}
			}

			// event.object element
			if lf, ok := listHandlerEventObject[data.FieldBranch]; ok {
				for _, f := range lf {
					f(data.Value)
				}
			}

			// event.object.customFields element
			if lf, ok := listHandlerEventObjectCustomFields[data.FieldBranch]; ok {
				for _, f := range lf {
					f(data.Value)
				}
			}

			// event.details element
			if lf, ok := listHandlerEventDetails[data.FieldBranch]; ok {
				for _, f := range lf {
					f(data.Value)
				}
			}

			// event.details.customFields element
			if lf, ok := listHandlerEventDetailsCustomFields[data.FieldBranch]; ok {
				for _, f := range lf {
					f(data.Value)
				}
			}

			//************************************************************************
			//********** Сбор всех объектов относящихся к полю Observables  **********
			// для всех полей входящих в observables, кроме содержимого
			//поля reports
			if lf, ok := listHandlerObservables[data.FieldBranch]; ok {
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
				so.HandlerReportValue(data.FieldBranch, data.Value)
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

			verifiedCase.SetEvent(event)
			verifiedCase.SetObservables(*observables)

			esm.ChanInputModule <- elasticsearchinteractions.SettingsInputChan{
				Command:        "add new object",
				VerifiedObject: verifiedCase.Get(),
			}

			/*bytes, err := json.Marshal(struct {
				datamodels.SourceMessageTheHive
				datamodels.ObservablesMessageTheHive
				event datamodels.EventMessageTheHive
			}{
				source,
				observables,
				event,
			})
			if err != nil {
				_, f, l, _ := runtime.Caller(0)
				logging <- datamodels.MessageLogging{
					MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-10),
					MsgType: "error",
				}

				return
			}*/

			//Тут надо отправить в модули Elasticsearch и MongoDB

			//esm.HandlerData(elasticsearchinteractions.SettingsInputChan{
			//	UUID: taskId,
			//	Data: bytes,
			//})

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

func checkDatetimeFieldsEventObject(e *datamodels.EventMessageTheHive) {
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
