package commonalert

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/supportingfunctions"
)

//****************** CommonAlertType ******************

func (e *CommonAlertType) Get() *CommonAlertType {
	return e
}

func (a *CommonAlertType) GetTlp() uint64 {
	return a.Tlp
}

// SetValueTlp устанавливает UINT64 значение для поля Tlp
func (a *CommonAlertType) SetValueTlp(v uint64) {
	a.Tlp = v
}

// SetAnyTlp устанавливает ЛЮБОЕ значение для поля Tlp
func (a *CommonAlertType) SetAnyTlp(i interface{}) {
	if v, ok := i.(float64); ok {
		a.Tlp = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		a.Tlp = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		a.Tlp = v
	}
}

func (a *CommonAlertType) GetDate() string {
	return a.Date
}

// SetValueDate устанавливает значение в формате RFC3339 для поля Date
func (a *CommonAlertType) SetValueDate(v string) {
	a.Date = v
}

// SetAnyDate устанавливает ЛЮБОЕ значение для поля Date
func (a *CommonAlertType) SetAnyDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	a.Date = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (a *CommonAlertType) GetCreatedAt() string {
	return a.CreatedAt
}

// SetValueCreatedAt устанавливает значение в формате RFC3339 для поля CreatedAt
func (a *CommonAlertType) SetValueCreatedAt(v string) {
	a.CreatedAt = v
}

// SetAnyCreatedAt устанавливает ЛЮБОЕ значение для поля CreatedAt
func (a *CommonAlertType) SetAnyCreatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	a.CreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (a *CommonAlertType) GetUpdatedAt() string {
	return a.UpdatedAt
}

// SetValueUpdatedAt устанавливает значение в формате RFC3339 для поля UpdatedAt
func (a *CommonAlertType) SetValueUpdatedAt(v string) {
	a.UpdatedAt = v
}

// SetAnyUpdatedAt устанавливает ЛЮБОЕ значение для поля UpdatedAt
func (a *CommonAlertType) SetAnyUpdatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	a.UpdatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (a *CommonAlertType) GetUpdatedBy() string {
	return a.UpdatedBy
}

// SetValueUpdatedBy устанавливает STRING значение для поля UpdatedBy
func (a *CommonAlertType) SetValueUpdatedBy(v string) {
	a.UpdatedBy = v
}

// SetAnyUpdatedBy устанавливает ЛЮБОЕ значение для поля UpdatedBy
func (a *CommonAlertType) SetAnyUpdatedBy(i interface{}) {
	a.UpdatedBy = fmt.Sprint(i)
}

func (a *CommonAlertType) GetUnderliningId() string {
	return a.UnderliningId
}

// SetValueUnderliningId устанавливает STRING значение для поля UnderliningId
func (a *CommonAlertType) SetValueUnderliningId(v string) {
	a.UnderliningId = v
}

// SetAnyUnderliningId устанавливает ЛЮБОЕ значение для поля UnderliningId
func (a *CommonAlertType) SetAnyUnderliningId(i interface{}) {
	a.UnderliningId = fmt.Sprint(i)
}

func (a *CommonAlertType) GetStatus() string {
	return a.Status
}

// SetValueStatus устанавливает STRING значение для поля Status
func (a *CommonAlertType) SetValueStatus(v string) {
	a.Status = v
}

// SetAnyStatus устанавливает ЛЮБОЕ значение для поля Status
func (a *CommonAlertType) SetAnyStatus(i interface{}) {
	a.Status = fmt.Sprint(i)
}

func (a *CommonAlertType) GetType() string {
	return a.Type
}

// SetValueType устанавливает STRING значение для поля Type
func (a *CommonAlertType) SetValueType(v string) {
	a.Type = v
}

// SetAnyType устанавливает ЛЮБОЕ значение для поля Type
func (a *CommonAlertType) SetAnyType(i interface{}) {
	a.Type = fmt.Sprint(i)
}

func (a *CommonAlertType) GetUnderliningType() string {
	return a.UnderliningType
}

// SetValueUnderliningType устанавливает STRING значение для поля UnderliningType
func (a *CommonAlertType) SetValueUnderliningType(v string) {
	a.UnderliningType = v
}

// SetAnyUnderliningType устанавливает ЛЮБОЕ значение для поля UnderliningType
func (a *CommonAlertType) SetAnyUnderliningType(i interface{}) {
	a.UnderliningType = fmt.Sprint(i)
}

func (a *CommonAlertType) GetDescription() string {
	return a.Description
}

// SetValueDescription устанавливает STRING значение для поля Description
func (a *CommonAlertType) SetValueDescription(v string) {
	a.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (a *CommonAlertType) SetAnyDescription(i interface{}) {
	str := fmt.Sprint(i)
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\n", "")

	a.Description = str
}

func (a *CommonAlertType) GetCaseTemplate() string {
	return a.CaseTemplate
}

// SetValueCaseTemplate устанавливает STRING значение для поля CaseTemplate
func (a *CommonAlertType) SetValueCaseTemplate(v string) {
	a.CaseTemplate = v
}

// SetAnyCaseTemplate устанавливает ЛЮБОЕ значение для поля CaseTemplate
func (a *CommonAlertType) SetAnyCaseTemplate(i interface{}) {
	a.CaseTemplate = fmt.Sprint(i)
}

func (a *CommonAlertType) GetSourceRef() string {
	return a.SourceRef
}

// SetValueSourceRef устанавливает STRING значение для поля SourceRef
func (a *CommonAlertType) SetValueSourceRef(v string) {
	a.SourceRef = v
}

// SetAnySourceRef устанавливает ЛЮБОЕ значение для поля SourceRef
func (a *CommonAlertType) SetAnySourceRef(i interface{}) {
	a.SourceRef = fmt.Sprint(i)
}

func (a *CommonAlertType) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'tlp': '%d'\n", ws, a.Tlp))
	str.WriteString(fmt.Sprintf("%s'date': '%s'\n", ws, a.Date))
	str.WriteString(fmt.Sprintf("%s'createdAt': '%s'\n", ws, a.CreatedAt))
	str.WriteString(fmt.Sprintf("%s'updatedAt': '%s'\n", ws, a.UpdatedAt))
	str.WriteString(fmt.Sprintf("%s'updatedBy': '%s'\n", ws, a.UpdatedBy))
	str.WriteString(fmt.Sprintf("%s'underliningId': '%s'\n", ws, a.UnderliningId))
	str.WriteString(fmt.Sprintf("%s'status': '%s'\n", ws, a.Status))
	str.WriteString(fmt.Sprintf("%s'type': '%s'\n", ws, a.Type))
	str.WriteString(fmt.Sprintf("%s'underliningType': '%s'\n", ws, a.UnderliningType))
	str.WriteString(fmt.Sprintf("%s'description': '%s'\n", ws, a.Description))
	str.WriteString(fmt.Sprintf("%s'caseTemplate': '%s'\n", ws, a.CaseTemplate))
	str.WriteString(fmt.Sprintf("%s'sourceRef': '%s'\n", ws, a.SourceRef))

	return str.String()
}
