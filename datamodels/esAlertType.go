package datamodels

import (
	"placeholder_elasticsearch/datamodels/commonalert"
	"placeholder_elasticsearch/datamodels/commonalertartifact"
)

// AlertMessageForEsAlert сообщение с информацией о Alert-те
// Tags - теги после обработки
// TagsAll - все теги
// CustomFields - настраиваемые поля
// Artifacts - артефакты
type AlertMessageForEsAlert struct {
	commonalert.CommonAlertType
	Tags         map[string][]string             `json:"tags" bson:"tags"`
	TagsAll      []string                        `json:"tagsAll" bson:"tagsAll"`
	CustomFields CustomFields                    `json:"customFields" bson:"customFields"`
	Artifacts    map[string][]ArtifactForEsAlert `json:"artifact" bson:"artifact"`
}

// ArtifactForEsAlert содержит артефакт к алерту
// Tags - теги после обработки
// TagsAll - все теги
type ArtifactForEsAlert struct {
	commonalertartifact.CommonArtifactType
	Tags    map[string][]string `json:"tags" bson:"tags"`
	TagsAll []string            `json:"tagsAll" bson:"tagsAll"`
}
