package datamodel

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/_moduledatacomparison/supportingfunctions"
)

func (om ObservablesMessageTheHive) ToStringBeautiful(num int) string {
	return fmt.Sprintf("%s'observables': \n%s", supportingfunctions.GetWhitespace(num), func(l []ObservableMessage) string {
		var str strings.Builder = strings.Builder{}
		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d.\n", supportingfunctions.GetWhitespace(num+1), k+1))
			str.WriteString(v.ToStringBeautiful(num + 2))
		}

		return str.String()
	}(om.Observables))
}

func (om ObservableMessage) ToStringBeautiful(num int) string {
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
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, om.Tags)))
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
