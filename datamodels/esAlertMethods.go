package datamodels

import (
	"fmt"
	"strconv"
	"strings"

	"placeholder_elasticsearch/datamodels/commonalert"
	"placeholder_elasticsearch/datamodels/commonalertartifact"
	"placeholder_elasticsearch/supportingfunctions"
)

func NewAlertMessageForEsAlert() *AlertMessageForEsAlert {
	return &AlertMessageForEsAlert{
		CommonAlertType: commonalert.CommonAlertType{
			Date:      "1970-01-01T00:00:00+00:00",
			CreatedAt: "1970-01-01T00:00:00+00:00",
			UpdatedAt: "1970-01-01T00:00:00+00:00",
		},
		Tags:         make(map[string][]string),
		TagsAll:      []string(nil),
		CustomFields: CustomFields{},
		Artifacts:    make(map[string][]ArtifactForEsAlert),
	}
}

// Get возвращает объект типа AlertMessageForEsAlert
func (a *AlertMessageForEsAlert) Get() *AlertMessageForEsAlert {
	return a
}

func (a *AlertMessageForEsAlert) GetFollow() bool {
	return a.Follow
}

// SetValueFollow устанавливает BOOL значение для поля Follow
func (a *AlertMessageForEsAlert) SetValueFollow(v bool) {
	a.Follow = v
}

// SetAnyFollow устанавливает ЛЮБОЕ значение для поля Follow
func (a *AlertMessageForEsAlert) SetAnyFollow(i interface{}) {
	if v, ok := i.(bool); ok {
		a.Follow = v
	}
}

func (a *AlertMessageForEsAlert) GetSeverity() uint64 {
	return a.Severity
}

// SetValueSeverity устанавливает UINT64 значение для поля Severity
func (a *AlertMessageForEsAlert) SetValueSeverity(v uint64) {
	a.Severity = v
}

// SetAnySeverity устанавливает ЛЮБОЕ значение для поля Severity
func (a *AlertMessageForEsAlert) SetAnySeverity(i interface{}) {
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

func (a *AlertMessageForEsAlert) GetTags() map[string][]string {
	return a.Tags
}

// SetValueTags добаляет значение в Tags по ключу
func (a *AlertMessageForEsAlert) SetValueTags(k, v string) bool {
	if _, ok := a.Tags[k]; !ok {
		a.Tags[k] = []string(nil)
	}

	v = supportingfunctions.TrimIsNotLetter(v)
	for _, value := range a.Tags[k] {
		if v == value {
			return false
		}
	}

	a.Tags[k] = append(a.Tags[k], v)

	return true
}

// SetAnyTags устанавливает ЛЮБОЕ значение для поля Tags
func (a *AlertMessageForEsAlert) SetAnyTags(k string, i interface{}) bool {
	return a.SetValueTags(k, fmt.Sprint(i))
}

func (a *AlertMessageForEsAlert) GetTagsAll() []string {
	return a.TagsAll
}

// SetValueTagsAll добавляет значение STRING в список поля TagsAll
func (a *AlertMessageForEsAlert) SetValueTagsAll(v string) {
	a.TagsAll = append(a.TagsAll, v)
}

// SetAnyTagsAll добавляет ЛЮБОЕ значение в список поля TagsAll
func (a *AlertMessageForEsAlert) SetAnyTagsAll(i interface{}) {
	a.TagsAll = append(a.TagsAll, fmt.Sprint(i))
}

func (a *AlertMessageForEsAlert) GetCustomFields() CustomFields {
	return a.CustomFields
}

// SetValueCustomFields устанавливает значение для поля CustomFields
func (a *AlertMessageForEsAlert) SetValueCustomFields(v CustomFields) {
	a.CustomFields = v
}

func (a *AlertMessageForEsAlert) GetArtifacts() map[string][]ArtifactForEsAlert {
	return a.Artifacts
}

func (a *AlertMessageForEsAlert) GetKeyArtifacts(k string) ([]ArtifactForEsAlert, bool) {
	if value, ok := a.Artifacts[k]; ok {
		return value, true
	}

	return nil, false
}

func (a *AlertMessageForEsAlert) SetKeyArtifacts(k string, artifacts []ArtifactForEsAlert) {
	a.Artifacts[k] = artifacts
}

// SetArtifacts устанавливает значение для поля Artifacts
func (a *AlertMessageForEsAlert) SetValueArtifacts(v map[string][]ArtifactForEsAlert) {
	a.Artifacts = v
}

// AddValueArtifact устанавливает значение для поля Artifacts
func (a *AlertMessageForEsAlert) AddValueArtifact(k string, v ArtifactForEsAlert) {
	if _, ok := a.Artifacts[k]; !ok {
		a.Artifacts[k] = []ArtifactForEsAlert(nil)
	}

	a.Artifacts[k] = append(a.Artifacts[k], v)
}

func (a *AlertMessageForEsAlert) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(a.CommonAlertType.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'follow': '%t'\n", ws, a.Follow))
	str.WriteString(fmt.Sprintf("%s'severity': '%d'\n", ws, a.Severity))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, ToStringBeautifulMapSlice(num, a.Tags)))
	str.WriteString(fmt.Sprintf("%s'tagsAll': \n%s", ws, ToStringBeautifulSlice(num, a.TagsAll)))
	str.WriteString(fmt.Sprintf("%s'customFields': \n%s", ws, CustomFieldsToStringBeautiful(a.CustomFields, num)))
	str.WriteString(fmt.Sprintf("%s'artifact': \n%s", ws, func(l map[string][]ArtifactForEsAlert) string {
		str := strings.Builder{}

		for key, value := range l {
			str.WriteString(fmt.Sprintf("%s%s:\n", supportingfunctions.GetWhitespace(num+1), key))
			for k, v := range value {
				str.WriteString(fmt.Sprintf("%s%d.\n", supportingfunctions.GetWhitespace(num+2), k))
				str.WriteString(v.ToStringBeautiful(num + 3))
			}
		}
		return str.String()
	}(a.Artifacts)))

	return str.String()
}

func NewArtifactForEsAlert() *ArtifactForEsAlert {
	return &ArtifactForEsAlert{
		CommonArtifactType: commonalertartifact.CommonArtifactType{
			CreatedAt: "1970-01-01T00:00:00+00:00",
			StartDate: "1970-01-01T00:00:00+00:00",
		},
		SnortSid:       []string(nil),
		SnortSidNumber: []int(nil),
		Tags:           make(map[string][]string),
		TagsAll:        []string(nil),
	}
}

// Get возвращает объект типа ArtifactForEsAlert
func (a *ArtifactForEsAlert) Get() *ArtifactForEsAlert {
	return a
}

func (a *ArtifactForEsAlert) GetSensorId() string {
	return a.SensorId
}

// SetSensorId добаляет значение в SensorId
func (a *ArtifactForEsAlert) SetValueSensorId(v string) {
	a.SensorId = v
}

// SetSensorId добавляет ЛЮБОЕ значение в SensorId
func (a *ArtifactForEsAlert) SetAnySensorId(i interface{}) {
	a.SensorId = fmt.Sprint(i)
}

func (a *ArtifactForEsAlert) GetSnortSid() []string {
	return a.SnortSid
}

// SetValueSnortSid добавляет значение STRING в список поля SnortSid
func (a *ArtifactForEsAlert) SetValueSnortSid(v string) {
	a.SnortSid = append(a.SnortSid, v)
}

// SetAnySnortSid добавляет ЛЮБОЕ значение в список поля SnortSid
func (a *ArtifactForEsAlert) SetAnySnortSid(i interface{}) {
	a.SnortSid = append(a.SnortSid, fmt.Sprint(i))
}

func (a *ArtifactForEsAlert) GetSnortSidNumber() []int {
	return a.SnortSidNumber
}

// SetValueSnortSidNumber добавляет значение INT в список поля SnortSidNumber
func (a *ArtifactForEsAlert) SetValueSnortSidNumber(v int) {
	a.SnortSidNumber = append(a.SnortSidNumber, v)
}

// SetAnySnortSidNumber добавляет ЛЮБОЕ значение в список поля SnortSidNumber
func (a *ArtifactForEsAlert) SetAnySnortSidNumber(i interface{}) {
	str := fmt.Sprint(i)
	if num, err := strconv.ParseUint(str, 10, 32); err == nil {
		a.SetValueSnortSidNumber(int(num))
	}
}

func (a *ArtifactForEsAlert) GetTags() map[string][]string {
	return a.Tags
}

// SetValueTags добаляет значение в Tags по ключу
func (a *ArtifactForEsAlert) SetValueTags(k, v string) bool {
	if _, ok := a.Tags[k]; !ok {
		a.Tags[k] = []string(nil)
	}

	for _, value := range a.Tags[k] {
		if v == value {
			return false
		}
	}

	a.Tags[k] = append(a.Tags[k], v)

	return true
}

// SetAnyTags устанавливает ЛЮБОЕ значение для поля Tags
func (a *ArtifactForEsAlert) SetAnyTags(k string, i interface{}) bool {
	return a.SetValueTags(k, fmt.Sprint(i))
}

func (a *ArtifactForEsAlert) GetTagsAll() []string {
	return a.TagsAll
}

// SetValueTagsAll добавляет значение STRING в список поля TagsAll
func (a *ArtifactForEsAlert) SetValueTagsAll(v string) {
	a.TagsAll = append(a.TagsAll, v)
}

// SetAnyTagsAll добавляет ЛЮБОЕ значение в список поля TagsAll
func (a *ArtifactForEsAlert) SetAnyTagsAll(i interface{}) {
	a.TagsAll = append(a.TagsAll, fmt.Sprint(i))
}

func (a *ArtifactForEsAlert) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(a.CommonArtifactType.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'sensorId': '%s'\n", ws, a.SensorId))
	str.WriteString(fmt.Sprintf("%s'snortSid': \n%s", ws, ToStringBeautifulSlice(num, a.SnortSid)))
	str.WriteString(fmt.Sprintf("%s'snortSidNumber': \n%s", ws, ToStringBeautifulSlice(num, a.SnortSidNumber)))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, ToStringBeautifulMapSlice(num, a.Tags)))
	str.WriteString(fmt.Sprintf("%s'tagsAll': \n%s", ws, ToStringBeautifulSlice(num, a.TagsAll)))

	return str.String()
}
