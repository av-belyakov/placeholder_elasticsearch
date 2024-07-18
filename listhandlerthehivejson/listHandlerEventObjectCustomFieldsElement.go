package listhandlerthehivejson

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/listhandlercommon"
	"placeholder_elasticsearch/supportingfunctions"
)

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

			//создаем элемент "geoip" если его нет
			listhandlercommon.NewCustomFieldsElement("geoip", "string", &eventObjectCustomFields)
			eventObjectCustomFields["geoip"].Set(0, supportingfunctions.TrimIsNotLetter(tmp[1]))
		}},

		//------------------ attack-type ------------------
		"event.object.customFields.attack-type.order": {func(i interface{}) {
			//создаем элемент "attack-type" если его нет
			listhandlercommon.NewCustomFieldsElement("attack-type", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["attack-type"].Get()
			eventObjectCustomFields["attack-type"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.attack-type.string": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("attack-type", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["attack-type"].Get()
			eventObjectCustomFields["attack-type"].Set(order, i)
		}},
		//------------------ class-attack ------------------
		"event.object.customFields.class-attack.order": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("class-attack", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["class-attack"].Get()
			eventObjectCustomFields["class-attack"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.class-attack.string": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("class-attack", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["class-attack"].Get()
			eventObjectCustomFields["class-attack"].Set(order, i)
		}},
		//------------------ ncircc-class-attack ------------------
		"event.object.customFields.ncircc-class-attack.order": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("ncircc-class-attack", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["ncircc-class-attack"].Get()
			eventObjectCustomFields["ncircc-class-attack"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.ncircc-class-attack.string": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("ncircc-class-attack", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["ncircc-class-attack"].Get()
			eventObjectCustomFields["ncircc-class-attack"].Set(order, i)
		}},
		//------------------ inbox1 ------------------
		"event.object.customFields.inbox1.order": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("inbox1", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["inbox1"].Get()
			eventObjectCustomFields["inbox1"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		//------------------ inner-letter ------------------
		"event.object.customFields.inner-letter.order": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("inner-letter", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["inner-letter"].Get()
			eventObjectCustomFields["inner-letter"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		//------------------ notification ------------------
		"event.object.customFields.notification.order": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("notification", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["notification"].Get()
			eventObjectCustomFields["notification"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		//------------------ report ------------------
		"event.object.customFields.report.order": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("report", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["report"].Get()
			eventObjectCustomFields["report"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		//------------------ first-time ------------------
		"event.object.customFields.first-time.order": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("first-time", "date", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["first-time"].Get()
			eventObjectCustomFields["first-time"].Set(i, str)
		}},
		"event.object.customFields.first-time.date": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("first-time", "date", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["first-time"].Get()
			eventObjectCustomFields["first-time"].Set(order, i)
		}},
		//------------------ last-time ------------------
		"event.object.customFields.last-time.order": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("last-time", "date", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["last-time"].Get()
			eventObjectCustomFields["last-time"].Set(i, str)
		}},
		"event.object.customFields.last-time.date": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("last-time", "date", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["last-time"].Get()
			eventObjectCustomFields["last-time"].Set(order, i)
		}},
		//------------------ sphere ------------------
		"event.object.customFields.sphere.order": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("sphere", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["sphere"].Get()
			eventObjectCustomFields["sphere"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.sphere.string": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("sphere", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["sphere"].Get()
			eventObjectCustomFields["sphere"].Set(order, i)
		}},
		//------------------ state ------------------
		"event.object.customFields.state.order": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("state", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["state"].Get()
			eventObjectCustomFields["state"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.state.string": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("state", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["state"].Get()
			eventObjectCustomFields["state"].Set(order, i)
		}},
		//------------------ ir-name ------------------
		"event.object.customFields.ir-name.order": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("ir-name", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["ir-name"].Get()
			eventObjectCustomFields["ir-name"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.ir-name.string": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("ir-name", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["ir-name"].Get()
			eventObjectCustomFields["ir-name"].Set(order, i)
		}},
		//------------------ id-soa ------------------
		"event.object.customFields.id-soa.order": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("id-soa", "string", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["id-soa"].Get()
			eventObjectCustomFields["id-soa"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.id-soa.string": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("id-soa", "string", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["id-soa"].Get()
			eventObjectCustomFields["id-soa"].Set(order, i)
		}},
		//--------------- is-incident ------------------
		"event.object.customFields.is-incident.order": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("is-incident", "boolen", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["is-incident"].Get()
			eventObjectCustomFields["is-incident"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.is-incident.boolean": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("is-incident", "boolean", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["is-incident"].Get()
			eventObjectCustomFields["is-incident"].Set(order, i)
		}},
		//--------------- work-admin ------------------
		"event.object.customFields.work-admin.order": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("work-admin", "boolen", &eventObjectCustomFields)
			_, _, _, str := eventObjectCustomFields["work-admin"].Get()
			eventObjectCustomFields["work-admin"].Set(i, supportingfunctions.TrimIsNotLetter(str))
		}},
		"event.object.customFields.work-admin.boolean": {func(i interface{}) {
			listhandlercommon.NewCustomFieldsElement("work-admin", "boolean", &eventObjectCustomFields)
			_, order, _, _ := eventObjectCustomFields["work-admin"].Get()
			eventObjectCustomFields["work-admin"].Set(order, i)
		}},
	}
}
