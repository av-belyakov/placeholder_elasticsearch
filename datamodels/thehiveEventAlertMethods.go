package datamodels

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/supportingfunctions"
)

// Get возвращает объект типа EventMessageTheHiveCase
func (e *EventMessageTheHiveAlert) Get() *EventMessageTheHiveAlert {
	return e
}

func (e *EventMessageTheHiveAlert) GetBase() bool {
	return e.Base
}

// SetValueBase устанавливает BOOL значение для поля Base
func (e *EventMessageTheHiveAlert) SetValueBase(v bool) {
	e.Base = v
}

// SetAnyBase устанавливает ЛЮБОЕ значение для поля Base
func (e *EventMessageTheHiveAlert) SetAnyBase(i interface{}) {
	if v, ok := i.(bool); ok {
		e.Base = v
	}
}

func (e *EventMessageTheHiveAlert) GetStartDate() string {
	return e.StartDate
}

// SetValueStartDate устанавливает значение в формате RFC3339 для поля StartDate
func (e *EventMessageTheHiveAlert) SetValueStartDate(v string) {
	e.StartDate = v
}

// SetAnyStartDate устанавливает ЛЮБОЕ значение для поля StartDate
func (e *EventMessageTheHiveAlert) SetAnyStartDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.StartDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *EventMessageTheHiveAlert) GetRootId() string {
	return e.RootId
}

// SetValueRootId устанавливает STRING значение для поля RootId
func (e *EventMessageTheHiveAlert) SetValueRootId(v string) {
	e.RootId = v
}

// SetAnyRootId устанавливает ЛЮБОЕ значение для поля RootId
func (e *EventMessageTheHiveAlert) SetAnyRootId(i interface{}) {
	e.RootId = fmt.Sprint(i)
}

func (e *EventMessageTheHiveAlert) GetOrganisation() string {
	return e.Organisation
}

// SetValueOrganisation устанавливает STRING значение для поля Organisation
func (e *EventMessageTheHiveAlert) SetValueOrganisation(v string) {
	e.Organisation = v
}

// SetAnyOrganisation устанавливает ЛЮБОЕ значение для поля Organisation
func (e *EventMessageTheHiveAlert) SetAnyOrganisation(i interface{}) {
	e.Organisation = fmt.Sprint(i)
}

func (e *EventMessageTheHiveAlert) GetOrganisationId() string {
	return e.OrganisationId
}

// SetValueOrganisationId устанавливает STRING значение для поля OrganisationId
func (e *EventMessageTheHiveAlert) SetValueOrganisationId(v string) {
	e.OrganisationId = v
}

// SetAnyOrganisationId устанавливает ЛЮБОЕ значение для поля OrganisationId
func (e *EventMessageTheHiveAlert) SetAnyOrganisationId(i interface{}) {
	e.OrganisationId = fmt.Sprint(i)
}

func (e *EventMessageTheHiveAlert) GetObjectId() string {
	return e.ObjectId
}

// SetValueObjectId устанавливает STRING значение для поля ObjectId
func (e *EventMessageTheHiveAlert) SetValueObjectId(v string) {
	e.ObjectId = v
}

// SetAnyObjectId устанавливает ЛЮБОЕ значение для поля ObjectId
func (e *EventMessageTheHiveAlert) SetAnyObjectId(i interface{}) {
	e.ObjectId = fmt.Sprint(i)
}

func (e *EventMessageTheHiveAlert) GetObjectType() string {
	return e.ObjectType
}

// SetValueObjectType устанавливает STRING значение для поля ObjectType
func (e *EventMessageTheHiveAlert) SetValueObjectType(v string) {
	e.ObjectType = v
}

// SetAnyObjectType устанавливает ЛЮБОЕ значение для поля ObjectType
func (e *EventMessageTheHiveAlert) SetAnyObjectType(i interface{}) {
	e.ObjectType = fmt.Sprint(i)
}

func (e *EventMessageTheHiveAlert) GetOperation() string {
	return e.Operation
}

// SetValueOperation устанавливает STRING значение для поля Operation
func (e *EventMessageTheHiveAlert) SetValueOperation(v string) {
	e.Operation = v
}

// SetAnyOperation устанавливает ЛЮБОЕ значение для поля Operation
func (e *EventMessageTheHiveAlert) SetAnyOperation(i interface{}) {
	e.Operation = fmt.Sprint(i)
}

func (e *EventMessageTheHiveAlert) GetRequestId() string {
	return e.RequestId
}

// SetValueRequestId устанавливает STRING значение для поля RequestId
func (e *EventMessageTheHiveAlert) SetValueRequestId(v string) {
	e.RequestId = v
}

// SetAnyRequestId устанавливает ЛЮБОЕ значение для поля RequestId
func (e *EventMessageTheHiveAlert) SetAnyRequestId(i interface{}) {
	e.RequestId = fmt.Sprint(i)
}

func (e *EventMessageTheHiveAlert) GetDetails() EventAlertDetails {
	return e.Details
}

// SetValueDetails устанавливает значение типа EventDetails для поля Details
func (e *EventMessageTheHiveAlert) SetValueDetails(v EventAlertDetails) {
	e.Details = v
}

func (e *EventMessageTheHiveAlert) GetObject() EventAlertObject {
	return e.Object
}

// SetValueObject устанавливает значение типа EventObject для поля Object
func (e *EventMessageTheHiveAlert) SetValueObject(v EventAlertObject) {
	e.Object = v
}

func (e *EventMessageTheHiveAlert) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%soperation: '%s'\n", ws, e.Operation))
	str.WriteString(fmt.Sprintf("%sobjectId: '%s'\n", ws, e.ObjectId))
	str.WriteString(fmt.Sprintf("%sobjectType: '%s'\n", ws, e.ObjectType))
	str.WriteString(fmt.Sprintf("%sbase: '%v'\n", ws, e.Base))
	str.WriteString(fmt.Sprintf("%sstartDate: '%s'\n", ws, e.StartDate))
	str.WriteString(fmt.Sprintf("%srootId: '%s'\n", ws, e.RootId))
	str.WriteString(fmt.Sprintf("%srequestId: '%s'\n", ws, e.RequestId))
	str.WriteString(fmt.Sprintf("%sorganisationId: '%s'\n", ws, e.OrganisationId))
	str.WriteString(fmt.Sprintf("%sorganisation: '%s'\n", ws, e.Organisation))
	str.WriteString(fmt.Sprintf("%sdetails:\n", ws))
	str.WriteString(e.Details.ToStringBeautiful(num + 1))
	str.WriteString(fmt.Sprintf("%sobject:\n", ws))
	str.WriteString(e.Object.ToStringBeautiful(num + 1))

	return str.String()
}

//****************** EventAlertDetails ******************

func (e *EventAlertDetails) Get() *EventAlertDetails {
	return e
}

func (e *EventAlertDetails) GetSourceRef() string {
	return e.SourceRef
}

// SetValueSourceRef устанавливает STRING значение для поля SourceRef
func (e *EventAlertDetails) SetValueSourceRef(v string) {
	e.SourceRef = v
}

// SetAnySourceRef устанавливает ЛЮБОЕ значение для поля SourceRef
func (e *EventAlertDetails) SetAnySourceRef(i interface{}) {
	e.SourceRef = fmt.Sprint(i)
}

func (e *EventAlertDetails) GetTitle() string {
	return e.Title
}

// SetValueTitle устанавливает STRING значение для поля Title
func (e *EventAlertDetails) SetValueTitle(v string) {
	e.Title = v
}

// SetAnyTitle устанавливает ЛЮБОЕ значение для поля Title
func (e *EventAlertDetails) SetAnyTitle(i interface{}) {
	e.Title = fmt.Sprint(i)
}

func (e *EventAlertDetails) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает STRING значение для поля Description
func (e *EventAlertDetails) SetValueDescription(v string) {
	e.Description = v
}

// SetAny устанавливает ЛЮБОЕ значение для поля Description
func (e *EventAlertDetails) SetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

func (e *EventAlertDetails) GetTags() []string {
	return e.Tags
}

// SetValueTags устанавливает STRING значение для поля Tags
func (e *EventAlertDetails) SetValueTags(v string) {
	e.Tags = append(e.Tags, v)
}

// SetAnyTags устанавливает ЛЮБОЕ значение для поля Tags
func (e *EventAlertDetails) SetAnyTags(i interface{}) {
	e.Tags = append(e.Tags, fmt.Sprint(i))
}

func (e *EventAlertDetails) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%ssourceRef: '%s'\n", ws, e.SourceRef))
	str.WriteString(fmt.Sprintf("%stitle: '%s'\n", ws, e.Title))
	str.WriteString(fmt.Sprintf("%sdescription: '%s'\n", ws, e.Description))
	str.WriteString(fmt.Sprintf("%stags: \n%s", ws, func(l []string) string {
		str := strings.Builder{}
		ws := supportingfunctions.GetWhitespace(num + 1)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d. '%s'\n", ws, k+1, v))
		}
		return str.String()
	}(e.Tags)))

	return str.String()
}

//****************** EventAlertObject ******************

func (e *EventAlertObject) Get() *EventAlertObject {
	return e
}

func (e *EventAlertObject) GetFollow() bool {
	return e.Follow
}

// SetValueFollow устанавливает BOOL значение для поля Follow
func (e *EventAlertObject) SetValueFollow(v bool) {
	e.Follow = v
}

// SetAnyFollow устанавливает ЛЮБОЕ значение для поля Follow
func (e *EventAlertObject) SetAnyFollow(i interface{}) {
	if v, ok := i.(bool); ok {
		e.Follow = v
	}
}

func (e *EventAlertObject) GetSeverity() uint64 {
	return e.Severity
}

// SetValueSeverity устанавливает INT значение для поля Severity
func (e *EventAlertObject) SetValueSeverity(v uint64) {
	e.Severity = v
}

// SetAnySeverity устанавливает ЛЮБОЕ значение для поля Severity
func (e *EventAlertObject) SetAnySeverity(i interface{}) {
	if v, ok := i.(float64); ok {
		e.Severity = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		e.Severity = v
	}
}

func (e *EventAlertObject) GetTlp() uint64 {
	return e.Tlp
}

// SetValueTlp устанавливает INT значение для поля Tlp
func (e *EventAlertObject) SetValueTlp(v uint64) {
	e.Tlp = v
}

// SetAnyTlp устанавливает ЛЮБОЕ значение для поля Tlp
func (e *EventAlertObject) SetAnyTlp(i interface{}) {
	if v, ok := i.(float64); ok {
		e.Tlp = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		e.Tlp = v
	}
}

func (e *EventAlertObject) GetPap() uint64 {
	return e.Pap
}

// SetValuePap устанавливает INT значение для поля Pap
func (e *EventAlertObject) SetValuePap(v uint64) {
	e.Pap = v
}

// SetAnyPap устанавливает ЛЮБОЕ значение для поля Pap
func (e *EventAlertObject) SetAnyPap(i interface{}) {
	if v, ok := i.(float64); ok {
		e.Pap = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		e.Pap = v
	}
}

func (e *EventAlertObject) GetUnderliningId() string {
	return e.UnderliningId
}

// SetValueUnderliningId устанавливает STRING значение для поля UnderliningId
func (e *EventAlertObject) SetValueUnderliningId(v string) {
	e.UnderliningId = v
}

// SetAnyUnderliningId устанавливает ЛЮБОЕ значение для поля UnderliningId
func (e *EventAlertObject) SetAnyUnderliningId(i interface{}) {
	e.UnderliningId = fmt.Sprint(i)
}

func (e *EventAlertObject) GetId() string {
	return e.Id
}

// SetValueId устанавливает STRING значение для поля Id
func (e *EventAlertObject) SetValueId(v string) {
	e.Id = v
}

// SetAnyId устанавливает ЛЮБОЕ значение для поля Id
func (e *EventAlertObject) SetAnyId(i interface{}) {
	e.Id = fmt.Sprint(i)
}

func (e *EventAlertObject) GetCreatedBy() string {
	return e.CreatedBy
}

// SetValueCreatedBy устанавливает STRING значение для поля CreatedBy
func (e *EventAlertObject) SetValueCreatedBy(v string) {
	e.CreatedBy = v
}

// SetAnyCreatedBy устанавливает ЛЮБОЕ значение для поля CreatedBy
func (e *EventAlertObject) SetAnyCreatedBy(i interface{}) {
	e.CreatedBy = fmt.Sprint(i)
}

func (e *EventAlertObject) GetUpdatedBy() string {
	return e.UpdatedBy
}

// SetValueUpdatedBy устанавливает STRING значение для поля UpdatedBy
func (e *EventAlertObject) SetValueUpdatedBy(v string) {
	e.UpdatedBy = v
}

// SetAnyUpdatedBy устанавливает ЛЮБОЕ значение для поля UpdatedBy
func (e *EventAlertObject) SetAnyUpdatedBy(i interface{}) {
	e.UpdatedBy = fmt.Sprint(i)
}

func (e *EventAlertObject) GetCreatedAt() string {
	return e.CreatedAt
}

// SetValueCreatedAt устанавливает значение в формате RFC3339 для поля CreatedAt
func (e *EventAlertObject) SetValueCreatedAt(v string) {
	e.CreatedAt = v
}

// SetAnyCreatedAt устанавливает ЛЮБОЕ значение для поля CreatedAt
func (e *EventAlertObject) SetAnyCreatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.CreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *EventAlertObject) GetUpdatedAt() string {
	return e.UpdatedAt
}

// SetValueUpdatedAt устанавливает значение  в формате RFC3339 для поля UpdatedAt
func (e *EventAlertObject) SetValueUpdatedAt(v string) {
	e.UpdatedAt = v
}

// SetAnyUpdatedAt устанавливает ЛЮБОЕ значение для поля UpdatedAt
func (e *EventAlertObject) SetAnyUpdatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.UpdatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *EventAlertObject) GetUnderliningType() string {
	return e.UnderliningType
}

// SetValueUnderliningType устанавливает STRING значение для поля UnderliningType
func (e *EventAlertObject) SetValueUnderliningType(v string) {
	e.UnderliningType = v
}

// SetAnyUnderliningType устанавливает ЛЮБОЕ значение для поля UnderliningType
func (e *EventAlertObject) SetAnyUnderliningType(i interface{}) {
	e.UnderliningType = fmt.Sprint(i)
}

func (e *EventAlertObject) GetTitle() string {
	return e.Title
}

// SetValueTitle устанавливает STRING значение для поля Title
func (e *EventAlertObject) SetValueTitle(v string) {
	e.Title = v
}

// SetAnyTitle устанавливает ЛЮБОЕ значение для поля Title
func (e *EventAlertObject) SetAnyTitle(i interface{}) {
	e.Title = fmt.Sprint(i)
}

func (e *EventAlertObject) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает STRING значение для поля Description
func (e *EventAlertObject) SetValueDescription(v string) {
	e.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (e *EventAlertObject) SetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

func (e *EventAlertObject) GetStatus() string {
	return e.Status
}

// SetValueStatus устанавливает STRING значение для поля Status
func (e *EventAlertObject) SetValueStatus(v string) {
	e.Status = v
}

// SetAnyStatus устанавливает ЛЮБОЕ значение для поля Status
func (e *EventAlertObject) SetAnyStatus(i interface{}) {
	e.Status = fmt.Sprint(i)
}

func (e *EventAlertObject) GetDate() string {
	return e.Date
}

// SetValueDate устанавливает значение в формате RFC3339 для поля Date
func (e *EventAlertObject) SetValueDate(v string) {
	e.Date = v
}

// SetAnyDate устанавливает ЛЮБОЕ значение для поля Date
func (e *EventAlertObject) SetAnyDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.Date = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *EventAlertObject) GetType() string {
	return e.Type
}

// SetValueType устанавливает STRING значение для поля Type
func (e *EventAlertObject) SetValueType(v string) {
	e.Type = v
}

// SetAnyType устанавливает ЛЮБОЕ значение для поля Type
func (e *EventAlertObject) SetAnyType(i interface{}) {
	e.Type = fmt.Sprint(i)
}

func (e *EventAlertObject) GetObjectType() string {
	return e.ObjectType
}

// SetValueObjectType устанавливает STRING значение для поля ObjectType
func (e *EventAlertObject) SetValueObjectType(v string) {
	e.ObjectType = v
}

// SetAnyObjectType устанавливает ЛЮБОЕ значение для поля ObjectType
func (e *EventAlertObject) SetAnyObjectType(i interface{}) {
	e.ObjectType = fmt.Sprint(i)
}

func (e *EventAlertObject) GetSource() string {
	return e.Source
}

// SetValueSource устанавливает STRING значение для поля Source
func (e *EventAlertObject) SetValueSource(v string) {
	e.Source = v
}

// SetAnySource устанавливает ЛЮБОЕ значение для поля Source
func (e *EventAlertObject) SetAnySource(i interface{}) {
	e.Source = fmt.Sprint(i)
}

func (e *EventAlertObject) GetSourceRef() string {
	return e.SourceRef
}

// SetValueSourceRef устанавливает STRING значение для поля SourceRef
func (e *EventAlertObject) SetValueSourceRef(v string) {
	e.SourceRef = v
}

// SetAnySourceRef устанавливает ЛЮБОЕ значение для поля SourceRef
func (e *EventAlertObject) SetAnySourceRef(i interface{}) {
	e.SourceRef = fmt.Sprint(i)
}

func (e *EventAlertObject) GetCase() string {
	return e.Case
}

// SetValueCase устанавливает STRING значение для поля Case
func (e *EventAlertObject) SetValueCase(v string) {
	e.Case = v
}

// SetAnyCase устанавливает ЛЮБОЕ значение для поля Case
func (e *EventAlertObject) SetAnyCase(i interface{}) {
	e.Case = fmt.Sprint(i)
}

func (e *EventAlertObject) GetCaseTemplate() string {
	return e.CaseTemplate
}

// SetValueCaseTemplate устанавливает STRING значение для поля CaseTemplate
func (e *EventAlertObject) SetValueCaseTemplate(v string) {
	e.CaseTemplate = v
}

// SetAnyCaseTemplate устанавливает ЛЮБОЕ значение для поля CaseTemplate
func (e *EventAlertObject) SetAnyCaseTemplate(i interface{}) {
	e.CaseTemplate = fmt.Sprint(i)
}

func (e *EventAlertObject) GetTags() []string {
	return e.Tags
}

// SetValueTags устанавливает STRING значение для поля Tags
func (e *EventAlertObject) SetValueTags(v string) {
	e.Tags = append(e.Tags, v)
}

// SetAnyTags устанавливает ЛЮБОЕ значение для поля Tags
func (e *EventAlertObject) SetAnyTags(i interface{}) {
	e.Tags = append(e.Tags, fmt.Sprint(i))
}

func (e *EventAlertObject) GetCustomFields() map[string]CustomerFields {
	return e.CustomFields
}

// SetValueCustomFields устанавливает STRING значение для поля CustomFields
func (e *EventAlertObject) SetValueCustomFields(v map[string]CustomerFields) {
	e.CustomFields = v
}

func (e *EventAlertObject) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s_id: '%s'\n", ws, e.UnderliningId))
	str.WriteString(fmt.Sprintf("%sid: '%s'\n", ws, e.Id))
	str.WriteString(fmt.Sprintf("%screatedBy: '%s'\n", ws, e.CreatedBy))
	str.WriteString(fmt.Sprintf("%supdatedBy: '%s'\n", ws, e.UpdatedBy))
	str.WriteString(fmt.Sprintf("%screatedAt: '%s'\n", ws, e.CreatedAt))
	str.WriteString(fmt.Sprintf("%supdatedAt: '%s'\n", ws, e.UpdatedAt))
	str.WriteString(fmt.Sprintf("%s_type: '%s'\n", ws, e.UnderliningType))
	str.WriteString(fmt.Sprintf("%sfolow: '%t'\n", ws, e.Follow))
	str.WriteString(fmt.Sprintf("%sseverity: '%d'\n", ws, e.Severity))
	str.WriteString(fmt.Sprintf("%stlp: '%d'\n", ws, e.Tlp))
	str.WriteString(fmt.Sprintf("%spap: '%d'\n", ws, e.Pap))
	str.WriteString(fmt.Sprintf("%stitle: '%s'\n", ws, e.Title))
	str.WriteString(fmt.Sprintf("%sdescription: '%s'\n", ws, e.Description))
	str.WriteString(fmt.Sprintf("%sstatus: '%s'\n", ws, e.Status))
	str.WriteString(fmt.Sprintf("%sdate: '%s'\n", ws, e.Date))
	str.WriteString(fmt.Sprintf("%stype: '%s'\n", ws, e.Type))
	str.WriteString(fmt.Sprintf("%sobjectType '%s'\n", ws, e.ObjectType))
	str.WriteString(fmt.Sprintf("%ssource: '%s'\n", ws, e.Source))
	str.WriteString(fmt.Sprintf("%ssourceRef: '%s'\n", ws, e.SourceRef))
	str.WriteString(fmt.Sprintf("%scase: '%s'\n", ws, e.Case))
	str.WriteString(fmt.Sprintf("%scaseTemplate: '%s'\n", ws, e.CaseTemplate))
	str.WriteString(fmt.Sprintf("%stags: \n%s", ws, func(l []string) string {
		str := strings.Builder{}
		ws := supportingfunctions.GetWhitespace(num + 1)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d. '%s'\n", ws, k+1, v))
		}
		return str.String()
	}(e.Tags)))
	str.WriteString(fmt.Sprintf("%scustomFields: \n%s", ws, CustomFieldsToStringBeautiful(e.CustomFields, num)))

	return str.String()
}
