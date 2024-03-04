package commonobjectevent

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/supportingfunctions"
)

func (o *CommonEventAlertObject) Get() *CommonEventAlertObject {
	return o
}

func (o *CommonEventAlertObject) GetTlp() uint64 {
	return o.Tlp
}

// SetValueTlp устанавливает INT значение для поля Tlp
func (o *CommonEventAlertObject) SetValueTlp(v uint64) {
	o.Tlp = v
}

// SetAnyTlp устанавливает ЛЮБОЕ значение для поля Tlp
func (o *CommonEventAlertObject) SetAnyTlp(i interface{}) {
	if v, ok := i.(float32); ok {
		o.Tlp = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		o.Tlp = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		o.Tlp = v
	}
}

func (o *CommonEventAlertObject) GetUnderliningId() string {
	return o.UnderliningId
}

// SetValueUnderliningId устанавливает STRING значение для поля UnderliningId
func (o *CommonEventAlertObject) SetValueUnderliningId(v string) {
	o.UnderliningId = v
}

// SetAnyUnderliningId устанавливает ЛЮБОЕ значение для поля UnderliningId
func (o *CommonEventAlertObject) SetAnyUnderliningId(i interface{}) {
	o.UnderliningId = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetId() string {
	return o.Id
}

// SetValueId устанавливает STRING значение для поля Id
func (o *CommonEventAlertObject) SetValueId(v string) {
	o.Id = v
}

// SetAnyId устанавливает ЛЮБОЕ значение для поля Id
func (o *CommonEventAlertObject) SetAnyId(i interface{}) {
	o.Id = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetCreatedBy() string {
	return o.CreatedBy
}

// SetValueCreatedBy устанавливает STRING значение для поля CreatedBy
func (o *CommonEventAlertObject) SetValueCreatedBy(v string) {
	o.CreatedBy = v
}

// SetAnyCreatedBy устанавливает ЛЮБОЕ значение для поля CreatedBy
func (o *CommonEventAlertObject) SetAnyCreatedBy(i interface{}) {
	o.CreatedBy = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetUpdatedBy() string {
	return o.UpdatedBy
}

// SetValueUpdatedBy устанавливает STRING значение для поля UpdatedBy
func (o *CommonEventAlertObject) SetValueUpdatedBy(v string) {
	o.UpdatedBy = v
}

// SetAnyUpdatedBy устанавливает ЛЮБОЕ значение для поля UpdatedBy
func (o *CommonEventAlertObject) SetAnyUpdatedBy(i interface{}) {
	o.UpdatedBy = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetCreatedAt() string {
	return o.CreatedAt
}

// SetValueCreatedAt устанавливает значение в формате RFC3339 для поля CreatedAt
func (o *CommonEventAlertObject) SetValueCreatedAt(v string) {
	o.CreatedAt = v
}

// SetAnyCreatedAt устанавливает ЛЮБОЕ значение для поля CreatedAt
func (o *CommonEventAlertObject) SetAnyCreatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	o.CreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (o *CommonEventAlertObject) GetUpdatedAt() string {
	return o.UpdatedAt
}

// SetValueUpdatedAt устанавливает значение  в формате RFC3339 для поля UpdatedAt
func (o *CommonEventAlertObject) SetValueUpdatedAt(v string) {
	o.UpdatedAt = v
}

// SetAnyUpdatedAt устанавливает ЛЮБОЕ значение для поля UpdatedAt
func (o *CommonEventAlertObject) SetAnyUpdatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	o.UpdatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (o *CommonEventAlertObject) GetUnderliningType() string {
	return o.UnderliningType
}

// SetValueUnderliningType устанавливает STRING значение для поля UnderliningType
func (o *CommonEventAlertObject) SetValueUnderliningType(v string) {
	o.UnderliningType = v
}

// SetAnyUnderliningType устанавливает ЛЮБОЕ значение для поля UnderliningType
func (o *CommonEventAlertObject) SetAnyUnderliningType(i interface{}) {
	o.UnderliningType = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetTitle() string {
	return o.Title
}

// SetValueTitle устанавливает STRING значение для поля Title
func (o *CommonEventAlertObject) SetValueTitle(v string) {
	o.Title = v
}

// SetAnyTitle устанавливает ЛЮБОЕ значение для поля Title
func (o *CommonEventAlertObject) SetAnyTitle(i interface{}) {
	o.Title = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetDescription() string {
	return o.Description
}

// SetValueDescription устанавливает STRING значение для поля Description
func (o *CommonEventAlertObject) SetValueDescription(v string) {
	v = strings.ReplaceAll(v, "\t", "")
	v = strings.ReplaceAll(v, "\n", "")

	o.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (o *CommonEventAlertObject) SetAnyDescription(i interface{}) {
	str := fmt.Sprint(i)
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\n", "")

	o.Description = str
}

func (o *CommonEventAlertObject) GetStatus() string {
	return o.Status
}

// SetValueStatus устанавливает STRING значение для поля Status
func (o *CommonEventAlertObject) SetValueStatus(v string) {
	o.Status = v
}

// SetAnyStatus устанавливает ЛЮБОЕ значение для поля Status
func (o *CommonEventAlertObject) SetAnyStatus(i interface{}) {
	o.Status = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetDate() string {
	return o.Date
}

// SetValueDate устанавливает значение в формате RFC3339 для поля Date
func (o *CommonEventAlertObject) SetValueDate(v string) {
	o.Date = v
}

// SetAnyDate устанавливает ЛЮБОЕ значение для поля Date
func (o *CommonEventAlertObject) SetAnyDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	o.Date = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (o *CommonEventAlertObject) GetType() string {
	return o.Type
}

// SetValueType устанавливает STRING значение для поля Type
func (o *CommonEventAlertObject) SetValueType(v string) {
	o.Type = v
}

// SetAnyType устанавливает ЛЮБОЕ значение для поля Type
func (o *CommonEventAlertObject) SetAnyType(i interface{}) {
	o.Type = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetObjectType() string {
	return o.ObjectType
}

// SetValueObjectType устанавливает STRING значение для поля ObjectType
func (o *CommonEventAlertObject) SetValueObjectType(v string) {
	o.ObjectType = v
}

// SetAnyObjectType устанавливает ЛЮБОЕ значение для поля ObjectType
func (o *CommonEventAlertObject) SetAnyObjectType(i interface{}) {
	o.ObjectType = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetSource() string {
	return o.Source
}

// SetValueSource устанавливает STRING значение для поля Source
func (o *CommonEventAlertObject) SetValueSource(v string) {
	o.Source = v
}

// SetAnySource устанавливает ЛЮБОЕ значение для поля Source
func (o *CommonEventAlertObject) SetAnySource(i interface{}) {
	o.Source = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetSourceRef() string {
	return o.SourceRef
}

// SetValueSourceRef устанавливает STRING значение для поля SourceRef
func (o *CommonEventAlertObject) SetValueSourceRef(v string) {
	o.SourceRef = v
}

// SetAnySourceRef устанавливает ЛЮБОЕ значение для поля SourceRef
func (o *CommonEventAlertObject) SetAnySourceRef(i interface{}) {
	o.SourceRef = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetCase() string {
	return o.Case
}

// SetValueCase устанавливает STRING значение для поля Case
func (o *CommonEventAlertObject) SetValueCase(v string) {
	o.Case = v
}

// SetAnyCase устанавливает ЛЮБОЕ значение для поля Case
func (o *CommonEventAlertObject) SetAnyCase(i interface{}) {
	o.Case = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) GetCaseTemplate() string {
	return o.CaseTemplate
}

// SetValueCaseTemplate устанавливает STRING значение для поля CaseTemplate
func (o *CommonEventAlertObject) SetValueCaseTemplate(v string) {
	o.CaseTemplate = v
}

// SetAnyCaseTemplate устанавливает ЛЮБОЕ значение для поля CaseTemplate
func (o *CommonEventAlertObject) SetAnyCaseTemplate(i interface{}) {
	o.CaseTemplate = fmt.Sprint(i)
}

func (o *CommonEventAlertObject) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'_id': '%s'\n", ws, o.UnderliningId))
	str.WriteString(fmt.Sprintf("%s'id': '%s'\n", ws, o.Id))
	str.WriteString(fmt.Sprintf("%s'createdBy': '%s'\n", ws, o.CreatedBy))
	str.WriteString(fmt.Sprintf("%s'updatedBy': '%s'\n", ws, o.UpdatedBy))
	str.WriteString(fmt.Sprintf("%s'createdAt': '%s'\n", ws, o.CreatedAt))
	str.WriteString(fmt.Sprintf("%s'updatedAt': '%s'\n", ws, o.UpdatedAt))
	str.WriteString(fmt.Sprintf("%s'_type': '%s'\n", ws, o.UnderliningType))
	str.WriteString(fmt.Sprintf("%s'tlp': '%d'\n", ws, o.Tlp))
	str.WriteString(fmt.Sprintf("%s'title': '%s'\n", ws, o.Title))
	str.WriteString(fmt.Sprintf("%s'description': '%s'\n", ws, o.Description))
	str.WriteString(fmt.Sprintf("%s'status': '%s'\n", ws, o.Status))
	str.WriteString(fmt.Sprintf("%s'date': '%s'\n", ws, o.Date))
	str.WriteString(fmt.Sprintf("%s'type': '%s'\n", ws, o.Type))
	str.WriteString(fmt.Sprintf("%s'objectType': '%s'\n", ws, o.ObjectType))
	str.WriteString(fmt.Sprintf("%s'source': '%s'\n", ws, o.Source))
	str.WriteString(fmt.Sprintf("%s'sourceRef': '%s'\n", ws, o.SourceRef))
	str.WriteString(fmt.Sprintf("%s'case': '%s'\n", ws, o.Case))
	str.WriteString(fmt.Sprintf("%s'caseTemplate': '%s'\n", ws, o.CaseTemplate))

	return str.String()
}
