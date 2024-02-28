package commonobjectevent

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/supportingfunctions"
)

//****************** CommonEventCaseObject ******************

func (e *CommonEventCaseObject) Get() *CommonEventCaseObject {
	return e
}

func (e *CommonEventCaseObject) GetFlag() bool {
	return e.Flag
}

// SetValueFlag устанавливает BOOL значение для поля Flag
func (e *CommonEventCaseObject) SetValueFlag(v bool) {
	e.Flag = v
}

// SetAnyFlag устанавливает ЛЮБОЕ значение для поля Flag
func (e *CommonEventCaseObject) SetAnyFlag(i interface{}) {
	if v, ok := i.(bool); ok {
		e.Flag = v
	}
}

func (e *CommonEventCaseObject) GetCaseId() uint64 {
	return e.CaseId
}

// SetValueCaseId устанавливает INT значение для поля CaseId
func (e *CommonEventCaseObject) SetValueCaseId(v uint64) {
	e.CaseId = v
}

// SetAnyCaseId устанавливает ЛЮБОЕ значение для поля CaseId
func (e *CommonEventCaseObject) SetAnyCaseId(i interface{}) {
	if v, ok := i.(float32); ok {
		e.CaseId = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		e.CaseId = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		e.CaseId = v
	}
}

func (e *CommonEventCaseObject) GetSeverity() uint64 {
	return e.Severity
}

// SetValueSeverity устанавливает INT значение для поля Severity
func (e *CommonEventCaseObject) SetValueSeverity(v uint64) {
	e.Severity = v
}

// SetAnySeverity устанавливает ЛЮБОЕ значение для поля Severity
func (e *CommonEventCaseObject) SetAnySeverity(i interface{}) {
	if v, ok := i.(float32); ok {
		e.Severity = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		e.Severity = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		e.Severity = v
	}
}

func (e *CommonEventCaseObject) GetTlp() uint64 {
	return e.Tlp
}

// SetValueTlp устанавливает INT значение для поля Tlp
func (e *CommonEventCaseObject) SetValueTlp(v uint64) {
	e.Tlp = v
}

// SetAnyTlp устанавливает ЛЮБОЕ значение для поля Tlp
func (e *CommonEventCaseObject) SetAnyTlp(i interface{}) {
	if v, ok := i.(float32); ok {
		e.Tlp = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		e.Tlp = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		e.Tlp = v
	}
}

func (e *CommonEventCaseObject) GetPap() uint64 {
	return e.Pap
}

// SetValuePap устанавливает INT значение для поля Pap
func (e *CommonEventCaseObject) SetValuePap(v uint64) {
	e.Pap = v
}

// SetAnyPap устанавливает ЛЮБОЕ значение для поля Pap
func (e *CommonEventCaseObject) SetAnyPap(i interface{}) {
	if v, ok := i.(float32); ok {
		e.Pap = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		e.Pap = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		e.Pap = v
	}
}

func (e *CommonEventCaseObject) GetStartDate() string {
	return e.StartDate
}

// SetValueStartDate устанавливает значение в формате RFC3339 для поля StartDate
func (e *CommonEventCaseObject) SetValueStartDate(v string) {
	e.StartDate = v
}

// SetAnyStartDate устанавливает ЛЮБОЕ значение для поля StartDate
func (e *CommonEventCaseObject) SetAnyStartDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.StartDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *CommonEventCaseObject) GetEndDate() string {
	return e.EndDate
}

// SetValueEndDate устанавливает значение в формате RFC3339 для поля EndDate
func (e *CommonEventCaseObject) SetValueEndDate(v string) {
	e.EndDate = v
}

// SetAnyEndDate устанавливает ЛЮБОЕ значение для поля EndDate
func (e *CommonEventCaseObject) SetAnyEndDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.EndDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *CommonEventCaseObject) GetCreatedAt() string {
	return e.CreatedAt
}

// SetValueCreatedAt устанавливает значение в формате RFC3339 для поля CreatedAt
func (e *CommonEventCaseObject) SetValueCreatedAt(v string) {
	e.CreatedAt = v
}

// SetAnyCreatedAt устанавливает ЛЮБОЕ значение для поля CreatedAt
func (e *CommonEventCaseObject) SetAnyCreatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.CreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *CommonEventCaseObject) GetUpdatedAt() string {
	return e.UpdatedAt
}

// SetValueUpdatedAt устанавливает значение  в формате RFC3339 для поля UpdatedAt
func (e *CommonEventCaseObject) SetValueUpdatedAt(v string) {
	e.UpdatedAt = v
}

// SetAnyUpdatedAt устанавливает ЛЮБОЕ значение для поля UpdatedAt
func (e *CommonEventCaseObject) SetAnyUpdatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.UpdatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *CommonEventCaseObject) GetUnderliningId() string {
	return e.UnderliningId
}

// SetValueUnderliningId устанавливает STRING значение для поля UnderliningId
func (e *CommonEventCaseObject) SetValueUnderliningId(v string) {
	e.UnderliningId = v
}

// SetAnyUnderliningId устанавливает ЛЮБОЕ значение для поля UnderliningId
func (e *CommonEventCaseObject) SetAnyUnderliningId(i interface{}) {
	e.UnderliningId = fmt.Sprint(i)
}

func (e *CommonEventCaseObject) GetId() string {
	return e.Id
}

// SetValueId устанавливает STRING значение для поля Id
func (e *CommonEventCaseObject) SetValueId(v string) {
	e.Id = v
}

// SetAnyId устанавливает ЛЮБОЕ значение для поля Id
func (e *CommonEventCaseObject) SetAnyId(i interface{}) {
	e.Id = fmt.Sprint(i)
}

func (e *CommonEventCaseObject) GetCreatedBy() string {
	return e.CreatedBy
}

// SetValueCreatedBy устанавливает STRING значение для поля CreatedBy
func (e *CommonEventCaseObject) SetValueCreatedBy(v string) {
	e.CreatedBy = v
}

// SetAnyCreatedBy устанавливает ЛЮБОЕ значение для поля CreatedBy
func (e *CommonEventCaseObject) SetAnyCreatedBy(i interface{}) {
	e.CreatedBy = fmt.Sprint(i)
}

func (e *CommonEventCaseObject) GetUpdatedBy() string {
	return e.UpdatedBy
}

// SetValueUpdatedBy устанавливает STRING значение для поля UpdatedBy
func (e *CommonEventCaseObject) SetValueUpdatedBy(v string) {
	e.UpdatedBy = v
}

// SetAnyUpdatedBy устанавливает ЛЮБОЕ значение для поля UpdatedBy
func (e *CommonEventCaseObject) SetAnyUpdatedBy(i interface{}) {
	e.UpdatedBy = fmt.Sprint(i)
}

func (e *CommonEventCaseObject) GetUnderliningType() string {
	return e.UnderliningType
}

// SetValueUnderliningType устанавливает STRING значение для поля UnderliningType
func (e *CommonEventCaseObject) SetValueUnderliningType(v string) {
	e.UnderliningType = v
}

// SetAnyUnderliningType устанавливает ЛЮБОЕ значение для поля UnderliningType
func (e *CommonEventCaseObject) SetAnyUnderliningType(i interface{}) {
	e.UnderliningType = fmt.Sprint(i)
}

func (e *CommonEventCaseObject) GetTitle() string {
	return e.Title
}

// SetValueTitle устанавливает STRING значение для поля Title
func (e *CommonEventCaseObject) SetValueTitle(v string) {
	e.Title = v
}

// SetAnyTitle устанавливает ЛЮБОЕ значение для поля Title
func (e *CommonEventCaseObject) SetAnyTitle(i interface{}) {
	e.Title = fmt.Sprint(i)
}

func (e *CommonEventCaseObject) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает STRING значение для поля Description
func (e *CommonEventCaseObject) SetValueDescription(v string) {
	e.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (e *CommonEventCaseObject) SetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

func (e *CommonEventCaseObject) GetImpactStatus() string {
	return e.ImpactStatus
}

// SetValueImpactStatus устанавливает STRING значение для поля ImpactStatus
func (e *CommonEventCaseObject) SetValueImpactStatus(v string) {
	e.ImpactStatus = v
}

// SetAnyImpactStatus устанавливает ЛЮБОЕ значение для поля ImpactStatus
func (e *CommonEventCaseObject) SetAnyImpactStatus(i interface{}) {
	e.ImpactStatus = fmt.Sprint(i)
}

func (e *CommonEventCaseObject) GetResolutionStatus() string {
	return e.ResolutionStatus
}

// SetValueResolutionStatus устанавливает STRING значение для поля ResolutionStatus
func (e *CommonEventCaseObject) SetValueResolutionStatus(v string) {
	e.ResolutionStatus = v
}

// SetAnyResolutionStatus устанавливает ЛЮБОЕ значение для поля ResolutionStatus
func (e *CommonEventCaseObject) SetAnyResolutionStatus(i interface{}) {
	e.ResolutionStatus = fmt.Sprint(i)
}

func (e *CommonEventCaseObject) GetStatus() string {
	return e.Status
}

// SetValueStatus устанавливает STRING значение для поля Status
func (e *CommonEventCaseObject) SetValueStatus(v string) {
	e.Status = v
}

// SetAnyStatus устанавливает ЛЮБОЕ значение для поля Status
func (e *CommonEventCaseObject) SetAnyStatus(i interface{}) {
	e.Status = fmt.Sprint(i)
}

func (e *CommonEventCaseObject) GetSummary() string {
	return e.Summary
}

// SetValueSummary устанавливает STRING значение для поля Summary
func (e *CommonEventCaseObject) SetValueSummary(v string) {
	e.Summary = v
}

// SetAnySummary устанавливает ЛЮБОЕ значение для поля Summary
func (e *CommonEventCaseObject) SetAnySummary(i interface{}) {
	e.Summary = fmt.Sprint(i)
}

func (e *CommonEventCaseObject) GetOwner() string {
	return e.Owner
}

// SetValueOwner устанавливает STRING значение для поля Owner
func (e *CommonEventCaseObject) SetValueOwner(v string) {
	e.Owner = v
}

// SetAnyOwner устанавливает ЛЮБОЕ значение для поля Owner
func (e *CommonEventCaseObject) SetAnyOwner(i interface{}) {
	e.Owner = fmt.Sprint(i)
}

func (eo CommonEventCaseObject) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'_id': '%s'\n", ws, eo.UnderliningId))
	str.WriteString(fmt.Sprintf("%s'id': '%s'\n", ws, eo.Id))
	str.WriteString(fmt.Sprintf("%s'createdBy': '%s'\n", ws, eo.CreatedBy))
	str.WriteString(fmt.Sprintf("%s'updatedBy': '%s'\n", ws, eo.UpdatedBy))
	str.WriteString(fmt.Sprintf("%s'createdAt': '%s'\n", ws, eo.CreatedAt))
	str.WriteString(fmt.Sprintf("%s'updatedAt': '%s'\n", ws, eo.UpdatedAt))
	str.WriteString(fmt.Sprintf("%s'_type': '%s'\n", ws, eo.UnderliningType))
	str.WriteString(fmt.Sprintf("%s'caseId': '%d'\n", ws, eo.CaseId))
	str.WriteString(fmt.Sprintf("%s'title': '%s'\n", ws, eo.Title))
	str.WriteString(fmt.Sprintf("%s'description': '%s'\n", ws, eo.Description))
	str.WriteString(fmt.Sprintf("%s'severity': '%d'\n", ws, eo.Severity))
	str.WriteString(fmt.Sprintf("%s'startDate': '%s'\n", ws, eo.StartDate))
	str.WriteString(fmt.Sprintf("%s'endDate': '%s'\n", ws, eo.EndDate))
	str.WriteString(fmt.Sprintf("%s'impactStatus': '%s'\n", ws, eo.ImpactStatus))
	str.WriteString(fmt.Sprintf("%s'resolutionStatus': '%s'\n", ws, eo.ResolutionStatus))
	str.WriteString(fmt.Sprintf("%s'flag': '%v'\n", ws, eo.Flag))
	str.WriteString(fmt.Sprintf("%s'tlp': '%d'\n", ws, eo.Tlp))
	str.WriteString(fmt.Sprintf("%s'pap': '%d'\n", ws, eo.Pap))
	str.WriteString(fmt.Sprintf("%s'status': '%s'\n", ws, eo.Status))
	str.WriteString(fmt.Sprintf("%s'summary': '%s'\n", ws, eo.Summary))
	str.WriteString(fmt.Sprintf("%s'owner': '%s'\n", ws, eo.Owner))

	return str.String()
}
