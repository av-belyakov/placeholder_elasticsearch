package datamodels

// EventMessageTheHiveAlert сообщение с информацией о событии
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
type EventMessageTheHiveAlert struct {
	Base           bool              `json:"base" bson:"base"`
	StartDate      string            `json:"startDate" bson:"startDate"` //в формате RFC3339
	RootId         string            `json:"rootId" bson:"rootId"`
	ObjectId       string            `json:"objectId" bson:"objectId"`
	ObjectType     string            `json:"objectType" bson:"objectType"`
	Organisation   string            `json:"organisation" bson:"organisation"`
	OrganisationId string            `json:"organisationId" bson:"organisationId"`
	Operation      string            `json:"operation" bson:"operation"`
	RequestId      string            `json:"requestId" bson:"requestId"`
	Details        EventAlertDetails `json:"details" bson:"details"`
	Object         EventAlertObject  `json:"object" bson:"object"`
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
// Tlp - tlp
// Pap - pap
// Severity - строгость
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
// Tags - теги
// CustomFields - настраиваемые поля
type EventAlertObject struct {
	Follow          bool                      `json:"follow" bson:"follow"`
	Severity        uint64                    `json:"severity" bson:"severity"`
	Tlp             uint64                    `json:"tlp" bson:"tlp"`
	Pap             uint64                    `json:"pap" bson:"pap"`
	UnderliningId   string                    `json:"_id" bson:"_id"`
	Id              string                    `json:"id" bson:"id"`
	CreatedBy       string                    `json:"createdBy" bson:"createdBy"`
	UpdatedBy       string                    `json:"updatedBy" bson:"updatedBy"`
	CreatedAt       string                    `json:"createdAt" bson:"createdAt"` //формат RFC3339
	UpdatedAt       string                    `json:"updatedAt" bson:"updatedAt"` //формат RFC3339
	UnderliningType string                    `json:"_type" bson:"_type"`
	Title           string                    `json:"title" bson:"title"`
	Description     string                    `json:"description" bson:"description"`
	Status          string                    `json:"status" bson:"status"`
	Date            string                    `json:"date" bson:"date"` //формат RFC3339
	Type            string                    `json:"type" bson:"type"`
	ObjectType      string                    `json:"objectType" bson:"objectType"`
	Source          string                    `json:"source" bson:"source"`
	SourceRef       string                    `json:"sourceRef" bson:"sourceRef"`
	Case            string                    `json:"case" bson:"case"`
	CaseTemplate    string                    `json:"caseTemplate" bson:"caseTemplate"`
	Tags            []string                  `json:"tags" bson:"tags"`
	CustomFields    map[string]CustomerFields `json:"customFields" bson:"customFields"`
	//"artifacts" : [ ], думаю эти не надо, всегда пустые
	//"similarCases" : [ ] думаю эти не надо, всегда пустые
}
