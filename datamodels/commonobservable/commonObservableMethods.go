package commonobservable

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/supportingfunctions"
)

//****************** CommonObservableType ******************

func (o *CommonObservableType) Get() *CommonObservableType {
	return o
}

func (o *CommonObservableType) GetIoc() bool {
	return o.Ioc
}

// SetValueIoc устанавливает BOOL значение для поля Ioc
func (o *CommonObservableType) SetValueIoc(v bool) {
	o.Ioc = v
}

// SetAnyIoc устанавливает ЛЮБОЕ значение для поля Ioc
func (o *CommonObservableType) SetAnyIoc(i interface{}) {
	if v, ok := i.(bool); ok {
		o.Ioc = v
	}
}

func (o *CommonObservableType) GetSighted() bool {
	return o.Sighted
}

// SetValueSighted устанавливает BOOL значение для поля Sighted
func (o *CommonObservableType) SetValueSighted(v bool) {
	o.Sighted = v
}

// SetAnySighted устанавливает ЛЮБОЕ значение для поля Sighted
func (o *CommonObservableType) SetAnySighted(i interface{}) {
	if v, ok := i.(bool); ok {
		o.Sighted = v
	}
}

func (o *CommonObservableType) GetIgnoreSimilarity() bool {
	return o.IgnoreSimilarity
}

// SetValueIgnoreSimilarity устанавливает BOOL значение для поля IgnoreSimilarity
func (o *CommonObservableType) SetValueIgnoreSimilarity(v bool) {
	o.IgnoreSimilarity = v
}

// SetAnyIgnoreSimilarity устанавливает ЛЮБОЕ значение для поля IgnoreSimilarity
func (o *CommonObservableType) SetAnyIgnoreSimilarity(i interface{}) {
	if v, ok := i.(bool); ok {
		o.IgnoreSimilarity = v
	}
}

func (o *CommonObservableType) GetTlp() uint64 {
	return o.Tlp
}

// SetValueTlp устанавливает INT значение для поля Tlp
func (o *CommonObservableType) SetValueTlp(v uint64) {
	o.Tlp = v
}

// SetAnyTlp устанавливает ЛЮБОЕ значение для поля Tlp
func (o *CommonObservableType) SetAnyTlp(i interface{}) {
	if v, ok := i.(float32); ok {
		o.Tlp = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		o.Tlp = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		o.Tlp = v
	}
}

func (o *CommonObservableType) GetUnderliningCreatedAt() string {
	return o.UnderliningCreatedAt
}

// SetValueUnderliningCreatedAt устанавливает значение в формате RFC3339 для поля CreatedAt
func (o *CommonObservableType) SetValueUnderliningCreatedAt(v string) {
	o.UnderliningCreatedAt = v
}

// SetAnyUnderliningCreatedAt устанавливает ЛЮБОЕ значение для поля CreatedAt
func (o *CommonObservableType) SetAnyUnderliningCreatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	o.UnderliningCreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (o *CommonObservableType) GetUnderliningUpdatedAt() string {
	return o.UnderliningUpdatedAt
}

// SetValueUnderliningUpdatedAt устанавливает значениев формате RFC3339 для поля UpdatedAt
func (o *CommonObservableType) SetValueUnderliningUpdatedAt(v string) {
	o.UnderliningUpdatedAt = v
}

// SetAnyUnderliningUpdatedAt устанавливает ЛЮБОЕ значение для поля UpdatedAt
func (o *CommonObservableType) SetAnyUnderliningUpdatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	o.UnderliningUpdatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (o *CommonObservableType) GetStartDate() string {
	return o.StartDate
}

// SetValueStartDate устанавливает значениев формате RFC3339 для поля StartDate
func (o *CommonObservableType) SetValueStartDate(v string) {
	o.StartDate = v
}

// SetAnyStartDate устанавливает ЛЮБОЕ значение для поля StartDate
func (o *CommonObservableType) SetAnyStartDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	o.StartDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (o *CommonObservableType) GetUnderliningCreatedBy() string {
	return o.UnderliningCreatedBy
}

// SetValueUnderliningCreatedBy устанавливает STRING значение для поля CreatedBy
func (o *CommonObservableType) SetValueUnderliningCreatedBy(v string) {
	o.UnderliningCreatedBy = v
}

// SetAnyUnderliningCreatedBy устанавливает ЛЮБОЕ значение для поля CreatedBy
func (o *CommonObservableType) SetAnyUnderliningCreatedBy(i interface{}) {
	o.UnderliningCreatedBy = fmt.Sprint(i)
}

func (o *CommonObservableType) GetUnderliningUpdatedBy() string {
	return o.UnderliningUpdatedBy
}

// SetValueUnderliningUpdatedBy устанавливает STRING значение для поля UpdatedBy
func (o *CommonObservableType) SetValueUnderliningUpdatedBy(v string) {
	o.UnderliningUpdatedBy = v
}

// SetAnyUnderliningUpdatedBy устанавливает ЛЮБОЕ значение для поля UpdatedBy
func (o *CommonObservableType) SetAnyUnderliningUpdatedBy(i interface{}) {
	o.UnderliningUpdatedBy = fmt.Sprint(i)
}

func (o *CommonObservableType) GetUnderliningId() string {
	return o.UnderliningId
}

// SetValueUnderliningId устанавливает STRING значение для поля UnderliningId
func (o *CommonObservableType) SetValueUnderliningId(v string) {
	o.UnderliningId = v
}

// SetAnyUnderliningId устанавливает ЛЮБОЕ значение для поля UnderliningId
func (o *CommonObservableType) SetAnyUnderliningId(i interface{}) {
	o.UnderliningId = fmt.Sprint(i)
}

func (o *CommonObservableType) GetUnderliningType() string {
	return o.UnderliningType
}

// SetValueUnderliningType устанавливает STRING значение для поля UnderliningType
func (o *CommonObservableType) SetValueUnderliningType(v string) {
	o.UnderliningType = v
}

// SetAnyUnderliningType устанавливает ЛЮБОЕ значение для поля UnderliningType
func (o *CommonObservableType) SetAnyUnderliningType(i interface{}) {
	o.UnderliningType = fmt.Sprint(i)
}

func (o *CommonObservableType) GetData() string {
	return o.Data
}

// SetValueData устанавливает STRING значение для поля Data
func (o *CommonObservableType) SetValueData(v string) {
	o.Data = v
}

// SetAnyData устанавливает ЛЮБОЕ значение для поля Data
func (o *CommonObservableType) SetAnyData(i interface{}) {
	o.Data = fmt.Sprint(i)
}

func (o *CommonObservableType) GetDataType() string {
	return o.DataType
}

// SetValueDataType устанавливает STRING значение для поля DataType
func (o *CommonObservableType) SetValueDataType(v string) {
	o.DataType = v
}

// SetAnyDataType устанавливает ЛЮБОЕ значение для поля DataType
func (o *CommonObservableType) SetAnyDataType(i interface{}) {
	o.DataType = fmt.Sprint(i)
}

func (o *CommonObservableType) GetMessage() string {
	return o.Message
}

// SetValueMessage устанавливает STRING значение для поля Message
func (o *CommonObservableType) SetValueMessage(v string) {
	o.Message = v
}

// SetAnyMessage устанавливает ЛЮБОЕ значение для поля Message
func (o *CommonObservableType) SetAnyMessage(i interface{}) {
	o.Message = fmt.Sprint(i)
}

func (om CommonObservableType) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'_createdAt': '%s'\n", ws, om.UnderliningCreatedAt))
	str.WriteString(fmt.Sprintf("%s'_createdBy': '%s'\n", ws, om.UnderliningCreatedBy))
	str.WriteString(fmt.Sprintf("%s'_id': '%s'\n", ws, om.UnderliningId))
	str.WriteString(fmt.Sprintf("%s'_type': '%s'\n", ws, om.UnderliningType))
	str.WriteString(fmt.Sprintf("%s'_updatedAt': '%s'\n", ws, om.UnderliningUpdatedAt))
	str.WriteString(fmt.Sprintf("%s'_updatedBy': '%s'\n", ws, om.UnderliningUpdatedBy))
	str.WriteString(fmt.Sprintf("%s'data': '%s'\n", ws, om.Data))
	str.WriteString(fmt.Sprintf("%s'dataType': '%s'\n", ws, om.DataType))
	str.WriteString(fmt.Sprintf("%s'ignoreSimilarity': '%v'\n", ws, om.IgnoreSimilarity))
	str.WriteString(fmt.Sprintf("%s'ioc': '%v'\n", ws, om.Ioc))
	str.WriteString(fmt.Sprintf("%s'message': '%s'\n", ws, om.Message))
	str.WriteString(fmt.Sprintf("%s'sighted': '%v'\n", ws, om.Sighted))
	str.WriteString(fmt.Sprintf("%s'startDate': '%s'\n", ws, om.StartDate))
	str.WriteString(fmt.Sprintf("%s'tlp': '%d'\n", ws, om.Tlp))

	return str.String()
}
