package datamodels

import (
	"fmt"
	"strings"
	"time"

	"placeholder_elasticsearch/supportingfunctions"
)

// VerifiedTheHiveCase объект представляет собой верифицированный 'кейс' TheHive
type VerifiedTheHiveCase struct {
	ID              string                  `json:"@id" bson:"@id"`
	ElasticsearchID string                  `json:"@es_id" bson:"@es_id"`
	CreateTimestamp string                  `json:"@timestamp" bson:"@timestamp"`
	Source          string                  `json:"source" bson:"source"`
	Event           EventMessageTheHiveCase `json:"event" bson:"event"`
	ObservablesMessageTheHive
	TtpsMessageTheHive
}

func NewVerifiedTheHiveCase() *VerifiedTheHiveCase {
	return &VerifiedTheHiveCase{
		CreateTimestamp: supportingfunctions.GetDateTimeFormatRFC3339(time.Now().UnixMilli()),
	}
}

func (hcase *VerifiedTheHiveCase) Get() *VerifiedTheHiveCase {
	return hcase
}

func (hcase *VerifiedTheHiveCase) GetID() string {
	return hcase.ID
}

func (hcase *VerifiedTheHiveCase) SetID(id string) {
	hcase.ID = id
}

func (hcase *VerifiedTheHiveCase) GetSource() string {
	return hcase.Source
}

func (hcase *VerifiedTheHiveCase) SetSource(source string) {
	hcase.Source = source
}

func (hcase *VerifiedTheHiveCase) GetCreateTimestamp() string {
	return hcase.CreateTimestamp
}

func (hcase *VerifiedTheHiveCase) SetCreateTimestamp(timestamp string) {
	hcase.CreateTimestamp = timestamp
}

func (hcase *VerifiedTheHiveCase) GetEvent() *EventMessageTheHiveCase {
	return &hcase.Event
}

func (hcase *VerifiedTheHiveCase) SetEvent(event EventMessageTheHiveCase) {
	hcase.Event = event
}

func (hcase *VerifiedTheHiveCase) GetObservables() *ObservablesMessageTheHive {
	return &hcase.ObservablesMessageTheHive
}

func (hcase *VerifiedTheHiveCase) SetObservables(observables ObservablesMessageTheHive) {
	hcase.ObservablesMessageTheHive = observables
}

func (hcase *VerifiedTheHiveCase) GetTtps() *TtpsMessageTheHive {
	return &hcase.TtpsMessageTheHive
}

func (hcase *VerifiedTheHiveCase) SetTtps(ttp TtpsMessageTheHive) {
	hcase.TtpsMessageTheHive = ttp
}

func (hcase *VerifiedTheHiveCase) ToStringBeautiful(num int) string {
	ws := supportingfunctions.GetWhitespace(num)

	strB := strings.Builder{}
	strB.WriteString(fmt.Sprintf("%s'createTimestatmp': '%s'\n", ws, hcase.CreateTimestamp))
	strB.WriteString(fmt.Sprintf("%s'source': '%s'\n", ws, hcase.Source))
	strB.WriteString(fmt.Sprintf("%s'event':\n", ws))
	strB.WriteString(hcase.Event.ToStringBeautiful(num + 1))
	strB.WriteString(hcase.ObservablesMessageTheHive.ToStringBeautiful(num))
	strB.WriteString(hcase.TtpsMessageTheHive.ToStringBeautiful(num))

	return strB.String()
}
