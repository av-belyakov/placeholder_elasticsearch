package listhandlerthehivejson

import "placeholder_elasticsearch/datamodels"

func newCustomFieldsElement(elem, objType string, customFields *map[string]datamodels.CustomerFields) {
	if _, ok := (*customFields)[elem]; !ok {
		switch objType {
		case "string":
			(*customFields)[elem] = &datamodels.CustomFieldStringType{}
		case "date":
			(*customFields)[elem] = &datamodels.CustomFieldDateType{}
		case "integer":
			(*customFields)[elem] = &datamodels.CustomFieldIntegerType{}
		}
	}
}
