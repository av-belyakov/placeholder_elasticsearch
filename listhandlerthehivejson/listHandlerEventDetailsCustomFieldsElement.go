package listhandlerthehivejson

import (
	"placeholder_elasticsearch/datamodels"
)

// func NewListHandlerEventDetailsCustomFieldsElement(eventDetailsCustomFields map[string]datamodels.CustomerFields) map[string][]func(interface{}) {
func NewListHandlerEventDetailsCustomFieldsElement(eventDetailsCustomFields datamodels.CustomFields) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		//--------------- attack-type ------------------
		"event.details.customFields.attack-type.order": {func(i interface{}) {
			//создаем элемент "attack-type" если его нет
			newCustomFieldsElement("attack-type", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields.CustomFields["attack-type"].Get()
			eventDetailsCustomFields.CustomFields["attack-type"].Set(i, str)
		}},
		"event.details.customFields.attack-type.string": {func(i interface{}) {
			newCustomFieldsElement("attack-type", "string", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields.CustomFields["attack-type"].Get()
			eventDetailsCustomFields.CustomFields["attack-type"].Set(order, i)
		}},
		//--------------- class-attack ------------------
		"event.details.customFields.class-attack.order": {func(i interface{}) {
			newCustomFieldsElement("class-attack", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields.CustomFields["class-attack"].Get()
			eventDetailsCustomFields.CustomFields["class-attack"].Set(i, str)
		}},
		"event.details.customFields.class-attack.string": {func(i interface{}) {
			newCustomFieldsElement("class-attack", "string", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields.CustomFields["class-attack"].Get()
			eventDetailsCustomFields.CustomFields["class-attack"].Set(order, i)
		}},
		//--------------- event-source ------------------
		"event.details.customFields.event-source.order": {func(i interface{}) {
			newCustomFieldsElement("event-source", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields.CustomFields["event-source"].Get()
			eventDetailsCustomFields.CustomFields["event-source"].Set(i, str)
		}},
		"event.details.customFields.event-source.string": {func(i interface{}) {
			newCustomFieldsElement("event-source", "string", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields.CustomFields["event-source"].Get()
			eventDetailsCustomFields.CustomFields["event-source"].Set(order, i)
		}},
		//--------------- ncircc-class-attack ------------------
		"event.details.customFields.ncircc-class-attack.order": {func(i interface{}) {
			newCustomFieldsElement("ncircc-class-attack", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields.CustomFields["ncircc-class-attack"].Get()
			eventDetailsCustomFields.CustomFields["ncircc-class-attack"].Set(i, str)
		}},
		"event.details.customFields.ncircc-class-attack.string": {func(i interface{}) {
			newCustomFieldsElement("ncircc-class-attack", "string", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields.CustomFields["ncircc-class-attack"].Get()
			eventDetailsCustomFields.CustomFields["ncircc-class-attack"].Set(order, i)
		}},
		//--------------- ncircc-bulletin-id ------------------
		"event.details.customFields.ncircc-bulletin-id.order": {func(i interface{}) {
			newCustomFieldsElement("ncircc-bulletin-id", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields.CustomFields["ncircc-bulletin-id"].Get()
			eventDetailsCustomFields.CustomFields["ncircc-bulletin-id"].Set(i, str)
		}},
		"event.details.customFields.ncircc-bulletin-id.string": {func(i interface{}) {
			newCustomFieldsElement("ncircc-bulletin-id", "string", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields.CustomFields["ncircc-bulletin-id"].Get()
			eventDetailsCustomFields.CustomFields["ncircc-bulletin-id"].Set(order, i)
		}},
		//--------------- sphere ------------------
		"event.details.customFields.sphere.order": {func(i interface{}) {
			newCustomFieldsElement("sphere", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields.CustomFields["sphere"].Get()
			eventDetailsCustomFields.CustomFields["sphere"].Set(i, str)
		}},
		"event.details.customFields.sphere.string": {func(i interface{}) {
			newCustomFieldsElement("sphere", "string", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields.CustomFields["sphere"].Get()
			eventDetailsCustomFields.CustomFields["sphere"].Set(order, i)
		}},
		//--------------- ir-name ------------------
		"event.details.customFields.ir-name.order": {func(i interface{}) {
			newCustomFieldsElement("ir-name", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields.CustomFields["ir-name"].Get()
			eventDetailsCustomFields.CustomFields["ir-name"].Set(i, str)
		}},
		"event.details.customFields.ir-name.string": {func(i interface{}) {
			newCustomFieldsElement("ir-name", "string", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields.CustomFields["ir-name"].Get()
			eventDetailsCustomFields.CustomFields["ir-name"].Set(order, i)
		}},
		//--------------- id-soa ------------------
		"event.details.customFields.id-soa.order": {func(i interface{}) {
			newCustomFieldsElement("id-soa", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields.CustomFields["id-soa"].Get()
			eventDetailsCustomFields.CustomFields["id-soa"].Set(i, str)
		}},
		"event.details.customFields.id-soa.string": {func(i interface{}) {
			newCustomFieldsElement("id-soa", "string", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields.CustomFields["id-soa"].Get()
			eventDetailsCustomFields.CustomFields["id-soa"].Set(order, i)
		}},
		//--------------- state ------------------
		"event.details.customFields.state.order": {func(i interface{}) {
			newCustomFieldsElement("state", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields.CustomFields["state"].Get()
			eventDetailsCustomFields.CustomFields["state"].Set(i, str)
		}},
		"event.details.customFields.state.string": {func(i interface{}) {
			newCustomFieldsElement("state", "string", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields.CustomFields["state"].Get()
			eventDetailsCustomFields.CustomFields["state"].Set(order, i)
		}},
		//--------------- external-letter ------------------
		"event.details.customFields.external-letter.order": {func(i interface{}) {
			newCustomFieldsElement("external-letter", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields.CustomFields["external-letter"].Get()
			eventDetailsCustomFields.CustomFields["external-letter"].Set(i, str)
		}},
		//--------------- inbox1 ------------------
		"event.details.customFields.inbox1.order": {func(i interface{}) {
			newCustomFieldsElement("inbox1", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields.CustomFields["inbox1"].Get()
			eventDetailsCustomFields.CustomFields["inbox1"].Set(i, str)
		}},
		//--------------- inner-letter ------------------
		"event.details.customFields.inner-letter.order": {func(i interface{}) {
			newCustomFieldsElement("inner-letter", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields.CustomFields["inner-letter"].Get()
			eventDetailsCustomFields.CustomFields["inner-letter"].Set(i, str)
		}},
		//--------------- notification ------------------
		"event.details.customFields.notification.order": {func(i interface{}) {
			newCustomFieldsElement("notification", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields.CustomFields["notification"].Get()
			eventDetailsCustomFields.CustomFields["notification"].Set(i, str)
		}},
		//--------------- report ------------------
		"event.details.customFields.report.order": {func(i interface{}) {
			newCustomFieldsElement("report", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields.CustomFields["report"].Get()
			eventDetailsCustomFields.CustomFields["report"].Set(i, str)
		}},
		//--------------- first-time ------------------
		"event.details.customFields.first-time.order": {func(i interface{}) {
			newCustomFieldsElement("first-time", "date", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields.CustomFields["first-time"].Get()
			eventDetailsCustomFields.CustomFields["first-time"].Set(i, str)
		}},
		"event.details.customFields.first-time.date": {func(i interface{}) {
			newCustomFieldsElement("first-time", "date", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields.CustomFields["first-time"].Get()
			eventDetailsCustomFields.CustomFields["first-time"].Set(order, i)
		}},
		//--------------- last-time ------------------
		"event.details.customFields.last-time.order": {func(i interface{}) {
			newCustomFieldsElement("last-time", "date", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields.CustomFields["last-time"].Get()
			eventDetailsCustomFields.CustomFields["last-time"].Set(i, str)
		}},
		"event.details.customFields.last-time.date": {func(i interface{}) {
			newCustomFieldsElement("last-time", "date", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields.CustomFields["last-time"].Get()
			eventDetailsCustomFields.CustomFields["last-time"].Set(order, i)
		}},
		//--------------- b2mid ------------------
		"event.details.customFields.b2mid.order": {func(i interface{}) {
			newCustomFieldsElement("b2mid", "integer", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields.CustomFields["b2mid"].Get()
			eventDetailsCustomFields.CustomFields["b2mid"].Set(i, str)
		}},
		"event.details.customFields.b2mid.integer": {func(i interface{}) {
			newCustomFieldsElement("b2mid", "integer", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields.CustomFields["b2mid"].Get()
			eventDetailsCustomFields.CustomFields["b2mid"].Set(order, i)
		}},
		//--------------- is-incident ------------------
		"event.details.customFields.is-incident.order": {func(i interface{}) {
			newCustomFieldsElement("is-incident", "integer", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields.CustomFields["is-incident"].Get()
			eventDetailsCustomFields.CustomFields["is-incident"].Set(i, str)
		}},
		"event.details.customFields.is-incident.boolean": {func(i interface{}) {
			newCustomFieldsElement("is-incident", "boolean", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields.CustomFields["is-incident"].Get()
			eventDetailsCustomFields.CustomFields["is-incident"].Set(order, i)
		}},
		//--------------- CNC ------------------
		"event.details.customFields.CNC.order": {func(i interface{}) {
			newCustomFieldsElement("CNC", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields.CustomFields["CNC"].Get()
			eventDetailsCustomFields.CustomFields["CNC"].Set(i, str)
		}},
		"event.details.customFields.CNC.string": {func(i interface{}) {
			newCustomFieldsElement("CNC", "string", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields.CustomFields["CNC"].Get()
			eventDetailsCustomFields.CustomFields["CNC"].Set(order, i)
		}},
	}
}
