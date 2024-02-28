package datamodels

import (
	commonevent "placeholder_elasticsearch/datamodels/commonevent"
	commonobjectevent "placeholder_elasticsearch/datamodels/commonobjectevent"
)

// EventMessageForEsAlert сообщение с информацией о событии
// Details - детальная информация о событии
// Object - объект события
type EventMessageForEsAlert struct {
	commonevent.CommonEventType
	Details EventMessageForEsAlertDetails `json:"details" bson:"details"`
	Object  EventMessageForEsAlertObject  `json:"object" bson:"object"`
}

// EventMessageForEsAlertDetails детальная информация о событии
// SourceRef - ссылка
// Title - заголовок
// Description - описание
// Tags - теги после обработки
// TagsAll - все теги
type EventMessageForEsAlertDetails struct {
	SourceRef   string              `json:"sourceRef" bson:"sourceRef"`
	Title       string              `json:"title" bson:"title"`
	Description string              `json:"description" bson:"description"`
	Tags        map[string][]string `json:"tags" bson:"tags"`
	TagsAll     []string            `json:"tagsAll" bson:"tagsAll"`
}

// EventMessageForEsAlertObject объект события
// Tags - теги после обработки
// TagsAll - все теги
// CustomFields - настраиваемые поля
type EventMessageForEsAlertObject struct {
	commonobjectevent.CommonEventAlertObject
	Tags         map[string][]string `json:"tags" bson:"tags"`
	TagsAll      []string            `json:"tagsAll" bson:"tagsAll"`
	CustomFields CustomFields        `json:"customFields" bson:"customFields"`
}
