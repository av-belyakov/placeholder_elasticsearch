package listhandlerthehivejson

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/datamodels"
)

// func NewListHandlerEventObjectCustomFieldsElement(eventObjectCustomFields map[string]datamodels.CustomerFields) map[string][]func(interface{}) {
func NewListHandlerEventObjectCustomFieldsElement(eventObjectCustomFields datamodels.CustomFields) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		//------------- для обработки тегов содержащих geoip -------------
		"event.object.tags": {func(i interface{}) {
			s := fmt.Sprint(i)
			if !strings.Contains(s, "geoip") {
				return
			}

			tmp := strings.Split(s, "=")
			if len(tmp) < 2 {
				return
			}

			countru := strings.Trim(tmp[1], "\"")

			//создаем элемент "geoip" если его нет
			newCustomFieldsElement("geoip", "string", &eventObjectCustomFields)
			eventObjectCustomFields.CustomFields["geoip"].Set(0, countru)
		}},

		//------------------ attack-type ------------------
		"event.object.customFields.attack-type.order": {func(i interface{}) {
			//создаем элемент "attack-type" если его нет
			newCustomFieldsElement("attack-type", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields.CustomFields["attack-type"].Get()
			eventObjectCustomFields.CustomFields["attack-type"].Set(i, str)
		}},
		"event.object.customFields.attack-type.string": {func(i interface{}) {
			newCustomFieldsElement("attack-type", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields.CustomFields["attack-type"].Get()
			eventObjectCustomFields.CustomFields["attack-type"].Set(order, i)
		}},
		//------------------ class-attack ------------------
		"event.object.customFields.class-attack.order": {func(i interface{}) {
			newCustomFieldsElement("class-attack", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields.CustomFields["class-attack"].Get()
			eventObjectCustomFields.CustomFields["class-attack"].Set(i, str)
		}},
		"event.object.customFields.class-attack.string": {func(i interface{}) {
			newCustomFieldsElement("class-attack", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields.CustomFields["class-attack"].Get()
			eventObjectCustomFields.CustomFields["class-attack"].Set(order, i)
		}},
		//------------------ ncircc-class-attack ------------------
		"event.object.customFields.ncircc-class-attack.order": {func(i interface{}) {
			newCustomFieldsElement("ncircc-class-attack", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields.CustomFields["ncircc-class-attack"].Get()
			eventObjectCustomFields.CustomFields["ncircc-class-attack"].Set(i, str)
		}},
		"event.object.customFields.ncircc-class-attack.string": {func(i interface{}) {
			newCustomFieldsElement("ncircc-class-attack", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields.CustomFields["ncircc-class-attack"].Get()
			eventObjectCustomFields.CustomFields["ncircc-class-attack"].Set(order, i)
		}},
		//------------------ inbox1 ------------------
		"event.object.customFields.inbox1.order": {func(i interface{}) {
			newCustomFieldsElement("inbox1", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields.CustomFields["inbox1"].Get()
			eventObjectCustomFields.CustomFields["inbox1"].Set(i, str)
		}},
		//------------------ inner-letter ------------------
		"event.object.customFields.inner-letter.order": {func(i interface{}) {
			newCustomFieldsElement("inner-letter", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields.CustomFields["inner-letter"].Get()
			eventObjectCustomFields.CustomFields["inner-letter"].Set(i, str)
		}},
		//------------------ notification ------------------
		"event.object.customFields.notification.order": {func(i interface{}) {
			newCustomFieldsElement("notification", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields.CustomFields["notification"].Get()
			eventObjectCustomFields.CustomFields["notification"].Set(i, str)
		}},
		//------------------ report ------------------
		"event.object.customFields.report.order": {func(i interface{}) {
			newCustomFieldsElement("report", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields.CustomFields["report"].Get()
			eventObjectCustomFields.CustomFields["report"].Set(i, str)
		}},
		//------------------ first-time ------------------
		"event.object.customFields.first-time.order": {func(i interface{}) {
			newCustomFieldsElement("first-time", "date", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields.CustomFields["first-time"].Get()
			eventObjectCustomFields.CustomFields["first-time"].Set(i, str)
		}},
		"event.object.customFields.first-time.date": {func(i interface{}) {
			newCustomFieldsElement("first-time", "date", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields.CustomFields["first-time"].Get()
			eventObjectCustomFields.CustomFields["first-time"].Set(order, i)
		}},
		//------------------ last-time ------------------
		"event.object.customFields.last-time.order": {func(i interface{}) {
			newCustomFieldsElement("last-time", "date", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields.CustomFields["last-time"].Get()
			eventObjectCustomFields.CustomFields["last-time"].Set(i, str)
		}},
		"event.object.customFields.last-time.date": {func(i interface{}) {
			newCustomFieldsElement("last-time", "date", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields.CustomFields["last-time"].Get()
			eventObjectCustomFields.CustomFields["last-time"].Set(order, i)
		}},
		//------------------ sphere ------------------
		"event.object.customFields.sphere.order": {func(i interface{}) {
			newCustomFieldsElement("sphere", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields.CustomFields["sphere"].Get()
			eventObjectCustomFields.CustomFields["sphere"].Set(i, str)
		}},
		"event.object.customFields.sphere.string": {func(i interface{}) {
			newCustomFieldsElement("sphere", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields.CustomFields["sphere"].Get()
			eventObjectCustomFields.CustomFields["sphere"].Set(order, i)
		}},
		//------------------ state ------------------
		"event.object.customFields.state.order": {func(i interface{}) {
			newCustomFieldsElement("state", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields.CustomFields["state"].Get()
			eventObjectCustomFields.CustomFields["state"].Set(i, str)
		}},
		"event.object.customFields.state.string": {func(i interface{}) {
			newCustomFieldsElement("state", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields.CustomFields["state"].Get()
			eventObjectCustomFields.CustomFields["state"].Set(order, i)
		}},
		//------------------ ir-name ------------------
		"event.object.customFields.ir-name.order": {func(i interface{}) {
			newCustomFieldsElement("ir-name", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields.CustomFields["ir-name"].Get()
			eventObjectCustomFields.CustomFields["ir-name"].Set(i, str)
		}},
		"event.object.customFields.ir-name.string": {func(i interface{}) {
			newCustomFieldsElement("ir-name", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields.CustomFields["ir-name"].Get()
			eventObjectCustomFields.CustomFields["ir-name"].Set(order, i)
		}},
		//------------------ id-soa ------------------
		"event.object.customFields.id-soa.order": {func(i interface{}) {
			newCustomFieldsElement("id-soa", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields.CustomFields["id-soa"].Get()
			eventObjectCustomFields.CustomFields["id-soa"].Set(i, str)
		}},
		"event.object.customFields.id-soa.string": {func(i interface{}) {
			newCustomFieldsElement("id-soa", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields.CustomFields["id-soa"].Get()
			eventObjectCustomFields.CustomFields["id-soa"].Set(order, i)
		}},
	}
}
