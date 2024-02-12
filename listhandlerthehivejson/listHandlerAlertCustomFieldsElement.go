package listhandlerthehivejson

import "placeholder_elasticsearch/datamodels"

func NewListHandlerAlertCustomFieldsElement(customFields map[string]datamodels.CustomerFields) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		//--------------- first-time ------------------
		"event.object.customFields.first-time.order": {func(i interface{}) {
			newCustomFieldsElement("first-time", "date", &customFields)
			_, _, _, str := customFields["first-time"].Get()
			customFields["first-time"].Set(i, str)
		}},
		"event.object.customFields.first-time.date": {func(i interface{}) {
			newCustomFieldsElement("first-time", "date", &customFields)
			_, order, _, _ := customFields["first-time"].Get()
			customFields["first-time"].Set(order, i)
		}},
		//--------------- last-time ------------------
		"event.object.customFields.last-time.order": {func(i interface{}) {
			newCustomFieldsElement("last-time", "date", &customFields)
			_, _, _, str := customFields["last-time"].Get()
			customFields["last-time"].Set(i, str)
		}},
		"event.object.customFields.last-time.date": {func(i interface{}) {
			newCustomFieldsElement("last-time", "date", &customFields)
			_, order, _, _ := customFields["last-time"].Get()
			customFields["last-time"].Set(order, i)
		}},
	}
}