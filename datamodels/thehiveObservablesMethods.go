package datamodels

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/datamodels/commonobservable"
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

func NewObservableMessage() *ObservableMessage {
	return &ObservableMessage{
		CommonObservableType: commonobservable.CommonObservableType{
			UnderliningCreatedAt: "1970-01-01T00:00:00+00:00",
			UnderliningUpdatedAt: "1970-01-01T00:00:00+00:00",
			StartDate:            "1970-01-01T00:00:00+00:00",
		},
		Tags:       []string(nil),
		Attachment: *NewAttachmentData(),
		Reports:    make(map[string]ReportTaxonomies),
	}
}

func (o *ObservableMessage) Get() *ObservableMessage {
	return o
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

	str.WriteString(om.CommonObservableType.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, ToStringBeautifulSlice(num, om.Tags)))
	str.WriteString(fmt.Sprintf("%s'attachment': \n%s", ws, om.Attachment.ToStringBeautiful(num)))
	str.WriteString(fmt.Sprintf("%s'reports': \n%s", ws, func(l map[string]ReportTaxonomies) string {
		var str strings.Builder = strings.Builder{}
		for key, value := range l {
			str.WriteString(fmt.Sprintf("%s'%s':\n", supportingfunctions.GetWhitespace(num+1), key))
			str.WriteString(fmt.Sprintf("%s'taxonomys':\n", supportingfunctions.GetWhitespace(num+2)))
			for k, v := range value.Taxonomies {
				str.WriteString(fmt.Sprintf("%s%d.\n", supportingfunctions.GetWhitespace(num+3), k+1))
				str.WriteString(fmt.Sprintf("%s'level': %v\n", supportingfunctions.GetWhitespace(num+4), v.Level))
				str.WriteString(fmt.Sprintf("%s'namespace': %v\n", supportingfunctions.GetWhitespace(num+4), v.Namespace))
				str.WriteString(fmt.Sprintf("%s'predicate': %v\n", supportingfunctions.GetWhitespace(num+4), v.Predicate))
				str.WriteString(fmt.Sprintf("%s'value': %v\n", supportingfunctions.GetWhitespace(num+4), v.Value))
			}
		}
		return str.String()
	}(om.Reports)))

	return str.String()
}

// ****************** AttachmentData ********************

func NewAttachmentData() *AttachmentData {
	return &AttachmentData{Hashes: []string(nil)}
}

func (a *AttachmentData) GetSize() uint64 {
	return a.Size
}

// SetValueSize устанавливает INT значение для поля Size
func (a *AttachmentData) SetValueSize(v uint64) {
	a.Size = v
}

// SetAnySize устанавливает ЛЮБОЕ значение для поля Size
func (a *AttachmentData) SetAnySize(i interface{}) {
	if v, ok := i.(float32); ok {
		a.Size = uint64(v)

		return
	}

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

func (a AttachmentData) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'size': '%d'\n", ws, a.Size))
	str.WriteString(fmt.Sprintf("%s'id': '%s'\n", ws, a.Id))
	str.WriteString(fmt.Sprintf("%s'name': '%s'\n", ws, a.Name))
	str.WriteString(fmt.Sprintf("%s'contentType': '%s'\n", ws, a.ContentType))
	str.WriteString(fmt.Sprintf("%s'hashes': \n%s", ws, ToStringBeautifulSlice(num, a.Hashes)))

	return str.String()
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
