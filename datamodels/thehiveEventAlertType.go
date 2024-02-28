package datamodels

import (
	commonevent "placeholder_elasticsearch/datamodels/commonevent"
	commonobjectevent "placeholder_elasticsearch/datamodels/commonobjectevent"
)

// EventMessageTheHiveAlert сообщение с информацией о событии
// Details - детальная информация о событии
// Object - объект события
type EventMessageTheHiveAlert struct {
	commonevent.CommonEventType
	Details EventAlertDetails `json:"details" bson:"details"`
	Object  EventAlertObject  `json:"object" bson:"object"`
}

// EventAlertDetails детальная информация о событии
// SourceRef - ссылка
// Title - заголовок
// Description - описание
// Tags - список тегов
type EventAlertDetails struct {
	SourceRef   string   `json:"sourceRef" bson:"sourceRef"`
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	Tags        []string `json:"tags" bson:"tags"`
}

// EventAlertObject объект события
// Follow - следовать
// Severity - строгость
// Pap - pap
// Tags - теги
// CustomFields - настраиваемые поля
type EventAlertObject struct {
	commonobjectevent.CommonEventAlertObject
	Follow       bool         `json:"follow" bson:"follow"`
	Severity     uint64       `json:"severity" bson:"severity"`
	Pap          uint64       `json:"pap" bson:"pap"`
	Tags         []string     `json:"tags" bson:"tags"`
	CustomFields CustomFields `json:"customFields" bson:"customFields"`
}
