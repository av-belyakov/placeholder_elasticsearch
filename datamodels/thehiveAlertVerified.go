package datamodels

import (
	"fmt"
	"strings"
	"time"

	"placeholder_elasticsearch/supportingfunctions"
)

// VerifiedTheHiveAlert объект представляет собой верифицированный Alert TheHive
type VerifiedTheHiveAlert struct {
	ID              string                   `json:"@id" bson:"@id"`
	ElasticsearchID string                   `json:"@es_id" bson:"@es_id"`
	CreateTimestamp string                   `json:"@timestamp" bson:"@timestamp"`
	Source          string                   `json:"source" bson:"source"`
	Event           EventMessageTheHiveAlert `json:"event" bson:"event"`
	Alert           AlertMessageTheHiveAlert `json:"alert" bson:"alert"`
}

func NewVerifiedTheHiveAlert() *VerifiedTheHiveAlert {
	return &VerifiedTheHiveAlert{
		CreateTimestamp: supportingfunctions.GetDateTimeFormatRFC3339(time.Now().UnixMilli()),
		Event:           *NewEventMessageTheHiveAlert(),
		Alert:           *NewAlertMessageTheHiveAlert(),
	}
}

func (va *VerifiedTheHiveAlert) Get() *VerifiedTheHiveAlert {
	return va
}

func (va *VerifiedTheHiveAlert) GetID() string {
	return va.ID
}

func (va *VerifiedTheHiveAlert) SetID(id string) {
	va.ID = id
}

func (va *VerifiedTheHiveAlert) GetElasticsearchID() string {
	return va.ElasticsearchID
}

func (va *VerifiedTheHiveAlert) SetElasticsearchID(id string) {
	va.ElasticsearchID = id
}

func (va *VerifiedTheHiveAlert) GetCreateTimestatmp() string {
	return va.CreateTimestamp
}

func (va *VerifiedTheHiveAlert) SetCreateTimestatmp(time string) {
	va.CreateTimestamp = time
}

func (va *VerifiedTheHiveAlert) GetSource() string {
	return va.Source
}

func (va *VerifiedTheHiveAlert) SetSource(source string) {
	va.Source = source
}

func (va *VerifiedTheHiveAlert) GetEvent() *EventMessageTheHiveAlert {
	return &va.Event
}

func (va *VerifiedTheHiveAlert) SetEvent(event EventMessageTheHiveAlert) {
	va.Event = event
}

func (va *VerifiedTheHiveAlert) GetAlert() *AlertMessageTheHiveAlert {
	return &va.Alert
}

func (va *VerifiedTheHiveAlert) SetAlert(alert AlertMessageTheHiveAlert) {
	va.Alert = alert
}

func (va *VerifiedTheHiveAlert) ToStringBeautiful(num int) string {
	ws := supportingfunctions.GetWhitespace(num)

	str := strings.Builder{}

	str.WriteString(fmt.Sprintf("%s'createTimestatmp': '%s'\n", ws, va.CreateTimestamp))
	str.WriteString(fmt.Sprintf("%s'source': '%s'\n", ws, va.Source))
	str.WriteString(fmt.Sprintf("%s'event':\n", ws))
	str.WriteString(va.Event.ToStringBeautiful(num + 1))
	str.WriteString(fmt.Sprintf("%s'alert':\n", ws))
	str.WriteString(va.Alert.ToStringBeautiful(num + 1))

	return str.String()
}
