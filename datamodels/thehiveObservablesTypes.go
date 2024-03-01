package datamodels

import "placeholder_elasticsearch/datamodels/commonobservable"

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
	commonobservable.CommonObservableType
	Tags       []string                    `json:"tags" bson:"tags"`
	Attachment AttachmentData              `json:"attachment" bson:"attachment"`
	Reports    map[string]ReportTaxonomies `json:"reports" bson:"reports"`
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
