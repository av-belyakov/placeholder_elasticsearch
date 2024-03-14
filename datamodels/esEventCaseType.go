package datamodels

import (
	commonevent "placeholder_elasticsearch/datamodels/commonevent"
	commonobjectevent "placeholder_elasticsearch/datamodels/commonobjectevent"
)

// EventMessageForEsCase сообщение с информацией о событии
// Details - детальная информация о событии
// Object - объект события
type EventMessageForEsCase struct {
	commonevent.CommonEventType
	Details EventCaseDetails     `json:"details,omitempty" bson:"details"`
	Object  EventForEsCaseObject `json:"object,omitempty" bson:"object"`
}

// EventForEsCaseObject объект события
// Tags - список тегов
// Tags - список всех тегов
// CustomFields - настраиваемые поля
type EventForEsCaseObject struct {
	commonobjectevent.CommonEventCaseObject
	Tags         map[string][]string `json:"tags" bson:"tags"`
	TagsAll      []string            `json:"tagsAll" bson:"tagsAll"`
	CustomFields CustomFields        `json:"customFields,omitempty" bson:"customFields"`
}
