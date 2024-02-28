package datamodels

import (
	"fmt"
	"strings"

	commonevent "placeholder_elasticsearch/datamodels/commonevent"
	commonobjectevent "placeholder_elasticsearch/datamodels/commonobjectevent"
	"placeholder_elasticsearch/supportingfunctions"
)

func NewEventMessageTheHiveCase() *EventMessageTheHiveCase {
	return &EventMessageTheHiveCase{
		CommonEventType: commonevent.CommonEventType{
			StartDate: "1970-01-01T00:00:00+00:00",
		},
		Details: *NewEventCaseDetails(),
		Object:  *NewEventCaseObject(),
	}
}

// Get возвращает объект типа EventMessageTheHiveCase
func (e *EventMessageTheHiveCase) Get() *EventMessageTheHiveCase {
	return e
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
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(em.CommonEventType.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'details':\n", ws))
	str.WriteString(em.Details.ToStringBeautiful(num + 1))
	str.WriteString(fmt.Sprintf("%s'object':\n", ws))
	str.WriteString(em.Object.ToStringBeautiful(num + 1))

	return str.String()
}

//****************** EventCaseDetails ******************

func NewEventCaseDetails() *EventCaseDetails {
	return &EventCaseDetails{
		EndDate:      "1970-01-01T00:00:00+00:00",
		CustomFields: CustomFields{},
	}
}

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

func (e *EventCaseDetails) GetCustomFields() CustomFields {
	return e.CustomFields
}

// SetValueCustomFields устанавливает STRING значение для поля CustomFields
func (e *EventCaseDetails) SetValueCustomFields(v CustomFields) {
	e.CustomFields = v
}

func (ed EventCaseDetails) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'endDate': '%s'\n", ws, ed.EndDate))
	str.WriteString(fmt.Sprintf("%s'resolutionStatus': '%s'\n", ws, ed.ResolutionStatus))
	str.WriteString(fmt.Sprintf("%s'summary': '%s'\n", ws, ed.Summary))
	str.WriteString(fmt.Sprintf("%s'status': '%s'\n", ws, ed.Status))
	str.WriteString(fmt.Sprintf("%s'impactStatus': '%s'\n", ws, ed.ImpactStatus))
	str.WriteString(fmt.Sprintf("%s'customFields': \n%s", ws, CustomFieldsToStringBeautiful(ed.CustomFields, num)))

	return str.String()
}

//****************** EventCaseObject ******************

func NewEventCaseObject() *EventCaseObject {
	return &EventCaseObject{
		CommonEventCaseObject: commonobjectevent.CommonEventCaseObject{
			StartDate: "1970-01-01T00:00:00+00:00",
			EndDate:   "1970-01-01T00:00:00+00:00",
			CreatedAt: "1970-01-01T00:00:00+00:00",
			UpdatedAt: "1970-01-01T00:00:00+00:00",
		},
		Tags:         []string(nil),
		CustomFields: CustomFields{},
	}
}

func (e *EventCaseObject) Get() *EventCaseObject {
	return e
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

func (e *EventCaseObject) GetCustomFields() CustomFields {
	return e.CustomFields
}

// SetValueCustomFields устанавливает STRING значение для поля CustomFields
func (e *EventCaseObject) SetValueCustomFields(v CustomFields) {
	e.CustomFields = v
}

func (eo EventCaseObject) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(eo.CommonEventCaseObject.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, ToStringBeautifulSlice(num, eo.Tags)))
	str.WriteString(fmt.Sprintf("%s'customFields': \n%s", ws, CustomFieldsToStringBeautiful(eo.CustomFields, num)))

	return str.String()
}
