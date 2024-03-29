package datamodels

import (
	"placeholder_elasticsearch/datamodels/commonalert"
	"placeholder_elasticsearch/datamodels/commonalertartifact"
)

// AlertMessageTheHiveAlert сообщение с информацией о Alert-те
// Follow - следовать
// Severity - строгость
// Tags - теги
// CustomFields - настраиваемые поля
// Artifacts - артефакты
type AlertMessageTheHiveAlert struct {
	Follow   bool   `json:"follow,omitempty" bson:"follow"`
	Severity uint64 `json:"severity,omitempty" bson:"severity"`
	commonalert.CommonAlertType
	Tags         []string        `json:"tags" bson:"tags"`
	CustomFields CustomFields    `json:"customFields,omitempty" bson:"customFields"`
	Artifacts    []AlertArtifact `json:"artifact" bson:"artifact"`
}

// AlertArtifact содержит артефакт к алерту
// Ioc - индикатор компрометации
// Sighted - видящий
// IgnoreSimilarity - игнорировать похожие
// Tlp - tlp
// UnderliningId - уникальный идентификатор
// Id - уникальный идентификатор
// UnderliningType - тип
// CreatedAt - время создания
// CreatedBy - кем создан
// StartDate - дата начала
// UpdatedAt - время обновления
// UpdatedBy - кем обновлен
// Data - данные
// DataType - тип данных
// Message - сообщение
// Tags - список тегов
type AlertArtifact struct {
	Sighted          bool `json:"sighted,omitempty" bson:"sighted"`
	IgnoreSimilarity bool `json:"ignoreSimilarity,omitempty" bson:"ignoreSimilarity"`
	commonalertartifact.CommonArtifactType
	UpdatedAt string   `json:"updatedAt" bson:"updatedAt"` //формат RFC3339
	UpdatedBy string   `json:"updatedBy,omitempty" bson:"updatedBy"`
	Tags      []string `json:"tags" bson:"tags"`
}
