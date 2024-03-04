package commonalert

// CommonAlertType общее описание Alert
// Tlp - номер группы разделяющие общие цели
// Date - дата (формат RFC3339)
// CreatedAt - дата создания (формат RFC3339)
// UpdatedAt - дата обновления (формат RFC3339)
// UpdatedBy - кем обновлен
// UnderliningId - уникальный идентификатор
// Id - уникальный идентификатор
// Status - статус
// Source - источник
// Type - тип
// Title - описание
// UnderliningType - тип
// Description - описание
// CaseTemplate - шаблон обращения
// SourceRef - ссылка на источник
type CommonAlertType struct {
	Tlp             uint64 `json:"tlp" bson:"tlp"`
	Date            string `json:"date" bson:"date"`           //формат RFC3339
	CreatedAt       string `json:"createdAt" bson:"createdAt"` //формат RFC3339
	UpdatedAt       string `json:"updatedAt" bson:"updatedAt"` //формат RFC3339
	UpdatedBy       string `json:"updatedBy" bson:"updatedBy"`
	UnderliningId   string `json:"_id" bson:"_id"`
	Id              string `json:"id" bson:"id"`
	Status          string `json:"status" bson:"status"`
	Source          string `json:"source" bson:"source"`
	Type            string `json:"type" bson:"type"`
	Title           string `json:"title" bson:"title"`
	UnderliningType string `json:"_type" bson:"_type"`
	Description     string `json:"description" bson:"description"`
	CaseTemplate    string `json:"caseTemplate" bson:"caseTemplate"`
	SourceRef       string `json:"sourceRef" bson:"sourceRef"`
}
