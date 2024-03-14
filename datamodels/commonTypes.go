package datamodels

import "strings"

// MessageLogging содержит информацию используемую при логировании
// MsgData - сообщение
// MsgType - тип сообщения
type MessageLogging struct {
	MsgData, MsgType string
}

// DataCounterSettings содержит информацию для подсчета
type DataCounterSettings struct {
	DataType string
	DataMsg  string
	Count    int
}

// ListEventObjectTags временное хранилище для тегов полученных из event.object.tags
type ListEventObjectTags []string

func NewListEventObjectTags() *ListEventObjectTags {
	return &ListEventObjectTags{}
}

func (leot *ListEventObjectTags) SetTag(v string) {
	if !strings.Contains(v, "ATs:") {
		return
	}

	*leot = append(*leot, v)
}

func (leot *ListEventObjectTags) GetListTags() ListEventObjectTags {
	return *leot
}

func (leot *ListEventObjectTags) CleanListTags() {
	leot = &ListEventObjectTags{}
}

// ChanOutputDecodeJSON содержит данные получаемые при декодировании JSON формата
// обрабатываемого обработчиком HandlerMessageFromHive
// ExclusionRuleWorked - информирует что сработало правило исключения значения из списка
// передаваемых данных
// UUID - уникальный идентификатор в формате UUID
// FieldName - наименование поля
// ValueType - тип передаваемого значения (string, int и т.д.)
// Value - любые передаваемые данные
// FieldBranch - 'путь' до значения в как в JSON формате, например 'event.details.customFields.class'
type ChanOutputDecodeJSON struct {
	ExclusionRuleWorked bool
	UUID                string
	FieldName           string
	ValueType           string
	Value               interface{}
	FieldBranch         string
}

type HitsChildren struct {
	Id     string             `json:"_id"`
	Index  string             `json:"_index"`
	Source VerifiedForEsAlert `json:"_source"`
}

type HitsParrent struct {
	Hits []HitsChildren `json:"hits"`
}

type ShardsPattern struct {
	Total      int `json:"total"`
	Failed     int `json:"failed"`
	Skipped    int `json:"skipped"`
	Successful int `json:"successful"`
}

type ElasticsearchPatternVerifiedForEsAlert struct {
	Hits HitsParrent `json:"hits"`
}

// AttachmentData прикрепленные данные
// Size - размер
// Id - идентификатор
// Name - наименование
// ContentType - тип контента
// Hashes - список хешей
type AttachmentData struct {
	Size        uint64   `json:"size" bson:"size"`
	Id          string   `json:"id" bson:"id"`
	Name        string   `json:"name" bson:"name"`
	ContentType string   `json:"contentType" bson:"contentType"`
	Hashes      []string `json:"hashes" bson:"hashes"`
}

// ReportTaxonomies
type ReportTaxonomies struct {
	Taxonomies []Taxonomy `json:"taxonomies,omitempty" bson:"taxonomies"`
}

// Taxonomy
type Taxonomy struct {
	Level     string `json:"level" bson:"level"`
	Namespace string `json:"namespace" bson:"namespace"`
	Predicate string `json:"predicate" bson:"predicate"`
	Value     string `json:"value" bson:"value"`
}

// CustomerFields
// Set принимает значения где первое значение метода это первое значение
// в структуре данных, второе значение метода это второе значение
// в структуре данных
// Get возвращает значения где 1 и 3 значение это
// наименование поля
type CustomerFields interface {
	Set(oneStructField interface{}, twoStructField interface{})
	Get() (fieldNameOne string, valueOne int, fieldNameTwo string, valueTwo string)
}

type CustomFields map[string]CustomerFields

type CustomFieldStringType struct {
	Order  int    `json:"order" bson:"order"`
	String string `json:"string" bson:"string"`
}

type CustomFieldDateType struct {
	Order int    `json:"order" bson:"order"`
	Date  string `json:"date" bson:"date"`
}

type CustomFieldIntegerType struct {
	Order   int `json:"order" bson:"order"`
	Integer int `json:"integer" bson:"integer"`
}

type CustomFieldBoolenType struct {
	Order   int  `json:"order" bson:"order"`
	Boolean bool `json:"boolen" bson:"boolen"`
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
	OccurDate            string              `json:"occurDate" bson:"occurDate"`   //формат RFC3339
	UnderliningCreatedAt string              `json:"_createdAt" bson:"_createdAt"` //формат RFC3339
	UnderliningId        string              `json:"_id" bson:"_id"`
	UnderliningCreatedBy string              `json:"_createdBy" bson:"_createdBy"`
	PatternId            string              `json:"patternId" bson:"patternId"`
	Tactic               string              `json:"tactic" bson:"tactic"`
	ExtraData            ExtraDataTtpMessage `json:"extraData" bson:"extraData"`
}

// ExtraDataTtpMessage дополнительные данные TTP сообщения
// Pattern - шаблон
// PatternParent - родительский шаблон
type ExtraDataTtpMessage struct {
	Pattern       PatternExtraData `json:"pattern" bson:"pattern"`
	PatternParent PatternExtraData `json:"patternParent" bson:"patternParent"`
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
	RemoteSupport        bool     `json:"remoteSupport" bson:"remoteSupport"`
	Revoked              bool     `json:"revoked" bson:"revoked"`
	UnderliningCreatedAt string   `json:"_createdAt" bson:"_createdAt"` //формат RFC3339
	UnderliningCreatedBy string   `json:"_createdBy" bson:"_createdBy"`
	UnderliningId        string   `json:"_id" bson:"_id"`
	UnderliningType      string   `json:"_type" bson:"_type"`
	Description          string   `json:"description" bson:"description"`
	Detection            string   `json:"detection" bson:"detection"`
	Name                 string   `json:"name" bson:"name"`
	PatternId            string   `json:"patternId" bson:"patternId"`
	PatternType          string   `json:"patternType" bson:"patternType"`
	URL                  string   `json:"url" bson:"url"`
	Version              string   `json:"version" bson:"version"`
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
