package listhandlerthehivejson

import "placeholder_elasticsearch/datamodels"

func NewListHandlerEventCaseDetailsElement(details *datamodels.EventCaseDetails) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		"event.details.endDate":          {details.SetAnyEndDate},
		"event.details.resolutionStatus": {details.SetAnyResolutionStatus},
		"event.details.summary":          {details.SetAnySummary},
		"event.details.status":           {details.SetAnyStatus},
		"event.details.impactStatus":     {details.SetAnyImpactStatus},
	}
}
