package datamodels

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/supportingfunctions"
)

// Get возвращает объект типа EventMessageTheHiveCase
func (e *EventMessageTheHiveCase) Get() *EventMessageTheHiveCase {
	return e
}

func (e *EventMessageTheHiveCase) GetBase() bool {
	return e.Base
}

// SetValueBase устанавливает BOOL значение для поля Base
func (e *EventMessageTheHiveCase) SetValueBase(v bool) {
	e.Base = v
}

// SetAnyBase устанавливает ЛЮБОЕ значение для поля Base
func (e *EventMessageTheHiveCase) SetAnyBase(i interface{}) {
	if v, ok := i.(bool); ok {
		e.Base = v
	}
}

func (e *EventMessageTheHiveCase) GetStartDate() string {
	return e.StartDate
}

// SetValueStartDate устанавливает значение в формате RFC3339 для поля StartDate
func (e *EventMessageTheHiveCase) SetValueStartDate(v string) {
	e.StartDate = v
}

// SetAnyStartDate устанавливает ЛЮБОЕ значение для поля StartDate
func (e *EventMessageTheHiveCase) SetAnyStartDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.StartDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *EventMessageTheHiveCase) GetRootId() string {
	return e.RootId
}

// SetValueRootId устанавливает STRING значение для поля RootId
func (e *EventMessageTheHiveCase) SetValueRootId(v string) {
	e.RootId = v
}

// SetAnyRootId устанавливает ЛЮБОЕ значение для поля RootId
func (e *EventMessageTheHiveCase) SetAnyRootId(i interface{}) {
	e.RootId = fmt.Sprint(i)
}

func (e *EventMessageTheHiveCase) GetOrganisation() string {
	return e.Organisation
}

// SetValueOrganisation устанавливает STRING значение для поля Organisation
func (e *EventMessageTheHiveCase) SetValueOrganisation(v string) {
	e.Organisation = v
}

// SetAnyOrganisation устанавливает ЛЮБОЕ значение для поля Organisation
func (e *EventMessageTheHiveCase) SetAnyOrganisation(i interface{}) {
	e.Organisation = fmt.Sprint(i)
}

func (e *EventMessageTheHiveCase) GetOrganisationId() string {
	return e.OrganisationId
}

// SetValueOrganisationId устанавливает STRING значение для поля OrganisationId
func (e *EventMessageTheHiveCase) SetValueOrganisationId(v string) {
	e.OrganisationId = v
}

// SetAnyOrganisationId устанавливает ЛЮБОЕ значение для поля OrganisationId
func (e *EventMessageTheHiveCase) SetAnyOrganisationId(i interface{}) {
	e.OrganisationId = fmt.Sprint(i)
}

func (e *EventMessageTheHiveCase) GetObjectId() string {
	return e.ObjectId
}

// SetValueObjectId устанавливает STRING значение для поля ObjectId
func (e *EventMessageTheHiveCase) SetValueObjectId(v string) {
	e.ObjectId = v
}

// SetAnyObjectId устанавливает ЛЮБОЕ значение для поля ObjectId
func (e *EventMessageTheHiveCase) SetAnyObjectId(i interface{}) {
	e.ObjectId = fmt.Sprint(i)
}

func (e *EventMessageTheHiveCase) GetObjectType() string {
	return e.ObjectType
}

// SetValueObjectType устанавливает STRING значение для поля ObjectType
func (e *EventMessageTheHiveCase) SetValueObjectType(v string) {
	e.ObjectType = v
}

// SetAnyObjectType устанавливает ЛЮБОЕ значение для поля ObjectType
func (e *EventMessageTheHiveCase) SetAnyObjectType(i interface{}) {
	e.ObjectType = fmt.Sprint(i)
}

func (e *EventMessageTheHiveCase) GetOperation() string {
	return e.Operation
}

// SetValueOperation устанавливает STRING значение для поля Operation
func (e *EventMessageTheHiveCase) SetValueOperation(v string) {
	e.Operation = v
}

// SetAnyOperation устанавливает ЛЮБОЕ значение для поля Operation
func (e *EventMessageTheHiveCase) SetAnyOperation(i interface{}) {
	e.Operation = fmt.Sprint(i)
}

func (e *EventMessageTheHiveCase) GetRequestId() string {
	return e.RequestId
}

// SetValueRequestId устанавливает STRING значение для поля RequestId
func (e *EventMessageTheHiveCase) SetValueRequestId(v string) {
	e.RequestId = v
}

// SetAnyRequestId устанавливает ЛЮБОЕ значение для поля RequestId
func (e *EventMessageTheHiveCase) SetAnyRequestId(i interface{}) {
	e.RequestId = fmt.Sprint(i)
}

func (e *EventMessageTheHiveCase) GetDetails() EventCaseDetails {
	return e.Details
}

// SetValueDetails устанавливает значение типа EventDetails для поля Details
func (e *EventMessageTheHiveCase) SetValueDetails(v EventCaseDetails) {
	e.Details = v
}

func (e *EventMessageTheHiveCase) GetObject() EventCaseObject {
	return e.Object
}

// SetValueObject устанавливает значение типа EventObject для поля Object
func (e *EventMessageTheHiveCase) SetValueObject(v EventCaseObject) {
	e.Object = v
}

func (em EventMessageTheHiveCase) ToStringBeautiful(num int) string {
	strB := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	strB.WriteString(fmt.Sprintf("%soperation: '%s'\n", ws, em.Operation))
	strB.WriteString(fmt.Sprintf("%sobjectId: '%s'\n", ws, em.ObjectId))
	strB.WriteString(fmt.Sprintf("%sobjectType: '%s'\n", ws, em.ObjectType))
	strB.WriteString(fmt.Sprintf("%sbase: '%v'\n", ws, em.Base))
	strB.WriteString(fmt.Sprintf("%sstartDate: '%s'\n", ws, em.StartDate))
	strB.WriteString(fmt.Sprintf("%srootId: '%s'\n", ws, em.RootId))
	strB.WriteString(fmt.Sprintf("%srequestId: '%s'\n", ws, em.RequestId))
	strB.WriteString(fmt.Sprintf("%sorganisationId: '%s'\n", ws, em.OrganisationId))
	strB.WriteString(fmt.Sprintf("%sorganisation: '%s'\n", ws, em.Organisation))
	strB.WriteString(fmt.Sprintf("%sdetails:\n", ws))
	strB.WriteString(em.Details.ToStringBeautiful(num + 1))
	strB.WriteString(fmt.Sprintf("%sobject:\n", ws))
	strB.WriteString(em.Object.ToStringBeautiful(num + 1))

	return strB.String()
}

//****************** EventCaseDetails ******************

func (e *EventCaseDetails) Get() *EventCaseDetails {
	return e
}

func (e *EventCaseDetails) GetEndDate() string {
	return e.EndDate
}

// SetValueEndDate устанавливает значение в формате RFC3339 для поля EndDate
func (e *EventCaseDetails) SetValueEndDate(v string) {
	e.EndDate = v
}

// SetAnyEndDate устанавливает ЛЮБОЕ значение для поля EndDate
func (e *EventCaseDetails) SetAnyEndDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.EndDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *EventCaseDetails) GetResolutionStatus() string {
	return e.ResolutionStatus
}

// SetValueResolutionStatus устанавливает STRING значение для поля ResolutionStatus
func (e *EventCaseDetails) SetValueResolutionStatus(v string) {
	e.ResolutionStatus = v
}

// SetAnyResolutionStatus устанавливает ЛЮБОЕ значение для поля ResolutionStatus
func (e *EventCaseDetails) SetAnyResolutionStatus(i interface{}) {
	e.ResolutionStatus = fmt.Sprint(i)
}

func (e *EventCaseDetails) GetSummary() string {
	return e.Summary
}

// SetValueSummary устанавливает STRING значение для поля Summary
func (e *EventCaseDetails) SetValueSummary(v string) {
	e.Summary = v
}

// SetAnySummary устанавливает ЛЮБОЕ значение для поля Summary
func (e *EventCaseDetails) SetAnySummary(i interface{}) {
	e.Summary = fmt.Sprint(i)
}

func (e *EventCaseDetails) GetStatus() string {
	return e.Status
}

// SetValueStatus устанавливает STRING значение для поля Status
func (e *EventCaseDetails) SetValueStatus(v string) {
	e.Status = v
}

// SetAnyStatus устанавливает ЛЮБОЕ значение для поля Status
func (e *EventCaseDetails) SetAnyStatus(i interface{}) {
	e.Status = fmt.Sprint(i)
}

func (e *EventCaseDetails) GetImpactStatus() string {
	return e.ImpactStatus
}

// SetValueImpactStatus устанавливает STRING значение для поля ImpactStatus
func (e *EventCaseDetails) SetValueImpactStatus(v string) {
	e.ImpactStatus = v
}

// SetAnyImpactStatus устанавливает ЛЮБОЕ значение для поля ImpactStatus
func (e *EventCaseDetails) SetAnyImpactStatus(i interface{}) {
	e.ImpactStatus = fmt.Sprint(i)
}

func (e *EventCaseDetails) GetCustomFields() map[string]CustomerFields {
	return e.CustomFields
}

// SetValueCustomFields устанавливает STRING значение для поля CustomFields
func (e *EventCaseDetails) SetValueCustomFields(v map[string]CustomerFields) {
	e.CustomFields = v
}

func (ed EventCaseDetails) ToStringBeautiful(num int) string {
	strB := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	strB.WriteString(fmt.Sprintf("%sendDate: '%s'\n", ws, ed.EndDate))
	strB.WriteString(fmt.Sprintf("%sresolutionStatus: '%s'\n", ws, ed.ResolutionStatus))
	strB.WriteString(fmt.Sprintf("%ssummary: '%s'\n", ws, ed.Summary))
	strB.WriteString(fmt.Sprintf("%sstatus: '%s'\n", ws, ed.Status))
	strB.WriteString(fmt.Sprintf("%simpactStatus: '%s'\n", ws, ed.ImpactStatus))
	strB.WriteString(fmt.Sprintf("%scustomFields: \n%s", ws, CustomFieldsToStringBeautiful(ed.CustomFields, num)))

	return strB.String()
}

//****************** EventCaseObject ******************

func (e *EventCaseObject) Get() *EventCaseObject {
	return e
}

func (e *EventCaseObject) GetFlag() bool {
	return e.Flag
}

// SetValueFlag устанавливает BOOL значение для поля Flag
func (e *EventCaseObject) SetValueFlag(v bool) {
	e.Flag = v
}

// SetAnyFlag устанавливает ЛЮБОЕ значение для поля Flag
func (e *EventCaseObject) SetAnyFlag(i interface{}) {
	if v, ok := i.(bool); ok {
		e.Flag = v
	}
}

func (e *EventCaseObject) GetCaseId() uint64 {
	return e.CaseId
}

// SetValueCaseId устанавливает INT значение для поля CaseId
func (e *EventCaseObject) SetValueCaseId(v uint64) {
	e.CaseId = v
}

// SetAnyCaseId устанавливает ЛЮБОЕ значение для поля CaseId
func (e *EventCaseObject) SetAnyCaseId(i interface{}) {
	if v, ok := i.(float64); ok {
		e.CaseId = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		e.CaseId = v
	}
}

func (e *EventCaseObject) GetSeverity() uint64 {
	return e.Severity
}

// SetValueSeverity устанавливает INT значение для поля Severity
func (e *EventCaseObject) SetValueSeverity(v uint64) {
	e.Severity = v
}

// SetAnySeverity устанавливает ЛЮБОЕ значение для поля Severity
func (e *EventCaseObject) SetAnySeverity(i interface{}) {
	if v, ok := i.(float64); ok {
		e.Severity = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		e.Severity = v
	}
}

func (e *EventCaseObject) GetTlp() uint64 {
	return e.Tlp
}

// SetValueTlp устанавливает INT значение для поля Tlp
func (e *EventCaseObject) SetValueTlp(v uint64) {
	e.Tlp = v
}

// SetAnyTlp устанавливает ЛЮБОЕ значение для поля Tlp
func (e *EventCaseObject) SetAnyTlp(i interface{}) {
	if v, ok := i.(float64); ok {
		e.Tlp = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		e.Tlp = v
	}
}

func (e *EventCaseObject) GetPap() uint64 {
	return e.Pap
}

// SetValuePap устанавливает INT значение для поля Pap
func (e *EventCaseObject) SetValuePap(v uint64) {
	e.Pap = v
}

// SetAnyPap устанавливает ЛЮБОЕ значение для поля Pap
func (e *EventCaseObject) SetAnyPap(i interface{}) {
	if v, ok := i.(float64); ok {
		e.Pap = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		e.Pap = v
	}
}

func (e *EventCaseObject) GetStartDate() string {
	return e.StartDate
}

// SetValueStartDate устанавливает значение в формате RFC3339 для поля StartDate
func (e *EventCaseObject) SetValueStartDate(v string) {
	e.StartDate = v
}

// SetAnyStartDate устанавливает ЛЮБОЕ значение для поля StartDate
func (e *EventCaseObject) SetAnyStartDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.StartDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *EventCaseObject) GetEndDate() string {
	return e.EndDate
}

// SetValueEndDate устанавливает значение в формате RFC3339 для поля EndDate
func (e *EventCaseObject) SetValueEndDate(v string) {
	e.EndDate = v
}

// SetAnyEndDate устанавливает ЛЮБОЕ значение для поля EndDate
func (e *EventCaseObject) SetAnyEndDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.EndDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *EventCaseObject) GetCreatedAt() string {
	return e.CreatedAt
}

// SetValueCreatedAt устанавливает значение в формате RFC3339 для поля CreatedAt
func (e *EventCaseObject) SetValueCreatedAt(v string) {
	e.CreatedAt = v
}

// SetAnyCreatedAt устанавливает ЛЮБОЕ значение для поля CreatedAt
func (e *EventCaseObject) SetAnyCreatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.CreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *EventCaseObject) GetUpdatedAt() string {
	return e.UpdatedAt
}

// SetValueUpdatedAt устанавливает значение  в формате RFC3339 для поля UpdatedAt
func (e *EventCaseObject) SetValueUpdatedAt(v string) {
	e.UpdatedAt = v
}

// SetAnyUpdatedAt устанавливает ЛЮБОЕ значение для поля UpdatedAt
func (e *EventCaseObject) SetAnyUpdatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.UpdatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *EventCaseObject) GetUnderliningId() string {
	return e.UnderliningId
}

// SetValueUnderliningId устанавливает STRING значение для поля UnderliningId
func (e *EventCaseObject) SetValueUnderliningId(v string) {
	e.UnderliningId = v
}

// SetAnyUnderliningId устанавливает ЛЮБОЕ значение для поля UnderliningId
func (e *EventCaseObject) SetAnyUnderliningId(i interface{}) {
	e.UnderliningId = fmt.Sprint(i)
}

func (e *EventCaseObject) GetId() string {
	return e.Id
}

// SetValueId устанавливает STRING значение для поля Id
func (e *EventCaseObject) SetValueId(v string) {
	e.Id = v
}

// SetAnyId устанавливает ЛЮБОЕ значение для поля Id
func (e *EventCaseObject) SetAnyId(i interface{}) {
	e.Id = fmt.Sprint(i)
}

func (e *EventCaseObject) GetCreatedBy() string {
	return e.CreatedBy
}

// SetValueCreatedBy устанавливает STRING значение для поля CreatedBy
func (e *EventCaseObject) SetValueCreatedBy(v string) {
	e.CreatedBy = v
}

// SetAnyCreatedBy устанавливает ЛЮБОЕ значение для поля CreatedBy
func (e *EventCaseObject) SetAnyCreatedBy(i interface{}) {
	e.CreatedBy = fmt.Sprint(i)
}

func (e *EventCaseObject) GetUpdatedBy() string {
	return e.UpdatedBy
}

// SetValueUpdatedBy устанавливает STRING значение для поля UpdatedBy
func (e *EventCaseObject) SetValueUpdatedBy(v string) {
	e.UpdatedBy = v
}

// SetAnyUpdatedBy устанавливает ЛЮБОЕ значение для поля UpdatedBy
func (e *EventCaseObject) SetAnyUpdatedBy(i interface{}) {
	e.UpdatedBy = fmt.Sprint(i)
}

func (e *EventCaseObject) GetUnderliningType() string {
	return e.UnderliningType
}

// SetValueUnderliningType устанавливает STRING значение для поля UnderliningType
func (e *EventCaseObject) SetValueUnderliningType(v string) {
	e.UnderliningType = v
}

// SetAnyUnderliningType устанавливает ЛЮБОЕ значение для поля UnderliningType
func (e *EventCaseObject) SetAnyUnderliningType(i interface{}) {
	e.UnderliningType = fmt.Sprint(i)
}

func (e *EventCaseObject) GetTitle() string {
	return e.Title
}

// SetValueTitle устанавливает STRING значение для поля Title
func (e *EventCaseObject) SetValueTitle(v string) {
	e.Title = v
}

// SetAnyTitle устанавливает ЛЮБОЕ значение для поля Title
func (e *EventCaseObject) SetAnyTitle(i interface{}) {
	e.Title = fmt.Sprint(i)
}

func (e *EventCaseObject) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает STRING значение для поля Description
func (e *EventCaseObject) SetValueDescription(v string) {
	e.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (e *EventCaseObject) SetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

func (e *EventCaseObject) GetImpactStatus() string {
	return e.ImpactStatus
}

// SetValueImpactStatus устанавливает STRING значение для поля ImpactStatus
func (e *EventCaseObject) SetValueImpactStatus(v string) {
	e.ImpactStatus = v
}

// SetAnyImpactStatus устанавливает ЛЮБОЕ значение для поля ImpactStatus
func (e *EventCaseObject) SetAnyImpactStatus(i interface{}) {
	e.ImpactStatus = fmt.Sprint(i)
}

func (e *EventCaseObject) GetResolutionStatus() string {
	return e.ResolutionStatus
}

// SetValueResolutionStatus устанавливает STRING значение для поля ResolutionStatus
func (e *EventCaseObject) SetValueResolutionStatus(v string) {
	e.ResolutionStatus = v
}

// SetAnyResolutionStatus устанавливает ЛЮБОЕ значение для поля ResolutionStatus
func (e *EventCaseObject) SetAnyResolutionStatus(i interface{}) {
	e.ResolutionStatus = fmt.Sprint(i)
}

func (e *EventCaseObject) GetStatus() string {
	return e.Status
}

// SetValueStatus устанавливает STRING значение для поля Status
func (e *EventCaseObject) SetValueStatus(v string) {
	e.Status = v
}

// SetAnyStatus устанавливает ЛЮБОЕ значение для поля Status
func (e *EventCaseObject) SetAnyStatus(i interface{}) {
	e.Status = fmt.Sprint(i)
}

func (e *EventCaseObject) GetSummary() string {
	return e.Summary
}

// SetValueSummary устанавливает STRING значение для поля Summary
func (e *EventCaseObject) SetValueSummary(v string) {
	e.Summary = v
}

// SetAnySummary устанавливает ЛЮБОЕ значение для поля Summary
func (e *EventCaseObject) SetAnySummary(i interface{}) {
	e.Summary = fmt.Sprint(i)
}

func (e *EventCaseObject) GetOwner() string {
	return e.Owner
}

// SetValueOwner устанавливает STRING значение для поля Owner
func (e *EventCaseObject) SetValueOwner(v string) {
	e.Owner = v
}

// SetAnyOwner устанавливает ЛЮБОЕ значение для поля Owner
func (e *EventCaseObject) SetAnyOwner(i interface{}) {
	e.Owner = fmt.Sprint(i)
}

func (e *EventCaseObject) GetTags() []string {
	return e.Tags
}

// SetValueTags устанавливает STRING значение для поля Tags
func (e *EventCaseObject) SetValueTags(v string) {
	e.Tags = append(e.Tags, v)
}

// SetAnyTags устанавливает ЛЮБОЕ значение для поля Tags
func (e *EventCaseObject) SetAnyTags(i interface{}) {
	e.Tags = append(e.Tags, fmt.Sprint(i))
}

func (e *EventCaseObject) GetCustomFields() map[string]CustomerFields {
	return e.CustomFields
}

// SetValueCustomFields устанавливает STRING значение для поля CustomFields
func (e *EventCaseObject) SetValueCustomFields(v map[string]CustomerFields) {
	e.CustomFields = v
}

func (eo EventCaseObject) ToStringBeautiful(num int) string {
	strB := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	strB.WriteString(fmt.Sprintf("%s_id: '%s'\n", ws, eo.UnderliningId))
	strB.WriteString(fmt.Sprintf("%sid: '%s'\n", ws, eo.Id))
	strB.WriteString(fmt.Sprintf("%screatedBy: '%s'\n", ws, eo.CreatedBy))
	strB.WriteString(fmt.Sprintf("%supdatedBy: '%s'\n", ws, eo.UpdatedBy))
	strB.WriteString(fmt.Sprintf("%screatedAt: '%s'\n", ws, eo.CreatedAt))
	strB.WriteString(fmt.Sprintf("%supdatedAt: '%s'\n", ws, eo.UpdatedAt))
	strB.WriteString(fmt.Sprintf("%s_type: '%s'\n", ws, eo.UnderliningType))
	strB.WriteString(fmt.Sprintf("%scaseId: '%d'\n", ws, eo.CaseId))
	strB.WriteString(fmt.Sprintf("%stitle: '%s'\n", ws, eo.Title))
	strB.WriteString(fmt.Sprintf("%sdescription: '%s'\n", ws, eo.Description))
	strB.WriteString(fmt.Sprintf("%sseverity: '%d'\n", ws, eo.Severity))
	strB.WriteString(fmt.Sprintf("%sstartDate: '%s'\n", ws, eo.StartDate))
	strB.WriteString(fmt.Sprintf("%sendDate: '%s'\n", ws, eo.EndDate))
	strB.WriteString(fmt.Sprintf("%simpactStatus: '%s'\n", ws, eo.ImpactStatus))
	strB.WriteString(fmt.Sprintf("%sresolutionStatus: '%s'\n", ws, eo.ResolutionStatus))
	strB.WriteString(fmt.Sprintf("%stags: \n%s", ws, func(l []string) string {
		str := strings.Builder{}
		ws := supportingfunctions.GetWhitespace(num + 1)

		for k, v := range l {
			strB.WriteString(fmt.Sprintf("%s%d. '%s'\n", ws, k+1, v))
		}
		return str.String()
	}(eo.Tags)))
	strB.WriteString(fmt.Sprintf("%sflag: '%v'\n", ws, eo.Flag))
	strB.WriteString(fmt.Sprintf("%stlp: '%d'\n", ws, eo.Tlp))
	strB.WriteString(fmt.Sprintf("%spap: '%d'\n", ws, eo.Pap))
	strB.WriteString(fmt.Sprintf("%sstatus: '%s'\n", ws, eo.Status))
	strB.WriteString(fmt.Sprintf("%ssummary: '%s'\n", ws, eo.Summary))
	strB.WriteString(fmt.Sprintf("%sowner: '%s'\n", ws, eo.Owner))
	strB.WriteString(fmt.Sprintf("%scustomFields: \n%s", ws, CustomFieldsToStringBeautiful(eo.CustomFields, num)))
	/*strB.WriteString(fmt.Sprintf("%spermissions: \n%s", ws, func(l []string) string {
		str := strings.Builder{}
		ws := supportingfunctions.GetWhitespace(num + 1)

		for k, v := range l {
			strB.WriteString(fmt.Sprintf("%s%d. '%s'\n", ws, k+1, v))
		}
		return str.String()
	}(eo.Permissions)))*/

	return strB.String()
}

func CustomFieldsToStringBeautiful(l map[string]CustomerFields, num int) string {
	strB := strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num + 2)

	for k, v := range l {
		strB.WriteString(fmt.Sprintf("%s%s:\n", supportingfunctions.GetWhitespace(num+1), k))

		nameOne, dataOne, nameTwo, dataTwo := v.Get()
		strB.WriteString(fmt.Sprintf("%s%s: %d\n", ws, nameOne, dataOne))
		strB.WriteString(fmt.Sprintf("%s%s: %s\n", ws, nameTwo, dataTwo))
	}
	return strB.String()
}
