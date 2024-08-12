package datamodel

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/_moduledatacomparison/supportingfunctions"
)

func (tm TtpsMessageTheHive) ToStringBeautiful(num int) string {
	return fmt.Sprintf("%s'ttp': \n%s", supportingfunctions.GetWhitespace(num), func(l []TtpMessage) string {
		var str strings.Builder = strings.Builder{}
		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d.\n", supportingfunctions.GetWhitespace(num+1), k+1))
			str.WriteString(v.ToStringBeautiful(num + 2))
		}

		return str.String()
	}(tm.Ttp))
}

func (tm TtpMessage) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'_createdAt': '%s'\n", ws, tm.UnderliningCreatedAt))
	str.WriteString(fmt.Sprintf("%s'_createdBy': '%s'\n", ws, tm.UnderliningCreatedBy))
	str.WriteString(fmt.Sprintf("%s'_id': '%s'\n", ws, tm.UnderliningId))
	str.WriteString(fmt.Sprintf("%s'occurDate': '%s'\n", ws, tm.OccurDate))
	str.WriteString(fmt.Sprintf("%s'patternId': '%s'\n", ws, tm.PatternId))
	str.WriteString(fmt.Sprintf("%s'tactic': '%s'\n", ws, tm.Tactic))
	str.WriteString(fmt.Sprintf("%s'extraData':\n", ws))
	str.WriteString(tm.ExtraData.ToStringBeautiful(num + 1))

	return str.String()
}

func (edtm ExtraDataTtpMessage) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'pattern':\n", ws))
	str.WriteString(edtm.Pattern.ToStringBeautiful(num + 1))
	str.WriteString(fmt.Sprintf("%s'patternParent':\n", ws))
	str.WriteString(edtm.PatternParent.ToStringBeautiful(num + 1))

	return str.String()
}

func (ped PatternExtraData) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'_createdAt': '%s'\n", ws, ped.UnderliningCreatedAt))
	str.WriteString(fmt.Sprintf("%s'_createdBy': '%s'\n", ws, ped.UnderliningCreatedBy))
	str.WriteString(fmt.Sprintf("%s'_id': '%s'\n", ws, ped.UnderliningId))
	str.WriteString(fmt.Sprintf("%s'_type': '%s'\n", ws, ped.UnderliningType))
	str.WriteString(fmt.Sprintf("%s'dataSources': \n%v", ws, func(l []string) string {
		var str strings.Builder = strings.Builder{}
		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d. '%s'\n", supportingfunctions.GetWhitespace(num+1), k+1, v))
		}

		return str.String()
	}(ped.DataSources)))
	/*str.WriteString(fmt.Sprintf("%sdefenseBypassed: \n%v", ws, func(l []string) string {
		var str strings.Builder = strings.Builder{}
		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d. '%s'\n", supportingfunctions.GetWhitespace(num+1), k+1, v))
		}
		return str.String()
	}(ped.DefenseBypassed)))*/
	str.WriteString(fmt.Sprintf("%s'description': '%s'\n", ws, ped.Description))
	/*str.WriteString(fmt.Sprintf("%sextraData: \n%s", ws, func(l map[string]interface{}) string {
		var str strings.Builder = string.Builder{}
		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%s: '%v'\n", supportingfunctions.GetWhitespace(num+1), k, v))
		}
		return str
	}(ped.ExtraData)))*/
	str.WriteString(fmt.Sprintf("%s'name': '%s'\n", ws, ped.Name))
	str.WriteString(fmt.Sprintf("%s'patternId': '%s'\n", ws, ped.PatternId))
	str.WriteString(fmt.Sprintf("%s'patternType': '%s'\n", ws, ped.PatternType))
	str.WriteString(fmt.Sprintf("%s'permissionsRequired': \n%s", ws, func(l []string) string {
		var str strings.Builder = strings.Builder{}
		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d. '%s'\n", supportingfunctions.GetWhitespace(num+1), k+1, v))
		}

		return str.String()
	}(ped.PermissionsRequired)))
	str.WriteString(fmt.Sprintf("%s'platforms': \n%s", ws, func(l []string) string {
		var str strings.Builder = strings.Builder{}
		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d. '%s'\n", supportingfunctions.GetWhitespace(num+1), k+1, v))
		}

		return str.String()
	}(ped.Platforms)))
	str.WriteString(fmt.Sprintf("%s'remoteSupport': '%v'\n", ws, ped.RemoteSupport))
	str.WriteString(fmt.Sprintf("%s'revoked': '%v'\n", ws, ped.Revoked))
	/*str.WriteString(fmt.Sprintf("%ssystemRequirements: \n%s", ws, func(l []string) string {
		var str strings.Builder = strings.Builder()
		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d. '%s'\n", supportingfunctions.GetWhitespace(num+1), k+1, v))
		}
		return str.String()
	}(ped.SystemRequirements)))*/
	str.WriteString(fmt.Sprintf("%s'tactics': \n%s", ws, func(l []string) string {
		var str strings.Builder = strings.Builder{}
		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s%d. '%s'\n", supportingfunctions.GetWhitespace(num+1), k+1, v))
		}

		return str.String()
	}(ped.Tactics)))
	str.WriteString(fmt.Sprintf("%s'url': '%s'\n", ws, ped.URL))
	str.WriteString(fmt.Sprintf("%s'version': '%s'\n", ws, ped.Version))

	return str.String()
}
