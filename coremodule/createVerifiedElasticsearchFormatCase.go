package coremodule

import (
	"fmt"
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"
	"placeholder_elasticsearch/listhandlerforesjson"
	"placeholder_elasticsearch/listhandlerthehivejson"
	"reflect"
	"runtime"
	"strings"
)

func NewVerifiedElasticsearchFormatCase(
	input <-chan datamodels.ChanOutputDecodeJSON,
	done <-chan bool,
	esm *elasticsearchinteractions.ElasticSearchModule,
	logging chan<- datamodels.MessageLogging,
) {
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
	sttp := listhandlerforesjson.NewSupportiveTtp()
	listHandlerTtp := listhandlerforesjson.NewListHandlerTtpElement(sttp)

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
	ttps := datamodels.NewTtpsMessageEs()
	ttps.SetValueTtp(sttp.GetTtps())

	verifiedCase.SetEvent(*event)
	verifiedCase.SetObservables(*observables)
	verifiedCase.SetTtps(*ttps)

	esm.ChanInputModule <- elasticsearchinteractions.SettingsInputChan{
		Section: "handling case",
		Command: "add new case",
		Data:    verifiedCase.Get(),
	}
}
