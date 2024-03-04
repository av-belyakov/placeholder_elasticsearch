package datamodels

import (
	"fmt"
	"strings"

	commonevent "placeholder_elasticsearch/datamodels/commonevent"
	commonobjectevent "placeholder_elasticsearch/datamodels/commonobjectevent"
	"placeholder_elasticsearch/supportingfunctions"
)

func NewEventMessageForEsAlert() *EventMessageForEsAlert {
	return &EventMessageForEsAlert{
		CommonEventType: commonevent.CommonEventType{
			StartDate: "1970-01-01T00:00:00+00:00",
		},
		Details: *NewEventMessageForEsAlertDetails(),
		Object:  *NewEventMessageForEsAlertObject(),
	}
}

// Get возвращает объект типа EventMessageForEsAlert
func (e *EventMessageForEsAlert) Get() *EventMessageForEsAlert {
	return e
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

	str.WriteString(e.CommonEventType.ToStringBeautiful(num))
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
		CommonEventAlertObject: commonobjectevent.CommonEventAlertObject{
			CreatedAt: "1970-01-01T00:00:00+00:00",
			UpdatedAt: "1970-01-01T00:00:00+00:00",
			Date:      "1970-01-01T00:00:00+00:00",
		},
		Tags:         make(map[string][]string),
		TagsAll:      []string(nil),
		CustomFields: CustomFields{},
	}
}

func (e *EventMessageForEsAlertObject) Get() *EventMessageForEsAlertObject {
	return e
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

	str.WriteString(e.CommonEventAlertObject.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, ToStringBeautifulMapSlice(num, e.Tags)))
	str.WriteString(fmt.Sprintf("%s'tagsAll': \n%s", ws, ToStringBeautifulSlice(num, e.TagsAll)))
	str.WriteString(fmt.Sprintf("%s'customFields': \n%s", ws, CustomFieldsToStringBeautiful(e.CustomFields, num)))

	return str.String()
}
