package datamodels

// TtpsMessageTheHive список TTP сообщений
type TtpsMessageTheHive struct {
	Ttp []TtpMessage `json:"ttp"`
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
	OccurDate            string              `json:"occurDate"`  //формат RFC3339
	UnderliningCreatedAt string              `json:"_createdAt"` //формат RFC3339
	UnderliningId        string              `json:"_id"`
	UnderliningCreatedBy string              `json:"_createdBy"`
	PatternId            string              `json:"patternId"`
	Tactic               string              `json:"tactic"`
	ExtraData            ExtraDataTtpMessage `json:"extraData"`
}

// ExtraDataTtpMessage дополнительные данные TTP сообщения
// Pattern - шаблон
// PatternParent - родительский шаблон
type ExtraDataTtpMessage struct {
	Pattern       PatternExtraData `json:"pattern"`
	PatternParent PatternExtraData `json:"patternParent"`
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
	RemoteSupport        bool     `json:"remoteSupport"`
	Revoked              bool     `json:"revoked"`
	UnderliningCreatedAt string   `json:"_createdAt"` //формат RFC3339
	UnderliningCreatedBy string   `json:"_createdBy"`
	UnderliningId        string   `json:"_id"`
	UnderliningType      string   `json:"_type"`
	Description          string   `json:"description"`
	Detection            string   `json:"detection"`
	Name                 string   `json:"name"`
	PatternId            string   `json:"patternId"`
	PatternType          string   `json:"patternType"`
	URL                  string   `json:"url"`
	Version              string   `json:"version"`
	Platforms            []string `json:"platforms"`
	PermissionsRequired  []string `json:"permissionsRequired"`
	DataSources          []string `json:"dataSources"`
	Tactics              []string `json:"tactics"`
	//данное поле редко используемое, думаю пока оно не требует реализации
	//DefenseBypassed     []string               `json:"defenseBypassed"` //надо проверить тип
	//данное поле редко используемое, думаю пока оно не требует реализации
	//SystemRequirements  []string               `json:"systemRequirements"` //надо проверить тип
	//данное поле редко используемое, думаю пока оно не требует реализации
	//ExtraData           map[string]interface{} `json:"extraData"`
}
