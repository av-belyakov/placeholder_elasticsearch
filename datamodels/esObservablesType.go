package datamodels

import "placeholder_elasticsearch/datamodels/commonobservable"

// ObservablesMessageEs список наблюдаемых сообщений
// Observables - наблюдаемые сообщения
type ObservablesMessageEs struct {
	Observables map[string][]ObservableMessageEs `json:"observables" bson:"observables"`
}

// ObservableMessageEs наблюдаемое сообщение
// SensorId - идентификатор сенсора
// SnortSid - список идентификаторов сигнатур (строка)
// SnortSidNumber - список идентификаторов сигнатур (число)
// Tags - список тегов
// TagsAll - список всех тегов
// Attachment - приложенные данные
// Reports - список отчетов (ИСКЛЮЧИЛ из-за черезмерно большого количества
// полей которое влечет за собой превышения лимита маппинга в Elsticsearch)
// что выражается в ошибке от СУБД типа "Limit of total fields [2000] has
// been exceeded while adding new fields"
type ObservableMessageEs struct {
	commonobservable.CommonObservableType
	SensorId       string              `json:"sensorId,omitempty" bson:"sensorId"`
	SnortSid       []string            `json:"snortSid,omitempty" bson:"snortSid"`
	SnortSidNumber []int               `json:"SnortSidNumber,omitempty" bson:"SnortSidNumber"`
	Tags           map[string][]string `json:"tags" bson:"tags"`
	TagsAll        []string            `json:"tagsAll" bson:"tagsAll"`
	Attachment     AttachmentData      `json:"attachment,omitempty" bson:"attachment"`
	// Reports    map[string]ReportTaxonomies `json:"reports" bson:"reports"`
}
