package datamodels

import (
	"fmt"
	"strings"

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
