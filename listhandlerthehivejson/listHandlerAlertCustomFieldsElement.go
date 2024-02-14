package listhandlerthehivejson

import "placeholder_elasticsearch/datamodels"

// func NewListHandlerAlertCustomFieldsElement(customFields map[string]datamodels.CustomerFields) map[string][]func(interface{}) {
func NewListHandlerAlertCustomFieldsElement(customFields datamodels.CustomFields) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		//--------------- first-time ------------------
		"event.object.customFields.first-time.order": {func(i interface{}) {
			newCustomFieldsElement("first-time", "date", &customFields)
			_, _, _, str := customFields.CustomFields["first-time"].Get()
			customFields.CustomFields["first-time"].Set(i, str)
		}},
		"event.object.customFields.first-time.date": {func(i interface{}) {
			newCustomFieldsElement("first-time", "date", &customFields)
			_, order, _, _ := customFields.CustomFields["first-time"].Get()
			customFields.CustomFields["first-time"].Set(order, i)
		}},
		//--------------- last-time ------------------
		"event.object.customFields.last-time.order": {func(i interface{}) {
			newCustomFieldsElement("last-time", "date", &customFields)
			_, _, _, str := customFields.CustomFields["last-time"].Get()
			customFields.CustomFields["last-time"].Set(i, str)
		}},
		"event.object.customFields.last-time.date": {func(i interface{}) {
			newCustomFieldsElement("last-time", "date", &customFields)
			_, order, _, _ := customFields.CustomFields["last-time"].Get()
			customFields.CustomFields["last-time"].Set(order, i)
		}},
	}
}
