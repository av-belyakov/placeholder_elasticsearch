package commonobjectevent

// CommonEventCaseObject общие поля для описания объекта события
// Flag - флаг
// CaseId - уникальный идентификатор дела
// Severity - строгость
// Tlp - tlp
// Pap - pap
// StartDate - начальная дата
// EndDate - конечная дата
// CreatedAt - дата создания
// UpdatedAt - дата обновления
// UnderliningId - уникальный идентификатор
// Id - уникальный идентификатор
// CreatedBy - кем создан
// UpdatedBy - кем обновлен
// UnderliningType - тип
// Title - заголовок
// Description - описание
// ImpactStatus - краткое описание воздействия
// ResolutionStatus - статус разрешения
// Status - статус
// Summary - резюме
// Owner - владелец
type CommonEventCaseObject struct {
	Flag             bool   `json:"flag" bson:"flag"`
	CaseId           uint64 `json:"caseId" bson:"caseId"`
	Severity         uint64 `json:"severity" bson:"severity"`
	Tlp              uint64 `json:"tlp" bson:"tlp"`
	Pap              uint64 `json:"pap" bson:"pap"`
	StartDate        string `json:"startDate" bson:"startDate"` //формат RFC3339
	EndDate          string `json:"endDate" bson:"endDate"`     //формат RFC3339
	CreatedAt        string `json:"createdAt" bson:"createdAt"` //формат RFC3339
	UpdatedAt        string `json:"updatedAt" bson:"updatedAt"` //формат RFC3339
	UnderliningId    string `json:"_id" bson:"_id"`
	Id               string `json:"id" bson:"id"`
	CreatedBy        string `json:"createdBy" bson:"createdBy"`
	UpdatedBy        string `json:"updatedBy" bson:"updatedBy"`
	UnderliningType  string `json:"_type" bson:"_type"`
	Title            string `json:"title" bson:"title"`
	Description      string `json:"description" bson:"description"`
	ImpactStatus     string `json:"impactStatus" bson:"impactStatus"`
	ResolutionStatus string `json:"resolutionStatus" bson:"resolutionStatus"`
	Status           string `json:"status" bson:"status"`
	Summary          string `json:"summary" bson:"summary"`
	Owner            string `json:"owner" bson:"owner"`
}
