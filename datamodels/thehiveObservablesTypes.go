package datamodels

// ObservablesMessageTheHive список наблюдаемых сообщений
// Observables - наблюдаемые сообщения
type ObservablesMessageTheHive struct {
	Observables []ObservableMessage `json:"observables" bson:"observables"`
}

// ObservableMessage наблюдаемое сообщение
// UnderliningCreatedAt - время создания
// UnderliningCreatedBy - кем создан
// UnderliningId - уникальный идентификатор
// UnderliningType - тип
// UnderliningUpdatedAt - время обновления
// UnderliningUpdatedBy - кем обновлен
// Data - данные
// DataType - тип данных
// IgnoreSimilarity - игнорировать сходство
// ExtraData - дополнительные данные
// Ioc - индикатор компрометации
// Message - сообщение
// Sighted - видящий
// StartDate - дата начала
// Tags - список тегов
// Tlp - tlp
// Reports - список отчетов
type ObservableMessage struct {
	Ioc                  bool                        `json:"ioc" bson:"ioc"`
	Sighted              bool                        `json:"sighted" bson:"sighted"`
	IgnoreSimilarity     bool                        `json:"ignoreSimilarity" bson:"ignoreSimilarity"`
	Tlp                  uint64                      `json:"tlp" bson:"tlp"`
	UnderliningCreatedAt string                      `json:"_createdAt" bson:"_createdAt"` //формат RFC3339
	UnderliningUpdatedAt string                      `json:"_updatedAt" bson:"_updatedAt"` //формат RFC3339
	StartDate            string                      `json:"startDate" bson:"startDate"`   //формат RFC3339
	UnderliningCreatedBy string                      `json:"_createdBy" bson:"_createdBy"`
	UnderliningUpdatedBy string                      `json:"_updatedBy" bson:"_updatedBy"`
	UnderliningId        string                      `json:"_id" bson:"_id"`
	UnderliningType      string                      `json:"_type" bson:"_type"`
	Data                 string                      `json:"data" bson:"data"`
	DataType             string                      `json:"dataType" bson:"dataType"`
	Message              string                      `json:"message" bson:"message"`
	Tags                 []string                    `json:"tags" bson:"tags"`
	Attachment           AttachmentData              `json:"attachment" bson:"attachment"`
	Reports              map[string]ReportTaxonomies `json:"reports" bson:"reports"`
	//данное поле редко используемое, думаю пока оно не требует реализации
	//ExtraData        map[string]interface{}                         `json:"extraData"`
}

// AttachmentData прикрепленные данные
type AttachmentData struct {
	Size        uint64   `json:"size" bson:"size"`
	Id          string   `json:"id" bson:"id"`
	Name        string   `json:"name" bson:"name"`
	ContentType string   `json:"contentType" bson:"contentType"`
	Hashes      []string `json:"hashes" bson:"hashes"`
}

// ReportTaxonomies
type ReportTaxonomies struct {
	Taxonomies []Taxonomy `json:"taxonomies" bson:"taxonomies"`
}

// Taxonomy
type Taxonomy struct {
	Level     string `json:"level" bson:"level"`
	Namespace string `json:"namespace" bson:"namespace"`
	Predicate string `json:"predicate" bson:"predicate"`
	Value     string `json:"value" bson:"value"`
}
