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
