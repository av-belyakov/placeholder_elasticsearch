package commonalert

// CommonAlertType общее описание Alert
// Tlp - номер группы разделяющие общие цели
// Date - дата (формат RFC3339)
// CreatedAt - дата создания (формат RFC3339)
// UpdatedAt - дата обновления (формат RFC3339)
// CreatedBy - кем создан
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
	Tlp             uint64 `json:"tlp,omitempty" bson:"tlp"`
	Date            string `json:"date,omitempty" bson:"date"`           //формат RFC3339
	CreatedAt       string `json:"createdAt,omitempty" bson:"createdAt"` //формат RFC3339
	UpdatedAt       string `json:"updatedAt,omitempty" bson:"updatedAt"` //формат RFC3339
	CreatedBy       string `json:"createdBy,omitempty" bson:"createdBy"`
	UpdatedBy       string `json:"updatedBy,omitempty" bson:"updatedBy"`
	UnderliningId   string `json:"_id,omitempty" bson:"_id"`
	Id              string `json:"id,omitempty" bson:"id"`
	Status          string `json:"status,omitempty" bson:"status"`
	Source          string `json:"source,omitempty" bson:"source"`
	Type            string `json:"type,omitempty" bson:"type"`
	Title           string `json:"title,omitempty" bson:"title"`
	UnderliningType string `json:"_type,omitempty" bson:"_type"`
	Description     string `json:"description,omitempty" bson:"description"`
	CaseTemplate    string `json:"caseTemplate,omitempty" bson:"caseTemplate"`
	SourceRef       string `json:"sourceRef,omitempty" bson:"sourceRef"`
}
