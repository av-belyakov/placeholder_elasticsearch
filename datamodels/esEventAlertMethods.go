package datamodels

import (
	"fmt"
	"placeholder_elasticsearch/supportingfunctions"
	"strings"
)

func NewEventMessageForEsAlert() *EventMessageForEsAlert {
	return &EventMessageForEsAlert{
		StartDate: "1970-01-01T00:00:00+00:00",
		Details:   *NewEventMessageForEsAlertDetails(),
		Object:    *NewEventMessageForEsAlertObject(),
	}
}

// Get возвращает объект типа EventMessageForEsAlert
func (e *EventMessageForEsAlert) Get() *EventMessageForEsAlert {
	return e
}

func (e *EventMessageForEsAlert) GetBase() bool {
	return e.Base
}

// SetValueBase устанавливает BOOL значение для поля Base
func (e *EventMessageForEsAlert) SetValueBase(v bool) {
	e.Base = v
}

// SetAnyBase устанавливает ЛЮБОЕ значение для поля Base
func (e *EventMessageForEsAlert) SetAnyBase(i interface{}) {
	if v, ok := i.(bool); ok {
		e.Base = v
	}
}

func (e *EventMessageForEsAlert) GetStartDate() string {
	return e.StartDate
}

// SetValueStartDate устанавливает значение в формате RFC3339 для поля StartDate
func (e *EventMessageForEsAlert) SetValueStartDate(v string) {
	e.StartDate = v
}

// SetAnyStartDate устанавливает ЛЮБОЕ значение для поля StartDate
func (e *EventMessageForEsAlert) SetAnyStartDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.StartDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *EventMessageForEsAlert) GetRootId() string {
	return e.RootId
}

// SetValueRootId устанавливает STRING значение для поля RootId
func (e *EventMessageForEsAlert) SetValueRootId(v string) {
	e.RootId = v
}

// SetAnyRootId устанавливает ЛЮБОЕ значение для поля RootId
func (e *EventMessageForEsAlert) SetAnyRootId(i interface{}) {
	e.RootId = fmt.Sprint(i)
}

func (e *EventMessageForEsAlert) GetOrganisation() string {
	return e.Organisation
}

// SetValueOrganisation устанавливает STRING значение для поля Organisation
func (e *EventMessageForEsAlert) SetValueOrganisation(v string) {
	e.Organisation = v
}

// SetAnyOrganisation устанавливает ЛЮБОЕ значение для поля Organisation
func (e *EventMessageForEsAlert) SetAnyOrganisation(i interface{}) {
	e.Organisation = fmt.Sprint(i)
}

func (e *EventMessageForEsAlert) GetOrganisationId() string {
	return e.OrganisationId
}

// SetValueOrganisationId устанавливает STRING значение для поля OrganisationId
func (e *EventMessageForEsAlert) SetValueOrganisationId(v string) {
	e.OrganisationId = v
}

// SetAnyOrganisationId устанавливает ЛЮБОЕ значение для поля OrganisationId
func (e *EventMessageForEsAlert) SetAnyOrganisationId(i interface{}) {
	e.OrganisationId = fmt.Sprint(i)
}

func (e *EventMessageForEsAlert) GetObjectId() string {
	return e.ObjectId
}

// SetValueObjectId устанавливает STRING значение для поля ObjectId
func (e *EventMessageForEsAlert) SetValueObjectId(v string) {
	e.ObjectId = v
}

// SetAnyObjectId устанавливает ЛЮБОЕ значение для поля ObjectId
func (e *EventMessageForEsAlert) SetAnyObjectId(i interface{}) {
	e.ObjectId = fmt.Sprint(i)
}

func (e *EventMessageForEsAlert) GetObjectType() string {
	return e.ObjectType
}

// SetValueObjectType устанавливает STRING значение для поля ObjectType
func (e *EventMessageForEsAlert) SetValueObjectType(v string) {
	e.ObjectType = v
}

// SetAnyObjectType устанавливает ЛЮБОЕ значение для поля ObjectType
func (e *EventMessageForEsAlert) SetAnyObjectType(i interface{}) {
	e.ObjectType = fmt.Sprint(i)
}

func (e *EventMessageForEsAlert) GetOperation() string {
	return e.Operation
}

// SetValueOperation устанавливает STRING значение для поля Operation
func (e *EventMessageForEsAlert) SetValueOperation(v string) {
	e.Operation = v
}

// SetAnyOperation устанавливает ЛЮБОЕ значение для поля Operation
func (e *EventMessageForEsAlert) SetAnyOperation(i interface{}) {
	e.Operation = fmt.Sprint(i)
}

func (e *EventMessageForEsAlert) GetRequestId() string {
	return e.RequestId
}

// SetValueRequestId устанавливает STRING значение для поля RequestId
func (e *EventMessageForEsAlert) SetValueRequestId(v string) {
	e.RequestId = v
}

// SetAnyRequestId устанавливает ЛЮБОЕ значение для поля RequestId
func (e *EventMessageForEsAlert) SetAnyRequestId(i interface{}) {
	e.RequestId = fmt.Sprint(i)
}

func (e *EventMessageForEsAlert) GetDetails() EventMessageForEsAlertDetails {
	return e.Details
}

// SetValueDetails устанавливает значение типа EventMessageForEsAlertDetails для поля Details
func (e *EventMessageForEsAlert) SetValueDetails(v EventMessageForEsAlertDetails) {
	e.Details = v
}

func (e *EventMessageForEsAlert) GetObject() EventMessageForEsAlertObject {
	return e.Object
}

// SetValueObject устанавливает значение типа EventMessageForEsAlertObject для поля Object
func (e *EventMessageForEsAlert) SetValueObject(v EventMessageForEsAlertObject) {
	e.Object = v
}

func (e *EventMessageForEsAlert) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'operation': '%s'\n", ws, e.Operation))
	str.WriteString(fmt.Sprintf("%s'objectId': '%s'\n", ws, e.ObjectId))
	str.WriteString(fmt.Sprintf("%s'objectType': '%s'\n", ws, e.ObjectType))
	str.WriteString(fmt.Sprintf("%s'base': '%v'\n", ws, e.Base))
	str.WriteString(fmt.Sprintf("%s'startDate': '%s'\n", ws, e.StartDate))
	str.WriteString(fmt.Sprintf("%s'rootId': '%s'\n", ws, e.RootId))
	str.WriteString(fmt.Sprintf("%s'requestId': '%s'\n", ws, e.RequestId))
	str.WriteString(fmt.Sprintf("%s'organisationId': '%s'\n", ws, e.OrganisationId))
	str.WriteString(fmt.Sprintf("%s'organisation': '%s'\n", ws, e.Organisation))
	str.WriteString(fmt.Sprintf("%s'details':\n", ws))
	str.WriteString(e.Details.ToStringBeautiful(num + 1))
	str.WriteString(fmt.Sprintf("%s'object':\n", ws))
	str.WriteString(e.Object.ToStringBeautiful(num + 1))

	return str.String()
}

//****************** EventAlertDetails ******************

func NewEventMessageForEsAlertDetails() *EventMessageForEsAlertDetails {
	return &EventMessageForEsAlertDetails{
		Tags:    make(map[string][]string),
		TagsAll: []string(nil),
	}
}

func (e *EventMessageForEsAlertDetails) Get() *EventMessageForEsAlertDetails {
	return e
}

func (e *EventMessageForEsAlertDetails) GetSourceRef() string {
	return e.SourceRef
}

// SetValueSourceRef устанавливает STRING значение для поля SourceRef
func (e *EventMessageForEsAlertDetails) SetValueSourceRef(v string) {
	e.SourceRef = v
}

// SetAnySourceRef устанавливает ЛЮБОЕ значение для поля SourceRef
func (e *EventMessageForEsAlertDetails) SetAnySourceRef(i interface{}) {
	e.SourceRef = fmt.Sprint(i)
}

func (e *EventMessageForEsAlertDetails) GetTitle() string {
	return e.Title
}

// SetValueTitle устанавливает STRING значение для поля Title
func (e *EventMessageForEsAlertDetails) SetValueTitle(v string) {
	e.Title = v
}

// SetAnyTitle устанавливает ЛЮБОЕ значение для поля Title
func (e *EventMessageForEsAlertDetails) SetAnyTitle(i interface{}) {
	e.Title = fmt.Sprint(i)
}

func (e *EventMessageForEsAlertDetails) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает STRING значение для поля Description
func (e *EventMessageForEsAlertDetails) SetValueDescription(v string) {
	e.Description = v
}

// SetAny устанавливает ЛЮБОЕ значение для поля Description
func (e *EventMessageForEsAlertDetails) SetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

func (e *EventMessageForEsAlertDetails) GetTags() map[string][]string {
	return e.Tags
}

// SetValueTags добаляет значение в Tags по ключу
func (e *EventMessageForEsAlertDetails) SetValueTags(k, v string) bool {
	if _, ok := e.Tags[k]; !ok {
		e.Tags[k] = []string(nil)
	}

	for _, value := range e.Tags[k] {
		if v == value {
			return false
		}
	}

	e.Tags[k] = append(e.Tags[k], v)

	return true
}

// SetAnyTags устанавливает ЛЮБОЕ значение для поля Tags
func (e *EventMessageForEsAlertDetails) SetAnyTags(k string, i interface{}) bool {
	return e.SetValueTags(k, fmt.Sprint(i))
}

// SetSliceValueTags устанавливает ma[string][]STRING значение для поля Tags
func (e *EventMessageForEsAlertDetails) SetSliceValueTags(v map[string][]string) {
	e.Tags = v
}

func (a *EventMessageForEsAlertDetails) GetTagsAll() []string {
	return a.TagsAll
}

// SetValueTagsAll добавляет значение STRING в список поля TagsAll
func (a *EventMessageForEsAlertDetails) SetValueTagsAll(v string) {
	a.TagsAll = append(a.TagsAll, v)
}

// SetAnyTagsAll добавляет ЛЮБОЕ значение в список поля TagsAll
func (a *EventMessageForEsAlertDetails) SetAnyTagsAll(i interface{}) {
	a.TagsAll = append(a.TagsAll, fmt.Sprint(i))
}

func (e *EventMessageForEsAlertDetails) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'sourceRef': '%s'\n", ws, e.SourceRef))
	str.WriteString(fmt.Sprintf("%s'title': '%s'\n", ws, e.Title))
	str.WriteString(fmt.Sprintf("%s'description': '%s'\n", ws, e.Description))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, ToStringBeautifulMapSlice(num, e.Tags)))
	str.WriteString(fmt.Sprintf("%s'tagsAll': \n%s", ws, ToStringBeautifulSlice(num, e.TagsAll)))

	return str.String()
}

//****************** EventAlertObject ******************

func NewEventMessageForEsAlertObject() *EventMessageForEsAlertObject {
	return &EventMessageForEsAlertObject{
		CreatedAt:    "1970-01-01T00:00:00+00:00",
		UpdatedAt:    "1970-01-01T00:00:00+00:00",
		Tags:         make(map[string][]string),
		TagsAll:      []string(nil),
		CustomFields: CustomFields{},
	}
}

func (e *EventMessageForEsAlertObject) Get() *EventMessageForEsAlertObject {
	return e
}

func (e *EventMessageForEsAlertObject) GetTlp() uint64 {
	return e.Tlp
}

// SetValueTlp устанавливает INT значение для поля Tlp
func (e *EventMessageForEsAlertObject) SetValueTlp(v uint64) {
	e.Tlp = v
}

// SetAnyTlp устанавливает ЛЮБОЕ значение для поля Tlp
func (e *EventMessageForEsAlertObject) SetAnyTlp(i interface{}) {
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

func (e *EventMessageForEsAlertObject) GetUnderliningId() string {
	return e.UnderliningId
}

// SetValueUnderliningId устанавливает STRING значение для поля UnderliningId
func (e *EventMessageForEsAlertObject) SetValueUnderliningId(v string) {
	e.UnderliningId = v
}

// SetAnyUnderliningId устанавливает ЛЮБОЕ значение для поля UnderliningId
func (e *EventMessageForEsAlertObject) SetAnyUnderliningId(i interface{}) {
	e.UnderliningId = fmt.Sprint(i)
}

func (e *EventMessageForEsAlertObject) GetId() string {
	return e.Id
}

// SetValueId устанавливает STRING значение для поля Id
func (e *EventMessageForEsAlertObject) SetValueId(v string) {
	e.Id = v
}

// SetAnyId устанавливает ЛЮБОЕ значение для поля Id
func (e *EventMessageForEsAlertObject) SetAnyId(i interface{}) {
	e.Id = fmt.Sprint(i)
}

func (e *EventMessageForEsAlertObject) GetCreatedBy() string {
	return e.CreatedBy
}

// SetValueCreatedBy устанавливает STRING значение для поля CreatedBy
func (e *EventMessageForEsAlertObject) SetValueCreatedBy(v string) {
	e.CreatedBy = v
}

// SetAnyCreatedBy устанавливает ЛЮБОЕ значение для поля CreatedBy
func (e *EventMessageForEsAlertObject) SetAnyCreatedBy(i interface{}) {
	e.CreatedBy = fmt.Sprint(i)
}

func (e *EventMessageForEsAlertObject) GetUpdatedBy() string {
	return e.UpdatedBy
}

// SetValueUpdatedBy устанавливает STRING значение для поля UpdatedBy
func (e *EventMessageForEsAlertObject) SetValueUpdatedBy(v string) {
	e.UpdatedBy = v
}

// SetAnyUpdatedBy устанавливает ЛЮБОЕ значение для поля UpdatedBy
func (e *EventMessageForEsAlertObject) SetAnyUpdatedBy(i interface{}) {
	e.UpdatedBy = fmt.Sprint(i)
}

func (e *EventMessageForEsAlertObject) GetCreatedAt() string {
	return e.CreatedAt
}

// SetValueCreatedAt устанавливает значение в формате RFC3339 для поля CreatedAt
func (e *EventMessageForEsAlertObject) SetValueCreatedAt(v string) {
	e.CreatedAt = v
}

// SetAnyCreatedAt устанавливает ЛЮБОЕ значение для поля CreatedAt
func (e *EventMessageForEsAlertObject) SetAnyCreatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.CreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *EventMessageForEsAlertObject) GetUpdatedAt() string {
	return e.UpdatedAt
}

// SetValueUpdatedAt устанавливает значение  в формате RFC3339 для поля UpdatedAt
func (e *EventMessageForEsAlertObject) SetValueUpdatedAt(v string) {
	e.UpdatedAt = v
}

// SetAnyUpdatedAt устанавливает ЛЮБОЕ значение для поля UpdatedAt
func (e *EventMessageForEsAlertObject) SetAnyUpdatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.UpdatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *EventMessageForEsAlertObject) GetUnderliningType() string {
	return e.UnderliningType
}

// SetValueUnderliningType устанавливает STRING значение для поля UnderliningType
func (e *EventMessageForEsAlertObject) SetValueUnderliningType(v string) {
	e.UnderliningType = v
}

// SetAnyUnderliningType устанавливает ЛЮБОЕ значение для поля UnderliningType
func (e *EventMessageForEsAlertObject) SetAnyUnderliningType(i interface{}) {
	e.UnderliningType = fmt.Sprint(i)
}

func (e *EventMessageForEsAlertObject) GetTitle() string {
	return e.Title
}

// SetValueTitle устанавливает STRING значение для поля Title
func (e *EventMessageForEsAlertObject) SetValueTitle(v string) {
	e.Title = v
}

// SetAnyTitle устанавливает ЛЮБОЕ значение для поля Title
func (e *EventMessageForEsAlertObject) SetAnyTitle(i interface{}) {
	e.Title = fmt.Sprint(i)
}

func (e *EventMessageForEsAlertObject) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает STRING значение для поля Description
func (e *EventMessageForEsAlertObject) SetValueDescription(v string) {
	e.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (e *EventMessageForEsAlertObject) SetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

func (e *EventMessageForEsAlertObject) GetStatus() string {
	return e.Status
}

// SetValueStatus устанавливает STRING значение для поля Status
func (e *EventMessageForEsAlertObject) SetValueStatus(v string) {
	e.Status = v
}

// SetAnyStatus устанавливает ЛЮБОЕ значение для поля Status
func (e *EventMessageForEsAlertObject) SetAnyStatus(i interface{}) {
	e.Status = fmt.Sprint(i)
}

func (e *EventMessageForEsAlertObject) GetDate() string {
	return e.Date
}

// SetValueDate устанавливает значение в формате RFC3339 для поля Date
func (e *EventMessageForEsAlertObject) SetValueDate(v string) {
	e.Date = v
}

// SetAnyDate устанавливает ЛЮБОЕ значение для поля Date
func (e *EventMessageForEsAlertObject) SetAnyDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	e.Date = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *EventMessageForEsAlertObject) GetType() string {
	return e.Type
}

// SetValueType устанавливает STRING значение для поля Type
func (e *EventMessageForEsAlertObject) SetValueType(v string) {
	e.Type = v
}

// SetAnyType устанавливает ЛЮБОЕ значение для поля Type
func (e *EventMessageForEsAlertObject) SetAnyType(i interface{}) {
	e.Type = fmt.Sprint(i)
}

func (e *EventMessageForEsAlertObject) GetObjectType() string {
	return e.ObjectType
}

// SetValueObjectType устанавливает STRING значение для поля ObjectType
func (e *EventMessageForEsAlertObject) SetValueObjectType(v string) {
	e.ObjectType = v
}

// SetAnyObjectType устанавливает ЛЮБОЕ значение для поля ObjectType
func (e *EventMessageForEsAlertObject) SetAnyObjectType(i interface{}) {
	e.ObjectType = fmt.Sprint(i)
}

func (e *EventMessageForEsAlertObject) GetSource() string {
	return e.Source
}

// SetValueSource устанавливает STRING значение для поля Source
func (e *EventMessageForEsAlertObject) SetValueSource(v string) {
	e.Source = v
}

// SetAnySource устанавливает ЛЮБОЕ значение для поля Source
func (e *EventMessageForEsAlertObject) SetAnySource(i interface{}) {
	e.Source = fmt.Sprint(i)
}

func (e *EventMessageForEsAlertObject) GetSourceRef() string {
	return e.SourceRef
}

// SetValueSourceRef устанавливает STRING значение для поля SourceRef
func (e *EventMessageForEsAlertObject) SetValueSourceRef(v string) {
	e.SourceRef = v
}

// SetAnySourceRef устанавливает ЛЮБОЕ значение для поля SourceRef
func (e *EventMessageForEsAlertObject) SetAnySourceRef(i interface{}) {
	e.SourceRef = fmt.Sprint(i)
}

func (e *EventMessageForEsAlertObject) GetCase() string {
	return e.Case
}

// SetValueCase устанавливает STRING значение для поля Case
func (e *EventMessageForEsAlertObject) SetValueCase(v string) {
	e.Case = v
}

// SetAnyCase устанавливает ЛЮБОЕ значение для поля Case
func (e *EventMessageForEsAlertObject) SetAnyCase(i interface{}) {
	e.Case = fmt.Sprint(i)
}

func (e *EventMessageForEsAlertObject) GetCaseTemplate() string {
	return e.CaseTemplate
}

// SetValueCaseTemplate устанавливает STRING значение для поля CaseTemplate
func (e *EventMessageForEsAlertObject) SetValueCaseTemplate(v string) {
	e.CaseTemplate = v
}

// SetAnyCaseTemplate устанавливает ЛЮБОЕ значение для поля CaseTemplate
func (e *EventMessageForEsAlertObject) SetAnyCaseTemplate(i interface{}) {
	e.CaseTemplate = fmt.Sprint(i)
}

func (e *EventMessageForEsAlertObject) GetTags() map[string][]string {
	return e.Tags
}

// SetValueTags добаляет значение в Tags по ключу
func (e *EventMessageForEsAlertObject) SetValueTags(k, v string) bool {
	if _, ok := e.Tags[k]; !ok {
		e.Tags[k] = []string(nil)
	}

	for _, value := range e.Tags[k] {
		if v == value {
			return false
		}
	}

	e.Tags[k] = append(e.Tags[k], v)

	return true
}

// SetAnyTags устанавливает ЛЮБОЕ значение для поля Tags
func (e *EventMessageForEsAlertObject) SetAnyTags(k string, i interface{}) bool {
	return e.SetValueTags(k, fmt.Sprint(i))
}

func (a *EventMessageForEsAlertObject) GetTagsAll() []string {
	return a.TagsAll
}

// SetValueTagsAll добавляет значение STRING в список поля TagsAll
func (a *EventMessageForEsAlertObject) SetValueTagsAll(v string) {
	a.TagsAll = append(a.TagsAll, v)
}

// SetAnyTagsAll добавляет ЛЮБОЕ значение в список поля TagsAll
func (a *EventMessageForEsAlertObject) SetAnyTagsAll(i interface{}) {
	a.TagsAll = append(a.TagsAll, fmt.Sprint(i))
}

func (e *EventMessageForEsAlertObject) GetCustomFields() CustomFields {
	return e.CustomFields
}

// SetValueCustomFields устанавливает STRING значение для поля CustomFields
func (e *EventMessageForEsAlertObject) SetValueCustomFields(v CustomFields) {
	e.CustomFields = v
}

func (e *EventMessageForEsAlertObject) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'_id': '%s'\n", ws, e.UnderliningId))
	str.WriteString(fmt.Sprintf("%s'id': '%s'\n", ws, e.Id))
	str.WriteString(fmt.Sprintf("%s'createdBy': '%s'\n", ws, e.CreatedBy))
	str.WriteString(fmt.Sprintf("%s'updatedBy': '%s'\n", ws, e.UpdatedBy))
	str.WriteString(fmt.Sprintf("%s'createdAt': '%s'\n", ws, e.CreatedAt))
	str.WriteString(fmt.Sprintf("%s'updatedAt': '%s'\n", ws, e.UpdatedAt))
	str.WriteString(fmt.Sprintf("%s'_type': '%s'\n", ws, e.UnderliningType))
	str.WriteString(fmt.Sprintf("%s'tlp': '%d'\n", ws, e.Tlp))
	str.WriteString(fmt.Sprintf("%s'title': '%s'\n", ws, e.Title))
	str.WriteString(fmt.Sprintf("%s'description': '%s'\n", ws, e.Description))
	str.WriteString(fmt.Sprintf("%s'status': '%s'\n", ws, e.Status))
	str.WriteString(fmt.Sprintf("%s'date': '%s'\n", ws, e.Date))
	str.WriteString(fmt.Sprintf("%s'type': '%s'\n", ws, e.Type))
	str.WriteString(fmt.Sprintf("%s'objectType': '%s'\n", ws, e.ObjectType))
	str.WriteString(fmt.Sprintf("%s'source': '%s'\n", ws, e.Source))
	str.WriteString(fmt.Sprintf("%s'sourceRef': '%s'\n", ws, e.SourceRef))
	str.WriteString(fmt.Sprintf("%s'case': '%s'\n", ws, e.Case))
	str.WriteString(fmt.Sprintf("%s'caseTemplate': '%s'\n", ws, e.CaseTemplate))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, ToStringBeautifulMapSlice(num, e.Tags)))
	str.WriteString(fmt.Sprintf("%s'tagsAll': \n%s", ws, ToStringBeautifulSlice(num, e.TagsAll)))
	str.WriteString(fmt.Sprintf("%s'customFields': \n%s", ws, CustomFieldsToStringBeautiful(e.CustomFields, num)))

	return str.String()
}
