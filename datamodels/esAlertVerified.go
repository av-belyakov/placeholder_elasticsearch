package datamodels

import (
	"fmt"
	"strings"
	"time"

	"placeholder_elasticsearch/supportingfunctions"
)

// VerifiedForEsAlert объект представляет собой верифицированный и
// модифицированный Alert для загрузки в СУБД Elasticsearch
type VerifiedForEsAlert struct {
	ID              string                 `json:"@id" bson:"@id"`
	ElasticsearchID string                 `json:"@es_id" bson:"@es_id"`
	CreateTimestamp string                 `json:"@timestamp" bson:"@timestamp"`
	Source          string                 `json:"source" bson:"source"`
	Event           EventMessageForEsAlert `json:"event" bson:"event"`
	Alert           AlertMessageForEsAlert `json:"alert,omitempty" bson:"alert"`
}

func NewVerifiedForEsAlert() *VerifiedForEsAlert {
	return &VerifiedForEsAlert{
		CreateTimestamp: supportingfunctions.GetDateTimeFormatRFC3339(time.Now().UnixMilli()),
		Event:           *NewEventMessageForEsAlert(),
		Alert:           *NewAlertMessageForEsAlert(),
	}
}

func (va *VerifiedForEsAlert) Get() *VerifiedForEsAlert {
	return va
}

func (va *VerifiedForEsAlert) GetID() string {
	return va.ID
}

func (va *VerifiedForEsAlert) SetID(id string) {
	va.ID = id
}

func (va *VerifiedForEsAlert) GetElasticsearchID() string {
	return va.ElasticsearchID
}

func (va *VerifiedForEsAlert) SetElasticsearchID(id string) {
	va.ElasticsearchID = id
}

func (va *VerifiedForEsAlert) GetCreateTimestatmp() string {
	return va.CreateTimestamp
}

func (va *VerifiedForEsAlert) SetCreateTimestatmp(time string) {
	va.CreateTimestamp = time
}

func (va *VerifiedForEsAlert) GetSource() string {
	return va.Source
}

func (va *VerifiedForEsAlert) SetSource(source string) {
	va.Source = source
}

func (va *VerifiedForEsAlert) GetEvent() *EventMessageForEsAlert {
	return &va.Event
}

func (va *VerifiedForEsAlert) SetEvent(event EventMessageForEsAlert) {
	va.Event = event
}

func (va *VerifiedForEsAlert) GetAlert() *AlertMessageForEsAlert {
	return &va.Alert
}

func (va *VerifiedForEsAlert) SetAlert(alert AlertMessageForEsAlert) {
	va.Alert = alert
}

func (va *VerifiedForEsAlert) ToStringBeautiful(num int) string {
	ws := supportingfunctions.GetWhitespace(num)

	strB := strings.Builder{}
	strB.WriteString(fmt.Sprintf("%s'createTimestatmp': '%s'\n", ws, va.CreateTimestamp))
	strB.WriteString(fmt.Sprintf("%s'source': '%s'\n", ws, va.Source))
	strB.WriteString(fmt.Sprintf("%s'event':\n", ws))
	strB.WriteString(va.Event.ToStringBeautiful(num + 1))
	strB.WriteString(fmt.Sprintf("%s'alert':\n", ws))
	strB.WriteString(va.Alert.ToStringBeautiful(num + 1))

	return strB.String()
}
