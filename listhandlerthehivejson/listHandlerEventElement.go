package listhandlerthehivejson

import "placeholder_elasticsearch/datamodels"

func NewListHandlerEventCaseElement(event *datamodels.EventMessageTheHiveCase) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		"event.rootId":         {event.SetAnyRootId},
		"event.objectId":       {event.SetAnyObjectId},
		"event.objectType":     {event.SetAnyObjectType},
		"event.base":           {event.SetAnyBase},
		"event.startDate":      {event.SetAnyStartDate},
		"event.requestId":      {event.SetAnyRequestId},
		"event.organisation":   {event.SetAnyOrganisation},
		"event.organisationId": {event.SetAnyOrganisationId},
		"event.operation":      {event.SetAnyOperation},
	}
}
