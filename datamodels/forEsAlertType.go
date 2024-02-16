package datamodels

// AlertMessageForEsAlert сообщение с информацией о Alert-те
// Tlp - номер группы разделяющие общие цели
// Date - дата (формат RFC3339)
// CreatedAt - дата создания (формат RFC3339)
// UpdatedAt - дата обновления (формат RFC3339)
// UpdatedBy - кем обновлен
// UnderliningId - уникальный идентификатор
// Status - статус
// Type - тип
// UnderliningType - тип
// Description - описание
// CaseTemplate - шаблон обращения
// SourceRef - ссылка на источник
// Tags - теги
// CustomFields - настраиваемые поля
// Artifacts - артефакты
type AlertMessageForEsAlert struct {
	Tlp             uint64                        `json:"tlp" bson:"tlp"`
	Date            string                        `json:"date" bson:"date"`           //формат RFC3339
	CreatedAt       string                        `json:"createdAt" bson:"createdAt"` //формат RFC3339
	UpdatedAt       string                        `json:"updatedAt" bson:"updatedAt"` //формат RFC3339
	UpdatedBy       string                        `json:"updatedBy" bson:"updatedBy"`
	UnderliningId   string                        `json:"_id" bson:"_id"`
	Status          string                        `json:"status" bson:"status"`
	Type            string                        `json:"type" bson:"type"`
	UnderliningType string                        `json:"_type" bson:"_type"`
	Description     string                        `json:"description" bson:"description"`
	CaseTemplate    string                        `json:"caseTemplate" bson:"caseTemplate"`
	SourceRef       string                        `json:"sourceRef" bson:"sourceRef"`
	Tags            map[string][]string           `json:"tags" bson:"tags"`
	CustomFields    CustomFields                  `json:"customFields" bson:"customFields"`
	Artifacts       map[string]ArtifactForEsAlert `json:"artifact" bson:"artifact"`
}

// ArtifactForEsAlert содержит артефакт к алерту
// Ioc - индикатор компрометации
// Tlp - tlp
// UnderliningId - уникальный идентификатор
// Id - уникальный идентификатор
// UnderliningType - тип
// CreatedAt - время создания
// CreatedBy - кем создан
// StartDate - дата начала
// Data - данные
// DataType - тип данных
// Message - сообщение
// Tags - список тегов
type ArtifactForEsAlert struct {
	Ioc             bool                `json:"ioc" bson:"ioc"`
	Tlp             uint64              `json:"tlp" bson:"tlp"`
	UnderliningId   string              `json:"_id" bson:"_id"`
	Id              string              `json:"id" bson:"id"`
	UnderliningType string              `json:"_type" bson:"_type"`
	CreatedAt       string              `json:"createdAt" bson:"createdAt"` //формат RFC3339
	StartDate       string              `json:"startDate" bson:"startDate"` //формат RFC3339
	CreatedBy       string              `json:"createdBy" bson:"createdBy"`
	Data            string              `json:"data" bson:"data"`
	DataType        string              `json:"dataType" bson:"dataType"`
	Message         string              `json:"message,omitempty" bson:"message"`
	Tags            map[string][]string `json:"tags" bson:"tags"`
}
