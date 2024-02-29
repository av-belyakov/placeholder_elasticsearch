package commonalertartifact

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/supportingfunctions"
)

//****************** CommonArtifactType ******************

func (a *CommonArtifactType) Get() *CommonArtifactType {
	return a
}

func (a *CommonArtifactType) GetIoc() bool {
	return a.Ioc
}

// SetValueIoc устанавливает BOOL значение для поля Ioc
func (a *CommonArtifactType) SetValueIoc(v bool) {
	a.Ioc = v
}

// SetAnyIoc устанавливает ЛЮБОЕ значение для поля Ioc
func (a *CommonArtifactType) SetAnyIoc(i interface{}) {
	if v, ok := i.(bool); ok {
		a.Ioc = v
	}
}

func (a *CommonArtifactType) GetTlp() uint64 {
	return a.Tlp
}

// SetValueTlp устанавливает UINT64 значение для поля Tlp
func (a *CommonArtifactType) SetValueTlp(v uint64) {
	a.Tlp = v
}

// SetAnyTlp устанавливает ЛЮБОЕ значение для поля Tlp
func (a *CommonArtifactType) SetAnyTlp(i interface{}) {
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

func (a *CommonArtifactType) GetUnderliningId() string {
	return a.UnderliningId
}

// SetValueUnderliningId устанавливает STRING значение для поля UnderliningId
func (a *CommonArtifactType) SetValueUnderliningId(v string) {
	a.UnderliningId = v
}

// SetAnyUnderliningId устанавливает ЛЮБОЕ значение для поля UnderliningId
func (a *CommonArtifactType) SetAnyUnderliningId(i interface{}) {
	a.UnderliningId = fmt.Sprint(i)
}

func (a *CommonArtifactType) GetId() string {
	return a.Id
}

// SetValueId устанавливает STRING значение для поля Id
func (a *CommonArtifactType) SetValueId(v string) {
	a.Id = v
}

// SetAnyId устанавливает ЛЮБОЕ значение для поля Id
func (a *CommonArtifactType) SetAnyId(i interface{}) {
	a.Id = fmt.Sprint(i)
}

func (a *CommonArtifactType) GetUnderliningType() string {
	return a.UnderliningType
}

// SetValueUnderliningType устанавливает STRING значение для поля UnderliningType
func (a *CommonArtifactType) SetValueUnderliningType(v string) {
	a.UnderliningType = v
}

// SetAnyUnderliningType устанавливает ЛЮБОЕ значение для поля UnderliningType
func (a *CommonArtifactType) SetAnyUnderliningType(i interface{}) {
	a.UnderliningType = fmt.Sprint(i)
}

func (a *CommonArtifactType) GetCreatedAt() string {
	return a.CreatedAt
}

// SetValueCreatedAt устанавливает значение в формате RFC3339 для поля CreatedAt
func (a *CommonArtifactType) SetValueCreatedAt(v string) {
	a.CreatedAt = v
}

// SetAnyCreatedAt устанавливает ЛЮБОЕ значение для поля CreatedAt
func (a *CommonArtifactType) SetAnyCreatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	a.CreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (a *CommonArtifactType) GetStartDate() string {
	return a.StartDate
}

// SetValueStartDate устанавливает значение  в формате RFC3339 для поля StartDate
func (a *CommonArtifactType) SetValueStartDate(v string) {
	a.StartDate = v
}

// SetAnyStartDate устанавливает ЛЮБОЕ значение для поля StartDate
func (a *CommonArtifactType) SetAnyStartDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	a.StartDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (a *CommonArtifactType) GetCreatedBy() string {
	return a.CreatedBy
}

// SetValueCreatedBy устанавливает STRING значение для поля CreatedBy
func (a *CommonArtifactType) SetValueCreatedBy(v string) {
	a.CreatedBy = v
}

// SetAnyCreatedBy устанавливает ЛЮБОЕ значение для поля CreatedBy
func (a *CommonArtifactType) SetAnyCreatedBy(i interface{}) {
	a.CreatedBy = fmt.Sprint(i)
}

func (a *CommonArtifactType) GetData() string {
	return a.Data
}

// SetValueData устанавливает STRING значение для поля Data
func (a *CommonArtifactType) SetValueData(v string) {
	a.Data = v
}

// SetAnyData устанавливает ЛЮБОЕ значение для поля Data
func (a *CommonArtifactType) SetAnyData(i interface{}) {
	a.Data = fmt.Sprint(i)
}

func (a *CommonArtifactType) GetDataType() string {
	return a.DataType
}

// SetValueDataType устанавливает STRING значение для поля DataType
func (a *CommonArtifactType) SetValueDataType(v string) {
	a.DataType = v
}

// SetAnyDataType устанавливает ЛЮБОЕ значение для поля DataType
func (a *CommonArtifactType) SetAnyDataType(i interface{}) {
	a.DataType = fmt.Sprint(i)
}

func (a *CommonArtifactType) GetMessage() string {
	return a.Message
}

// SetValueMessage устанавливает STRING значение для поля Message
func (a *CommonArtifactType) SetValueMessage(v string) {
	a.Message = v
}

// SetAnyMessage устанавливает ЛЮБОЕ значение для поля Message
func (a *CommonArtifactType) SetAnyMessage(i interface{}) {
	a.Message = fmt.Sprint(i)
}

func (a *CommonArtifactType) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'ioc': '%t'\n", ws, a.Ioc))
	str.WriteString(fmt.Sprintf("%s'tlp': '%d'\n", ws, a.Tlp))
	str.WriteString(fmt.Sprintf("%s'underliningId': '%s'\n", ws, a.UnderliningId))
	str.WriteString(fmt.Sprintf("%s'id': '%s'\n", ws, a.Id))
	str.WriteString(fmt.Sprintf("%s'underliningType': '%s'\n", ws, a.UnderliningType))
	str.WriteString(fmt.Sprintf("%s'createdAt': '%s'\n", ws, a.CreatedAt))
	str.WriteString(fmt.Sprintf("%s'startDate': '%s'\n", ws, a.StartDate))
	str.WriteString(fmt.Sprintf("%s'createdBy': '%s'\n", ws, a.CreatedBy))
	str.WriteString(fmt.Sprintf("%s'data': '%s'\n", ws, a.Data))
	str.WriteString(fmt.Sprintf("%s'dataType': '%s'\n", ws, a.DataType))
	str.WriteString(fmt.Sprintf("%s'message': '%s'\n", ws, a.Message))

	return str.String()
}
