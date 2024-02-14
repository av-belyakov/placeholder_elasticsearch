package coremodule

import (
	"fmt"
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"
	"placeholder_elasticsearch/listhandlerthehivejson"
	"placeholder_elasticsearch/mongodbinteractions"
	"strings"
)

func NewVerifiedTheHiveFormatAlert(
	input <-chan datamodels.ChanOutputDecodeJSON,
	done <-chan bool,
	esm *elasticsearchinteractions.ElasticSearchModule,
	mongodbm *mongodbinteractions.MongoDBModule,
	logging chan<- datamodels.MessageLogging,
) {
	var (
		rootId string

		event        datamodels.EventMessageTheHiveAlert = datamodels.EventMessageTheHiveAlert{}
		eventObject  datamodels.EventAlertObject         = datamodels.EventAlertObject{}
		eventDetails datamodels.EventAlertDetails        = datamodels.EventAlertDetails{}
		//eventObjectCustomFields map[string]datamodels.CustomerFields = make(map[string]datamodels.CustomerFields)
		eventObjectCustomFields datamodels.CustomFields = datamodels.CustomFields{}

		alert datamodels.AlertMessageTheHiveAlert = datamodels.AlertMessageTheHiveAlert{}
		//alertObjectCustomFields map[string]datamodels.CustomerFields = make(map[string]datamodels.CustomerFields)
		alertObjectCustomFields datamodels.CustomFields = datamodels.CustomFields{}

		//вспомогательный объект
		sa listhandlerthehivejson.SupportiveAlertArtifacts = *listhandlerthehivejson.NewSupportiveAlertArtifacts()
	)

	//финальный объект
	verifiedAlert := datamodels.VerifiedTheHiveAlert{}

	// ------ EVENT ------
	listHandlerEvent := listhandlerthehivejson.NewListHandlerEventAlertElement(&event)

	// ------ EVENT OBJECT ------
	listHandlerEventObject := listhandlerthehivejson.NewListHandlerEventAlertObjectElement(&eventObject)

	// ------ EVENT OBJECT CUSTOMFIELDS ------
	listHandlerEventObjectCustomFields := listhandlerthehivejson.NewListHandlerAlertCustomFieldsElement(eventObjectCustomFields)

	// ------ EVENT DETAILS ------
	listHandlerEventDetails := listhandlerthehivejson.NewListHandlerEventAlertDetailsElement(&eventDetails)

	// ------ ALERT ------
	listHandlerAlert := listhandlerthehivejson.NewListHandlerAlertElement(&alert)

	// ------ ALERT CUSTOMFIELDS ------
	listHandlerAlertCustomFields := listhandlerthehivejson.NewListHandlerAlertCustomFieldsElement(alertObjectCustomFields)

	// ------ ALERT ARTIFACTS ------
	listHandlerAlertArtifacts := listhandlerthehivejson.NewListHandlerAlertArtifactsElement(&sa)

	for {
		select {
		case data := <-input:
			var handlerIsExist bool

			verifiedAlert.SetID(data.UUID)

			if source, ok := searchEventSource(data.FieldBranch, data.Value); ok {
				verifiedAlert.SetSource(source)
			}

			if data.FieldBranch == "event.rootId" {
				rootId = fmt.Sprint(data.Value)
			}

			//************ Обработчики для Event ************
			//event element
			if lf, ok := listHandlerEvent[data.FieldBranch]; ok {
				handlerIsExist = true

				for _, f := range lf {
					f(data.Value)
				}
			}

			//event.object element
			if lf, ok := listHandlerEventObject[data.FieldBranch]; ok {
				handlerIsExist = true

				for _, f := range lf {
					f(data.Value)
				}
			}

			//event.object.customFields element
			if lf, ok := listHandlerEventObjectCustomFields[data.FieldBranch]; ok {
				handlerIsExist = true

				for _, f := range lf {
					f(data.Value)
				}
			}

			//event.details element
			if lf, ok := listHandlerEventDetails[data.FieldBranch]; ok {
				handlerIsExist = true

				for _, f := range lf {
					f(data.Value)
				}
			}

			//************ Обработчики для Alert ************
			//alert element
			if lf, ok := listHandlerAlert[data.FieldBranch]; ok {
				handlerIsExist = true

				for _, f := range lf {
					f(data.Value)
				}
			}

			//alert.customFields
			if lf, ok := listHandlerAlertCustomFields[data.FieldBranch]; ok {
				handlerIsExist = true

				for _, f := range lf {
					f(data.Value)
				}
			}

			//alert.artifacts
			if strings.Contains(data.FieldBranch, "alert.artifacts.") {
				if lf, ok := listHandlerAlertArtifacts[data.FieldBranch]; ok {
					handlerIsExist = true

					for _, f := range lf {
						f(data.Value)
					}
				}
			}

			// записываем в лог-файл поля, которые не были обработаны
			if handlerIsExist {
				logging <- datamodels.MessageLogging{
					MsgData: fmt.Sprintf("event rootId: '%s', field: '%s', value: '%v'", rootId, data.FieldBranch, data.Value),
					MsgType: "alert_raw_fields",
				}
			}

		case <-done:
			//Собираем объект Alert
			eventObject.SetValueCustomFields(eventObjectCustomFields)

			event.SetValueObject(eventObject)
			event.SetValueDetails(eventDetails)

			alert.SetValueCustomFields(alertObjectCustomFields)
			alert.SetValueArtifacts(sa.GetArtifacts())

			verifiedAlert.SetEvent(event)
			verifiedAlert.SetAlert(alert)

			mongodbm.ChanInputModule <- mongodbinteractions.SettingsInputChan{
				Section: "handling alert",
				Command: "add new alert",
				Data:    verifiedAlert.Get(),
			}

			esm.ChanInputModule <- elasticsearchinteractions.SettingsInputChan{
				Section: "handling alert",
				Command: "add new alert",
				Data:    verifiedAlert.Get(),
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
