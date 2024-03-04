package listhandlerthehivejson

import "placeholder_elasticsearch/datamodels"

func NewListHandlerAlertElement(alert *datamodels.AlertMessageTheHiveAlert) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		"alert.follow":       {alert.SetAnyFollow},
		"alert.tlp":          {alert.SetAnyTlp},
		"alert.severity":     {alert.SetAnySeverity},
		"alert.date":         {alert.SetAnyDate},
		"alert.createdAt":    {alert.SetAnyCreatedAt},
		"alert.updatedAt":    {alert.SetAnyUpdatedAt},
		"alert.createdBy":    {alert.SetAnyCreatedBy},
		"alert.updatedBy":    {alert.SetAnyUpdatedBy},
		"alert.source":       {alert.SetAnySource},
		"alert.title":        {alert.SetAnyTitle},
		"alert.id":           {alert.SetAnyId},
		"alert._id":          {alert.SetAnyUnderliningId},
		"alert.status":       {alert.SetAnyStatus},
		"alert.type":         {alert.SetAnyType},
		"alert._type":        {alert.SetAnyUnderliningType},
		"alert.description":  {alert.SetAnyDescription},
		"alert.caseTemplate": {alert.SetAnyCaseTemplate},
		"alert.sourceRef":    {alert.SetAnySourceRef},
		"alert.tags":         {alert.SetAnyTags},
	}
}
