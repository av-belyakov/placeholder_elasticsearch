package datamodel

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/_moduledatacomparison/supportingfunctions"
)

func (em EventMessageTheHiveCase) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(em.CommonEventType.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'details':\n", ws))
	str.WriteString(em.Details.ToStringBeautiful(num + 1))
	str.WriteString(fmt.Sprintf("%s'object':\n", ws))
	str.WriteString(em.Object.ToStringBeautiful(num + 1))

	return str.String()
}

func (ed EventCaseDetails) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'endDate': '%s'\n", ws, ed.EndDate))
	str.WriteString(fmt.Sprintf("%s'resolutionStatus': '%s'\n", ws, ed.ResolutionStatus))
	str.WriteString(fmt.Sprintf("%s'summary': '%s'\n", ws, ed.Summary))
	str.WriteString(fmt.Sprintf("%s'status': '%s'\n", ws, ed.Status))
	str.WriteString(fmt.Sprintf("%s'impactStatus': '%s'\n", ws, ed.ImpactStatus))
	str.WriteString(fmt.Sprintf("%s'customFields': \n%s", ws, CustomFieldsToStringBeautiful(ed.CustomFields, num)))

	return str.String()
}

func (eo EventCaseObject) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(eo.CommonEventCaseObject.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'tags': \n%s", ws, supportingfunctions.ToStringBeautifulSlice(num, eo.Tags)))
	str.WriteString(fmt.Sprintf("%s'customFields': \n%s", ws, CustomFieldsToStringBeautiful(eo.CustomFields, num)))

	return str.String()
}
