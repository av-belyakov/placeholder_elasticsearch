package listhandlerthehivejson

import "placeholder_elasticsearch/datamodels"

func NewListHandlerEventAlertElement(event *datamodels.EventMessageTheHiveAlert) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		"event.base":           {event.SetAnyBase},
		"event.startDate":      {event.SetAnyStartDate},
		"event.rootId":         {event.SetAnyRootId},
		"event.objectId":       {event.SetAnyObjectId},
		"event.objectType":     {event.SetAnyObjectType},
		"event.organisation":   {event.SetAnyOrganisation},
		"event.organisationId": {event.SetAnyOrganisationId},
		"event.operation":      {event.SetAnyOperation},
		"event.requestId":      {event.SetAnyRequestId},
	}
}
