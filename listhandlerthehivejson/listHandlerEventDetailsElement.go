package listhandlerthehivejson

import "placeholder_elasticsearch/datamodels"

func NewListHandlerEventDetailsElement(eventDetails *datamodels.EventDetails) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		"event.details.endDate":          {eventDetails.SetAnyEndDate},
		"event.details.resolutionStatus": {eventDetails.SetAnyResolutionStatus},
		"event.details.summary":          {eventDetails.SetAnySummary},
		"event.details.status":           {eventDetails.SetAnyStatus},
		"event.details.impactStatus":     {eventDetails.SetAnyImpactStatus},
	}
}
