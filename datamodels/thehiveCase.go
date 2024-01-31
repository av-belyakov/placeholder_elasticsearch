package datamodels

import (
	"fmt"
	"time"

	"placeholder_elasticsearch/supportingfunctions"
)

// VerifiedTheHiveCase объект представляет собой верифицированный 'кейс' TheHive
type VerifiedTheHiveCase struct {
	ID               string              `json:"@id" bson:"@id"`
	CreateTimestatmp string              `json:"@timestamp" bson:"@timestamp"`
	Source           string              `json:"source" bson:"source"`
	Event            EventMessageTheHive `json:"event" bson:"event"`
	ObservablesMessageTheHive
	TtpsMessageTheHive
}

func NewVerifiedTheHiveCase() *VerifiedTheHiveCase {
	return &VerifiedTheHiveCase{
		CreateTimestatmp: supportingfunctions.GetDateTimeFormat(time.Now().UnixMilli()),
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

func (hcase *VerifiedTheHiveCase) GetEvent() *EventMessageTheHive {
	return &hcase.Event
}

func (hcase *VerifiedTheHiveCase) SetEvent(event EventMessageTheHive) {
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

	str := fmt.Sprintf("%sCreateTimestatmp: '%s'\n", ws, hcase.CreateTimestatmp)
	str += fmt.Sprintf("%sSource: '%s'\n", ws, hcase.Source)
	str += fmt.Sprintf("%sEvent:\n", ws)
	str += hcase.Event.ToStringBeautiful(num + 1)
	str += fmt.Sprintf("%sObservables:\n", ws)
	str += hcase.ObservablesMessageTheHive.ToStringBeautiful(num + 1)

	return str
}
