package datamodels

import "time"

// VerifiedTheHiveCase объект представляет собой верифицированный 'кейс' TheHive
type VerifiedTheHiveCase struct {
	CreateTimestatmp string                    `json:"@timestamp"`
	Source           string                    `json:"source"`
	Event            EventMessageTheHive       `json:"event"`
	Observables      ObservablesMessageTheHive `json:"observables"`
}

func NewVerifiedTheHiveCase() *VerifiedTheHiveCase {
	return &VerifiedTheHiveCase{
		CreateTimestatmp: time.UnixMilli(time.Now().UnixMilli()).Format(time.RFC3339),
	}
}

func (hcase *VerifiedTheHiveCase) Get() *VerifiedTheHiveCase {
	return hcase
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
	return &hcase.Observables
}

func (hcase *VerifiedTheHiveCase) SetObservables(observables ObservablesMessageTheHive) {
	hcase.Observables = observables
}
