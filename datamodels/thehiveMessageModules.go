package datamodels

import (
	"encoding/json"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"

	"placeholder_elasticsearch/supportingfunctions"
)

// NewResponseMessage формирует новый тип ResponseMessageFromMispToTheHave с предустановленными значениями
func NewResponseMessage() *ResponseMessageFromMispToTheHave {
	return &ResponseMessageFromMispToTheHave{
		Success: true,
		Service: "MISP",
		Commands: []ResponseCommandForTheHive{
			{
				Command: "addtag",
				String:  "Webhook: send=\"MISP\"",
				//String:  "Webhook: send=\"MISP-WORLD\"",
				//String: "Webhook: send=\"MISP-CENTER\"",
			},
		},
	}
}

func (rm *ResponseMessageFromMispToTheHave) ResponseMessageAddNewCommand(rcm ResponseCommandForTheHive) {
	rm.Commands = append(rm.Commands, rcm)
}

func (rm *ResponseMessageFromMispToTheHave) GetResponseMessageFromMispToTheHave() ResponseMessageFromMispToTheHave {
	return *rm
}

// Get возвращает MainMessageTheHive
func (mm *MainMessageTheHive) Get() *MainMessageTheHive {
	return mm
}

func (mm MainMessageTheHive) ToStringBeautiful(num int) string {
	strB := strings.Builder{}

	strB.WriteString(mm.SourceMessageTheHive.ToStringBeautiful(num + 1))
	strB.WriteString(fmt.Sprintln("event:"))
	strB.WriteString(mm.EventMessageTheHiveCase.ToStringBeautiful(num + 1))
	strB.WriteString(fmt.Sprintln("observables:"))
	strB.WriteString(mm.ObservablesMessageTheHive.ToStringBeautiful(num + 1))
	strB.WriteString(fmt.Sprintln("ttps:"))
	strB.WriteString(mm.TtpsMessageTheHive.ToStringBeautiful(num + 1))

	return strB.String()
}

// GetSource возвращает содержимое поля Source
func (s *SourceMessageTheHive) GetSource() string {
	return s.Source
}

// SetValueSource устанавливает СТРОКОВОЕ значение для поля Source
func (s *SourceMessageTheHive) SetValueSource(v string) {
	s.Source = v
}

// SetAnySource устанавливает ЛЮБОЕ значение для поля Source
func (s *SourceMessageTheHive) SetAnySource(i interface{}) {
	s.Source = fmt.Sprint(i)
}

func (s SourceMessageTheHive) ToStringBeautiful(num int) string {
	return fmt.Sprintf("source: '%s'\n", s.Source)
}

func (fields *CustomFields) Get() CustomFields {
	return *fields
}

func (fields *CustomFields) Set(v CustomFields) {
	if *fields == nil {
		*fields = make(CustomFields)
	}

	for key, value := range v {
		(*fields)[key] = value
	}
}

func (fields *CustomFields) UnmarshalJSON(data []byte) error {
	type tmpCustomFieldType map[string]*json.RawMessage

	list := map[string]string{
		"attack-type":         "string",
		"class-attack":        "string",
		"event-source":        "string",
		"external-letter":     "string",
		"geoip":               "string",
		"ncircc-class-attack": "string",
		"ncircc-bulletin-id":  "string",
		"inbox1":              "string",
		"inner-letter":        "string",
		"ir-name":             "string",
		"id-soa":              "string",
		"notification":        "string",
		"sphere":              "string",
		"state":               "string",
		"report":              "string",
		"first-time":          "date",
		"last-time":           "date",
		"b2mid":               "integer",
		"is-incident":         "boolen",
		"work-admin":          "boolen",
	}

	tmp := tmpCustomFieldType{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	//newResult := make(map[string]CustomerFields, len(tmp.CustomFields))
	newResult := make(CustomFields, len(tmp))
	for k, v := range tmp {
		name, ok := list[k]
		if !ok {
			continue
		}

		switch name {
		case "string":
			custField := &CustomFieldStringType{}
			if err := json.Unmarshal(*v, custField); err != nil {
				return err
			}

			newResult[k] = custField

		case "date":
			custField := &CustomFieldDateType{}
			if err := json.Unmarshal(*v, custField); err != nil {
				return err
			}

			newResult[k] = custField

		case "integer":
			custField := &CustomFieldIntegerType{}
			if err := json.Unmarshal(*v, custField); err != nil {
				return err
			}

			newResult[k] = custField

		case "boolen":
			custField := &CustomFieldBoolenType{}
			if err := json.Unmarshal(*v, custField); err != nil {
				return err
			}

			newResult[k] = custField
		}
	}

	fields.Set(newResult)

	return nil
}

func (fields *CustomFields) UnmarshalBSON(data []byte) error {
	type tmpCustomFieldType map[string]*bson.Raw

	list := map[string]string{
		"attack-type":         "string",
		"class-attack":        "string",
		"event-source":        "string",
		"external-letter":     "string",
		"geoip":               "string",
		"ncircc-class-attack": "string",
		"ncircc-bulletin-id":  "string",
		"inbox1":              "string",
		"inner-letter":        "string",
		"ir-name":             "string",
		"id-soa":              "string",
		"notification":        "string",
		"sphere":              "string",
		"state":               "string",
		"report":              "string",
		"first-time":          "date",
		"last-time":           "date",
		"b2mid":               "integer",
		"is-incident":         "boolen",
		"work-admin":          "boolen",
	}

	tmp := tmpCustomFieldType{}
	if err := bson.Unmarshal(data, &tmp); err != nil {
		return err
	}

	//newResult := make(map[string]CustomerFields, len(tmp.CustomFields))
	newResult := make(CustomFields, len(tmp))
	for k, v := range tmp {
		name, ok := list[k]
		if !ok {
			continue
		}

		switch name {
		case "string":
			custField := &CustomFieldStringType{}
			if err := bson.Unmarshal(*v, custField); err != nil {
				return err
			}

			newResult[k] = custField

		case "date":
			custField := &CustomFieldDateType{}
			if err := bson.Unmarshal(*v, custField); err != nil {
				return err
			}

			newResult[k] = custField

		case "integer":
			custField := &CustomFieldIntegerType{}
			if err := bson.Unmarshal(*v, custField); err != nil {
				return err
			}

			newResult[k] = custField

		case "boolen":
			custField := &CustomFieldBoolenType{}
			if err := bson.Unmarshal(*v, custField); err != nil {
				return err
			}

			newResult[k] = custField
		}
	}

	fields.Set(newResult)

	return nil
}

// Get возвращает значения CustomFieldStringType, где 1 и 3 значение это
// наименование поля
func (cf *CustomFieldStringType) Get() (string, int, string, string) {
	return "order", cf.Order, "string", cf.String
}

// Set устанавливает значения CustomFieldStringType
func (cf *CustomFieldStringType) Set(order, str interface{}) {
	cf.Order = supportingfunctions.ConversionAnyToInt(order)
	cf.String = fmt.Sprint(str)
}

// Get возвращает значения CustomFieldDateType, где 1 и 3 значение это
// наименование поля
func (cf *CustomFieldDateType) Get() (string, int, string, string) {
	return "order", cf.Order, "date", cf.Date
}

// Set устанавливает значения CustomFieldDateType
func (cf *CustomFieldDateType) Set(order, date interface{}) {
	cf.Order = supportingfunctions.ConversionAnyToInt(order)

	if str, ok := date.(string); ok {
		cf.Date = str

		return
	}

	tmp := supportingfunctions.ConversionAnyToInt(date)
	cf.Date = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

// Get возвращает значения CustomFieldIntegerType, где 1 и 3 значение это
// наименование поля
func (cf *CustomFieldIntegerType) Get() (string, int, string, string) {
	return "order", cf.Order, "integer", fmt.Sprint(cf.Integer)
}

// Set устанавливает значения CustomFieldIntegerType
func (cf *CustomFieldIntegerType) Set(order, integer interface{}) {
	cf.Order = supportingfunctions.ConversionAnyToInt(order)

	if i, ok := integer.(int); ok {
		cf.Integer = i
	}
}

// Get возвращает значения CustomFieldBoolenType, где 1 и 3 значение это
// наименование поля
func (cf *CustomFieldBoolenType) Get() (string, int, string, string) {
	return "order", cf.Order, "boolen", fmt.Sprint(cf.Boolean)
}

// Set устанавливает значения CustomFieldBoolenType
func (cf *CustomFieldBoolenType) Set(order, boolen interface{}) {
	cf.Order = supportingfunctions.ConversionAnyToInt(order)

	if i, ok := boolen.(bool); ok {
		cf.Boolean = i
	}
}
