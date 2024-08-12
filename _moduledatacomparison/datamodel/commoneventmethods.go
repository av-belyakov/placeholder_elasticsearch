package datamodel

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/_moduledatacomparison/supportingfunctions"
)

func (em CommonEventType) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'operation': '%s'\n", ws, em.Operation))
	str.WriteString(fmt.Sprintf("%s'objectId': '%s'\n", ws, em.ObjectId))
	str.WriteString(fmt.Sprintf("%s'objectType': '%s'\n", ws, em.ObjectType))
	str.WriteString(fmt.Sprintf("%s'base': '%v'\n", ws, em.Base))
	str.WriteString(fmt.Sprintf("%s'startDate': '%s'\n", ws, em.StartDate))
	str.WriteString(fmt.Sprintf("%s'rootId': '%s'\n", ws, em.RootId))
	str.WriteString(fmt.Sprintf("%s'requestId': '%s'\n", ws, em.RequestId))
	str.WriteString(fmt.Sprintf("%s'organisationId': '%s'\n", ws, em.OrganisationId))
	str.WriteString(fmt.Sprintf("%s'organisation': '%s'\n", ws, em.Organisation))

	return str.String()
}

func (eo CommonEventCaseObject) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'_id': '%s'\n", ws, eo.UnderliningId))
	str.WriteString(fmt.Sprintf("%s'id': '%s'\n", ws, eo.Id))
	str.WriteString(fmt.Sprintf("%s'createdBy': '%s'\n", ws, eo.CreatedBy))
	str.WriteString(fmt.Sprintf("%s'updatedBy': '%s'\n", ws, eo.UpdatedBy))
	str.WriteString(fmt.Sprintf("%s'createdAt': '%s'\n", ws, eo.CreatedAt))
	str.WriteString(fmt.Sprintf("%s'updatedAt': '%s'\n", ws, eo.UpdatedAt))
	str.WriteString(fmt.Sprintf("%s'_type': '%s'\n", ws, eo.UnderliningType))
	str.WriteString(fmt.Sprintf("%s'caseId': '%d'\n", ws, eo.CaseId))
	str.WriteString(fmt.Sprintf("%s'title': '%s'\n", ws, eo.Title))
	str.WriteString(fmt.Sprintf("%s'description': '%s'\n", ws, eo.Description))
	str.WriteString(fmt.Sprintf("%s'severity': '%d'\n", ws, eo.Severity))
	str.WriteString(fmt.Sprintf("%s'startDate': '%s'\n", ws, eo.StartDate))
	str.WriteString(fmt.Sprintf("%s'endDate': '%s'\n", ws, eo.EndDate))
	str.WriteString(fmt.Sprintf("%s'impactStatus': '%s'\n", ws, eo.ImpactStatus))
	str.WriteString(fmt.Sprintf("%s'resolutionStatus': '%s'\n", ws, eo.ResolutionStatus))
	str.WriteString(fmt.Sprintf("%s'flag': '%v'\n", ws, eo.Flag))
	str.WriteString(fmt.Sprintf("%s'tlp': '%d'\n", ws, eo.Tlp))
	str.WriteString(fmt.Sprintf("%s'pap': '%d'\n", ws, eo.Pap))
	str.WriteString(fmt.Sprintf("%s'status': '%s'\n", ws, eo.Status))
	str.WriteString(fmt.Sprintf("%s'summary': '%s'\n", ws, eo.Summary))
	str.WriteString(fmt.Sprintf("%s'owner': '%s'\n", ws, eo.Owner))

	return str.String()
}
