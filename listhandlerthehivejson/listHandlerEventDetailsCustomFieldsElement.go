package listhandlerthehivejson

import "placeholder_elasticsearch/datamodels"

func NewListHandlerEventDetailsCustomFieldsElement(eventDetailsCustomFields map[string]datamodels.CustomerFields) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		//--- attack-type ---
		"event.details.customFields.attack-type.order": {func(i interface{}) {
			//создаем элемент "attack-type" если его нет
			newCustomFieldsElement("attack-type", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["attack-type"].Get()
			eventDetailsCustomFields["attack-type"].Set(i, str)
		}},
		"event.details.customFields.attack-type.string": {func(i interface{}) {
			newCustomFieldsElement("attack-type", "string", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields["attack-type"].Get()
			eventDetailsCustomFields["attack-type"].Set(order, i)
		}},
		//--- class-attack ---
		"event.details.customFields.class-attack.order": {func(i interface{}) {
			newCustomFieldsElement("class-attack", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["class-attack"].Get()
			eventDetailsCustomFields["class-attack"].Set(i, str)
		}},
		"event.details.customFields.class-attack.string": {func(i interface{}) {
			newCustomFieldsElement("class-attack", "string", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields["class-attack"].Get()
			eventDetailsCustomFields["class-attack"].Set(order, i)
		}},
		//--- event-source ---
		"event.details.customFields.event-source.order": {func(i interface{}) {
			newCustomFieldsElement("event-source", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["event-source"].Get()
			eventDetailsCustomFields["event-source"].Set(i, str)
		}},
		"event.details.customFields.event-source.string": {func(i interface{}) {
			newCustomFieldsElement("event-source", "string", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields["event-source"].Get()
			eventDetailsCustomFields["event-source"].Set(order, i)
		}},
		//--- ncircc-class-attack ---
		"event.details.customFields.ncircc-class-attack.order": {func(i interface{}) {
			newCustomFieldsElement("ncircc-class-attack", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["ncircc-class-attack"].Get()
			eventDetailsCustomFields["ncircc-class-attack"].Set(i, str)
		}},
		"event.details.customFields.ncircc-class-attack.string": {func(i interface{}) {
			newCustomFieldsElement("ncircc-class-attack", "string", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields["ncircc-class-attack"].Get()
			eventDetailsCustomFields["ncircc-class-attack"].Set(order, i)
		}},
		//--- ncircc-bulletin-id ---
		"event.details.customFields.ncircc-bulletin-id.order": {func(i interface{}) {
			newCustomFieldsElement("ncircc-bulletin-id", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["ncircc-bulletin-id"].Get()
			eventDetailsCustomFields["ncircc-bulletin-id"].Set(i, str)
		}},
		"event.details.customFields.ncircc-bulletin-id.string": {func(i interface{}) {
			newCustomFieldsElement("ncircc-bulletin-id", "string", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields["ncircc-bulletin-id"].Get()
			eventDetailsCustomFields["ncircc-bulletin-id"].Set(order, i)
		}},
		//--- sphere ---
		"event.details.customFields.sphere.order": {func(i interface{}) {
			newCustomFieldsElement("sphere", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["sphere"].Get()
			eventDetailsCustomFields["sphere"].Set(i, str)
		}},
		"event.details.customFields.sphere.string": {func(i interface{}) {
			newCustomFieldsElement("sphere", "string", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields["sphere"].Get()
			eventDetailsCustomFields["sphere"].Set(order, i)
		}},
		//--- ir-name ---
		"event.details.customFields.ir-name.order": {func(i interface{}) {
			newCustomFieldsElement("ir-name", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["ir-name"].Get()
			eventDetailsCustomFields["ir-name"].Set(i, str)
		}},
		"event.details.customFields.ir-name.string": {func(i interface{}) {
			newCustomFieldsElement("ir-name", "string", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields["ir-name"].Get()
			eventDetailsCustomFields["ir-name"].Set(order, i)
		}},
		//--- id-soa ---
		"event.details.customFields.id-soa.order": {func(i interface{}) {
			newCustomFieldsElement("id-soa", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["id-soa"].Get()
			eventDetailsCustomFields["id-soa"].Set(i, str)
		}},
		"event.details.customFields.id-soa.string": {func(i interface{}) {
			newCustomFieldsElement("id-soa", "string", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields["id-soa"].Get()
			eventDetailsCustomFields["id-soa"].Set(order, i)
		}},
		//--- state ---
		"event.details.customFields.state.order": {func(i interface{}) {
			newCustomFieldsElement("state", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["state"].Get()
			eventDetailsCustomFields["state"].Set(i, str)
		}},
		"event.details.customFields.state.string": {func(i interface{}) {
			newCustomFieldsElement("state", "string", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields["state"].Get()
			eventDetailsCustomFields["state"].Set(order, i)
		}},
		//--- external-letter ---
		"event.details.customFields.external-letter.order": {func(i interface{}) {
			newCustomFieldsElement("external-letter", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["external-letter"].Get()
			eventDetailsCustomFields["external-letter"].Set(i, str)
		}},
		//--- inbox1 ---
		"event.details.customFields.inbox1.order": {func(i interface{}) {
			newCustomFieldsElement("inbox1", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["inbox1"].Get()
			eventDetailsCustomFields["inbox1"].Set(i, str)
		}},
		//--- inner-letter ---
		"event.details.customFields.inner-letter.order": {func(i interface{}) {
			newCustomFieldsElement("inner-letter", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["inner-letter"].Get()
			eventDetailsCustomFields["inner-letter"].Set(i, str)
		}},
		//--- notification ---
		"event.details.customFields.notification.order": {func(i interface{}) {
			newCustomFieldsElement("notification", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["notification"].Get()
			eventDetailsCustomFields["notification"].Set(i, str)
		}},
		//--- report ---
		"event.details.customFields.report.order": {func(i interface{}) {
			newCustomFieldsElement("report", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["report"].Get()
			eventDetailsCustomFields["report"].Set(i, str)
		}},
		//--- first-time ---
		"event.details.customFields.first-time.order": {func(i interface{}) {
			newCustomFieldsElement("first-time", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["first-time"].Get()
			eventDetailsCustomFields["first-time"].Set(i, str)
		}},
		"event.details.customFields.first-time.date": {func(i interface{}) {
			newCustomFieldsElement("first-time", "date", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields["first-time"].Get()
			eventDetailsCustomFields["first-time"].Set(order, i)
		}},
		//--- last-time ---
		"event.details.customFields.last-time.order": {func(i interface{}) {
			newCustomFieldsElement("last-time", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["last-time"].Get()
			eventDetailsCustomFields["last-time"].Set(i, str)
		}},
		"event.details.customFields.last-time.date": {func(i interface{}) {
			newCustomFieldsElement("last-time", "date", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields["last-time"].Get()
			eventDetailsCustomFields["last-time"].Set(order, i)
		}},
		//--- b2mid ---
		"event.details.customFields.b2mid.order": {func(i interface{}) {
			newCustomFieldsElement("b2mid", "integer", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["b2mid"].Get()
			eventDetailsCustomFields["b2mid"].Set(i, str)
		}},
		"event.details.customFields.b2mid.integer": {func(i interface{}) {
			newCustomFieldsElement("b2mid", "integer", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields["b2mid"].Get()
			eventDetailsCustomFields["b2mid"].Set(order, i)
		}},
		//--- is-incident ---
		"event.details.customFields.is-incident.order": {func(i interface{}) {
			newCustomFieldsElement("is-incident", "integer", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["is-incident"].Get()
			eventDetailsCustomFields["is-incident"].Set(i, str)
		}},
		"event.details.customFields.is-incident.boolean": {func(i interface{}) {
			newCustomFieldsElement("is-incident", "boolean", &eventDetailsCustomFields)
			_, order, _, _ := eventDetailsCustomFields["is-incident"].Get()
			eventDetailsCustomFields["is-incident"].Set(order, i)
		}},
		//--- CNC ---
		"event.details.customFields.CNC.order": {func(i interface{}) {
			newCustomFieldsElement("CNC", "integer", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["CNC"].Get()
			eventDetailsCustomFields["CNC"].Set(i, str)
		}},
		"event.details.customFields.CNC.string": {func(i interface{}) {
			newCustomFieldsElement("CNC", "string", &eventDetailsCustomFields)
			_, _, _, str := eventDetailsCustomFields["CNC"].Get()
			eventDetailsCustomFields["CNC"].Set(i, str)
		}},
	}
}
