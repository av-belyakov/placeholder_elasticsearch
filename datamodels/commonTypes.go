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
