package datamodels

// AlertMessageTheHiveAlert сообщение с информацией о Alert-те
// Follow - следовать
// Tlp - номер группы разделяющие общие цели
// Severity - строгость
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
// Artifacts - артифакты
type AlertMessageTheHiveAlert struct {
	Follow          bool     `json:"follow" bson:"follow"`
	Tlp             uint64   `json:"tlp" bson:"tlp"`
	Severity        uint64   `json:"severity" bson:"severity"`
	Date            string   `json:"date" bson:"date"`           //формат RFC3339
	CreatedAt       string   `json:"createdAt" bson:"createdAt"` //формат RFC3339
	UpdatedAt       string   `json:"updatedAt" bson:"updatedAt"` //формат RFC3339
	UpdatedBy       string   `json:"updatedBy" bson:"updatedBy"`
	UnderliningId   string   `json:"_id" bson:"_id"`
	Status          string   `json:"status" bson:"status"`
	Type            string   `json:"type" bson:"type"`
	UnderliningType string   `json:"_type" bson:"_type"`
	Description     string   `json:"description" bson:"description"`
	CaseTemplate    string   `json:"caseTemplate" bson:"caseTemplate"`
	SourceRef       string   `json:"sourceRef" bson:"sourceRef"`
	Tags            []string `json:"tags" bson:"tags"`
	CustomFields
	//CustomFields    map[string]CustomerFields `json:"customFields" bson:"customFields"`
	Artifacts []AlertArtifact `json:"artifact" bson:"artifact"`
}

// AlertArtifact содержит артифакт к алерту
// Ioc - индикатор компрометации
// Sighted - видящий
// IgnoreSimilarity - игнорировать похожие
// Tlp - tlp
// UnderliningId - уникальный идентификатор
// Id - уникальный идентификатор
// UnderliningType - тип
// CreatedAt - время создания
// CreatedBy - кем создан
// StartDate - дата начала
// UpdatedAt - время обновления
// UpdatedBy - кем обновлен
// Data - данные
// DataType - тип данных
// Message - сообщение
// Tags - список тегов
type AlertArtifact struct {
	Ioc              bool     `json:"ioc" bson:"ioc"`
	Sighted          bool     `json:"sighted" bson:"sighted"`
	IgnoreSimilarity bool     `json:"ignoreSimilarity" bson:"ignoreSimilarity"`
	Tlp              uint64   `json:"tlp" bson:"tlp"`
	UnderliningId    string   `json:"_id" bson:"_id"`
	Id               string   `json:"id" bson:"id"`
	UnderliningType  string   `json:"_type" bson:"_type"`
	CreatedAt        string   `json:"createdAt" bson:"createdAt"` //формат RFC3339
	UpdatedAt        string   `json:"updatedAt" bson:"updatedAt"` //формат RFC3339
	StartDate        string   `json:"startDate" bson:"startDate"` //формат RFC3339
	CreatedBy        string   `json:"createdBy" bson:"createdBy"`
	UpdatedBy        string   `json:"updatedBy" bson:"updatedBy"`
	Data             string   `json:"data" bson:"data"`
	DataType         string   `json:"dataType" bson:"dataType"`
	Message          string   `json:"message" bson:"message"`
	Tags             []string `json:"tags" bson:"tags"`
}
