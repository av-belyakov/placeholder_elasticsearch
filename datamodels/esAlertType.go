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
	Follow   bool   `json:"follow,omitempty" bson:"follow"`
	Severity uint64 `json:"severity,omitempty" bson:"severity"`
	commonalert.CommonAlertType
	Tags         map[string][]string             `json:"tags" bson:"tags"`
	TagsAll      []string                        `json:"tagsAll" bson:"tagsAll"`
	CustomFields CustomFields                    `json:"customFields,omitempty" bson:"customFields"`
	Artifacts    map[string][]ArtifactForEsAlert `json:"artifact" bson:"artifact"`
}

// ArtifactForEsAlert содержит артефакт к алерту
// SensorId - сенсор id
// SnortSid - список snort сигнатур (строка)
// SnortSidNumber - список snort сигнатур (число)
// Tags - теги после обработки
// TagsAll - все теги
type ArtifactForEsAlert struct {
	commonalertartifact.CommonArtifactType
	SensorId       string              `json:"sensorId,omitempty" bson:"sensorId"`
	SnortSid       []string            `json:"snortSid,omitempty" bson:"snortSid"`
	SnortSidNumber []int               `json:"SnortSidNumber,omitempty" bson:"SnortSidNumber"`
	Tags           map[string][]string `json:"tags" bson:"tags"`
	TagsAll        []string            `json:"tagsAll" bson:"tagsAll"`
}
