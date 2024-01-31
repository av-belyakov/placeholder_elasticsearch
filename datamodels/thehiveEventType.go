package datamodels

// EventMessageTheHive сообщение с информацией о событии
// Operation - операция
// ObjectId - уникальный идентификатор объекта
// ObjectType - тип объекта
// Base - основа
// StartDate - начальная дата
// RootId - главный уникальный идентификатор
// RequestId - уникальный идентификатор запроса
// EventDetails - детальная информация о событии
// Object - объект события
// OrganisationId - уникальный идентификатор организации
// Organisation - наименование организации
type EventMessageTheHive struct {
	Base           bool         `json:"base" bson:"base"`
	StartDate      string       `json:"startDate" bson:"startDate"` //в формате RFC3339
	RootId         string       `json:"rootId" bson:"rootId"`
	Organisation   string       `json:"organisation" bson:"organisation"`
	OrganisationId string       `json:"organisationId" bson:"organisationId"`
	ObjectId       string       `json:"objectId" bson:"objectId"`
	ObjectType     string       `json:"objectType" bson:"objectType"`
	Operation      string       `json:"operation" bson:"operation"`
	RequestId      string       `json:"requestId" bson:"requestId"`
	Details        EventDetails `json:"details" bson:"details"`
	Object         EventObject  `json:"object" bson:"object"`
}

// EventDetails детальная информация о событии
// EndDate - конечное дата и время
// ResolutionStatus - статус постановления
// Summary - резюме
// Status - статус
// ImpactStatus - краткое описание воздействия
// CustomFields - настраиваемые поля
type EventDetails struct {
	EndDate          string                    `json:"endDate" bson:"endDate"` //формат RFC3339
	ResolutionStatus string                    `json:"resolutionStatus" bson:"resolutionStatus"`
	Summary          string                    `json:"summary" bson:"summary"`
	Status           string                    `json:"status" bson:"status"`
	ImpactStatus     string                    `json:"impactStatus" bson:"impactStatus"`
	CustomFields     map[string]CustomerFields `json:"customFields" bson:"customFields"`
}

// EventObject объект события
// UnderliningId - уникальный идентификатор
// Id - уникальный идентификатор
// CreatedBy - кем создан
// UpdatedBy - кем обновлен
// CreatedAt - дата создания
// UpdatedAt - дата обновления
// UnderliningType - тип
// CaseId - уникальный идентификатор дела
// Title - заголовок
// Description - описание
// Severity - строгость
// StartDate - начальная дата
// EndDate - конечная дата
// ImpactStatus - краткое описание воздействия
// ResolutionStatus - статус разрешения
// Tags - список тегов
// Flag - флаг
// Tlp - tlp
// Pap - pap
// Status - статус
// Summary - резюме
// Owner - владелец
// CustomFields - настраиваемые поля
// Stats - статистика
// Permissions - разрешения
type EventObject struct {
	Flag             bool                      `json:"flag" bson:"flag"`
	CaseId           uint64                    `json:"caseId" bson:"caseId"`
	Severity         uint64                    `json:"severity" bson:"severity"`
	Tlp              uint64                    `json:"tlp" bson:"tlp"`
	Pap              uint64                    `json:"pap" bson:"pap"`
	StartDate        string                    `json:"startDate" bson:"startDate"` //формат RFC3339
	EndDate          string                    `json:"endDate" bson:"endDate"`     //формат RFC3339
	CreatedAt        string                    `json:"createdAt" bson:"createdAt"` //формат RFC3339
	UpdatedAt        string                    `json:"updatedAt" bson:"updatedAt"` //формат RFC3339
	UnderliningId    string                    `json:"_id" bson:"_id"`
	Id               string                    `json:"id" bson:"id"`
	CreatedBy        string                    `json:"createdBy" bson:"createdBy"`
	UpdatedBy        string                    `json:"updatedBy" bson:"updatedBy"`
	UnderliningType  string                    `json:"_type" bson:"_type"`
	Title            string                    `json:"title" bson:"title"`
	Description      string                    `json:"description" bson:"description"`
	ImpactStatus     string                    `json:"impactStatus" bson:"impactStatus"`
	ResolutionStatus string                    `json:"resolutionStatus" bson:"resolutionStatus"`
	Status           string                    `json:"status" bson:"status"`
	Summary          string                    `json:"summary" bson:"summary"`
	Owner            string                    `json:"owner" bson:"owner"`
	Tags             []string                  `json:"tags" bson:"tags"`
	CustomFields     map[string]CustomerFields `json:"customFields" bson:"customFields"`
	//данное поле редко используемое, думаю пока оно не требует реализации
	//Stats            map[string]interface{} `json:"stats"`
	//данное поле редко используемое, думаю пока оно не требует реализации
	//Permissions  []string              `json:"permissions"`
}
