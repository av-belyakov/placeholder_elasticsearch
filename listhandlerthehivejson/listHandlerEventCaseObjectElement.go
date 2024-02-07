package listhandlerthehivejson

import "placeholder_elasticsearch/datamodels"

func NewListHandlerEventCaseObjectElement(object *datamodels.EventCaseObject) map[string][]func(interface{}) {
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
		"event.object.tags":             {object.SetAnyTags},

		//ниже следующие поля редко используются, думаю пока они не требуют реализации
		//"event.object.stats.impactStatus":    {},
		//"event.object.permissions.id":        {},
		//"event.object.permissions.createdAt": {},
		//"event.object.permissions.pap":       {},
	}
}
