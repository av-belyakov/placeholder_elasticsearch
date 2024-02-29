package commonalertartifact

// CommonArtifactType содержит общие поля для описания объекта Artifact
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
type CommonArtifactType struct {
	Ioc             bool   `json:"ioc" bson:"ioc"`
	Tlp             uint64 `json:"tlp" bson:"tlp"`
	UnderliningId   string `json:"_id" bson:"_id"`
	Id              string `json:"id" bson:"id"`
	UnderliningType string `json:"_type" bson:"_type"`
	CreatedAt       string `json:"createdAt" bson:"createdAt"` //формат RFC3339
	StartDate       string `json:"startDate" bson:"startDate"` //формат RFC3339
	CreatedBy       string `json:"createdBy" bson:"createdBy"`
	Data            string `json:"data" bson:"data"`
	DataType        string `json:"dataType" bson:"dataType"`
	Message         string `json:"message,omitempty" bson:"message"`
}
