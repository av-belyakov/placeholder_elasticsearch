package listhandlerforesjson

import (
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/listhandlercommon"
	"strings"
)

func NewListHandlerEventCaseObjectElement(object *datamodels.EventForEsCaseObject) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		"event.object.flag":             {object.SetAnyFlag},
		"event.object.caseId":           {object.SetAnyCaseId},
		"event.object.severity":         {object.SetAnySeverity},
		"event.object.tlp":              {object.SetAnyTlp},
		"event.object.pap":              {object.SetAnyPap},
		"event.object.startDate":        {object.SetAnyStartDate},
		"event.object.endDate":          {object.SetAnyEndDate},
		"event.object.createdAt":        {object.SetAnyCreatedAt},
		"event.object.updatedAt":        {object.SetAnyUpdatedAt},
		"event.object._id":              {object.SetAnyUnderliningId},
		"event.object.id":               {object.SetAnyId},
		"event.object.createdBy":        {object.SetAnyCreatedBy},
		"event.object.updatedBy":        {object.SetAnyUpdatedBy},
		"event.object._type":            {object.SetAnyUnderliningType},
		"event.object.title":            {object.SetAnyTitle},
		"event.object.description":      {object.SetAnyDescription},
		"event.object.impactStatus":     {object.SetAnyImpactStatus},
		"event.object.resolutionStatus": {object.SetAnyResolutionStatus},
		"event.object.status":           {object.SetAnyStatus},
		"event.object.summary":          {object.SetAnySummary},
		"event.object.owner":            {object.SetAnyOwner},
		"event.object.tags": {
			func(i interface{}) {
				key, value := listhandlercommon.HandlerTag(i)
				if value == "" {
					return
				}

				value = strings.TrimSpace(value)
				value = strings.Trim(value, "\"")
				object.SetAnyTags(key, value)
			},
			object.SetAnyTagsAll},
	}
}
