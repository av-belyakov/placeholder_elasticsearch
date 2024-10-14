package datamodels

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/supportingfunctions"
)

// ****************** AttachmentData ********************

func NewAttachmentData() *AttachmentData {
	return &AttachmentData{Hashes: []string(nil)}
}

func (a *AttachmentData) GetSize() uint64 {
	return a.Size
}

// SetValueSize устанавливает INT значение для поля Size
func (a *AttachmentData) SetValueSize(v uint64) {
	a.Size = v
}

// SetAnySize устанавливает ЛЮБОЕ значение для поля Size
func (a *AttachmentData) SetAnySize(i interface{}) {
	if v, ok := i.(float32); ok {
		a.Size = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		a.Size = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		a.Size = v
	}
}

func (a *AttachmentData) GetId() string {
	return a.Id
}

// SetValueId устанавливает STRING значение для поля Id
func (a *AttachmentData) SetValueId(v string) {
	a.Id = v
}

// SetAnyId устанавливает ЛЮБОЕ значение для поля Id
func (a *AttachmentData) SetAnyId(i interface{}) {
	a.Id = fmt.Sprint(i)
}

func (a *AttachmentData) GetName() string {
	return a.Name
}

// SetValueName устанавливает STRING значение для поля Name
func (a *AttachmentData) SetValueName(v string) {
	a.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (a *AttachmentData) SetAnyName(i interface{}) {
	a.Name = fmt.Sprint(i)
}

func (a *AttachmentData) GetContentType() string {
	return a.ContentType
}

// SetValueContentType устанавливает STRING значение для поля ContentType
func (a *AttachmentData) SetValueContentType(v string) {
	a.ContentType = v
}

// SetAnyContentType устанавливает ЛЮБОЕ значение для поля ContentType
func (a *AttachmentData) SetAnyContentType(i interface{}) {
	a.ContentType = fmt.Sprint(i)
}

func (a *AttachmentData) GetHashes() []string {
	return a.Hashes
}

// SetValueHashes устанавливает STRING значение для поля Hashes
func (a *AttachmentData) SetValueHashes(v string) {
	a.Hashes = append(a.Hashes, v)
}

// SetAnyHashes устанавливает ЛЮБОЕ значение для поля Hashes
func (a *AttachmentData) SetAnyHashes(i interface{}) {
	a.Hashes = append(a.Hashes, fmt.Sprint(i))
}

func (a AttachmentData) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'size': '%d'\n", ws, a.Size))
	str.WriteString(fmt.Sprintf("%s'id': '%s'\n", ws, a.Id))
	str.WriteString(fmt.Sprintf("%s'name': '%s'\n", ws, a.Name))
	str.WriteString(fmt.Sprintf("%s'contentType': '%s'\n", ws, a.ContentType))
	str.WriteString(fmt.Sprintf("%s'hashes': \n%s", ws, ToStringBeautifulSlice(num, a.Hashes)))

	return str.String()
}

// ********************* ReportTaxonomys *******************
func (t *ReportTaxonomies) GetTaxonomys() []Taxonomy {
	return t.Taxonomies
}

func (t *ReportTaxonomies) GetReportTaxonomys() ReportTaxonomies {
	return *t
}

func (t *ReportTaxonomies) AddTaxonomy(taxonomy Taxonomy) {
	t.Taxonomies = append(t.Taxonomies, taxonomy)
}

// *********************** Taxonomy ************************
func (t *Taxonomy) GetLevel() string {
	return t.Level
}

// SetValueLevel устанавливает STRING значение для поля Level
func (t *Taxonomy) SetValueLevel(v string) {
	t.Level = v
}

// SetAnyLevel устанавливает ЛЮБОЕ значение для поля Level
func (t *Taxonomy) SetAnyLevel(i interface{}) {
	t.Level = fmt.Sprint(i)
}

func (t *Taxonomy) GetNamespace() string {
	return t.Namespace
}

// SetValueNamespace устанавливает STRING значение для поля Namespace
func (t *Taxonomy) SetValueNamespace(v string) {
	t.Namespace = v
}

// SetAnyNamespace устанавливает ЛЮБОЕ значение для поля Namespace
func (t *Taxonomy) SetAnyNamespace(i interface{}) {
	t.Namespace = fmt.Sprint(i)
}

func (t *Taxonomy) GetPredicate() string {
	return t.Predicate
}

// SetValuePredicate устанавливает STRING значение для поля Predicate
func (t *Taxonomy) SetValuePredicate(v string) {
	t.Predicate = v
}

// SetAnyPredicate устанавливает ЛЮБОЕ значение для поля Predicate
func (t *Taxonomy) SetAnyPredicate(i interface{}) {
	t.Predicate = fmt.Sprint(i)
}

func (t *Taxonomy) GetValue() string {
	return t.Value
}

// SetValueValue устанавливает STRING значение для поля Value
func (t *Taxonomy) SetValueValue(v string) {
	t.Value = v
}

// SetAnyValue устанавливает ЛЮБОЕ значение для поля Value
func (t *Taxonomy) SetAnyValue(i interface{}) {
	t.Value = fmt.Sprint(i)
}

// *********************** TtpMessage ************************

func NewTtpMessage() *TtpMessage {
	return &TtpMessage{
		OccurDate:            "1970-01-01T00:00:00+00:00",
		UnderliningCreatedAt: "1970-01-01T00:00:00+00:00",
		ExtraData: ExtraDataTtpMessage{
			Pattern:       *NewPatternExtraData(),
			PatternParent: *NewPatternExtraData(),
		},
	}
}

func (ttpm *TtpMessage) Get() *TtpMessage {
	return ttpm
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

func (tm *TtpMessage) GetExtraData() ExtraDataTtpMessage {
	return tm.ExtraData
}

func (ed *ExtraDataTtpMessage) GetPattern() PatternExtraData {
	return ed.Pattern
}

// SetValueDetails устанавливает значение типа PatternExtraData для поля Pattern
func (ed *ExtraDataTtpMessage) SetValuePattern(v PatternExtraData) {
	ed.Pattern = v
}

func (ed *ExtraDataTtpMessage) GetPatternParent() PatternExtraData {
	return ed.PatternParent
}

// SetValuePatternParent устанавливает значение типа PatternExtraData для поля PatternParent
func (ed *ExtraDataTtpMessage) SetValuePatternParent(v PatternExtraData) {
	ed.PatternParent = v
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

func (tm *TtpMessage) GetPattern() *PatternExtraData {
	return &tm.ExtraData.Pattern
}

func (tm *TtpMessage) GetPatternParent() *PatternExtraData {
	return &tm.ExtraData.PatternParent
}

func NewPatternExtraData() *PatternExtraData {
	return &PatternExtraData{
		UnderliningCreatedAt: "1970-01-01T00:00:00+00:00",
		Platforms:            []string(nil),
		PermissionsRequired:  []string(nil),
		DataSources:          []string(nil),
		Tactics:              []string(nil),
	}
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
	v = strings.ReplaceAll(v, "\t", "")
	v = strings.ReplaceAll(v, "\n", "")

	ped.Detection = v
}

// SetAnyDetection устанавливает ЛЮБОЕ значение для поля Detection
func (ped *PatternExtraData) SetAnyDetection(i interface{}) {
	str := fmt.Sprint(i)
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\n", "")

	ped.Detection = str
}

func (ped *PatternExtraData) GetDescription() string {
	return ped.Description
}

// SetValueDescription устанавливает значение поля Description
func (ped *PatternExtraData) SetValueDescription(v string) {
	v = strings.ReplaceAll(v, "\t", "")
	v = strings.ReplaceAll(v, "\n", "")

	ped.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (ped *PatternExtraData) SetAnyDescription(i interface{}) {
	str := fmt.Sprint(i)
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\n", "")

	ped.Description = str
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

// NewAdditionalInformation формирует структуру содержащую дополнительную информацию
func NewAdditionalInformation() *AdditionalInformation {
	return &AdditionalInformation{
		Sensors:     []SensorInformation(nil),
		IpAddresses: []IpAddressesInformation(nil)}
}

// Get возвращает дополнительную информацию
func (ai *AdditionalInformation) Get() *AdditionalInformation {
	return ai
}

// GetSensors возвращает дополнительную информацию по сенсорам
func (ai *AdditionalInformation) GetSensors() []SensorInformation {
	return ai.Sensors
}

// AddSensor добавляет дополнительную информациб по сенсору
func (ai *AdditionalInformation) AddSensor(e SensorInformation) {
	ai.Sensors = append(ai.Sensors, e)
}

// GetIpAddresses возвращает дополнительную информацию по ip адресам
func (ai *AdditionalInformation) GetIpAddresses() []IpAddressesInformation {
	return ai.IpAddresses
}

// AddIpAddress добавляет дополнительную информациб по ip адресу
func (ai *AdditionalInformation) AddIpAddress(e IpAddressesInformation) {
	ai.IpAddresses = append(ai.IpAddresses, e)
}

// NewSensorInformation формирует структуру с информацией о сенсоре
func NewSensorInformation() *SensorInformation {
	return &SensorInformation{}
}

// GetSensorId возвращает значение SensorId
func (si *SensorInformation) GetSensorId() string {
	return si.SensorId
}

// SetSensorId устанавливает значение SensorId
func (si *SensorInformation) SetSensorId(v string) {
	si.SensorId = v
}

// GetHostId возвращает значение HostId
func (si *SensorInformation) GetHostId() string {
	return si.HostId
}

// SetHostId устанавливает значение HostId
func (si *SensorInformation) SetHostId(v string) {
	si.HostId = v
}

// GetGeoCode возвращает значение GeoCode
func (si *SensorInformation) GetGeoCode() string {
	return si.GeoCode
}

// SetGeoCode устанавливает значение GeoCode
func (si *SensorInformation) SetGeoCode(v string) {
	si.GeoCode = v
}

// GetObjectArea возвращает значение ObjectArea
func (si *SensorInformation) GetObjectArea() string {
	return si.ObjectArea
}

// SetObjectArea устанавливает значение ObjectArea
func (si *SensorInformation) SetObjectArea(v string) {
	si.ObjectArea = v
}

// GetSubjectRF возвращает значение SubjectRF
func (si *SensorInformation) GetSubjectRF() string {
	return si.SubjectRF
}

// SetSubjectRF устанавливает значение SubjectRF
func (si *SensorInformation) SetSubjectRF(v string) {
	si.SubjectRF = v
}

// GetINN возвращает значение INN
func (si *SensorInformation) GetINN() string {
	return si.INN
}

// SetINN устанавливает значение INN
func (si *SensorInformation) SetINN(v string) {
	si.INN = v
}

// GetHomeNet возвращает значение HomeNet
func (si *SensorInformation) GetHomeNet() string {
	return si.HomeNet
}

// SetHomeNet устанавливает значение HomeNet
func (si *SensorInformation) SetHomeNet(v string) {
	si.HomeNet = v
}

// GetOrgName возвращает значение OrgName
func (si *SensorInformation) GetOrgName() string {
	return si.OrgName
}

// SetOrgName устанавливает значение OrgName
func (si *SensorInformation) SetOrgName(v string) {
	si.OrgName = v
}

// GetFullOrgName возвращает значение FullOrgName
func (si *SensorInformation) GetFullOrgName() string {
	return si.FullOrgName
}

// SetFullOrgName устанавливает значение FullOrgName
func (si *SensorInformation) SetFullOrgName(v string) {
	si.FullOrgName = v
}

// NewIpAddressesInformation формирует структуру с информацией об ip адресе
func NewIpAddressesInformation() *IpAddressesInformation {
	return &IpAddressesInformation{}
}

// GetIp возвращает ip адрес
func (ipi *IpAddressesInformation) GetIp() string {
	return ipi.Ip
}

// SetIp устанавливает ip адрес
func (ipi *IpAddressesInformation) SetIp(v string) {
	ipi.Ip = v
}

// GetCity возвращает название города
func (ipi *IpAddressesInformation) GetCity() string {
	return ipi.City
}

// SetCity устанавливает название города
func (ipi *IpAddressesInformation) SetCity(v string) {
	ipi.City = v
}

// GetCountry возвращает название страны
func (ipi *IpAddressesInformation) GetCountry() string {
	return ipi.Country
}

// SetCountry устанавливает название страны
func (ipi *IpAddressesInformation) SetCountry(v string) {
	ipi.Country = v
}

// GetCountryCode возвращает код страны
func (ipi *IpAddressesInformation) GetCountryCode() string {
	return ipi.CountryCode
}

// SetCountryCode устанавливает код страны
func (ipi *IpAddressesInformation) SetCountryCode(v string) {
	ipi.CountryCode = v
}
