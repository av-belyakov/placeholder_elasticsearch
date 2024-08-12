package datamodel

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/_moduledatacomparison/supportingfunctions"
)

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
