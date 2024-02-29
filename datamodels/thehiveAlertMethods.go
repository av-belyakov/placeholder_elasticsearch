package datamodels

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/datamodels/commonalert"
	"placeholder_elasticsearch/datamodels/commonalertartifact"
	"placeholder_elasticsearch/supportingfunctions"
)

func NewAlertMessageTheHiveAlert() *AlertMessageTheHiveAlert {
	return &AlertMessageTheHiveAlert{
		CommonAlertType: commonalert.CommonAlertType{
			CreatedAt: "1970-01-01T00:00:00+00:00",
			UpdatedAt: "1970-01-01T00:00:00+00:00",
		},
		Tags:         []string(nil),
		CustomFields: CustomFields{},
		Artifacts:    []AlertArtifact(nil),
	}
}

// Get возвращает объект типа AlertMessageTheHiveAlert
func (a *AlertMessageTheHiveAlert) Get() *AlertMessageTheHiveAlert {
	return a
}

func (a *AlertMessageTheHiveAlert) GetFollow() bool {
	return a.Follow
}

// SetValueFollow устанавливает BOOL значение для поля Follow
func (a *AlertMessageTheHiveAlert) SetValueFollow(v bool) {
	a.Follow = v
}

// SetAnyFollow устанавливает ЛЮБОЕ значение для поля Follow
func (a *AlertMessageTheHiveAlert) SetAnyFollow(i interface{}) {
	if v, ok := i.(bool); ok {
		a.Follow = v
	}
}

func (a *AlertMessageTheHiveAlert) GetSeverity() uint64 {
	return a.Severity
}

// SetValueSeverity устанавливает UINT64 значение для поля Severity
func (a *AlertMessageTheHiveAlert) SetValueSeverity(v uint64) {
	a.Severity = v
}

// SetAnySeverity устанавливает ЛЮБОЕ значение для поля Severity
func (a *AlertMessageTheHiveAlert) SetAnySeverity(i interface{}) {
	if v, ok := i.(float64); ok {
		a.Severity = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		a.Severity = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		a.Severity = v
	}
}

func (a *AlertMessageTheHiveAlert) GetTags() []string {
	return a.Tags
}

// SetValueTags устанавливает STRING значение для поля Tags
func (a *AlertMessageTheHiveAlert) SetValueTags(v string) {
	a.Tags = append(a.Tags, v)
}

// SetAnyTags устанавливает ЛЮБОЕ значение для поля Tags
func (a *AlertMessageTheHiveAlert) SetAnyTags(i interface{}) {
	a.Tags = append(a.Tags, fmt.Sprint(i))
}

func (a *AlertMessageTheHiveAlert) GetCustomFields() CustomFields {
	return a.CustomFields
}

// SetValueCustomFields устанавливает значение для поля CustomFields
func (a *AlertMessageTheHiveAlert) SetValueCustomFields(v CustomFields) {
	a.CustomFields = v
}

func (a *AlertMessageTheHiveAlert) GetArtifacts() []AlertArtifact {
	return a.Artifacts
}

// SetArtifacts устанавливает значение для поля Artifacts
func (a *AlertMessageTheHiveAlert) SetValueArtifacts(v []AlertArtifact) {
	a.Artifacts = v
}

// AddValueArtifact устанавливает значение для поля Artifacts
func (a *AlertMessageTheHiveAlert) AddValueArtifact(v AlertArtifact) {
	a.Artifacts = append(a.Artifacts, v)
}

func (a *AlertMessageTheHiveAlert) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'follow': '%t'\n", ws, a.Follow))
	str.WriteString(fmt.Sprintf("%s'severity': '%d'\n", ws, a.Severity))
	str.WriteString(a.CommonAlertType.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, ToStringBeautifulSlice(num, a.Tags)))
	str.WriteString(fmt.Sprintf("%s'customFields': \n%s", ws, CustomFieldsToStringBeautiful(a.CustomFields, num)))
	str.WriteString(fmt.Sprintf("%s'artifact': \n%s", ws, func(l []AlertArtifact) string {
		str := strings.Builder{}
		ws := supportingfunctions.GetWhitespace(num + 1)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d.\n", ws, k+1))
			str.WriteString(v.ToStringBeautiful(num + 2))
		}
		return str.String()
	}(a.Artifacts)))

	return str.String()
}

func NewAlertArtifact() *AlertArtifact {
	return &AlertArtifact{
		CommonArtifactType: commonalertartifact.CommonArtifactType{
			CreatedAt: "1970-01-01T00:00:00+00:00",
			StartDate: "1970-01-01T00:00:00+00:00",
		},
		UpdatedAt: "1970-01-01T00:00:00+00:00",
		Tags:      []string(nil),
	}
}

// Get возвращает объект типа AlertArtifact
func (a *AlertArtifact) Get() *AlertArtifact {
	return a
}

func (a *AlertArtifact) GetSighted() bool {
	return a.Sighted
}

// SetValueSighted устанавливает BOOL значение для поля Sighted
func (a *AlertArtifact) SetValueSighted(v bool) {
	a.Sighted = v
}

// SetAnySighted устанавливает ЛЮБОЕ значение для поля Sighted
func (a *AlertArtifact) SetAnySighted(i interface{}) {
	if v, ok := i.(bool); ok {
		a.Sighted = v
	}
}

func (a *AlertArtifact) GetIgnoreSimilarity() bool {
	return a.IgnoreSimilarity
}

// SetValueIgnoreSimilarity устанавливает BOOL значение для поля IgnoreSimilarity
func (a *AlertArtifact) SetValueIgnoreSimilarity(v bool) {
	a.IgnoreSimilarity = v
}

// SetAnyIgnoreSimilarity устанавливает ЛЮБОЕ значение для поля IgnoreSimilarity
func (a *AlertArtifact) SetAnyIgnoreSimilarity(i interface{}) {
	if v, ok := i.(bool); ok {
		a.IgnoreSimilarity = v
	}
}

func (a *AlertArtifact) GetUpdatedAt() string {
	return a.UpdatedAt
}

// SetValueUpdatedAt устанавливает значение  в формате RFC3339 для поля UpdatedAt
func (a *AlertArtifact) SetValueUpdatedAt(v string) {
	a.UpdatedAt = v
}

// SetAnyUpdatedAt устанавливает ЛЮБОЕ значение для поля UpdatedAt
func (a *AlertArtifact) SetAnyUpdatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	a.UpdatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (a *AlertArtifact) GetUpdatedBy() string {
	return a.UpdatedBy
}

// SetValueUpdatedBy устанавливает STRING значение для поля UpdatedBy
func (a *AlertArtifact) SetValueUpdatedBy(v string) {
	a.UpdatedBy = v
}

// SetAnyUpdatedBy устанавливает ЛЮБОЕ значение для поля UpdatedBy
func (a *AlertArtifact) SetAnyUpdatedBy(i interface{}) {
	a.UpdatedBy = fmt.Sprint(i)
}

func (a *AlertArtifact) GetTags() []string {
	return a.Tags
}

// SetValueTags устанавливает STRING значение для поля Tags
func (a *AlertArtifact) SetValueTags(v string) {
	a.Tags = append(a.Tags, v)
}

// SetAnyTags устанавливает ЛЮБОЕ значение для поля Tags
func (a *AlertArtifact) SetAnyTags(i interface{}) {
	a.Tags = append(a.Tags, fmt.Sprint(i))
}

func (a *AlertArtifact) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'sighted': '%t'\n", ws, a.Sighted))
	str.WriteString(fmt.Sprintf("%s'ignoreSimilarity': '%t'\n", ws, a.IgnoreSimilarity))
	str.WriteString(a.CommonArtifactType.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'updatedAt': '%s'\n", ws, a.UpdatedAt))
	str.WriteString(fmt.Sprintf("%s'updatedBy': '%s'\n", ws, a.UpdatedBy))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, ToStringBeautifulSlice(num, a.Tags)))

	return str.String()
}
