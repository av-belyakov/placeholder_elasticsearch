package datamodels

import (
	"placeholder_elasticsearch/datamodels/commonalert"
	"placeholder_elasticsearch/datamodels/commonalertartifact"
)

// AlertMessageForEsAlert сообщение с информацией о Alert-те
// Follow - следовать
// Severity - строгость
// Tags - теги после обработки
// TagsAll - все теги
// CustomFields - настраиваемые поля
// Artifacts - артефакты
type AlertMessageForEsAlert struct {
	Follow   bool   `json:"follow" bson:"follow"`
	Severity uint64 `json:"severity" bson:"severity"`
	commonalert.CommonAlertType
	Tags         map[string][]string             `json:"tags" bson:"tags"`
	TagsAll      []string                        `json:"tagsAll" bson:"tagsAll"`
	CustomFields CustomFields                    `json:"customFields" bson:"customFields"`
	Artifacts    map[string][]ArtifactForEsAlert `json:"artifact" bson:"artifact"`
}

// ArtifactForEsAlert содержит артефакт к алерту
// SensorId - сенсор id
// SnortSid - список snort сигнатур
// Tags - теги после обработки
// TagsAll - все теги
type ArtifactForEsAlert struct {
	commonalertartifact.CommonArtifactType
	SensorId string              `json:"sensorId" bson:"sensorId"`
	SnortSid []string            `json:"snortSid" bson:"snortSid"`
	Tags     map[string][]string `json:"tags" bson:"tags"`
	TagsAll  []string            `json:"tagsAll" bson:"tagsAll"`
}
