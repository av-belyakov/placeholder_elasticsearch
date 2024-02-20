package listhandlerforesjson

import (
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/listhandlercommon"
)

func NewListHandlerAlertElement(alert *datamodels.AlertMessageForEsAlert) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		"alert.tlp":          {alert.SetAnyTlp},
		"alert.date":         {alert.SetAnyDate},
		"alert.createdAt":    {alert.SetAnyCreatedAt},
		"alert.updatedAt":    {alert.SetAnyUpdatedAt},
		"alert.updatedBy":    {alert.SetAnyUpdatedBy},
		"alert._id":          {alert.SetAnyUnderliningId},
		"alert.status":       {alert.SetAnyStatus},
		"alert.type":         {alert.SetAnyType},
		"alert._type":        {alert.SetAnyUnderliningType},
		"alert.description":  {alert.SetAnyDescription},
		"alert.caseTemplate": {alert.SetAnyCaseTemplate},
		"alert.sourceRef":    {alert.SetAnySourceRef},
		"alert.tags": {
			func(i interface{}) {
				key, value := listhandlercommon.HandlerTag(i)
				if value == "" {
					return
				}

				alert.SetAnyTags(key, value)
			},
			alert.SetAnyTagsAll},
	}
}
