package coremodule

import (
	"fmt"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/listhandlerthehivejson"
)

func NewVerifiedTheHiveFormat(
	input <-chan datamodels.ChanOutputDecodeJSON,
	done <-chan bool,
	//esm *elasticsearchinteractions.ModuleElasticSearch,
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

	for {
		select {
		case data := <-input:
			if source, ok := searchEventSource(data.FieldBranch, data.Value); ok {
				verifiedCase.SetSource(source)
			}

			//*************************************************************
			//********** Сбор всех объектов относящихся к event  **********
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

		case <-done:
			//Собираем объект Event
			eventObject.SetValueCustomFields(eventObjectCustomFields)
			eventDetails.SetValueCustomFields(eventDetailsCustomFields)
			event.SetValueObject(eventObject)
			event.SetValueDetails(eventDetails)

			verifiedCase.SetEvent(event)

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

			fmt.Println("func 'NewVerifiedTheHiveFormat' PROCESSING IS STOPED")

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
