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
	listHandlerEventObjectCustomFields := listhandlercommon.NewListHandlerAlertCustomFieldsElement(eventObjectCustomFields)

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
			}

			if data.FieldBranch == "event.rootId" {
				rootId = fmt.Sprint(data.Value)
			}

			//************ Обработчики для Event ************
			//event element
			if lf, ok := listHandlerEvent[data.FieldBranch]; ok {
				for _, f := range lf {
					f(data.Value)
				}
			}

			//event.object element
			if lf, ok := listHandlerEventObject[data.FieldBranch]; ok {
				for _, f := range lf {
					f(data.Value)
				}
			}

			//event.object.customFields element
			if lf, ok := listHandlerEventObjectCustomFields[data.FieldBranch]; ok {
				for _, f := range lf {
					f(data.Value)
				}
			}

			//event.details element
			if lf, ok := listHandlerEventDetails[data.FieldBranch]; ok {
				for _, f := range lf {
					f(data.Value)
				}
			}

			//************ Обработчики для Alert ************
			//alert element
			if lf, ok := listHandlerAlert[data.FieldBranch]; ok {
				for _, f := range lf {
					f(data.Value)
				}
			}

			//alert.customFields
			if lf, ok := listHandlerAlertCustomFields[data.FieldBranch]; ok {
				for _, f := range lf {
					f(data.Value)
				}
			}

			//alert.artifacts
			if strings.Contains(data.FieldBranch, "alert.artifacts.") {
				if lf, ok := listHandlerAlertArtifacts[data.FieldBranch]; ok {
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
