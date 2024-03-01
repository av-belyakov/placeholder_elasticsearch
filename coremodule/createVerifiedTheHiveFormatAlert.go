package coremodule

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/listhandlercommon"
	"placeholder_elasticsearch/listhandlerthehivejson"
	"placeholder_elasticsearch/mongodbinteractions"
)

func NewVerifiedTheHiveFormatAlert(
	input <-chan datamodels.ChanOutputDecodeJSON,
	done <-chan bool,
	mongodbm *mongodbinteractions.MongoDBModule,
	logging chan<- datamodels.MessageLogging,
) {
	var (
		rootId string
		//список не обработанных полей
		listRawFields map[string]string = make(map[string]string)

		event        *datamodels.EventMessageTheHiveAlert = datamodels.NewEventMessageTheHiveAlert()
		eventObject  *datamodels.EventAlertObject         = datamodels.NewEventAlertObject()
		eventDetails *datamodels.EventAlertDetails        = datamodels.NewEventAlertDetails()

		alert *datamodels.AlertMessageTheHiveAlert = datamodels.NewAlertMessageTheHiveAlert()

		eventObjectCustomFields datamodels.CustomFields = datamodels.CustomFields{}
		alertObjectCustomFields datamodels.CustomFields = datamodels.CustomFields{}

		//вспомогательный объект
		sa listhandlerthehivejson.SupportiveAlertArtifacts = *listhandlerthehivejson.NewSupportiveAlertArtifacts()
	)

	//финальный объект
	verifiedAlert := datamodels.NewVerifiedTheHiveAlert()

	// ------ EVENT ------
	listHandlerEvent := listhandlerthehivejson.NewListHandlerEventAlertElement(event)

	// ------ EVENT OBJECT ------
	listHandlerEventObject := listhandlerthehivejson.NewListHandlerEventAlertObjectElement(eventObject)

	// ------ EVENT OBJECT CUSTOMFIELDS ------
	listHandlerEventObjectCustomFields := listhandlercommon.NewListHandlerAlertCustomFieldsElement(eventObjectCustomFields)

	// ------ EVENT DETAILS ------
	listHandlerEventDetails := listhandlerthehivejson.NewListHandlerEventAlertDetailsElement(eventDetails)

	// ------ ALERT ------
	listHandlerAlert := listhandlerthehivejson.NewListHandlerAlertElement(alert)

	// ------ ALERT CUSTOMFIELDS ------
	listHandlerAlertCustomFields := listhandlercommon.NewListHandlerAlertCustomFieldsElement(alertObjectCustomFields)

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

			if !handlerIsExist {
				// записываем в лог-файл поля, которые не были обработаны
				listRawFields[data.FieldBranch] = fmt.Sprint(data.Value)
			}

		case <-done:
			//Собираем объект Alert
			eventObject.SetValueCustomFields(eventObjectCustomFields)

			event.SetValueObject(*eventObject)
			event.SetValueDetails(*eventDetails)

			alert.SetValueCustomFields(alertObjectCustomFields)
			alert.SetValueArtifacts(sa.GetArtifacts())

			verifiedAlert.SetEvent(*event)
			verifiedAlert.SetAlert(*alert)

			mongodbm.ChanInputModule <- mongodbinteractions.SettingsInputChan{
				Section: "handling alert",
				Command: "add new alert",
				Data:    verifiedAlert.Get(),
			}

			// отправляем список полей которые не смогли обработать
			if len(listRawFields) > 0 {
				logging <- datamodels.MessageLogging{
					MsgData: joinRawFieldsToString(listRawFields, "rootId", rootId),
					MsgType: "alert_raw_fields",
				}
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
