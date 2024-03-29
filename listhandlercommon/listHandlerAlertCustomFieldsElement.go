package listhandlercommon

import "placeholder_elasticsearch/datamodels"

func NewListHandlerAlertCustomFieldsElement(customFields datamodels.CustomFields) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		//--------------- first-time ------------------
		"alert.customFields.first-time.order": {func(i interface{}) {
			NewCustomFieldsElement("first-time", "date", &customFields)
			_, _, _, str := customFields["first-time"].Get()
			customFields["first-time"].Set(i, str)
		}},
		"alert.customFields.first-time.date": {func(i interface{}) {
			NewCustomFieldsElement("first-time", "date", &customFields)
			_, order, _, _ := customFields["first-time"].Get()
			customFields["first-time"].Set(order, i)
		}},
		//--------------- last-time ------------------
		"alert.customFields.last-time.order": {func(i interface{}) {
			NewCustomFieldsElement("last-time", "date", &customFields)
			_, _, _, str := customFields["last-time"].Get()
			customFields["last-time"].Set(i, str)
		}},
		"alert.customFields.last-time.date": {func(i interface{}) {
			NewCustomFieldsElement("last-time", "date", &customFields)
			_, order, _, _ := customFields["last-time"].Get()
			customFields["last-time"].Set(order, i)
		}},
	}
}
