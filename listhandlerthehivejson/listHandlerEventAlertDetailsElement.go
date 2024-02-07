package listhandlerthehivejson

import "placeholder_elasticsearch/datamodels"

func NewListHandlerEventAlertDetailsElement(details *datamodels.EventAlertDetails) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		"event.details.sourceRef":   {details.SetAnySourceRef},
		"event.details.title":       {details.SetAnyTitle},
		"event.details.description": {details.SetAnyDescription},
		"event.details.tags":        {details.SetAnyTags},
	}
}
