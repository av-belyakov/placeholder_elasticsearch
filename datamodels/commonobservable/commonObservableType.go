package commonobservable

// CommonObservableType общие поля наблюдаемого сообщения
// Ioc - индикатор компрометации
// Sighted - видящий
// IgnoreSimilarity - игнорировать сходство
// Tlp - tlp
// UnderliningCreatedAt - время создания
// UnderliningUpdatedAt - время обновления
// StartDate - дата начала
// UnderliningCreatedBy - кем создан
// UnderliningUpdatedBy - кем обновлен
// UnderliningId - уникальный идентификатор
// UnderliningType - тип
// Data - данные
// DataType - тип данных
// Message - сообщение
type CommonObservableType struct {
	Ioc                  bool   `json:"ioc" bson:"ioc"`
	Sighted              bool   `json:"sighted" bson:"sighted"`
	IgnoreSimilarity     bool   `json:"ignoreSimilarity" bson:"ignoreSimilarity"`
	Tlp                  uint64 `json:"tlp" bson:"tlp"`
	UnderliningCreatedAt string `json:"_createdAt" bson:"_createdAt"` //формат RFC3339
	UnderliningUpdatedAt string `json:"_updatedAt" bson:"_updatedAt"` //формат RFC3339
	StartDate            string `json:"startDate" bson:"startDate"`   //формат RFC3339
	UnderliningCreatedBy string `json:"_createdBy,omitempty" bson:"_createdBy"`
	UnderliningUpdatedBy string `json:"_updatedBy,omitempty" bson:"_updatedBy"`
	UnderliningId        string `json:"_id,omitempty" bson:"_id"`
	UnderliningType      string `json:"_type,omitempty" bson:"_type"`
	Data                 string `json:"data,omitempty" bson:"data"`
	DataType             string `json:"dataType,omitempty" bson:"dataType"`
	Message              string `json:"message,omitempty" bson:"message"`
}
