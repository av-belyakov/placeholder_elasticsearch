package datamodel

import (
	"encoding/json"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"

	"placeholder_elasticsearch/_moduledatacomparison/supportingfunctions"
)

func CustomFieldsToStringBeautiful(l CustomFields, num int) string {
	strB := strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num + 2)

	for k, v := range l {
		strB.WriteString(fmt.Sprintf("%s'%s':\n", supportingfunctions.GetWhitespace(num+1), k))

		nameOne, dataOne, nameTwo, dataTwo := v.Get()
		strB.WriteString(fmt.Sprintf("%s'%s': %d\n", ws, nameOne, dataOne))
		strB.WriteString(fmt.Sprintf("%s'%s': %s\n", ws, nameTwo, dataTwo))
	}

	return strB.String()
}

func (a AttachmentData) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'size': '%d'\n", ws, a.Size))
	str.WriteString(fmt.Sprintf("%s'id': '%s'\n", ws, a.Id))
	str.WriteString(fmt.Sprintf("%s'name': '%s'\n", ws, a.Name))
	str.WriteString(fmt.Sprintf("%s'contentType': '%s'\n", ws, a.ContentType))
	str.WriteString(fmt.Sprintf("%s'hashes': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, a.Hashes)))

	return str.String()
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
