package listhandlerthehivejson

import "placeholder_elasticsearch/datamodels"

// newCustomFieldsElement данный метод, на основании значения objType
// определяет ссылку на какой пользовательский тип datamodels.CustomFieldStringType,
// datamodels.CustomFieldDateType или datamodels.CustomFieldIntegerType добавить
// добавить к map[string]datamodels.CustomerFields. Это сделанно для того что бы
// была возможность подобрать польховательский тип на основе типов свойств.
// Например, нужен тип с полями order int и string типа string, а в другой раз
// тип с полями date string и order int.
func newCustomFieldsElement(elem, objType string, customFields *map[string]datamodels.CustomerFields) {
	if _, ok := (*customFields)[elem]; !ok {
		switch objType {
		case "string":
			(*customFields)[elem] = &datamodels.CustomFieldStringType{}
		case "date":
			(*customFields)[elem] = &datamodels.CustomFieldDateType{}
		case "integer":
			(*customFields)[elem] = &datamodels.CustomFieldIntegerType{}
		default:
			(*customFields)[elem] = &datamodels.CustomFieldStringType{}
		}
	}
}
