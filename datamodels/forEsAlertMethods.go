package datamodels

import (
	"fmt"
	"placeholder_elasticsearch/supportingfunctions"
	"strings"
)

func NewAlertMessageForEsAlert() *AlertMessageForEsAlert {
	return &AlertMessageForEsAlert{
		CreatedAt:    "1970-01-01T00:00:00+00:00",
		UpdatedAt:    "1970-01-01T00:00:00+00:00",
		Tags:         make(map[string][]string),
		CustomFields: CustomFields{},
		Artifacts:    make(map[string]ArtifactForEsAlert),
	}
}

// Get возвращает объект типа AlertMessageForEsAlert
func (a *AlertMessageForEsAlert) Get() *AlertMessageForEsAlert {
	return a
}

func (a *AlertMessageForEsAlert) GetTlp() uint64 {
	return a.Tlp
}

// SetValueTlp устанавливает UINT64 значение для поля Tlp
func (a *AlertMessageForEsAlert) SetValueTlp(v uint64) {
	a.Tlp = v
}

// SetAnyTlp устанавливает ЛЮБОЕ значение для поля Tlp
func (a *AlertMessageForEsAlert) SetAnyTlp(i interface{}) {
	if v, ok := i.(float64); ok {
		a.Tlp = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		a.Tlp = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		a.Tlp = v
	}
}

func (a *AlertMessageForEsAlert) GetDate() string {
	return a.Date
}

// SetValueDate устанавливает значение в формате RFC3339 для поля Date
func (a *AlertMessageForEsAlert) SetValueDate(v string) {
	a.Date = v
}

// SetAnyDate устанавливает ЛЮБОЕ значение для поля Date
func (a *AlertMessageForEsAlert) SetAnyDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	a.Date = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (a *AlertMessageForEsAlert) GetCreatedAt() string {
	return a.CreatedAt
}

// SetValueCreatedAt устанавливает значение в формате RFC3339 для поля CreatedAt
func (a *AlertMessageForEsAlert) SetValueCreatedAt(v string) {
	a.CreatedAt = v
}

// SetAnyCreatedAt устанавливает ЛЮБОЕ значение для поля CreatedAt
func (a *AlertMessageForEsAlert) SetAnyCreatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	a.CreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (a *AlertMessageForEsAlert) GetUpdatedAt() string {
	return a.UpdatedAt
}

// SetValueUpdatedAt устанавливает значение в формате RFC3339 для поля UpdatedAt
func (a *AlertMessageForEsAlert) SetValueUpdatedAt(v string) {
	a.UpdatedAt = v
}

// SetAnyUpdatedAt устанавливает ЛЮБОЕ значение для поля UpdatedAt
func (a *AlertMessageForEsAlert) SetAnyUpdatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	a.UpdatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (a *AlertMessageForEsAlert) GetUpdatedBy() string {
	return a.UpdatedBy
}

// SetValueUpdatedBy устанавливает STRING значение для поля UpdatedBy
func (a *AlertMessageForEsAlert) SetValueUpdatedBy(v string) {
	a.UpdatedBy = v
}

// SetAnyUpdatedBy устанавливает ЛЮБОЕ значение для поля UpdatedBy
func (a *AlertMessageForEsAlert) SetAnyUpdatedBy(i interface{}) {
	a.UpdatedBy = fmt.Sprint(i)
}

func (a *AlertMessageForEsAlert) GetUnderliningId() string {
	return a.UnderliningId
}

// SetValueUnderliningId устанавливает STRING значение для поля UnderliningId
func (a *AlertMessageForEsAlert) SetValueUnderliningId(v string) {
	a.UnderliningId = v
}

// SetAnyUnderliningId устанавливает ЛЮБОЕ значение для поля UnderliningId
func (a *AlertMessageForEsAlert) SetAnyUnderliningId(i interface{}) {
	a.UnderliningId = fmt.Sprint(i)
}

func (a *AlertMessageForEsAlert) GetStatus() string {
	return a.Status
}

// SetValueStatus устанавливает STRING значение для поля Status
func (a *AlertMessageForEsAlert) SetValueStatus(v string) {
	a.Status = v
}

// SetAnyStatus устанавливает ЛЮБОЕ значение для поля Status
func (a *AlertMessageForEsAlert) SetAnyStatus(i interface{}) {
	a.Status = fmt.Sprint(i)
}

func (a *AlertMessageForEsAlert) GetType() string {
	return a.Type
}

// SetValueType устанавливает STRING значение для поля Type
func (a *AlertMessageForEsAlert) SetValueType(v string) {
	a.Type = v
}

// SetAnyType устанавливает ЛЮБОЕ значение для поля Type
func (a *AlertMessageForEsAlert) SetAnyType(i interface{}) {
	a.Type = fmt.Sprint(i)
}

func (a *AlertMessageForEsAlert) GetUnderliningType() string {
	return a.UnderliningType
}

// SetValueUnderliningType устанавливает STRING значение для поля UnderliningType
func (a *AlertMessageForEsAlert) SetValueUnderliningType(v string) {
	a.UnderliningType = v
}

// SetAnyUnderliningType устанавливает ЛЮБОЕ значение для поля UnderliningType
func (a *AlertMessageForEsAlert) SetAnyUnderliningType(i interface{}) {
	a.UnderliningType = fmt.Sprint(i)
}

func (a *AlertMessageForEsAlert) GetDescription() string {
	return a.Description
}

// SetValueDescription устанавливает STRING значение для поля Description
func (a *AlertMessageForEsAlert) SetValueDescription(v string) {
	a.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (a *AlertMessageForEsAlert) SetAnyDescription(i interface{}) {
	a.Description = fmt.Sprint(i)
}

func (a *AlertMessageForEsAlert) GetCaseTemplate() string {
	return a.CaseTemplate
}

// SetValueCaseTemplate устанавливает STRING значение для поля CaseTemplate
func (a *AlertMessageForEsAlert) SetValueCaseTemplate(v string) {
	a.CaseTemplate = v
}

// SetAnyCaseTemplate устанавливает ЛЮБОЕ значение для поля CaseTemplate
func (a *AlertMessageForEsAlert) SetAnyCaseTemplate(i interface{}) {
	a.CaseTemplate = fmt.Sprint(i)
}

func (a *AlertMessageForEsAlert) GetSourceRef() string {
	return a.SourceRef
}

// SetValueSourceRef устанавливает STRING значение для поля SourceRef
func (a *AlertMessageForEsAlert) SetValueSourceRef(v string) {
	a.SourceRef = v
}

// SetAnySourceRef устанавливает ЛЮБОЕ значение для поля SourceRef
func (a *AlertMessageForEsAlert) SetAnySourceRef(i interface{}) {
	a.SourceRef = fmt.Sprint(i)
}

func (a *AlertMessageForEsAlert) GetTags() map[string][]string {
	return a.Tags
}

// SetValueTags добаляет значение в Tags по ключу
func (a *AlertMessageForEsAlert) SetValueTags(k, v string) bool {
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
func (a *AlertMessageForEsAlert) SetAnyTags(k string, i interface{}) bool {
	return a.SetValueTags(k, fmt.Sprint(i))
}

func (a *AlertMessageForEsAlert) GetCustomFields() CustomFields {
	return a.CustomFields
}

// SetValueCustomFields устанавливает значение для поля CustomFields
func (a *AlertMessageForEsAlert) SetValueCustomFields(v CustomFields) {
	a.CustomFields = v
}

func (a *AlertMessageForEsAlert) GetArtifacts() map[string]ArtifactForEsAlert {
	return a.Artifacts
}

// SetArtifacts устанавливает значение для поля Artifacts
func (a *AlertMessageForEsAlert) SetValueArtifacts(v map[string]ArtifactForEsAlert) {
	a.Artifacts = v
}

// AddValueArtifact устанавливает значение для поля Artifacts
func (a *AlertMessageForEsAlert) AddValueArtifact(k string, v ArtifactForEsAlert) {
	a.Artifacts[k] = v
}

func (a *AlertMessageForEsAlert) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%stlp: '%d'\n", ws, a.Tlp))
	str.WriteString(fmt.Sprintf("%sdate: '%s'\n", ws, a.Date))
	str.WriteString(fmt.Sprintf("%screatedAt: '%s'\n", ws, a.CreatedAt))
	str.WriteString(fmt.Sprintf("%supdatedAt: '%s'\n", ws, a.UpdatedAt))
	str.WriteString(fmt.Sprintf("%supdatedBy: '%s'\n", ws, a.UpdatedBy))
	str.WriteString(fmt.Sprintf("%sunderliningId: '%s'\n", ws, a.UnderliningId))
	str.WriteString(fmt.Sprintf("%sstatus: '%s'\n", ws, a.Status))
	str.WriteString(fmt.Sprintf("%stype: '%s'\n", ws, a.Type))
	str.WriteString(fmt.Sprintf("%sunderliningType: '%s'\n", ws, a.UnderliningType))
	str.WriteString(fmt.Sprintf("%sdescription: '%s'\n", ws, a.Description))
	str.WriteString(fmt.Sprintf("%scaseTemplate: '%s'\n", ws, a.CaseTemplate))
	str.WriteString(fmt.Sprintf("%ssourceRef: '%s'\n", ws, a.SourceRef))
	str.WriteString(fmt.Sprintf("%stags: \n%s", ws, ToStringBeautifulMapSlice(num, a.Tags)))
	str.WriteString(fmt.Sprintf("%scustomFields: \n%s", ws, CustomFieldsToStringBeautiful(a.CustomFields, num)))
	str.WriteString(fmt.Sprintf("%sartifact: \n%s", ws, func(l map[string]ArtifactForEsAlert) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%s:\n", supportingfunctions.GetWhitespace(num+1), k))
			str.WriteString(v.ToStringBeautiful(num + 2))
		}
		return str.String()
	}(a.Artifacts)))

	return str.String()
}

func NewArtifactForEsAlert() *ArtifactForEsAlert {
	return &ArtifactForEsAlert{
		CreatedAt: "1970-01-01T00:00:00+00:00",
		StartDate: "1970-01-01T00:00:00+00:00",
		Tags:      make(map[string][]string),
	}
}

// Get возвращает объект типа ArtifactForEsAlert
func (a *ArtifactForEsAlert) Get() *ArtifactForEsAlert {
	return a
}

func (a *ArtifactForEsAlert) GetIoc() bool {
	return a.Ioc
}

// SetValueIoc устанавливает BOOL значение для поля Ioc
func (a *ArtifactForEsAlert) SetValueIoc(v bool) {
	a.Ioc = v
}

// SetAnyIoc устанавливает ЛЮБОЕ значение для поля Ioc
func (a *ArtifactForEsAlert) SetAnyIoc(i interface{}) {
	if v, ok := i.(bool); ok {
		a.Ioc = v
	}
}

func (a *ArtifactForEsAlert) GetTlp() uint64 {
	return a.Tlp
}

// SetValueTlp устанавливает UINT64 значение для поля Tlp
func (a *ArtifactForEsAlert) SetValueTlp(v uint64) {
	a.Tlp = v
}

// SetAnyTlp устанавливает ЛЮБОЕ значение для поля Tlp
func (a *ArtifactForEsAlert) SetAnyTlp(i interface{}) {
	if v, ok := i.(float32); ok {
		a.Tlp = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		a.Tlp = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		a.Tlp = v
	}
}

func (a *ArtifactForEsAlert) GetUnderliningId() string {
	return a.UnderliningId
}

// SetValueUnderliningId устанавливает STRING значение для поля UnderliningId
func (a *ArtifactForEsAlert) SetValueUnderliningId(v string) {
	a.UnderliningId = v
}

// SetAnyUnderliningId устанавливает ЛЮБОЕ значение для поля UnderliningId
func (a *ArtifactForEsAlert) SetAnyUnderliningId(i interface{}) {
	a.UnderliningId = fmt.Sprint(i)
}

func (a *ArtifactForEsAlert) GetId() string {
	return a.Id
}

// SetValueId устанавливает STRING значение для поля Id
func (a *ArtifactForEsAlert) SetValueId(v string) {
	a.Id = v
}

// SetAnyId устанавливает ЛЮБОЕ значение для поля Id
func (a *ArtifactForEsAlert) SetAnyId(i interface{}) {
	a.Id = fmt.Sprint(i)
}

func (a *ArtifactForEsAlert) GetUnderliningType() string {
	return a.UnderliningType
}

// SetValueUnderliningType устанавливает STRING значение для поля UnderliningType
func (a *ArtifactForEsAlert) SetValueUnderliningType(v string) {
	a.UnderliningType = v
}

// SetAnyUnderliningType устанавливает ЛЮБОЕ значение для поля UnderliningType
func (a *ArtifactForEsAlert) SetAnyUnderliningType(i interface{}) {
	a.UnderliningType = fmt.Sprint(i)
}

func (a *ArtifactForEsAlert) GetCreatedAt() string {
	return a.CreatedAt
}

// SetValueCreatedAt устанавливает значение в формате RFC3339 для поля CreatedAt
func (a *ArtifactForEsAlert) SetValueCreatedAt(v string) {
	a.CreatedAt = v
}

// SetAnyCreatedAt устанавливает ЛЮБОЕ значение для поля CreatedAt
func (a *ArtifactForEsAlert) SetAnyCreatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	a.CreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (a *ArtifactForEsAlert) GetStartDate() string {
	return a.StartDate
}

// SetValueStartDate устанавливает значение  в формате RFC3339 для поля StartDate
func (a *ArtifactForEsAlert) SetValueStartDate(v string) {
	a.StartDate = v
}

// SetAnyStartDate устанавливает ЛЮБОЕ значение для поля StartDate
func (a *ArtifactForEsAlert) SetAnyStartDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	a.StartDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (a *ArtifactForEsAlert) GetCreatedBy() string {
	return a.CreatedBy
}

// SetValueCreatedBy устанавливает STRING значение для поля CreatedBy
func (a *ArtifactForEsAlert) SetValueCreatedBy(v string) {
	a.CreatedBy = v
}

// SetAnyCreatedBy устанавливает ЛЮБОЕ значение для поля CreatedBy
func (a *ArtifactForEsAlert) SetAnyCreatedBy(i interface{}) {
	a.CreatedBy = fmt.Sprint(i)
}

func (a *ArtifactForEsAlert) GetData() string {
	return a.Data
}

// SetValueData устанавливает STRING значение для поля Data
func (a *ArtifactForEsAlert) SetValueData(v string) {
	a.Data = v
}

// SetAnyData устанавливает ЛЮБОЕ значение для поля Data
func (a *ArtifactForEsAlert) SetAnyData(i interface{}) {
	a.Data = fmt.Sprint(i)
}

func (a *ArtifactForEsAlert) GetDataType() string {
	return a.DataType
}

// SetValueDataType устанавливает STRING значение для поля DataType
func (a *ArtifactForEsAlert) SetValueDataType(v string) {
	a.DataType = v
}

// SetAnyDataType устанавливает ЛЮБОЕ значение для поля DataType
func (a *ArtifactForEsAlert) SetAnyDataType(i interface{}) {
	a.DataType = fmt.Sprint(i)
}

func (a *ArtifactForEsAlert) GetMessage() string {
	return a.Message
}

// SetValueMessage устанавливает STRING значение для поля Message
func (a *ArtifactForEsAlert) SetValueMessage(v string) {
	a.Message = v
}

// SetAnyMessage устанавливает ЛЮБОЕ значение для поля Message
func (a *ArtifactForEsAlert) SetAnyMessage(i interface{}) {
	a.Message = fmt.Sprint(i)
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

func (a *ArtifactForEsAlert) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%sioc: '%t'\n", ws, a.Ioc))
	str.WriteString(fmt.Sprintf("%stlp: '%d'\n", ws, a.Tlp))
	str.WriteString(fmt.Sprintf("%sunderliningId: '%s'\n", ws, a.UnderliningId))
	str.WriteString(fmt.Sprintf("%sid: '%s'\n", ws, a.Id))
	str.WriteString(fmt.Sprintf("%sunderliningType: '%s'\n", ws, a.UnderliningType))
	str.WriteString(fmt.Sprintf("%screatedAt: '%s'\n", ws, a.CreatedAt))
	str.WriteString(fmt.Sprintf("%sstartDate: '%s'\n", ws, a.StartDate))
	str.WriteString(fmt.Sprintf("%screatedBy: '%s'\n", ws, a.CreatedBy))
	str.WriteString(fmt.Sprintf("%sdata: '%s'\n", ws, a.Data))
	str.WriteString(fmt.Sprintf("%sdataType: '%s'\n", ws, a.DataType))
	str.WriteString(fmt.Sprintf("%smessage: '%s'\n", ws, a.Message))
	str.WriteString(fmt.Sprintf("%stags: \n%s", ws, ToStringBeautifulMapSlice(num, a.Tags)))

	return str.String()
}
