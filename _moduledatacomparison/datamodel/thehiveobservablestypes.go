package datamodel

// ObservablesMessageTheHive список наблюдаемых сообщений
// Observables - наблюдаемые сообщения
type ObservablesMessageTheHive struct {
	Observables []ObservableMessage `json:"observables" bson:"observables"`
}

// ObservableMessage наблюдаемое сообщение
// Tags - список тегов
// Attachment - приложенные данные
// Reports - список отчетов
type ObservableMessage struct {
	Ioc                  bool                        `json:"ioc,omitempty" bson:"ioc"`
	Sighted              bool                        `json:"sighted,omitempty" bson:"sighted"`
	IgnoreSimilarity     bool                        `json:"ignoreSimilarity,omitempty" bson:"ignoreSimilarity"`
	Tlp                  uint64                      `json:"tlp,omitempty" bson:"tlp"`
	UnderliningCreatedAt string                      `json:"_createdAt,omitempty" bson:"_createdAt"` //формат RFC3339
	UnderliningUpdatedAt string                      `json:"_updatedAt,omitempty" bson:"_updatedAt"` //формат RFC3339
	StartDate            string                      `json:"startDate,omitempty" bson:"startDate"`   //формат RFC3339
	UnderliningCreatedBy string                      `json:"_createdBy,omitempty" bson:"_createdBy"`
	UnderliningUpdatedBy string                      `json:"_updatedBy,omitempty" bson:"_updatedBy"`
	UnderliningId        string                      `json:"_id,omitempty" bson:"_id"`
	UnderliningType      string                      `json:"_type,omitempty" bson:"_type"`
	Data                 string                      `json:"data,omitempty" bson:"data"`
	DataType             string                      `json:"dataType,omitempty" bson:"dataType"`
	Message              string                      `json:"message,omitempty" bson:"message"`
	Tags                 []string                    `json:"tags" bson:"tags"`
	Attachment           AttachmentData              `json:"attachment,omitempty" bson:"attachment"`
	Reports              map[string]ReportTaxonomies `json:"reports" bson:"reports"`
	//данное поле редко используемое, думаю пока оно не требует реализации
	//ExtraData        map[string]interface{}                         `json:"extraData"`
}
