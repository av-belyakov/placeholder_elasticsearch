package datamodels

import (
	"fmt"
	"strings"
	"time"

	"placeholder_elasticsearch/supportingfunctions"
)

// VerifiedEsCase объект представляет собой верифицированный 'кейс' Elasticsearch
type VerifiedEsCase struct {
	ID              string                `json:"@id" bson:"@id"`
	ElasticsearchID string                `json:"@es_id" bson:"@es_id"`
	CreateTimestamp string                `json:"@timestamp" bson:"@timestamp"`
	Source          string                `json:"source" bson:"source"`
	Event           EventMessageForEsCase `json:"event" bson:"event"`
	ObservablesMessageEs
	TtpsMessageEs
}

func NewVerifiedEsCase() *VerifiedEsCase {
	return &VerifiedEsCase{
		CreateTimestamp: supportingfunctions.GetDateTimeFormatRFC3339(time.Now().UnixMilli()),
		Event:           *NewEventMessageForEsCase(),
	}
}

func (c *VerifiedEsCase) Get() *VerifiedEsCase {
	return c
}

func (c *VerifiedEsCase) GetID() string {
	return c.ID
}

func (c *VerifiedEsCase) SetID(id string) {
	c.ID = id
}

func (c *VerifiedEsCase) GetSource() string {
	return c.Source
}

func (c *VerifiedEsCase) SetSource(source string) {
	c.Source = source
}

func (c *VerifiedEsCase) GetCreateTimestamp() string {
	return c.CreateTimestamp
}

func (c *VerifiedEsCase) SetCreateTimestamp(timestamp string) {
	c.CreateTimestamp = timestamp
}

func (c *VerifiedEsCase) GetEvent() *EventMessageForEsCase {
	return &c.Event
}

func (c *VerifiedEsCase) SetEvent(event EventMessageForEsCase) {
	c.Event = event
}

func (c *VerifiedEsCase) GetObservables() *ObservablesMessageEs {
	return &c.ObservablesMessageEs
}

func (c *VerifiedEsCase) SetObservables(observables ObservablesMessageEs) {
	c.ObservablesMessageEs = observables
}

func (c *VerifiedEsCase) GetTtps() *TtpsMessageEs {
	return &c.TtpsMessageEs
}

func (c *VerifiedEsCase) SetTtps(ttp TtpsMessageEs) {
	c.TtpsMessageEs = ttp
}

func (c *VerifiedEsCase) ToStringBeautiful(num int) string {
	ws := supportingfunctions.GetWhitespace(num)

	str := strings.Builder{}
	str.WriteString(fmt.Sprintf("%s'createTimestatmp': '%s'\n", ws, c.CreateTimestamp))
	str.WriteString(fmt.Sprintf("%s'source': '%s'\n", ws, c.Source))
	str.WriteString(fmt.Sprintf("%s'event':\n", ws))
	str.WriteString(c.Event.ToStringBeautiful(num + 1))
	str.WriteString(fmt.Sprintf("%s'observables':\n", ws))
	str.WriteString(c.ObservablesMessageEs.ToStringBeautiful(num + 1))
	str.WriteString(fmt.Sprintf("%s'ttp':\n", ws))
	str.WriteString(c.TtpsMessageEs.ToStringBeautiful(num + 1))

	return str.String()
}
