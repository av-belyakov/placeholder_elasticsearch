package datamodels

// EventMessageForEsAlert сообщение с информацией о событии
// Base - основа
// StartDate - начальная дата
// RootId - главный уникальный идентификатор
// ObjectId - уникальный идентификатор объекта
// ObjectType - тип объекта
// Organisation - наименование организации
// OrganisationId - уникальный идентификатор организации
// Operation - операция
// RequestId - уникальный идентификатор запроса
// Details - детальная информация о событии
// Object - объект события
type EventMessageForEsAlert struct {
	Base           bool                          `json:"base" bson:"base"`
	StartDate      string                        `json:"startDate" bson:"startDate"` //в формате RFC3339
	RootId         string                        `json:"rootId" bson:"rootId"`
	ObjectId       string                        `json:"objectId" bson:"objectId"`
	ObjectType     string                        `json:"objectType" bson:"objectType"`
	Organisation   string                        `json:"organisation" bson:"organisation"`
	OrganisationId string                        `json:"organisationId" bson:"organisationId"`
	Operation      string                        `json:"operation" bson:"operation"`
	RequestId      string                        `json:"requestId" bson:"requestId"`
	Details        EventMessageForEsAlertDetails `json:"details" bson:"details"`
	Object         EventMessageForEsAlertObject  `json:"object" bson:"object"`
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
// Tlp - tlp
// UnderliningId - уникальный идентификатор
// Id - уникальный идентификатор
// CreatedBy - кем создан
// UpdatedBy - кем обновлен
// CreatedAt - дата создания (формат RFC3339)
// UpdatedAt - дата обновления (формат RFC3339)
// UnderliningType - тип
// Title - заголовок
// Description - описание
// Tags - список тегов
// Status - статус
// CustomFields - настраиваемые поля
// Date - дата (формат RFC3339)
// Type - тип
// Source - источник
// SourceRef - ссылка на источник
// Case - кейс
// CaseTemplate - шаблон обращения
// ObjectType - тип объекта
// Tags - теги после обработки
// TagsAll - все теги
// CustomFields - настраиваемые поля
type EventMessageForEsAlertObject struct {
	Tlp             uint64              `json:"tlp" bson:"tlp"`
	UnderliningId   string              `json:"_id" bson:"_id"`
	Id              string              `json:"id" bson:"id"`
	CreatedBy       string              `json:"createdBy" bson:"createdBy"`
	UpdatedBy       string              `json:"updatedBy,omitempty" bson:"updatedBy,omitempty"`
	CreatedAt       string              `json:"createdAt" bson:"createdAt"`                     //формат RFC3339
	UpdatedAt       string              `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"` //формат RFC3339
	UnderliningType string              `json:"_type" bson:"_type"`
	Title           string              `json:"title" bson:"title"`
	Description     string              `json:"description" bson:"description"`
	Status          string              `json:"status" bson:"status"`
	Date            string              `json:"date" bson:"date"` //формат RFC3339
	Type            string              `json:"type" bson:"type"`
	ObjectType      string              `json:"objectType" bson:"objectType"`
	Source          string              `json:"source" bson:"source"`
	SourceRef       string              `json:"sourceRef" bson:"sourceRef"`
	Case            string              `json:"case,omitempty" bson:"case,omitempty"`
	CaseTemplate    string              `json:"caseTemplate,omitempty" bson:"caseTemplate,omitempty"`
	Tags            map[string][]string `json:"tags" bson:"tags"`
	TagsAll         []string            `json:"tagsAll" bson:"tagsAll"`
	CustomFields    CustomFields        `json:"customFields" bson:"customFields"`
	//"artifacts" : [ ], думаю эти не надо, всегда пустые
	//"similarCases" : [ ] думаю эти не надо, всегда пустые
}
