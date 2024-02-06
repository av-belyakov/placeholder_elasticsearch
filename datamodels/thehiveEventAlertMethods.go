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

func (em EventMessageTheHiveAlert) ToStringBeautiful(num int) string {
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

func (ed EventAlertDetails) ToStringBeautiful(num int) string {
	strB := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	strB.WriteString(fmt.Sprintf("%ssourceRef: '%s'\n", ws, ed.SourceRef))
	strB.WriteString(fmt.Sprintf("%stitle: '%s'\n", ws, ed.Title))
	strB.WriteString(fmt.Sprintf("%sdescription: '%s'\n", ws, ed.Description))
	strB.WriteString(fmt.Sprintf("%stags: \n%s", ws, func(l []string) string {
		str := strings.Builder{}
		ws := supportingfunctions.GetWhitespace(num + 1)

		for k, v := range l {
			strB.WriteString(fmt.Sprintf("%s%d. '%s'\n", ws, k+1, v))
		}
		return str.String()
	}(ed.Tags)))

	return strB.String()
}
