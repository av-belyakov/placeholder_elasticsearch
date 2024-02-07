package datamodels

import (
	"fmt"
	"placeholder_elasticsearch/supportingfunctions"
	"strings"
)

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

func (a *AlertMessageTheHiveAlert) GetTlp() uint64 {
	return a.Tlp
}

// SetValueTlp устанавливает UINT64 значение для поля Tlp
func (a *AlertMessageTheHiveAlert) SetValueTlp(v uint64) {
	a.Tlp = v
}

// SetAnyTlp устанавливает ЛЮБОЕ значение для поля Tlp
func (a *AlertMessageTheHiveAlert) SetAnyTlp(i interface{}) {
	if v, ok := i.(float64); ok {
		a.Tlp = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		a.Tlp = v
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

	if v, ok := i.(uint64); ok {
		a.Severity = v
	}
}

func (a *AlertMessageTheHiveAlert) GetDate() string {
	return a.Date
}

// SetValueDate устанавливает значение в формате RFC3339 для поля Date
func (a *AlertMessageTheHiveAlert) SetValueDate(v string) {
	a.Date = v
}

// SetAnyDate устанавливает ЛЮБОЕ значение для поля Date
func (a *AlertMessageTheHiveAlert) SetAnyDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	a.Date = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (a *AlertMessageTheHiveAlert) GetCreatedAt() string {
	return a.CreatedAt
}

// SetValueCreatedAt устанавливает значение в формате RFC3339 для поля CreatedAt
func (a *AlertMessageTheHiveAlert) SetValueCreatedAt(v string) {
	a.CreatedAt = v
}

// SetAnyCreatedAt устанавливает ЛЮБОЕ значение для поля CreatedAt
func (a *AlertMessageTheHiveAlert) SetAnyCreatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	a.CreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (a *AlertMessageTheHiveAlert) GetUpdatedAt() string {
	return a.UpdatedAt
}

// SetValueUpdatedAt устанавливает значение в формате RFC3339 для поля UpdatedAt
func (a *AlertMessageTheHiveAlert) SetValueUpdatedAt(v string) {
	a.UpdatedAt = v
}

// SetAnyUpdatedAt устанавливает ЛЮБОЕ значение для поля UpdatedAt
func (a *AlertMessageTheHiveAlert) SetAnyUpdatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	a.UpdatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (a *AlertMessageTheHiveAlert) GetUpdatedBy() string {
	return a.UpdatedBy
}

// SetValueUpdatedBy устанавливает STRING значение для поля UpdatedBy
func (a *AlertMessageTheHiveAlert) SetValueUpdatedBy(v string) {
	a.UpdatedBy = v
}

// SetAnyUpdatedBy устанавливает ЛЮБОЕ значение для поля UpdatedBy
func (a *AlertMessageTheHiveAlert) SetAnyUpdatedBy(i interface{}) {
	a.UpdatedBy = fmt.Sprint(i)
}

func (a *AlertMessageTheHiveAlert) GetUnderliningId() string {
	return a.UnderliningId
}

// SetValueUnderliningId устанавливает STRING значение для поля UnderliningId
func (a *AlertMessageTheHiveAlert) SetValueUnderliningId(v string) {
	a.UnderliningId = v
}

// SetAnyUnderliningId устанавливает ЛЮБОЕ значение для поля UnderliningId
func (a *AlertMessageTheHiveAlert) SetAnyUnderliningId(i interface{}) {
	a.UnderliningId = fmt.Sprint(i)
}

func (a *AlertMessageTheHiveAlert) GetStatus() string {
	return a.Status
}

// SetValueStatus устанавливает STRING значение для поля Status
func (a *AlertMessageTheHiveAlert) SetValueStatus(v string) {
	a.Status = v
}

// SetAnyStatus устанавливает ЛЮБОЕ значение для поля Status
func (a *AlertMessageTheHiveAlert) SetAnyStatus(i interface{}) {
	a.Status = fmt.Sprint(i)
}

func (a *AlertMessageTheHiveAlert) GetType() string {
	return a.Type
}

// SetValueType устанавливает STRING значение для поля Type
func (a *AlertMessageTheHiveAlert) SetValueType(v string) {
	a.Type = v
}

// SetAnyType устанавливает ЛЮБОЕ значение для поля Type
func (a *AlertMessageTheHiveAlert) SetAnyType(i interface{}) {
	a.Type = fmt.Sprint(i)
}

func (a *AlertMessageTheHiveAlert) GetUnderliningType() string {
	return a.UnderliningType
}

// SetValueUnderliningType устанавливает STRING значение для поля UnderliningType
func (a *AlertMessageTheHiveAlert) SetValueUnderliningType(v string) {
	a.UnderliningType = v
}

// SetAnyUnderliningType устанавливает ЛЮБОЕ значение для поля UnderliningType
func (a *AlertMessageTheHiveAlert) SetAnyUnderliningType(i interface{}) {
	a.UnderliningType = fmt.Sprint(i)
}

func (a *AlertMessageTheHiveAlert) GetDescription() string {
	return a.Description
}

// SetValueDescription устанавливает STRING значение для поля Description
func (a *AlertMessageTheHiveAlert) SetValueDescription(v string) {
	a.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (a *AlertMessageTheHiveAlert) SetAnyDescription(i interface{}) {
	a.Description = fmt.Sprint(i)
}

func (a *AlertMessageTheHiveAlert) GetCaseTemplate() string {
	return a.CaseTemplate
}

// SetValueCaseTemplate устанавливает STRING значение для поля CaseTemplate
func (a *AlertMessageTheHiveAlert) SetValueCaseTemplate(v string) {
	a.CaseTemplate = v
}

// SetAnyCaseTemplate устанавливает ЛЮБОЕ значение для поля CaseTemplate
func (a *AlertMessageTheHiveAlert) SetAnyCaseTemplate(i interface{}) {
	a.CaseTemplate = fmt.Sprint(i)
}

func (a *AlertMessageTheHiveAlert) GetSourceRef() string {
	return a.SourceRef
}

// SetValueSourceRef устанавливает STRING значение для поля SourceRef
func (a *AlertMessageTheHiveAlert) SetValueSourceRef(v string) {
	a.SourceRef = v
}

// SetAnySourceRef устанавливает ЛЮБОЕ значение для поля SourceRef
func (a *AlertMessageTheHiveAlert) SetAnySourceRef(i interface{}) {
	a.SourceRef = fmt.Sprint(i)
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

func (a *AlertMessageTheHiveAlert) GetCustomFields() map[string]CustomerFields {
	return a.CustomFields
}

// SetValueCustomFields устанавливает STRING значение для поля CustomFields
func (a *AlertMessageTheHiveAlert) SetValueCustomFields(v map[string]CustomerFields) {
	a.CustomFields = v
}

func (a *AlertMessageTheHiveAlert) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%sfollow: '%t'\n", ws, a.Follow))
	str.WriteString(fmt.Sprintf("%stlp: '%d'\n", ws, a.Tlp))
	str.WriteString(fmt.Sprintf("%sseverity: '%d'\n", ws, a.Severity))
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
	str.WriteString(fmt.Sprintf("%stags: \n%s", ws, func(l []string) string {
		str := strings.Builder{}
		ws := supportingfunctions.GetWhitespace(num + 1)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d. '%s'\n", ws, k+1, v))
		}
		return str.String()
	}(a.Tags)))
	str.WriteString(fmt.Sprintf("%scustomFields: \n%s", ws, CustomFieldsToStringBeautiful(a.CustomFields, num)))
	str.WriteString(fmt.Sprintf("%sartifact: \n%s", ws, func(l []AlertArtifact) string {
		str := strings.Builder{}
		ws := supportingfunctions.GetWhitespace(num + 1)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d.\n", ws, k+1))
			str.WriteString(v.ToStringBeautiful(num))
		}
		return str.String()
	}(a.Artifacts)))

	return str.String()
}

// Get возвращает объект типа AlertArtifact
func (a *AlertArtifact) Get() *AlertArtifact {
	return a
}

func (a *AlertArtifact) GetIoc() bool {
	return a.Ioc
}

// SetValueIoc устанавливает BOOL значение для поля Ioc
func (a *AlertArtifact) SetValueIoc(v bool) {
	a.Ioc = v
}

// SetAnyIoc устанавливает ЛЮБОЕ значение для поля Ioc
func (a *AlertArtifact) SetAnyIoc(i interface{}) {
	if v, ok := i.(bool); ok {
		a.Ioc = v
	}
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

func (a *AlertArtifact) GetTlp() uint64 {
	return a.Tlp
}

// SetValueTlp устанавливает UINT64 значение для поля Tlp
func (a *AlertArtifact) SetValueTlp(v uint64) {
	a.Tlp = v
}

// SetAnyTlp устанавливает ЛЮБОЕ значение для поля Tlp
func (a *AlertArtifact) SetAnyTlp(i interface{}) {
	if v, ok := i.(float64); ok {
		a.Tlp = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		a.Tlp = v
	}
}

func (a *AlertArtifact) GetUnderliningId() string {
	return a.UnderliningId
}

// SetValueUnderliningId устанавливает STRING значение для поля UnderliningId
func (a *AlertArtifact) SetValueUnderliningId(v string) {
	a.UnderliningId = v
}

// SetAnyUnderliningId устанавливает ЛЮБОЕ значение для поля UnderliningId
func (a *AlertArtifact) SetAnyUnderliningId(i interface{}) {
	a.UnderliningId = fmt.Sprint(i)
}

func (a *AlertArtifact) GetId() string {
	return a.Id
}

// SetValueId устанавливает STRING значение для поля Id
func (a *AlertArtifact) SetValueId(v string) {
	a.Id = v
}

// SetAnyId устанавливает ЛЮБОЕ значение для поля Id
func (a *AlertArtifact) SetAnyId(i interface{}) {
	a.Id = fmt.Sprint(i)
}

func (a *AlertArtifact) GetUnderliningType() string {
	return a.UnderliningType
}

// SetValueUnderliningType устанавливает STRING значение для поля UnderliningType
func (a *AlertArtifact) SetValueUnderliningType(v string) {
	a.UnderliningType = v
}

// SetAnyUnderliningType устанавливает ЛЮБОЕ значение для поля UnderliningType
func (a *AlertArtifact) SetAnyUnderliningType(i interface{}) {
	a.UnderliningType = fmt.Sprint(i)
}

func (a *AlertArtifact) GetCreatedAt() string {
	return a.CreatedAt
}

// SetValueCreatedAt устанавливает значение в формате RFC3339 для поля CreatedAt
func (a *AlertArtifact) SetValueCreatedAt(v string) {
	a.CreatedAt = v
}

// SetAnyCreatedAt устанавливает ЛЮБОЕ значение для поля CreatedAt
func (a *AlertArtifact) SetAnyCreatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	a.CreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
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

func (a *AlertArtifact) GetStartDate() string {
	return a.StartDate
}

// SetValueStartDate устанавливает значение  в формате RFC3339 для поля StartDate
func (a *AlertArtifact) SetValueStartDate(v string) {
	a.StartDate = v
}

// SetAnyStartDate устанавливает ЛЮБОЕ значение для поля StartDate
func (a *AlertArtifact) SetAnyStartDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	a.StartDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (a *AlertArtifact) GetCreatedBy() string {
	return a.CreatedBy
}

// SetValueCreatedBy устанавливает STRING значение для поля CreatedBy
func (a *AlertArtifact) SetValueCreatedBy(v string) {
	a.CreatedBy = v
}

// SetAnyCreatedBy устанавливает ЛЮБОЕ значение для поля CreatedBy
func (a *AlertArtifact) SetAnyCreatedBy(i interface{}) {
	a.CreatedBy = fmt.Sprint(i)
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

func (a *AlertArtifact) GetData() string {
	return a.Data
}

// SetValueData устанавливает STRING значение для поля Data
func (a *AlertArtifact) SetValueData(v string) {
	a.Data = v
}

// SetAnyData устанавливает ЛЮБОЕ значение для поля Data
func (a *AlertArtifact) SetAnyData(i interface{}) {
	a.Data = fmt.Sprint(i)
}

func (a *AlertArtifact) GetDataType() string {
	return a.DataType
}

// SetValueDataType устанавливает STRING значение для поля DataType
func (a *AlertArtifact) SetValueDataType(v string) {
	a.DataType = v
}

// SetAnyDataType устанавливает ЛЮБОЕ значение для поля DataType
func (a *AlertArtifact) SetAnyDataType(i interface{}) {
	a.DataType = fmt.Sprint(i)
}

func (a *AlertArtifact) GetMessage() string {
	return a.Message
}

// SetValueMessage устанавливает STRING значение для поля Message
func (a *AlertArtifact) SetValueMessage(v string) {
	a.Message = v
}

// SetAnyMessage устанавливает ЛЮБОЕ значение для поля Message
func (a *AlertArtifact) SetAnyMessage(i interface{}) {
	a.Message = fmt.Sprint(i)
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

	str.WriteString(fmt.Sprintf("%sioc: '%t'\n", ws, a.Ioc))
	str.WriteString(fmt.Sprintf("%ssighted: '%t'\n", ws, a.Sighted))
	str.WriteString(fmt.Sprintf("%signoreSimilarity: '%t'\n", ws, a.IgnoreSimilarity))
	str.WriteString(fmt.Sprintf("%stlp: '%d'\n", ws, a.Tlp))
	str.WriteString(fmt.Sprintf("%sunderliningId: '%s'\n", ws, a.UnderliningId))
	str.WriteString(fmt.Sprintf("%sid: '%s'\n", ws, a.Id))
	str.WriteString(fmt.Sprintf("%sunderliningType: '%s'\n", ws, a.UnderliningType))
	str.WriteString(fmt.Sprintf("%screatedAt: '%s'\n", ws, a.CreatedAt))
	str.WriteString(fmt.Sprintf("%supdatedAt: '%s'\n", ws, a.UpdatedAt))
	str.WriteString(fmt.Sprintf("%sstartDate: '%s'\n", ws, a.StartDate))
	str.WriteString(fmt.Sprintf("%screatedBy: '%s'\n", ws, a.CreatedBy))
	str.WriteString(fmt.Sprintf("%supdatedBy: '%s'\n", ws, a.UpdatedBy))
	str.WriteString(fmt.Sprintf("%sdata: '%s'\n", ws, a.Data))
	str.WriteString(fmt.Sprintf("%sdataType: '%s'\n", ws, a.DataType))
	str.WriteString(fmt.Sprintf("%smessage: '%s'\n", ws, a.Message))
	str.WriteString(fmt.Sprintf("%stags: \n%s", ws, func(l []string) string {
		str := strings.Builder{}
		ws := supportingfunctions.GetWhitespace(num + 1)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d. '%s'\n", ws, k+1, v))
		}
		return str.String()
	}(a.Tags)))

	return str.String()
}
