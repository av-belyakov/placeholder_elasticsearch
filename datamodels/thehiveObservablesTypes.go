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
