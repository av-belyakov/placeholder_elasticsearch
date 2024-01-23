package listhandlerthehivejson

import "placeholder_elasticsearch/datamodels"

func NewListHandlerEventObjectElement(eventObject *datamodels.EventObject) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		"event.object.flag":             {eventObject.SetAnyFlag},
		"event.object.caseId":           {eventObject.SetAnyCaseId},
		"event.object.severity":         {eventObject.SetAnySeverity},
		"event.object.tlp":              {eventObject.SetAnyTlp},
		"event.object.pap":              {eventObject.SetAnyPap},
		"event.object.startDate":        {eventObject.SetAnyStartDate},
		"event.object.endDate":          {eventObject.SetAnyEndDate},
		"event.object.createdAt":        {eventObject.SetAnyCreatedAt},
		"event.object.updatedAt":        {eventObject.SetAnyUpdatedAt},
		"event.object._id":              {eventObject.SetAnyUnderliningId},
		"event.object.id":               {eventObject.SetAnyId},
		"event.object.createdBy":        {eventObject.SetAnyCreatedBy},
		"event.object.updatedBy":        {eventObject.SetAnyUpdatedBy},
		"event.object._type":            {eventObject.SetAnyUnderliningType},
		"event.object.title":            {eventObject.SetAnyTitle},
		"event.object.description":      {eventObject.SetAnyDescription},
		"event.object.impactStatus":     {eventObject.SetAnyImpactStatus},
		"event.object.resolutionStatus": {eventObject.SetAnyResolutionStatus},
		"event.object.status":           {eventObject.SetAnyStatus},
		"event.object.summary":          {eventObject.SetAnySummary},
		"event.object.owner":            {eventObject.SetAnyOwner},
		"event.object.tags":             {eventObject.SetAnyTags},

		//ниже следующие поля редко используются, думаю пока они не требуют реализации
		//"event.object.stats.impactStatus":    {},
		//"event.object.permissions.id":        {},
		//"event.object.permissions.createdAt": {},
		//"event.object.permissions.pap":       {},
	}
}
