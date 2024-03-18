package datamodels

import (
	commonevent "placeholder_elasticsearch/datamodels/commonevent"
	commonobjectevent "placeholder_elasticsearch/datamodels/commonobjectevent"
)

// EventMessageTheHiveCase сообщение с информацией о событии
// Details - детальная информация о событии
// Object - объект события
type EventMessageTheHiveCase struct {
	commonevent.CommonEventType
	Details EventCaseDetails `json:"details" bson:"details"`
	Object  EventCaseObject  `json:"object" bson:"object"`
}

// EventDetails детальная информация о событии
// EndDate - конечное дата и время
// ResolutionStatus - статус постановления
// Summary - резюме
// Status - статус
// ImpactStatus - краткое описание воздействия
// CustomFields - настраиваемые поля
type EventCaseDetails struct {
	EndDate          string       `json:"endDate,omitempty" bson:"endDate"` //формат RFC3339
	ResolutionStatus string       `json:"resolutionStatus,omitempty" bson:"resolutionStatus"`
	Summary          string       `json:"summary,omitempty" bson:"summary"`
	Status           string       `json:"status,omitempty" bson:"status"`
	ImpactStatus     string       `json:"impactStatus,omitempty" bson:"impactStatus"`
	CustomFields     CustomFields `json:"customFields,omitempty" bson:"customFields"`
}

// EventObject объект события
// Tags - список тегов
// CustomFields - настраиваемые поля
type EventCaseObject struct {
	commonobjectevent.CommonEventCaseObject
	Tags         []string     `json:"tags" bson:"tags"`
	CustomFields CustomFields `json:"customFields,omitempty" bson:"customFields"`
}
