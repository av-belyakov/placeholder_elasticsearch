package datamodels

import (
	"fmt"

	"placeholder_elasticsearch/supportingfunctions"
)

func NewTtpsMessageTheHive() *TtpsMessageTheHive {
	return &TtpsMessageTheHive{}
}

func (ttps *TtpsMessageTheHive) SetTtps(list []TtpMessage) {
	ttps.Ttp = list
}

func (ttps *TtpsMessageTheHive) GetTtps() []TtpMessage {
	return ttps.Ttp
}

func (ttps *TtpsMessageTheHive) Set(v TtpMessage) {
	ttps.Ttp = append(ttps.Ttp, v)
}

func (ttpm *TtpMessage) GetOccurDate() string {
	return ttpm.OccurDate
}

// SetValueOccurDate устанавливает дату в формате RFC3339
func (ttpm *TtpMessage) SetValueOccurDate(v string) {
	ttpm.OccurDate = v
}

// SetAnyOccurDate устанавливает ЛЮБОЕ значение для поля OccurDate
func (ttpm *TtpMessage) SetAnyOccurDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	ttpm.OccurDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (ttpm *TtpMessage) GetUnderliningCreatedAt() string {
	return ttpm.UnderliningCreatedAt
}

// SetValueUnderliningCreatedAt устанавливает дату в формате RFC3339
func (ttpm *TtpMessage) SetValueUnderliningCreatedAt(v string) {
	ttpm.UnderliningCreatedAt = v
}

// SetAnyUnderliningCreatedAt устанавливает ЛЮБОЕ значение для поля UnderliningCreatedAt
func (ttpm *TtpMessage) SetAnyUnderliningCreatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	ttpm.UnderliningCreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (ttpm *TtpMessage) GetUnderliningId() string {
	return ttpm.UnderliningId
}

// SetValueUnderliningId устанавливает значение поля UnderliningId
func (ttpm *TtpMessage) SetValueUnderliningId(v string) {
	ttpm.UnderliningId = v
}

// SetAnyUnderliningId устанавливает ЛЮБОЕ значение для поля UnderliningId
func (ttpm *TtpMessage) SetAnyUnderliningId(i interface{}) {
	ttpm.UnderliningId = fmt.Sprint(i)
}

func (ttpm *TtpMessage) GetUnderliningCreatedBy() string {
	return ttpm.UnderliningCreatedBy
}

// SetValueUnderliningCreatedBy устанавливает значение поля CreatedBy
func (ttpm *TtpMessage) SetValueUnderliningCreatedBy(v string) {
	ttpm.UnderliningCreatedBy = v
}

// SetAnyUnderliningCreatedBy устанавливает ЛЮБОЕ значение для поля UnderliningCreatedBy
func (ttpm *TtpMessage) SetAnyUnderliningCreatedBy(i interface{}) {
	ttpm.UnderliningCreatedBy = fmt.Sprint(i)
}

func (ttpm *TtpMessage) GetPatternId() string {
	return ttpm.PatternId
}

// SetValuePatternId устанавливает значение поля PatternId
func (ttpm *TtpMessage) SetValuePatternId(v string) {
	ttpm.PatternId = v
}

// SetAnyPatternId устанавливает ЛЮБОЕ значение для поля PatternId
func (ttpm *TtpMessage) SetAnyPatternId(i interface{}) {
	ttpm.PatternId = fmt.Sprint(i)
}

func (ttpm *TtpMessage) GetTactic() string {
	return ttpm.Tactic
}

// SetValueTactic устанавливает значение поля Tactic
func (ttpm *TtpMessage) SetValueTactic(v string) {
	ttpm.Tactic = v
}

// SetAnyTactic устанавливает ЛЮБОЕ значение для поля Tactic
func (ttpm *TtpMessage) SetAnyTactic(i interface{}) {
	ttpm.Tactic = fmt.Sprint(i)
}

func (tm TtpsMessageTheHive) ToStringBeautiful(num int) string {
	return fmt.Sprintf("%sttp: \n%s", supportingfunctions.GetWhitespace(num), func(l []TtpMessage) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("%s%d.\n", supportingfunctions.GetWhitespace(num+1), k+1)
			str += v.ToStringBeautiful(num + 2)
		}
		return str
	}(tm.Ttp))
}

func (tm TtpMessage) ToStringBeautiful(num int) string {
	var str string

	ws := supportingfunctions.GetWhitespace(num)

	str += fmt.Sprintf("%s_createdAt: '%s'\n", ws, tm.UnderliningCreatedAt)
	str += fmt.Sprintf("%s_createdBy: '%s'\n", ws, tm.UnderliningCreatedBy)
	str += fmt.Sprintf("%s_id: '%s'\n", ws, tm.UnderliningId)
	str += fmt.Sprintf("%soccurDate: '%s'\n", ws, tm.OccurDate)
	str += fmt.Sprintf("%spatternId: '%s'\n", ws, tm.PatternId)
	str += fmt.Sprintf("%stactic: '%s'\n", ws, tm.Tactic)
	str += fmt.Sprintf("%sextraData:\n", ws)
	str += tm.ExtraData.ToStringBeautiful(num + 1)

	return str
}

func (edtm ExtraDataTtpMessage) ToStringBeautiful(num int) string {
	var str string

	ws := supportingfunctions.GetWhitespace(num)

	str += fmt.Sprintf("%sPattern:\n", ws)
	str += edtm.Pattern.ToStringBeautiful(num + 1)
	str += fmt.Sprintf("%sPatternParent:\n", ws)
	str += edtm.PatternParent.ToStringBeautiful(num + 1)

	return str
}

func (ped *PatternExtraData) GetRemoteSupport() bool {
	return ped.RemoteSupport
}

// SetValueRemoteSupport устанавливает BOOL значение для поля RemoteSupport
func (ped *PatternExtraData) SetValueRemoteSupport(v bool) {
	ped.RemoteSupport = v
}

// SetAnyRemoteSupport устанавливает ЛЮБОЕ значение для поля RemoteSupport
func (ped *PatternExtraData) SetAnyRemoteSupport(i interface{}) {
	if v, ok := i.(bool); ok {
		ped.RemoteSupport = v
	}
}

func (ped *PatternExtraData) GetRevoked() bool {
	return ped.Revoked
}

// SetValueRevoked устанавливает BOOL значение для поля Revoked
func (ped *PatternExtraData) SetValueRevoked(v bool) {
	ped.Revoked = v
}

// SetAnyRemoteSupport устанавливает ЛЮБОЕ значение для поля Revoked
func (ped *PatternExtraData) SetAnyRevoked(i interface{}) {
	if v, ok := i.(bool); ok {
		ped.Revoked = v
	}
}

func (ped *PatternExtraData) GetUnderliningCreatedAt() string {
	return ped.UnderliningCreatedAt
}

// SetValueUnderliningCreatedAt устанавливает дату в формате RFC3339
func (ped *PatternExtraData) SetValueUnderliningCreatedAt(v string) {
	ped.UnderliningCreatedAt = v
}

// SetAnyUnderliningCreatedAt устанавливает ЛЮБОЕ значение для поля UnderliningCreatedAt
func (ped *PatternExtraData) SetAnyUnderliningCreatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	ped.UnderliningCreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (ped *PatternExtraData) GetUnderliningCreatedBy() string {
	return ped.UnderliningCreatedBy
}

// SetValueUnderliningCreatedBy устанавливает значение поля UnderliningCreatedBy
func (ped *PatternExtraData) SetValueUnderliningCreatedBy(v string) {
	ped.UnderliningCreatedBy = v
}

// SetAnyUnderliningCreatedBy устанавливает ЛЮБОЕ значение для поля UnderliningCreatedBy
func (ped *PatternExtraData) SetAnyUnderliningCreatedBy(i interface{}) {
	ped.UnderliningCreatedBy = fmt.Sprint(i)
}

func (ped *PatternExtraData) GetUnderliningId() string {
	return ped.UnderliningId
}

// SetValueUnderliningId устанавливает значение поля UnderliningId
func (ped *PatternExtraData) SetValueUnderliningId(v string) {
	ped.UnderliningId = v
}

// SetAnyUnderliningId устанавливает ЛЮБОЕ значение для поля UnderliningId
func (ped *PatternExtraData) SetAnyUnderliningId(i interface{}) {
	ped.UnderliningId = fmt.Sprint(i)
}

func (ped *PatternExtraData) GetUnderliningType() string {
	return ped.UnderliningType
}

// SetValueUnderliningType устанавливает значение поля UnderliningType
func (ped *PatternExtraData) SetValueUnderliningType(v string) {
	ped.UnderliningType = v
}

// SetAnyUnderliningType устанавливает ЛЮБОЕ значение для поля UnderliningType
func (ped *PatternExtraData) SetAnyUnderliningType(i interface{}) {
	ped.UnderliningType = fmt.Sprint(i)
}

func (ped *PatternExtraData) GetDetection() string {
	return ped.Detection
}

// SetAnyDetection устанавливает значение поля Detection
func (ped *PatternExtraData) SetValueDetection(v string) {
	ped.Detection = v
}

// SetAnyDetection устанавливает ЛЮБОЕ значение для поля Detection
func (ped *PatternExtraData) SetAnyDetection(i interface{}) {
	ped.Detection = fmt.Sprint(i)
}

func (ped *PatternExtraData) GetDescription() string {
	return ped.Description
}

// SetValueDescription устанавливает значение поля Description
func (ped *PatternExtraData) SetValueDescription(v string) {
	ped.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (ped *PatternExtraData) SetAnyDescription(i interface{}) {
	ped.Description = fmt.Sprint(i)
}

func (ped *PatternExtraData) GetName() string {
	return ped.Name
}

// SetValueName устанавливает значение поля Name
func (ped *PatternExtraData) SetValueName(v string) {
	ped.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (ped *PatternExtraData) SetAnyName(i interface{}) {
	ped.Name = fmt.Sprint(i)
}

func (ped *PatternExtraData) GetPatternId() string {
	return ped.PatternId
}

// SetValuePatternId устанавливает значение поля PatternId
func (ped *PatternExtraData) SetValuePatternId(v string) {
	ped.PatternId = v
}

// SetAnyPatternId устанавливает ЛЮБОЕ значение для поля PatternId
func (ped *PatternExtraData) SetAnyPatternId(i interface{}) {
	ped.PatternId = fmt.Sprint(i)
}

func (ped *PatternExtraData) GetPatternType() string {
	return ped.PatternType
}

// SetValuePatternType устанавливает значение поля PatternType
func (ped *PatternExtraData) SetValuePatternType(v string) {
	ped.PatternType = v
}

// SetAnyPatternType устанавливает ЛЮБОЕ значение для поля PatternType
func (ped *PatternExtraData) SetAnyPatternType(i interface{}) {
	ped.PatternType = fmt.Sprint(i)
}

func (ped *PatternExtraData) GetURL() string {
	return ped.URL
}

// SetValueURL устанавливает значение поля URL
func (ped *PatternExtraData) SetValueURL(v string) {
	ped.URL = v
}

// SetAnyURL устанавливает ЛЮБОЕ значение для поля URL
func (ped *PatternExtraData) SetAnyURL(i interface{}) {
	ped.URL = fmt.Sprint(i)
}

func (ped *PatternExtraData) GetVersion() string {
	return ped.Version
}

// SetValueVersion устанавливает значение поля Version
func (ped *PatternExtraData) SetValueVersion(v string) {
	ped.Version = v
}

// SetAnyVersion устанавливает ЛЮБОЕ значение для поля Version
func (ped *PatternExtraData) SetAnyVersion(i interface{}) {
	ped.Version = fmt.Sprint(i)
}

func (ped *PatternExtraData) GetPlatforms() []string {
	return ped.Platforms
}

// SetValuePlatforms устанавливает STRING значение для поля Platforms
func (ped *PatternExtraData) SetValuePlatforms(v string) {
	ped.Platforms = append(ped.Platforms, v)
}

// SetAnyPlatforms устанавливает ЛЮБОЕ значение для поля Platforms
func (ped *PatternExtraData) SetAnyPlatforms(i interface{}) {
	ped.Platforms = append(ped.Platforms, fmt.Sprint(i))
}

func (ped *PatternExtraData) GetPermissionsRequired() []string {
	return ped.PermissionsRequired
}

// SetValuePermissionsRequired устанавливает STRING значение для поля PermissionsRequired
func (ped *PatternExtraData) SetValuePermissionsRequired(v string) {
	ped.PermissionsRequired = append(ped.PermissionsRequired, v)
}

// SetAnyPermissionsRequired устанавливает ЛЮБОЕ значение для поля PermissionsRequired
func (ped *PatternExtraData) SetAnyPermissionsRequired(i interface{}) {
	ped.PermissionsRequired = append(ped.PermissionsRequired, fmt.Sprint(i))
}

func (ped *PatternExtraData) GetDataSources() []string {
	return ped.DataSources
}

// SetValueDataSources устанавливает STRING значение для поля DataSources
func (ped *PatternExtraData) SetValueDataSources(v string) {
	ped.DataSources = append(ped.DataSources, v)
}

// SetAnyDataSources устанавливает ЛЮБОЕ значение для поля DataSources
func (ped *PatternExtraData) SetAnyDataSources(i interface{}) {
	ped.DataSources = append(ped.DataSources, fmt.Sprint(i))
}

func (ped *PatternExtraData) GetTactics() []string {
	return ped.Tactics
}

// SetValueTactics устанавливает STRING значение для поля Tactics
func (ped *PatternExtraData) SetValueTactics(v string) {
	ped.Tactics = append(ped.Tactics, v)
}

// SetAnyTactics устанавливает ЛЮБОЕ значение для поля Tactics
func (ped *PatternExtraData) SetAnyTactics(i interface{}) {
	ped.Tactics = append(ped.Tactics, fmt.Sprint(i))
}

func (ped PatternExtraData) ToStringBeautiful(num int) string {
	var str string

	ws := supportingfunctions.GetWhitespace(num)

	str += fmt.Sprintf("%s_createdAt: '%s'\n", ws, ped.UnderliningCreatedAt)
	str += fmt.Sprintf("%s_createdBy: '%s'\n", ws, ped.UnderliningCreatedBy)
	str += fmt.Sprintf("%s_id: '%s'\n", ws, ped.UnderliningId)
	str += fmt.Sprintf("%s_type: '%s'\n", ws, ped.UnderliningType)
	str += fmt.Sprintf("%sdataSources: \n%v", ws, func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("%s%d. '%s'\n", supportingfunctions.GetWhitespace(num+1), k+1, v)
		}
		return str
	}(ped.DataSources))
	/*str += fmt.Sprintf("%sdefenseBypassed: \n%v", ws, func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("%s%d. '%s'\n", supportingfunctions.GetWhitespace(num+1), k+1, v)
		}
		return str
	}(ped.DefenseBypassed))*/
	str += fmt.Sprintf("%sdescription: '%s'\n", ws, ped.Description)
	/*str += fmt.Sprintf("%sextraData: \n%s", ws, func(l map[string]interface{}) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("%s%s: '%v'\n", supportingfunctions.GetWhitespace(num+1), k, v)
		}
		return str
	}(ped.ExtraData))*/
	str += fmt.Sprintf("%sname: '%s'\n", ws, ped.Name)
	str += fmt.Sprintf("%spatternId: '%s'\n", ws, ped.PatternId)
	str += fmt.Sprintf("%spatternType: '%s'\n", ws, ped.PatternType)
	str += fmt.Sprintf("%spermissionsRequired: \n%s", ws, func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("%s%d. '%s'\n", supportingfunctions.GetWhitespace(num+1), k+1, v)
		}
		return str
	}(ped.PermissionsRequired))
	str += fmt.Sprintf("%splatforms: \n%s", ws, func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("%s%d. '%s'\n", supportingfunctions.GetWhitespace(num+1), k+1, v)
		}
		return str
	}(ped.Platforms))
	str += fmt.Sprintf("%sremoteSupport: '%v'\n", ws, ped.RemoteSupport)
	str += fmt.Sprintf("%srevoked: '%v'\n", ws, ped.Revoked)
	/*str += fmt.Sprintf("%ssystemRequirements: \n%s", ws, func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("%s%d. '%s'\n", supportingfunctions.GetWhitespace(num+1), k+1, v)
		}
		return str
	}(ped.SystemRequirements))*/
	str += fmt.Sprintf("%stactics: \n%s", ws, func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("%s%d. '%s'\n", supportingfunctions.GetWhitespace(num+1), k+1, v)
		}
		return str
	}(ped.Tactics))
	str += fmt.Sprintf("%sURL: '%s'\n", ws, ped.URL)
	str += fmt.Sprintf("%sversion: '%s'\n", ws, ped.Version)

	return str
}
