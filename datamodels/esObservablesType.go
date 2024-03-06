package datamodels

import "placeholder_elasticsearch/datamodels/commonobservable"

// ObservablesMessageEs список наблюдаемых сообщений
// Observables - наблюдаемые сообщения
type ObservablesMessageEs struct {
	Observables map[string][]ObservableMessageEs `json:"observables" bson:"observables"`
}

// ObservableMessageEs наблюдаемое сообщение
// Tags - список тегов
// TagsAll - список всех тегов
// Attachment - приложенные данные
// Reports - список отчетов
type ObservableMessageEs struct {
	commonobservable.CommonObservableType
	Tags       map[string][]string         `json:"tags" bson:"tags"`
	TagsAll    []string                    `json:"tagsAll" bson:"tagsAll"`
	Attachment AttachmentData              `json:"attachment" bson:"attachment"`
	Reports    map[string]ReportTaxonomies `json:"reports" bson:"reports"`
}
