package datamodel

// CommonEventType общие поля для описания события
// Base - основа
// StartDate - начальная дата
// RootId - главный уникальный идентификатор
// Organisation - наименование организации
// OrganisationId - уникальный идентификатор организации
// ObjectId - уникальный идентификатор объекта
// ObjectType - тип объекта
// Operation - операция
// RequestId - уникальный идентификатор запроса
type CommonEventType struct {
	Base           bool   `json:"base,omitempty" bson:"base"`
	StartDate      string `json:"startDate,omitempty" bson:"startDate"` //в формате RFC3339
	RootId         string `json:"rootId,omitempty" bson:"rootId"`
	Organisation   string `json:"organisation,omitempty" bson:"organisation"`
	OrganisationId string `json:"organisationId,omitempty" bson:"organisationId"`
	ObjectId       string `json:"objectId,omitempty" bson:"objectId"`
	ObjectType     string `json:"objectType,omitempty" bson:"objectType"`
	Operation      string `json:"operation,omitempty" bson:"operation"`
	RequestId      string `json:"requestId,omitempty" bson:"requestId"`
}
