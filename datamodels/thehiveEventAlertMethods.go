package datamodels

import (
	"fmt"
	"strings"

	commonevent "placeholder_elasticsearch/datamodels/commonevent"
	commonobjectevent "placeholder_elasticsearch/datamodels/commonobjectevent"
	"placeholder_elasticsearch/supportingfunctions"
)

func NewEventMessageTheHiveAlert() *EventMessageTheHiveAlert {
	return &EventMessageTheHiveAlert{
		CommonEventType: commonevent.CommonEventType{
			StartDate: "1970-01-01T00:00:00+00:00",
		},
		Details: *NewEventAlertDetails(),
		Object:  *NewEventAlertObject(),
	}
}

// Get возвращает объект типа EventMessageTheHiveCase
func (e *EventMessageTheHiveAlert) Get() *EventMessageTheHiveAlert {
	return e
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

	str.WriteString(e.CommonEventType.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'details':\n", ws))
	str.WriteString(e.Details.ToStringBeautiful(num + 1))
	str.WriteString(fmt.Sprintf("%s'object':\n", ws))
	str.WriteString(e.Object.ToStringBeautiful(num + 1))

	return str.String()
}

//****************** EventAlertDetails ******************

func NewEventAlertDetails() *EventAlertDetails {
	return &EventAlertDetails{Tags: []string(nil)}
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
	str := fmt.Sprint(i)
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\n", "")

	e.Description = str
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

// SetSliceValueTags устанавливает []STRING значение для поля Tags
func (e *EventAlertDetails) SetSliceValueTags(v []string) {
	e.Tags = v
}

func (e *EventAlertDetails) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'sourceRef': '%s'\n", ws, e.SourceRef))
	str.WriteString(fmt.Sprintf("%s'title': '%s'\n", ws, e.Title))
	str.WriteString(fmt.Sprintf("%s'description': '%s'\n", ws, e.Description))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, ToStringBeautifulSlice(num, e.Tags)))

	return str.String()
}

//****************** EventAlertObject ******************

func NewEventAlertObject() *EventAlertObject {
	return &EventAlertObject{
		CommonEventAlertObject: commonobjectevent.CommonEventAlertObject{
			CreatedAt: "1970-01-01T00:00:00+00:00",
			UpdatedAt: "1970-01-01T00:00:00+00:00",
		},
		Tags:         []string(nil),
		CustomFields: CustomFields{},
	}
}

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

func (e *EventAlertObject) GetPap() uint64 {
	return e.Pap
}

// SetValuePap устанавливает INT значение для поля Pap
func (e *EventAlertObject) SetValuePap(v uint64) {
	e.Pap = v
}

// SetAnyPap устанавливает ЛЮБОЕ значение для поля Pap
func (e *EventAlertObject) SetAnyPap(i interface{}) {
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

// SetSliceValueTags устанавливает []STRING значение для поля Tags
func (e *EventAlertObject) SetSliceValueTags(v []string) {
	e.Tags = v
}

func (e *EventAlertObject) GetCustomFields() CustomFields {
	return e.CustomFields
}

// SetValueCustomFields устанавливает STRING значение для поля CustomFields
func (e *EventAlertObject) SetValueCustomFields(v CustomFields) {
	e.CustomFields = v
}

func (e *EventAlertObject) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'follow': '%t'\n", ws, e.Follow))
	str.WriteString(fmt.Sprintf("%s'severity': '%d'\n", ws, e.Severity))
	str.WriteString(fmt.Sprintf("%s'pap': '%d'\n", ws, e.Pap))
	str.WriteString(e.CommonEventAlertObject.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, ToStringBeautifulSlice(num, e.Tags)))
	str.WriteString(fmt.Sprintf("%s'customFields': \n%s", ws, CustomFieldsToStringBeautiful(e.CustomFields, num)))

	return str.String()
}

/*func (e *EventMessageTheHiveAlert) MarshalBSON() ([]byte, error) {
	return bson.Marshal(struct {
		Base           bool              `bson:"base"`
		StartDate      string            `bson:"startDate"` //в формате RFC3339
		RootId         string            `bson:"rootId"`
		Organisation   string            `bson:"organisation"`
		OrganisationId string            `bson:"organisationId"`
		ObjectId       string            `bson:"objectId"`
		ObjectType     string            `bson:"objectType"`
		Operation      string            `bson:"operation"`
		RequestId      string            `bson:"requestId"`
		Details        EventAlertDetails `bson:"details"`
		Object         EventAlertObject  `bson:"object"`
	}{
		Base:           e.Base,
		StartDate:      e.StartDate,
		RootId:         e.RootId,
		Organisation:   e.Organisation,
		OrganisationId: e.OrganisationId,
		ObjectId:       e.ObjectId,
		ObjectType:     e.ObjectType,
		Operation:      e.Operation,
		RequestId:      e.RequestId,
		Details:        e.Details,
		Object:         e.Object,
	})
	struct {
		Follow          bool         `bson:"follow"`
		Severity        uint64       `bson:"severity"`
		Tlp             uint64       `bson:"tlp"`
		UnderliningId   string       `bson:"_id"`
		Id              string       `bson:"id"`
		CreatedBy       string       `bson:"createdBy"`
		UpdatedBy       string       `bson:"updatedBy,omitempty"`
		CreatedAt       string       `bson:"createdAt"`           //формат RFC3339
		UpdatedAt       string       `bson:"updatedAt,omitempty"` //формат RFC3339
		UnderliningType string       `bson:"_type"`
		Title           string       `bson:"title"`
		Description     string       `bson:"description"`
		Status          string       `bson:"status"`
		Date            string       `bson:"date"` //формат RFC3339
		Type            string       `bson:"type"`
		ObjectType      string       `bson:"objectType"`
		Source          string       `bson:"source"`
		SourceRef       string       `bson:"sourceRef"`
		Case            string       `bson:"case,omitempty"`
		CaseTemplate    string       `bson:"caseTemplate,omitempty"`
		Tags            []string     `bson:"tags"`
		CustomFields    CustomFields `bson:"customFields"`
	}{
		Follow:          e.Follow,
		Severity:        e.Severity,
		Tlp:             e.Tlp,
		UnderliningId:   e.UnderliningId,
		Id:              e.Id,
		CreatedBy:       e.CreatedBy,
		UpdatedBy:       e.UpdatedBy,
		CreatedAt:       e.CreatedAt,
		UpdatedAt:       e.UpdatedAt,
		UnderliningType: e.UnderliningType,
		Title:           e.Title,
		Description:     e.Description,
		Status:          e.Status,
		Date:            e.Date,
		Type:            e.Type,
		ObjectType:      e.ObjectType,
		Source:          e.Source,
		SourceRef:       e.SourceRef,
		Case:            e.Case,
		CaseTemplate:    e.CaseTemplate,
		Tags:            e.Tags,
		CustomFields:    e.CustomFields,
	}
}*/
