package datamodel

// AttachmentData прикрепленные данные
// Size - размер
// Id - идентификатор
// Name - наименование
// ContentType - тип контента
// Hashes - список хешей
type AttachmentData struct {
	Size        uint64   `json:"size,omitempty" bson:"size"`
	Id          string   `json:"id,omitempty" bson:"id"`
	Name        string   `json:"name,omitempty" bson:"name"`
	ContentType string   `json:"contentType,omitempty" bson:"contentType"`
	Hashes      []string `json:"hashes,omitempty" bson:"hashes"`
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
