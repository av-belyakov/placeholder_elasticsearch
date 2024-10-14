package datamodels

import "strings"

// MessageLogging содержит информацию используемую при логировании
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

// NewListEventObjectTags новый список тегов для event.object
func NewListEventObjectTags() *ListEventObjectTags {
	return &ListEventObjectTags{}
}

// SetTag устанавливает теги
func (leot *ListEventObjectTags) SetTag(v string) {
	if !strings.Contains(v, "ATs:") {
		return
	}

	*leot = append(*leot, v)
}

// GetListTags возвращает список тегов
func (leot *ListEventObjectTags) GetListTags() ListEventObjectTags {
	return *leot
}

// CleanListTags очищает список тегов
func (leot *ListEventObjectTags) CleanListTags() {
	leot = &ListEventObjectTags{}
}

// ChanOutputDecodeJSON содержит данные получаемые при декодировании JSON формата
// обрабатываемого обработчиком HandlerMessageFromHive
type ChanOutputDecodeJSON struct {
	ExclusionRuleWorked bool        //информирует что сработало правило исключения значения из списка передаваемых данных
	UUID                string      //уникальный идентификатор в формате UUID
	FieldName           string      //наименование поля
	FieldBranch         string      //'путь' до значения в как в JSON формате, например 'event.details.customFields.class'
	ValueType           string      //тип передаваемого значения (string, int и т.д.)
	Value               interface{} //любые передаваемые данные
}

type HitsParrent struct {
	Hits []HitsChildren `json:"hits"`
}

type HitsChildren struct {
	Id     string             `json:"_id"`
	Index  string             `json:"_index"`
	Source VerifiedForEsAlert `json:"_source"`
}

type ShardsPattern struct {
	Total      int `json:"total"`
	Failed     int `json:"failed"`
	Skipped    int `json:"skipped"`
	Successful int `json:"successful"`
}

// AttachmentData прикрепленные данные
type AttachmentData struct {
	Size        uint64   `json:"size,omitempty" bson:"size"`               //размер
	Id          string   `json:"id,omitempty" bson:"id"`                   //идентификатор
	Name        string   `json:"name,omitempty" bson:"name"`               //наименование
	ContentType string   `json:"contentType,omitempty" bson:"contentType"` //тип контента
	Hashes      []string `json:"hashes,omitempty" bson:"hashes"`           //список хешей
}

// ReportTaxonomies
type ReportTaxonomies struct {
	Taxonomies []Taxonomy `json:"taxonomies,omitempty" bson:"taxonomies"`
}

// Taxonomy
type Taxonomy struct {
	Level     string `json:"level,omitempty" bson:"level"`
	Namespace string `json:"namespace,omitempty" bson:"namespace"`
	Predicate string `json:"predicate,omitempty" bson:"predicate"`
	Value     string `json:"value,omitempty" bson:"value"`
}

// CustomerFields
type CustomerFields interface {
	Set(oneStructField interface{}, twoStructField interface{})                     //принимает значения где первое значение метода это первое значение в структуре данных, второе значение метода это второе значение в структуре данных
	Get() (fieldNameOne string, valueOne int, fieldNameTwo string, valueTwo string) //возвращает значения где 1 и 3 значение это наименование поля
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
type TtpMessage struct {
	OccurDate            string              `json:"occurDate,omitempty" bson:"occurDate"`   //формат RFC3339 дата возникновения
	UnderliningCreatedAt string              `json:"_createdAt,omitempty" bson:"_createdAt"` //формат RFC3339 время создания
	UnderliningId        string              `json:"_id,omitempty" bson:"_id"`               //уникальный идентификатор
	UnderliningCreatedBy string              `json:"_createdBy,omitempty" bson:"_createdBy"` //кем создан
	PatternId            string              `json:"patternId,omitempty" bson:"patternId"`   //уникальный идентификатор шаблона
	Tactic               string              `json:"tactic,omitempty" bson:"tactic"`         //тактика
	ExtraData            ExtraDataTtpMessage `json:"extraData" bson:"extraData"`             //дополнительные данные
}

// ExtraDataTtpMessage дополнительные данные TTP сообщения
type ExtraDataTtpMessage struct {
	Pattern       PatternExtraData `json:"pattern,omitempty" bson:"pattern"`             //шаблон
	PatternParent PatternExtraData `json:"patternParent,omitempty" bson:"patternParent"` //родительский шаблон
}

// PatternExtraData шаблон дополнительных данных
type PatternExtraData struct {
	RemoteSupport        bool     `json:"remoteSupport,omitempty" bson:"remoteSupport"`   //удаленная поддержка
	Revoked              bool     `json:"revoked,omitempty" bson:"revoked"`               //аннулированный
	UnderliningCreatedAt string   `json:"_createdAt,omitempty" bson:"_createdAt"`         //формат RFC3339 время создания
	UnderliningCreatedBy string   `json:"_createdBy,omitempty" bson:"_createdBy"`         //кем создан
	UnderliningId        string   `json:"_id,omitempty" bson:"_id"`                       //уникальный идентификатор
	UnderliningType      string   `json:"_type,omitempty" bson:"_type"`                   //тип
	Description          string   `json:"description,omitempty" bson:"description"`       //описание
	Detection            string   `json:"detection,omitempty" bson:"detection"`           //обнаружен
	Name                 string   `json:"name,omitempty" bson:"name"`                     //наименование
	PatternId            string   `json:"patternId,omitempty" bson:"patternId"`           //уникальный идентификатор шаблона
	PatternType          string   `json:"patternType,omitempty" bson:"patternType"`       //тип шаблона
	URL                  string   `json:"url,omitempty" bson:"url"`                       //URL
	Version              string   `json:"version,omitempty" bson:"version"`               //версия
	Platforms            []string `json:"platforms" bson:"platforms"`                     //список платформ
	PermissionsRequired  []string `json:"permissionsRequired" bson:"permissionsRequired"` //требуемые разрешения
	DataSources          []string `json:"dataSources" bson:"dataSources"`                 //источники данных
	Tactics              []string `json:"tactics" bson:"tactics"`                         //тактики
	//данное поле редко используемое, думаю пока оно не требует реализации
	//DefenseBypassed     []string               `json:"defenseBypassed"` //надо проверить тип
	//данное поле редко используемое, думаю пока оно не требует реализации
	//SystemRequirements  []string               `json:"systemRequirements"` //надо проверить тип
	//данное поле редко используемое, думаю пока оно не требует реализации
	//ExtraData           map[string]interface{} `json:"extraData"`
}

// AdditionalInformation дополнительная информация добавляемая к информации по кейсам
type AdditionalInformation struct {
	Sensors     []SensorInformation      `json:"@sensorAdditionalInformation"`
	IpAddresses []IpAddressesInformation `json:"@ipAddressAdditionalInformation"`
}

// SensorInformation содержит дополнительную информацию о сенсоре
type SensorInformation struct {
	SensorId    string `json:"sensorId" bson:"sensorId"`       //идентификатор сенсора
	HostId      string `json:"hostId" bson:"hostId"`           //идентификатор сенсора, специальный, для поиска информации в НКЦКИ
	GeoCode     string `json:"geoCode" bson:"geoCode"`         //географический код страны
	ObjectArea  string `json:"objectArea" bson:"objectArea"`   //сфера деятельности объекта
	SubjectRF   string `json:"subjectRF" bson:"subjectRF"`     //субъект Российской Федерации
	INN         string `json:"inn" bson:"inn"`                 //налоговый идентификатор
	HomeNet     string `json:"homeNet" bson:"homeNet"`         //перечень домашних сетей
	OrgName     string `json:"orgName" bson:"orgName"`         //наименование организации
	FullOrgName string `json:"fullOrgName" bson:"fullOrgName"` //полное наименование организации
}

// IpAddressesInformation дополнительная информация об ip адресе
type IpAddressesInformation struct {
	Ip          string `json:"ip"`          //ip адрес по которомы выполнялся поиск
	City        string `json:"city"`        //город
	Country     string `json:"country"`     //страна
	CountryCode string `json:"countryCode"` //код страны
}
