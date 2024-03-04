package coremodule

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"
	"placeholder_elasticsearch/listhandlercommon"
	"placeholder_elasticsearch/listhandlerforesjson"
)

func NewVerifiedElasticsearchFormatAlert(
	input <-chan datamodels.ChanOutputDecodeJSON,
	done <-chan bool,
	esm *elasticsearchinteractions.ElasticSearchModule,
	logging chan<- datamodels.MessageLogging,
) {
	var (
		rootId string
		// список не обработанных полей
		listRawFields map[string]string = make(map[string]string)

		event        *datamodels.EventMessageForEsAlert        = datamodels.NewEventMessageForEsAlert()
		eventObject  *datamodels.EventMessageForEsAlertObject  = datamodels.NewEventMessageForEsAlertObject()
		eventDetails *datamodels.EventMessageForEsAlertDetails = datamodels.NewEventMessageForEsAlertDetails()

		alert *datamodels.AlertMessageForEsAlert = datamodels.NewAlertMessageForEsAlert()

		sa listhandlerforesjson.SupportiveAlertArtifacts = *listhandlerforesjson.NewSupportiveAlertArtifacts()

		eventObjectCustomFields datamodels.CustomFields = datamodels.CustomFields{}
		alertObjectCustomFields datamodels.CustomFields = datamodels.CustomFields{}
	)

	// финальный объект
	verifiedAlert := datamodels.NewVerifiedForEsAlert()

	// ------ EVENT ------
	listHandlerEvent := listhandlerforesjson.NewListHandlerEventAlertElement(event)

	// ------ EVENT OBJECT ------
	listHandlerEventObject := listhandlerforesjson.NewListHandlerEventAlertObjectElement(eventObject)

	// ------ EVENT OBJECT CUSTOMFIELDS ------
	listHandlerEventObjectCustomFields := listhandlercommon.NewListHandlerEventObjectCustomFieldsElement(eventObjectCustomFields)

	// ------ EVENT DETAILS ------
	listHandlerEventDetails := listhandlerforesjson.NewListHandlerEventAlertDetailsElement(eventDetails)

	// ------ ALERT ------
	listHandlerAlert := listhandlerforesjson.NewListHandlerAlertElement(alert)

	// ------ ALERT CUSTOMFIELDS ------
	listHandlerAlertCustomFields := listhandlercommon.NewListHandlerAlertCustomFieldsElement(alertObjectCustomFields)

	// ------ ALERT ARTIFACTS ------
	listHandlerAlertArtifacts := listhandlerforesjson.NewListHandlerAlertArtifactsElement(&sa)

	for {
		select {
		case data := <-input:
			var handlerIsExist bool

			verifiedAlert.SetID(data.UUID)

			if source, ok := searchEventSource(data.FieldBranch, data.Value); ok {
				verifiedAlert.SetSource(source)

				continue
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

				continue
			}

			//event.object element
			if lf, ok := listHandlerEventObject[data.FieldBranch]; ok {
				handlerIsExist = true

				for _, f := range lf {
					f(data.Value)
				}

				continue
			}

			//event.object.customFields element
			if lf, ok := listHandlerEventObjectCustomFields[data.FieldBranch]; ok {
				handlerIsExist = true

				for _, f := range lf {
					f(data.Value)
				}

				continue
			}

			//event.details element
			if lf, ok := listHandlerEventDetails[data.FieldBranch]; ok {
				handlerIsExist = true

				for _, f := range lf {
					f(data.Value)
				}

				continue
			}

			//************ Обработчики для Alert ************
			//alert element
			if lf, ok := listHandlerAlert[data.FieldBranch]; ok {
				handlerIsExist = true

				for _, f := range lf {
					f(data.Value)
				}

				continue
			}

			//alert.customFields
			if lf, ok := listHandlerAlertCustomFields[data.FieldBranch]; ok {
				handlerIsExist = true

				for _, f := range lf {
					f(data.Value)
				}

				continue
			}

			//alert.artifacts
			if strings.Contains(data.FieldBranch, "alert.artifacts.") {
				handlerIsExist = true

				if lf, ok := listHandlerAlertArtifacts[data.FieldBranch]; ok {
					for _, f := range lf {
						f(data.Value)
					}
				}

				continue
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

			esm.ChanInputModule <- elasticsearchinteractions.SettingsInputChan{
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
