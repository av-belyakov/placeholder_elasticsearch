package listhandlerforesjson

import (
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/listhandlercommon"
)

func NewListHandlerEventAlertObjectElement(object *datamodels.EventMessageForEsAlertObject) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		"event.object.tlp":          {object.SetAnyTlp},
		"event.object._id":          {object.SetAnyUnderliningId},
		"event.object.id":           {object.SetAnyId},
		"event.object.createdBy":    {object.SetAnyCreatedBy},
		"event.object.updatedBy":    {object.SetAnyUpdatedBy},
		"event.object.createdAt":    {object.SetAnyCreatedAt},
		"event.object.updatedAt":    {object.SetAnyUpdatedAt},
		"event.object._type":        {object.SetAnyUnderliningType},
		"event.object.title":        {object.SetAnyTitle},
		"event.object.description":  {object.SetAnyDescription},
		"event.object.status":       {object.SetAnyStatus},
		"event.object.date":         {object.SetAnyDate},
		"event.object.type":         {object.SetAnyType},
		"event.object.objectType":   {object.SetAnyObjectType},
		"event.object.source":       {object.SetAnySource},
		"event.object.sourceRef":    {object.SetAnySourceRef},
		"event.object.case":         {object.SetAnyCase},
		"event.object.caseTemplate": {object.SetAnyCaseTemplate},
		"event.object.tags": {
			func(i interface{}) {
				key, value := listhandlercommon.HandlerTag(i)
				if value == "" {
					return
				}

				object.SetAnyTags(key, value)
			},
			object.SetAnyTagsAll},
	}
}
