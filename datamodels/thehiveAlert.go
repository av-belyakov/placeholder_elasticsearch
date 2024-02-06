package datamodels

import (
	"fmt"
	"strings"
	"time"

	"placeholder_elasticsearch/supportingfunctions"
)

// VerifiedTheHiveAlert объект представляет собой верифицированный Alert TheHive
type VerifiedTheHiveAlert struct {
	ID               string                   `json:"@id" bson:"@id"`
	ElasticsearchID  string                   `json:"@es_id" bson:"@es_id"`
	CreateTimestatmp string                   `json:"@timestamp" bson:"@timestamp"`
	Source           string                   `json:"source" bson:"source"`
	Event            EventMessageTheHiveAlert `json:"event" bson:"event"`
	//Alert `json:"alert" bson:"alert"`
}

func NewVerifiedTheHiveAlert() *VerifiedTheHiveCase {
	return &VerifiedTheHiveCase{
		CreateTimestatmp: supportingfunctions.GetDateTimeFormatRFC3339(time.Now().UnixMilli()),
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

func (va *VerifiedTheHiveAlert) ToStringBeautiful(num int) string {
	ws := supportingfunctions.GetWhitespace(num)

	strB := strings.Builder{}
	strB.WriteString(fmt.Sprintf("%sCreateTimestatmp: '%s'\n", ws, va.CreateTimestatmp))
	strB.WriteString(fmt.Sprintf("%sSource: '%s'\n", ws, va.Source))
	strB.WriteString(fmt.Sprintf("%sEvent:\n", ws))
	strB.WriteString(va.Event.ToStringBeautiful(num + 1))
	strB.WriteString(fmt.Sprintf("%sAlert:\n", ws))
	//strB.WriteString(va. .ToStringBeautiful(num + 1))

	return strB.String()
}
