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
	Ioc             bool   `json:"ioc,omitempty" bson:"ioc"`
	Tlp             uint64 `json:"tlp,omitempty" bson:"tlp"`
	UnderliningId   string `json:"_id,omitempty" bson:"_id"`
	Id              string `json:"id,omitempty" bson:"id"`
	UnderliningType string `json:"_type,omitempty" bson:"_type"`
	CreatedAt       string `json:"createdAt,omitempty" bson:"createdAt"` //формат RFC3339
	StartDate       string `json:"startDate,omitempty" bson:"startDate"` //формат RFC3339
	CreatedBy       string `json:"createdBy,omitempty" bson:"createdBy"`
	Data            string `json:"data,omitempty" bson:"data"`
	DataType        string `json:"dataType,omitempty" bson:"dataType"`
	Message         string `json:"message,omitempty" bson:"message"`
}
