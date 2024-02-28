package commonevent

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/supportingfunctions"
)

//****************** CommonEventCase ******************

func (e *CommonEventType) Get() *CommonEventType {
	return e
}

func (e *CommonEventType) GetBase() bool {
	return e.Base
}

// SetValueBase устанавливает BOOL значение для поля Base
func (e *CommonEventType) SetValueBase(v bool) {
	e.Base = v
}

// SetAnyBase устанавливает ЛЮБОЕ значение для поля Base
func (e *CommonEventType) SetAnyBase(i interface{}) {
	if v, ok := i.(bool); ok {
		e.Base = v
	}
}

func (e *CommonEventType) GetStartDate() string {
	return e.StartDate
}

// SetValueStartDate устанавливает значение в формате RFC3339 для поля StartDate
func (e *CommonEventType) SetValueStartDate(v string) {
	e.StartDate = v
}

// SetAnyStartDate устанавливает ЛЮБОЕ значение для поля StartDate
func (e *CommonEventType) SetAnyStartDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.StartDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *CommonEventType) GetRootId() string {
	return e.RootId
}

// SetValueRootId устанавливает STRING значение для поля RootId
func (e *CommonEventType) SetValueRootId(v string) {
	e.RootId = v
}

// SetAnyRootId устанавливает ЛЮБОЕ значение для поля RootId
func (e *CommonEventType) SetAnyRootId(i interface{}) {
	e.RootId = fmt.Sprint(i)
}

func (e *CommonEventType) GetOrganisation() string {
	return e.Organisation
}

// SetValueOrganisation устанавливает STRING значение для поля Organisation
func (e *CommonEventType) SetValueOrganisation(v string) {
	e.Organisation = v
}

// SetAnyOrganisation устанавливает ЛЮБОЕ значение для поля Organisation
func (e *CommonEventType) SetAnyOrganisation(i interface{}) {
	e.Organisation = fmt.Sprint(i)
}

func (e *CommonEventType) GetOrganisationId() string {
	return e.OrganisationId
}

// SetValueOrganisationId устанавливает STRING значение для поля OrganisationId
func (e *CommonEventType) SetValueOrganisationId(v string) {
	e.OrganisationId = v
}

// SetAnyOrganisationId устанавливает ЛЮБОЕ значение для поля OrganisationId
func (e *CommonEventType) SetAnyOrganisationId(i interface{}) {
	e.OrganisationId = fmt.Sprint(i)
}

func (e *CommonEventType) GetObjectId() string {
	return e.ObjectId
}

// SetValueObjectId устанавливает STRING значение для поля ObjectId
func (e *CommonEventType) SetValueObjectId(v string) {
	e.ObjectId = v
}

// SetAnyObjectId устанавливает ЛЮБОЕ значение для поля ObjectId
func (e *CommonEventType) SetAnyObjectId(i interface{}) {
	e.ObjectId = fmt.Sprint(i)
}

func (e *CommonEventType) GetObjectType() string {
	return e.ObjectType
}

// SetValueObjectType устанавливает STRING значение для поля ObjectType
func (e *CommonEventType) SetValueObjectType(v string) {
	e.ObjectType = v
}

// SetAnyObjectType устанавливает ЛЮБОЕ значение для поля ObjectType
func (e *CommonEventType) SetAnyObjectType(i interface{}) {
	e.ObjectType = fmt.Sprint(i)
}

func (e *CommonEventType) GetOperation() string {
	return e.Operation
}

// SetValueOperation устанавливает STRING значение для поля Operation
func (e *CommonEventType) SetValueOperation(v string) {
	e.Operation = v
}

// SetAnyOperation устанавливает ЛЮБОЕ значение для поля Operation
func (e *CommonEventType) SetAnyOperation(i interface{}) {
	e.Operation = fmt.Sprint(i)
}

func (e *CommonEventType) GetRequestId() string {
	return e.RequestId
}

// SetValueRequestId устанавливает STRING значение для поля RequestId
func (e *CommonEventType) SetValueRequestId(v string) {
	e.RequestId = v
}

// SetAnyRequestId устанавливает ЛЮБОЕ значение для поля RequestId
func (e *CommonEventType) SetAnyRequestId(i interface{}) {
	e.RequestId = fmt.Sprint(i)
}

func (em CommonEventType) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'operation': '%s'\n", ws, em.Operation))
	str.WriteString(fmt.Sprintf("%s'objectId': '%s'\n", ws, em.ObjectId))
	str.WriteString(fmt.Sprintf("%s'objectType': '%s'\n", ws, em.ObjectType))
	str.WriteString(fmt.Sprintf("%s'base': '%v'\n", ws, em.Base))
	str.WriteString(fmt.Sprintf("%s'startDate': '%s'\n", ws, em.StartDate))
	str.WriteString(fmt.Sprintf("%s'rootId': '%s'\n", ws, em.RootId))
	str.WriteString(fmt.Sprintf("%s'requestId': '%s'\n", ws, em.RequestId))
	str.WriteString(fmt.Sprintf("%s'organisationId': '%s'\n", ws, em.OrganisationId))
	str.WriteString(fmt.Sprintf("%s'organisation': '%s'\n", ws, em.Organisation))

	return str.String()
}
