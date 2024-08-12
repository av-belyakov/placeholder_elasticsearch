package datamodel

// TtpsMessageTheHive список TTP сообщений
type TtpsMessageTheHive struct {
	Ttp []TtpMessage `json:"ttp" bson:"ttp"`
}

// TtpMessage TTP сообщения
// OccurDate - дата возникновения
// UnderliningCreatedAt - время создания
// UnderliningId - уникальный идентификатор
// UnderliningCreatedBy - кем создан
// PatternId - уникальный идентификатор шаблона
// Tactic - тактика
// ExtraData - дополнительные данные
type TtpMessage struct {
	OccurDate            string              `json:"occurDate,omitempty" bson:"occurDate"`   //формат RFC3339
	UnderliningCreatedAt string              `json:"_createdAt,omitempty" bson:"_createdAt"` //формат RFC3339
	UnderliningId        string              `json:"_id,omitempty" bson:"_id"`
	UnderliningCreatedBy string              `json:"_createdBy,omitempty" bson:"_createdBy"`
	PatternId            string              `json:"patternId,omitempty" bson:"patternId"`
	Tactic               string              `json:"tactic,omitempty" bson:"tactic"`
	ExtraData            ExtraDataTtpMessage `json:"extraData" bson:"extraData"`
}

// ExtraDataTtpMessage дополнительные данные TTP сообщения
// Pattern - шаблон
// PatternParent - родительский шаблон
type ExtraDataTtpMessage struct {
	Pattern       PatternExtraData `json:"pattern,omitempty" bson:"pattern"`
	PatternParent PatternExtraData `json:"patternParent,omitempty" bson:"patternParent"`
}

// PatternExtraData шаблон дополнительных данных
// RemoteSupport - удаленная поддержка
// Revoked - аннулированный
// UnderliningCreatedAt - время создания
// UnderliningCreatedBy - кем создан
// UnderliningId - уникальный идентификатор
// UnderliningType - тип
// DataSources - источники данных
// DefenseBypassed - чем выполнен обход защиты
// Description - описание
// ExtraData - дополнительные данные
// Name - наименование
// PatternId - уникальный идентификатор шаблона
// PatternType - тип шаблона
// PermissionsRequired - требуемые разрешения
// Platforms - список платформ
// SystemRequirements - системные требования
// Tactics - список тактик
// URL - URL
// Version - версия
type PatternExtraData struct {
	RemoteSupport        bool     `json:"remoteSupport,omitempty" bson:"remoteSupport"`
	Revoked              bool     `json:"revoked,omitempty" bson:"revoked"`
	UnderliningCreatedAt string   `json:"_createdAt,omitempty" bson:"_createdAt"` //формат RFC3339
	UnderliningCreatedBy string   `json:"_createdBy,omitempty" bson:"_createdBy"`
	UnderliningId        string   `json:"_id,omitempty" bson:"_id"`
	UnderliningType      string   `json:"_type,omitempty" bson:"_type"`
	Description          string   `json:"description,omitempty" bson:"description"`
	Detection            string   `json:"detection,omitempty" bson:"detection"`
	Name                 string   `json:"name,omitempty" bson:"name"`
	PatternId            string   `json:"patternId,omitempty" bson:"patternId"`
	PatternType          string   `json:"patternType,omitempty" bson:"patternType"`
	URL                  string   `json:"url,omitempty" bson:"url"`
	Version              string   `json:"version,omitempty" bson:"version"`
	Platforms            []string `json:"platforms" bson:"platforms"`
	PermissionsRequired  []string `json:"permissionsRequired" bson:"permissionsRequired"`
	DataSources          []string `json:"dataSources" bson:"dataSources"`
	Tactics              []string `json:"tactics" bson:"tactics"`
	//данное поле редко используемое, думаю пока оно не требует реализации
	//DefenseBypassed     []string               `json:"defenseBypassed"` //надо проверить тип
	//данное поле редко используемое, думаю пока оно не требует реализации
	//SystemRequirements  []string               `json:"systemRequirements"` //надо проверить тип
	//данное поле редко используемое, думаю пока оно не требует реализации
	//ExtraData           map[string]interface{} `json:"extraData"`
}
