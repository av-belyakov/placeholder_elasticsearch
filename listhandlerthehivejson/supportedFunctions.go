package listhandlerthehivejson

import "placeholder_elasticsearch/datamodels"

// newCustomFieldsElement данный метод, на основании значения objType
// определяет ссылку на какой пользовательский тип datamodels.CustomFieldStringType,
// datamodels.CustomFieldDateType или datamodels.CustomFieldIntegerType добавить
// добавить к map[string]datamodels.CustomerFields. Это сделанно для того что бы
// была возможность подобрать польховательский тип на основе типов свойств.
// Например, нужен тип с полями order int и string типа string, а в другой раз
// тип с полями date string и order int.
//func newCustomFieldsElement(elem, objType string, customFields *map[string]datamodels.CustomerFields) {
func newCustomFieldsElement(elem, objType string, customFields *datamodels.CustomFields) {
	if _, ok := (*customFields).CustomFields[elem]; !ok {
		switch objType {
		case "string":
			(*customFields).CustomFields[elem] = &datamodels.CustomFieldStringType{}
		case "date":
			(*customFields).CustomFields[elem] = &datamodels.CustomFieldDateType{
				Date: "1970-01-01T03:00:00+03:00",
			}
		case "integer":
			(*customFields).CustomFields[elem] = &datamodels.CustomFieldIntegerType{}
		default:
			(*customFields).CustomFields[elem] = &datamodels.CustomFieldStringType{}
		}
	}
}
