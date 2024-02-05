package datamodels

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/supportingfunctions"
)

func NewObservablesMessageTheHive() *ObservablesMessageTheHive {
	return &ObservablesMessageTheHive{}
}

func (o *ObservablesMessageTheHive) SetObservables(list []ObservableMessage) {
	o.Observables = list
}

func (o *ObservablesMessageTheHive) GetObservables() []ObservableMessage {
	return o.Observables
}

func (o *ObservablesMessageTheHive) Set(v ObservableMessage) {
	o.Observables = append(o.Observables, v)
}

func (o *ObservableMessage) GetIoc() bool {
	return o.Ioc
}

// SetValueIoc устанавливает BOOL значение для поля Ioc
func (o *ObservableMessage) SetValueIoc(v bool) {
	o.Ioc = v
}

// SetAnyIoc устанавливает ЛЮБОЕ значение для поля Ioc
func (o *ObservableMessage) SetAnyIoc(i interface{}) {
	if v, ok := i.(bool); ok {
		o.Ioc = v
	}
}

func (o *ObservableMessage) GetSighted() bool {
	return o.Sighted
}

// SetValueSighted устанавливает BOOL значение для поля Sighted
func (o *ObservableMessage) SetValueSighted(v bool) {
	o.Sighted = v
}

// SetAnySighted устанавливает ЛЮБОЕ значение для поля Sighted
func (o *ObservableMessage) SetAnySighted(i interface{}) {
	if v, ok := i.(bool); ok {
		o.Sighted = v
	}
}

func (o *ObservableMessage) GetIgnoreSimilarity() bool {
	return o.IgnoreSimilarity
}

// SetValueIgnoreSimilarity устанавливает BOOL значение для поля IgnoreSimilarity
func (o *ObservableMessage) SetValueIgnoreSimilarity(v bool) {
	o.IgnoreSimilarity = v
}

// SetAnyIgnoreSimilarity устанавливает ЛЮБОЕ значение для поля IgnoreSimilarity
func (o *ObservableMessage) SetAnyIgnoreSimilarity(i interface{}) {
	if v, ok := i.(bool); ok {
		o.IgnoreSimilarity = v
	}
}

func (o *ObservableMessage) GetTlp() uint64 {
	return o.Tlp
}

// SetValueTlp устанавливает INT значение для поля Tlp
func (o *ObservableMessage) SetValueTlp(v uint64) {
	o.Tlp = v
}

// SetAnyTlp устанавливает ЛЮБОЕ значение для поля Tlp
func (o *ObservableMessage) SetAnyTlp(i interface{}) {
	if v, ok := i.(float64); ok {
		o.Tlp = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		o.Tlp = v
	}
}

func (o *ObservableMessage) GetUnderliningCreatedAt() string {
	return o.UnderliningCreatedAt
}

// SetValueUnderliningCreatedAt устанавливает значение в формате RFC3339 для поля CreatedAt
func (o *ObservableMessage) SetValueUnderliningCreatedAt(v string) {
	o.UnderliningCreatedAt = v
}

// SetAnyUnderliningCreatedAt устанавливает ЛЮБОЕ значение для поля CreatedAt
func (o *ObservableMessage) SetAnyUnderliningCreatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	o.UnderliningCreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (o *ObservableMessage) GetUnderliningUpdatedAt() string {
	return o.UnderliningUpdatedAt
}

// SetValueUnderliningUpdatedAt устанавливает значениев формате RFC3339 для поля UpdatedAt
func (o *ObservableMessage) SetValueUnderliningUpdatedAt(v string) {
	o.UnderliningUpdatedAt = v
}

// SetAnyUnderliningUpdatedAt устанавливает ЛЮБОЕ значение для поля UpdatedAt
func (o *ObservableMessage) SetAnyUnderliningUpdatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	o.UnderliningUpdatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (o *ObservableMessage) GetStartDate() string {
	return o.StartDate
}

// SetValueStartDate устанавливает значениев формате RFC3339 для поля StartDate
func (o *ObservableMessage) SetValueStartDate(v string) {
	o.StartDate = v
}

// SetAnyStartDate устанавливает ЛЮБОЕ значение для поля StartDate
func (o *ObservableMessage) SetAnyStartDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	o.StartDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (o *ObservableMessage) GetUnderliningCreatedBy() string {
	return o.UnderliningCreatedBy
}

// SetValueUnderliningCreatedBy устанавливает STRING значение для поля CreatedBy
func (o *ObservableMessage) SetValueUnderliningCreatedBy(v string) {
	o.UnderliningCreatedBy = v
}

// SetAnyUnderliningCreatedBy устанавливает ЛЮБОЕ значение для поля CreatedBy
func (o *ObservableMessage) SetAnyUnderliningCreatedBy(i interface{}) {
	o.UnderliningCreatedBy = fmt.Sprint(i)
}

func (o *ObservableMessage) GetUnderliningUpdatedBy() string {
	return o.UnderliningUpdatedBy
}

// SetValueUnderliningUpdatedBy устанавливает STRING значение для поля UpdatedBy
func (o *ObservableMessage) SetValueUnderliningUpdatedBy(v string) {
	o.UnderliningUpdatedBy = v
}

// SetAnyUnderliningUpdatedBy устанавливает ЛЮБОЕ значение для поля UpdatedBy
func (o *ObservableMessage) SetAnyUnderliningUpdatedBy(i interface{}) {
	o.UnderliningUpdatedBy = fmt.Sprint(i)
}

func (o *ObservableMessage) GetUnderliningId() string {
	return o.UnderliningId
}

// SetValueUnderliningId устанавливает STRING значение для поля UnderliningId
func (o *ObservableMessage) SetValueUnderliningId(v string) {
	o.UnderliningId = v
}

// SetAnyUnderliningId устанавливает ЛЮБОЕ значение для поля UnderliningId
func (o *ObservableMessage) SetAnyUnderliningId(i interface{}) {
	o.UnderliningId = fmt.Sprint(i)
}

func (o *ObservableMessage) GetUnderliningType() string {
	return o.UnderliningType
}

// SetValueUnderliningType устанавливает STRING значение для поля UnderliningType
func (o *ObservableMessage) SetValueUnderliningType(v string) {
	o.UnderliningType = v
}

// SetAnyUnderliningType устанавливает ЛЮБОЕ значение для поля UnderliningType
func (o *ObservableMessage) SetAnyUnderliningType(i interface{}) {
	o.UnderliningType = fmt.Sprint(i)
}

func (o *ObservableMessage) GetData() string {
	return o.Data
}

// SetValueData устанавливает STRING значение для поля Data
func (o *ObservableMessage) SetValueData(v string) {
	o.Data = v
}

// SetAnyData устанавливает ЛЮБОЕ значение для поля Data
func (o *ObservableMessage) SetAnyData(i interface{}) {
	o.Data = fmt.Sprint(i)
}

func (o *ObservableMessage) GetDataType() string {
	return o.DataType
}

// SetValueDataType устанавливает STRING значение для поля DataType
func (o *ObservableMessage) SetValueDataType(v string) {
	o.DataType = v
}

// SetAnyDataType устанавливает ЛЮБОЕ значение для поля DataType
func (o *ObservableMessage) SetAnyDataType(i interface{}) {
	o.DataType = fmt.Sprint(i)
}

func (o *ObservableMessage) GetMessage() string {
	return o.Message
}

// SetValueMessage устанавливает STRING значение для поля Message
func (o *ObservableMessage) SetValueMessage(v string) {
	o.Message = v
}

// SetAnyMessage устанавливает ЛЮБОЕ значение для поля Message
func (o *ObservableMessage) SetAnyMessage(i interface{}) {
	o.Message = fmt.Sprint(i)
}

func (o *ObservableMessage) GetTags() []string {
	return o.Tags
}

// SetValueTags устанавливает STRING значение для поля Tags
func (o *ObservableMessage) SetValueTags(v string) {
	o.Tags = append(o.Tags, v)
}

// SetAnyTags устанавливает ЛЮБОЕ значение для поля Tags
func (o *ObservableMessage) SetAnyTags(i interface{}) {
	o.Tags = append(o.Tags, fmt.Sprint(i))
}

func (o *ObservableMessage) GetAttachment() *AttachmentData {
	return &o.Attachment
}

func (o *ObservableMessage) GetReports() map[string]ReportTaxonomies {
	return o.Reports
}

// SetValueReports устанавливает значение для поля Reports
func (o *ObservableMessage) SetValueReports(v map[string]ReportTaxonomies) {
	o.Reports = v
}

func (om ObservablesMessageTheHive) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num)

	for k, v := range om.Observables {
		str.WriteString(fmt.Sprintf("%s%d.\n", ws, k+1))
		str.WriteString(v.ToStringBeautiful(num + 1))
	}

	return str.String()
}

func (om ObservableMessage) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s_createdAt: '%s'\n", ws, om.UnderliningCreatedAt))
	str.WriteString(fmt.Sprintf("%s_createdBy: '%s'\n", ws, om.UnderliningCreatedBy))
	str.WriteString(fmt.Sprintf("%s_id: '%s'\n", ws, om.UnderliningId))
	str.WriteString(fmt.Sprintf("%s_type: '%s'\n", ws, om.UnderliningType))
	str.WriteString(fmt.Sprintf("%s_updatedAt: '%s'\n", ws, om.UnderliningUpdatedAt))
	str.WriteString(fmt.Sprintf("%s_updatedBy: '%s'\n", ws, om.UnderliningUpdatedBy))
	str.WriteString(fmt.Sprintf("%sdata: '%s'\n", ws, om.Data))
	str.WriteString(fmt.Sprintf("%sdataType: '%s'\n", ws, om.DataType))
	str.WriteString(fmt.Sprintf("%signoreSimilarity: '%v'\n", ws, om.IgnoreSimilarity))
	//данное поле редко используемое, думаю пока оно не требует реализации
	/*str.WriteString(fmt.Sprintf("%sextraData: \n%s", ws, func(l map[string]interface{}) string {
		var str strings.Builder = strings.Builds{}
		ws := supportingfunctions.GetWhitespace(num + 1)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%s: '%v'\n", ws, k, v))
		}
		return str.String()
	}(om.ExtraData)))*/
	str.WriteString(fmt.Sprintf("%sioc: '%v'\n", ws, om.Ioc))
	str.WriteString(fmt.Sprintf("%smessage: '%s'\n", ws, om.Message))
	str.WriteString(fmt.Sprintf("%ssighted: '%v'\n", ws, om.Sighted))
	str.WriteString(fmt.Sprintf("%sstartDate: '%s'\n", ws, om.StartDate))
	str.WriteString(fmt.Sprintf("%stags: \n%s", ws, func(l []string) string {
		var str strings.Builder = strings.Builder{}
		ws := supportingfunctions.GetWhitespace(num + 1)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d. '%s'\n", ws, k+1, v))
		}
		return str.String()
	}(om.Tags)))
	str.WriteString(fmt.Sprintf("%stlp: '%d'\n", ws, om.Tlp))
	str.WriteString(fmt.Sprintf("%sreports: \n%s", ws, func(l map[string]ReportTaxonomies) string {
		var str strings.Builder = strings.Builder{}
		for key, value := range l {
			str.WriteString(fmt.Sprintf("%s%s:\n", supportingfunctions.GetWhitespace(num+1), key))
			str.WriteString(fmt.Sprintf("%staxonomys:\n", supportingfunctions.GetWhitespace(num+2)))
			for k, v := range value.Taxonomies {
				str.WriteString(fmt.Sprintf("%s%d.\n", supportingfunctions.GetWhitespace(num+3), k+1))
				str.WriteString(fmt.Sprintf("%sLevel: %v\n", supportingfunctions.GetWhitespace(num+4), v.Level))
				str.WriteString(fmt.Sprintf("%sNamespace: %v\n", supportingfunctions.GetWhitespace(num+4), v.Namespace))
				str.WriteString(fmt.Sprintf("%sPredicate: %v\n", supportingfunctions.GetWhitespace(num+4), v.Predicate))
				str.WriteString(fmt.Sprintf("%sValue: %v\n", supportingfunctions.GetWhitespace(num+4), v.Value))
			}
		}
		return str.String()
	}(om.Reports)))

	return str.String()
}

// ****************** AttachmentData ********************
func (a *AttachmentData) GetSize() uint64 {
	return a.Size
}

// SetValueSize устанавливает INT значение для поля Size
func (a *AttachmentData) SetValueSize(v uint64) {
	a.Size = v
}

// SetAnySize устанавливает ЛЮБОЕ значение для поля Size
func (a *AttachmentData) SetAnySize(i interface{}) {
	if v, ok := i.(float64); ok {
		a.Size = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		a.Size = v
	}
}

func (a *AttachmentData) GetId() string {
	return a.Id
}

// SetValueId устанавливает STRING значение для поля Id
func (a *AttachmentData) SetValueId(v string) {
	a.Id = v
}

// SetAnyId устанавливает ЛЮБОЕ значение для поля Id
func (a *AttachmentData) SetAnyId(i interface{}) {
	a.Id = fmt.Sprint(i)
}

func (a *AttachmentData) GetName() string {
	return a.Name
}

// SetValueName устанавливает STRING значение для поля Name
func (a *AttachmentData) SetValueName(v string) {
	a.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (a *AttachmentData) SetAnyName(i interface{}) {
	a.Name = fmt.Sprint(i)
}

func (a *AttachmentData) GetContentType() string {
	return a.ContentType
}

// SetValueContentType устанавливает STRING значение для поля ContentType
func (a *AttachmentData) SetValueContentType(v string) {
	a.ContentType = v
}

// SetAnyContentType устанавливает ЛЮБОЕ значение для поля ContentType
func (a *AttachmentData) SetAnyContentType(i interface{}) {
	a.ContentType = fmt.Sprint(i)
}

func (a *AttachmentData) GetHashes() []string {
	return a.Hashes
}

// SetValueHashes устанавливает STRING значение для поля Hashes
func (a *AttachmentData) SetValueHashes(v string) {
	a.Hashes = append(a.Hashes, v)
}

// SetAnyHashes устанавливает ЛЮБОЕ значение для поля Hashes
func (a *AttachmentData) SetAnyHashes(i interface{}) {
	a.Hashes = append(a.Hashes, fmt.Sprint(i))
}

// ********************* ReportTaxonomys *******************
func (t *ReportTaxonomies) GetTaxonomys() []Taxonomy {
	return t.Taxonomies
}

func (t *ReportTaxonomies) GetReportTaxonomys() ReportTaxonomies {
	return *t
}

func (t *ReportTaxonomies) AddTaxonomy(taxonomy Taxonomy) {
	t.Taxonomies = append(t.Taxonomies, taxonomy)
}

// *********************** Taxonomy ************************
func (t *Taxonomy) GetLevel() string {
	return t.Level
}

// SetValueLevel устанавливает STRING значение для поля Level
func (t *Taxonomy) SetValueLevel(v string) {
	t.Level = v
}

// SetAnyLevel устанавливает ЛЮБОЕ значение для поля Level
func (t *Taxonomy) SetAnyLevel(i interface{}) {
	t.Level = fmt.Sprint(i)
}

func (t *Taxonomy) GetNamespace() string {
	return t.Namespace
}

// SetValueNamespace устанавливает STRING значение для поля Namespace
func (t *Taxonomy) SetValueNamespace(v string) {
	t.Namespace = v
}

// SetAnyNamespace устанавливает ЛЮБОЕ значение для поля Namespace
func (t *Taxonomy) SetAnyNamespace(i interface{}) {
	t.Namespace = fmt.Sprint(i)
}

func (t *Taxonomy) GetPredicate() string {
	return t.Predicate
}

// SetValuePredicate устанавливает STRING значение для поля Predicate
func (t *Taxonomy) SetValuePredicate(v string) {
	t.Predicate = v
}

// SetAnyPredicate устанавливает ЛЮБОЕ значение для поля Predicate
func (t *Taxonomy) SetAnyPredicate(i interface{}) {
	t.Predicate = fmt.Sprint(i)
}

func (t *Taxonomy) GetValue() string {
	return t.Value
}

// SetValueValue устанавливает STRING значение для поля Value
func (t *Taxonomy) SetValueValue(v string) {
	t.Value = v
}

// SetAnyValue устанавливает ЛЮБОЕ значение для поля Value
func (t *Taxonomy) SetAnyValue(i interface{}) {
	t.Value = fmt.Sprint(i)
}
