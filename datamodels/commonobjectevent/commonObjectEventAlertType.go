package commonobjectevent

// CommonEventAlertObject объект события для Alert
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
type CommonEventAlertObject struct {
	Tlp             uint64 `json:"tlp" bson:"tlp"`
	UnderliningId   string `json:"_id" bson:"_id"`
	Id              string `json:"id" bson:"id"`
	CreatedBy       string `json:"createdBy" bson:"createdBy"`
	UpdatedBy       string `json:"updatedBy,omitempty" bson:"updatedBy,omitempty"`
	CreatedAt       string `json:"createdAt" bson:"createdAt"`                     //формат RFC3339
	UpdatedAt       string `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"` //формат RFC3339
	UnderliningType string `json:"_type" bson:"_type"`
	Title           string `json:"title" bson:"title"`
	Description     string `json:"description" bson:"description"`
	Status          string `json:"status" bson:"status"`
	Date            string `json:"date" bson:"date"` //формат RFC3339
	Type            string `json:"type" bson:"type"`
	ObjectType      string `json:"objectType" bson:"objectType"`
	Source          string `json:"source" bson:"source"`
	SourceRef       string `json:"sourceRef" bson:"sourceRef"`
	Case            string `json:"case,omitempty" bson:"case,omitempty"`
	CaseTemplate    string `json:"caseTemplate,omitempty" bson:"caseTemplate,omitempty"`
}
