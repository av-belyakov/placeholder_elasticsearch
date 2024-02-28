package datamodels

import (
	"fmt"
	"strings"

	commonevent "placeholder_elasticsearch/datamodels/commonevent"
	commonobjectevent "placeholder_elasticsearch/datamodels/commonobjectevent"
	"placeholder_elasticsearch/supportingfunctions"
)

func NewEventMessageForEsCase() *EventMessageForEsCase {
	return &EventMessageForEsCase{
		CommonEventType: commonevent.CommonEventType{
			StartDate: "1970-01-01T00:00:00+00:00",
		},
		Details: *NewEventCaseDetails(),
		Object:  *NewEventForEsCaseObject(),
	}
}

// Get возвращает объект типа EventMessageTheHiveCase
func (e *EventMessageForEsCase) Get() *EventMessageForEsCase {
	return e
}

func (e *EventMessageForEsCase) GetDetails() EventCaseDetails {
	return e.Details
}

// SetValueDetails устанавливает значение типа EventDetails для поля Details
func (e *EventMessageForEsCase) SetValueDetails(v EventCaseDetails) {
	e.Details = v
}

func (e *EventMessageForEsCase) GetObject() EventForEsCaseObject {
	return e.Object
}

// SetValueObject устанавливает значение типа EventForEsCaseObject для поля Object
func (e *EventMessageForEsCase) SetValueObject(v EventForEsCaseObject) {
	e.Object = v
}

func (em EventMessageForEsCase) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(em.CommonEventType.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'details':\n", ws))
	str.WriteString(em.Details.ToStringBeautiful(num + 1))
	str.WriteString(fmt.Sprintf("%s'object':\n", ws))
	str.WriteString(em.Object.ToStringBeautiful(num + 1))

	return str.String()
}

//****************** EventForEsCaseObject ******************

func NewEventForEsCaseObject() *EventForEsCaseObject {
	return &EventForEsCaseObject{
		CommonEventCaseObject: commonobjectevent.CommonEventCaseObject{
			StartDate: "1970-01-01T00:00:00+00:00",
			EndDate:   "1970-01-01T00:00:00+00:00",
			CreatedAt: "1970-01-01T00:00:00+00:00",
			UpdatedAt: "1970-01-01T00:00:00+00:00",
		},
		Tags:         make(map[string][]string),
		TagsAll:      []string(nil),
		CustomFields: CustomFields{},
	}
}

func (e *EventForEsCaseObject) Get() *EventForEsCaseObject {
	return e
}

func (e *EventForEsCaseObject) GetTags() map[string][]string {
	return e.Tags
}

// SetValueTags добаляет значение в Tags по ключу
func (e *EventForEsCaseObject) SetValueTags(k, v string) bool {
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
func (e *EventForEsCaseObject) SetAnyTags(k string, i interface{}) bool {
	return e.SetValueTags(k, fmt.Sprint(i))
}

func (a *EventForEsCaseObject) GetTagsAll() []string {
	return a.TagsAll
}

// SetValueTagsAll добавляет значение STRING в список поля TagsAll
func (a *EventForEsCaseObject) SetValueTagsAll(v string) {
	a.TagsAll = append(a.TagsAll, v)
}

// SetAnyTagsAll добавляет ЛЮБОЕ значение в список поля TagsAll
func (a *EventForEsCaseObject) SetAnyTagsAll(i interface{}) {
	a.TagsAll = append(a.TagsAll, fmt.Sprint(i))
}

func (e *EventForEsCaseObject) GetCustomFields() CustomFields {
	return e.CustomFields
}

// SetValueCustomFields устанавливает STRING значение для поля CustomFields
func (e *EventForEsCaseObject) SetValueCustomFields(v CustomFields) {
	e.CustomFields = v
}

func (eo EventForEsCaseObject) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(eo.CommonEventCaseObject.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, ToStringBeautifulMapSlice(num, eo.Tags)))
	str.WriteString(fmt.Sprintf("%s'tagsAll': \n%s", ws, ToStringBeautifulSlice(num, eo.TagsAll)))
	str.WriteString(fmt.Sprintf("%s'customFields': \n%s", ws, CustomFieldsToStringBeautiful(eo.CustomFields, num)))

	return str.String()
}
