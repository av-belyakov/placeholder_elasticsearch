package datamodels

// ObservablesMessageTheHive список наблюдаемых сообщений
// Observables - наблюдаемые сообщения
type ObservablesMessageTheHive struct {
	Observables []ObservableMessage `json:"observables"`
}

// ObservableMessage наблюдаемое сообщение
// CreatedAt - время создания
// CreatedBy - кем создан
// UnderliningId - уникальный идентификатор
// UnderliningType - тип
// UpdatedAt - время обновления
// UpdatedBy - кем обновлен
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
	Ioc                  bool                        `json:"ioc"`
	Sighted              bool                        `json:"sighted"`
	IgnoreSimilarity     bool                        `json:"ignoreSimilarity"`
	Tlp                  uint64                      `json:"tlp"`
	UnderliningCreatedAt string                      `json:"_createdAt"` //формат RFC3339
	UnderliningUpdatedAt string                      `json:"_updatedAt"` //формат RFC3339
	StartDate            string                      `json:"startDate"`  //формат RFC3339
	UnderliningCreatedBy string                      `json:"_createdBy"`
	UnderliningUpdatedBy string                      `json:"_updatedBy"`
	UnderliningId        string                      `json:"_id"`
	UnderliningType      string                      `json:"_type"`
	Data                 string                      `json:"data"`
	DataType             string                      `json:"dataType"`
	Message              string                      `json:"message"`
	Tags                 []string                    `json:"tags"`
	Attachment           AttachmentData              `json:"attachment"`
	Reports              map[string]ReportTaxonomies `json:"reports"`
	//данное поле редко используемое, думаю пока оно не требует реализации
	//ExtraData        map[string]interface{}                         `json:"extraData"`
}

// AttachmentData прикрепленные данные
type AttachmentData struct {
	Size        uint64   `json:"size"`
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	ContentType string   `json:"contentType"`
	Hashes      []string `json:"hashes"`
}

// ReportTaxonomies
type ReportTaxonomies struct {
	Taxonomies []Taxonomy `json:"taxonomies"`
}

// Taxonomy
type Taxonomy struct {
	Level     string `json:"level"`
	Namespace string `json:"namespace"`
	Predicate string `json:"predicate"`
	Value     string `json:"value"`
}
