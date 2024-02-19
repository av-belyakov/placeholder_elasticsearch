package listhandlerforesjson

import (
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/listhandlercommon"
)

func NewListHandlerEventAlertDetailsElement(details *datamodels.EventMessageForEsAlertDetails) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		"event.details.sourceRef":   {details.SetAnySourceRef},
		"event.details.title":       {details.SetAnyTitle},
		"event.details.description": {details.SetAnyDescription},
		"event.details.tags": {
			func(i interface{}) {
				key, value := listhandlercommon.HandlerTag(i)
				details.SetAnyTags(key, value)
			},
			details.SetAnyTagsAll,
		},
	}
}
