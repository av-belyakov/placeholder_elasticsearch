package datamodels

import (
	"fmt"
	"strconv"
	"strings"

	"placeholder_elasticsearch/datamodels/commonobservable"
	"placeholder_elasticsearch/supportingfunctions"
)

func NewObservablesMessageEs() *ObservablesMessageEs {
	return &ObservablesMessageEs{
		Observables: make(map[string][]ObservableMessageEs),
	}
}

func (o *ObservablesMessageEs) GetObservables() map[string][]ObservableMessageEs {
	return o.Observables
}

func (o *ObservablesMessageEs) GetKeyObservables(k string) ([]ObservableMessageEs, bool) {
	if value, ok := o.Observables[k]; ok {
		return value, true
	}

	return nil, false
}

func (o *ObservablesMessageEs) SetKeyObservables(k string, observables []ObservableMessageEs) {
	o.Observables[k] = observables
}

// SetObservables устанавливает значение для поля Observables
func (o *ObservablesMessageEs) SetValueObservables(v map[string][]ObservableMessageEs) {
	o.Observables = v
}

// AddValueObservable устанавливает значение для поля Observable
func (o *ObservablesMessageEs) AddValueObservable(k string, v ObservableMessageEs) {
	if _, ok := o.Observables[k]; !ok {
		o.Observables[k] = []ObservableMessageEs(nil)
	}

	o.Observables[k] = append(o.Observables[k], v)
}

func (om ObservablesMessageEs) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}
	str.WriteString(fmt.Sprintf("%s'observables': \n", supportingfunctions.GetWhitespace(num)))

	for key, value := range om.Observables {
		str.WriteString(fmt.Sprintf("%s%s:\n", supportingfunctions.GetWhitespace(num+1), key))
		for k, v := range value {
			str.WriteString(fmt.Sprintf("%s%d.\n", supportingfunctions.GetWhitespace(num+2), k))
			str.WriteString(v.ToStringBeautiful(num + 3))
		}
	}

	return str.String()
}

func NewObservableMessageEs() *ObservableMessageEs {
	return &ObservableMessageEs{
		CommonObservableType: commonobservable.CommonObservableType{
			UnderliningCreatedAt: "1970-01-01T00:00:00+00:00",
			UnderliningUpdatedAt: "1970-01-01T00:00:00+00:00",
			StartDate:            "1970-01-01T00:00:00+00:00",
		},
		Tags:       make(map[string][]string),
		TagsAll:    []string(nil),
		Attachment: *NewAttachmentData(),
		//Reports:    make(map[string]ReportTaxonomies),
	}
}

func (o *ObservableMessageEs) Get() *ObservableMessageEs {
	return o
}

func (o *ObservableMessageEs) GetSensorId() string {
	return o.SensorId
}

// SetSensorId добаляет значение в SensorId
func (o *ObservableMessageEs) SetValueSensorId(v string) {
	o.SensorId = v
}

// SetSensorId добавляет ЛЮБОЕ значение в SensorId
func (o *ObservableMessageEs) SetAnySensorId(i interface{}) {
	o.SensorId = fmt.Sprint(i)
}

func (o *ObservableMessageEs) GetSnortSid() []string {
	return o.SnortSid
}

// SetValueSnortSid добавляет значение STRING в список поля SnortSid
func (o *ObservableMessageEs) SetValueSnortSid(v string) {
	if o.SnortSid == nil {
		o.SnortSid = []string(nil)
	}

	o.SnortSid = append(o.SnortSid, v)
}

func (o *ObservableMessageEs) GetSnortSidNumber() []int {
	return o.SnortSidNumber
}

// SetAnySnortSid добавляет ЛЮБОЕ значение в список поля SnortSid
func (o *ObservableMessageEs) SetAnySnortSid(i interface{}) {
	o.SetValueSnortSid(fmt.Sprint(i))
}

// SetValueSnortSidNumber добавляет значение INT в список поля SnortSidNumber
func (o *ObservableMessageEs) SetValueSnortSidNumber(v int) {
	if o.SnortSidNumber == nil {
		o.SnortSidNumber = []int(nil)
	}

	o.SnortSidNumber = append(o.SnortSidNumber, v)
}

// SetAnySnortSidNumber добавляет ЛЮБОЕ значение в список поля SnortSidNumber
func (o *ObservableMessageEs) SetAnySnortSidNumber(i interface{}) {
	str := fmt.Sprint(i)
	if num, err := strconv.ParseUint(str, 10, 32); err == nil {
		o.SetValueSnortSidNumber(int(num))
	}
}

func (o *ObservableMessageEs) GetTags() map[string][]string {
	return o.Tags
}

// SetValueTags добаляет значение в Tags по ключу
func (o *ObservableMessageEs) SetValueTags(k, v string) bool {
	if o.Tags == nil {
		o.Tags = make(map[string][]string)
	}

	if _, ok := o.Tags[k]; !ok {
		o.Tags[k] = []string(nil)
	}

	for _, value := range o.Tags[k] {
		if v == value {
			return false
		}
	}

	o.Tags[k] = append(o.Tags[k], v)

	return true
}

// SetAnyTags устанавливает ЛЮБОЕ значение для поля Tags
func (o *ObservableMessageEs) SetAnyTags(k string, i interface{}) bool {
	return o.SetValueTags(k, fmt.Sprint(i))
}

func (o *ObservableMessageEs) GetTagsAll() []string {
	return o.TagsAll
}

// SetValueTags устанавливает STRING значение для поля TagsAll
func (o *ObservableMessageEs) SetValueTagsAll(v string) {
	if o.TagsAll == nil {
		o.TagsAll = []string(nil)
	}

	o.TagsAll = append(o.TagsAll, v)
}

// SetAnyTagsAll устанавливает ЛЮБОЕ значение для поля Tags
func (o *ObservableMessageEs) SetAnyTagsAll(i interface{}) {
	o.SetValueTagsAll(fmt.Sprint(i))
}

func (o *ObservableMessageEs) GetAttachment() *AttachmentData {
	return &o.Attachment
}

/*
ИСКЛЮЧИЛ из-за черезмерно большого количества
 полей которое влечет за собой превышения лимита маппинга в Elsticsearch
 что выражается в ошибке от СУБД типа "Limit of total fields [2000] has been exceeded while adding new fields"
func (o *ObservableMessageEs) GetReports() map[string]ReportTaxonomies {
	return o.Reports
}

// GetTaxonomies возвращает список Taxonomy по заданному ключу
func (o *ObservableMessageEs) GetTaxonomies(key string) (*ReportTaxonomies, bool) {
	if res, ok := o.Reports[key]; ok {
		return &res, true
	}

	return &ReportTaxonomies{}, false
}

// SetValueReports устанавливает значение для поля Reports
func (o *ObservableMessageEs) SetValueReports(v map[string]ReportTaxonomies) {
	o.Reports = v
}

// AddValueReports добавляет список значений ReportTaxonomies
func (o *ObservableMessageEs) AddValueReports(key string, rt ReportTaxonomies) {
	o.Reports[key] = rt
}*/

func (om ObservableMessageEs) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(om.CommonObservableType.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'sensorId': '%s'\n", ws, om.SensorId))
	str.WriteString(fmt.Sprintf("%s'snortSid': \n%s", ws, ToStringBeautifulSlice(num, om.SnortSid)))
	str.WriteString(fmt.Sprintf("%s'snortSidNumber': \n%s", ws, ToStringBeautifulSlice(num, om.SnortSidNumber)))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, ToStringBeautifulMapSlice(num, om.Tags)))
	str.WriteString(fmt.Sprintf("%s'tagsAll': \n%s", ws, ToStringBeautifulSlice(num, om.TagsAll)))
	str.WriteString(fmt.Sprintf("%s'attachment': \n%s", ws, om.Attachment.ToStringBeautiful(num+1)))
	/*str.WriteString(fmt.Sprintf("%s'reports': \n%s", ws, func(l map[string]ReportTaxonomies) string {
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
	}(om.Reports)))*/

	return str.String()
}
