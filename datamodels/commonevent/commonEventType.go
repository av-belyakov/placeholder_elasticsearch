package commonevent

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
	Base           bool   `json:"base" bson:"base"`
	StartDate      string `json:"startDate" bson:"startDate"` //в формате RFC3339
	RootId         string `json:"rootId" bson:"rootId"`
	Organisation   string `json:"organisation" bson:"organisation"`
	OrganisationId string `json:"organisationId" bson:"organisationId"`
	ObjectId       string `json:"objectId" bson:"objectId"`
	ObjectType     string `json:"objectType" bson:"objectType"`
	Operation      string `json:"operation" bson:"operation"`
	RequestId      string `json:"requestId" bson:"requestId"`
}
