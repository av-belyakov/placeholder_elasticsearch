package listhandlerthehivejson

func NewListHandlerTtpElement(sttp *SupportiveTtp) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		//--- occurDate ---
		"ttp.occurDate": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.occurDate",
				i,
				sttp.GetTtpTmp().SetAnyOccurDate,
			)
		}},
		//--- _createdAt ---
		"ttp._createdAt": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp._createdAt",
				i,
				sttp.GetTtpTmp().SetAnyUnderliningCreatedAt,
			)
		}},
		//--- _id ---
		"ttp._id": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp._id",
				i,
				sttp.GetTtpTmp().SetAnyUnderliningId,
			)
		}},
		//--- _createdBy ---
		"ttp._createdBy": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp._createdBy",
				i,
				sttp.GetTtpTmp().SetAnyUnderliningCreatedBy,
			)
		}},
		//--- patternId ---
		"ttp.patternId": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.patternId",
				i,
				sttp.GetTtpTmp().SetAnyPatternId,
			)
		}},
		//--- tactic ---
		"ttp.tactic": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.tactic",
				i,
				sttp.GetTtpTmp().SetAnyTactic,
			)
		}},
		//****************** pattern *******************
		//--- ttp.pattern.remoteSupport ---
		"ttp.pattern.remoteSupport": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.pattern.remoteSupport",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyRemoteSupport,
			)
		}},
		//--- ttp.pattern.revoked ---
		"ttp.pattern.revoked": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.pattern.revoked",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyRevoked,
			)
		}},
		//--- ttp.pattern._createdAt ---
		"ttp.pattern._createdAt": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.pattern._createdAt",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyUnderliningCreatedAt,
			)
		}},
		//--- ttp.pattern._createdBy ---
		"ttp.pattern._createdBy": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.pattern._createdBy",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyUnderliningCreatedBy,
			)
		}},
		//--- ttp.pattern._id ---
		"ttp.pattern._id": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.pattern._id",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyUnderliningId,
			)
		}},
		//--- ttp.pattern._type ---
		"ttp.pattern._type": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.pattern._type",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyUnderliningType,
			)
		}},
		//--- ttp.pattern.detection ---
		"ttp.pattern.detection": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.pattern.detection",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyDetection,
			)
		}},
		//--- ttp.pattern.description ---
		"ttp.pattern.description": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.pattern.description",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyDescription,
			)
		}},
		// --- ttp.pattern.name ---
		"ttp.pattern.name": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.pattern.name",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyName,
			)
		}},
		// --- ttp.pattern.patternId ---
		"ttp.pattern.patternId": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.pattern.patternId",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyPatternId,
			)
		}},
		// --- ttp.pattern.patternType ---
		"ttp.pattern.patternType": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.pattern.patternType",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyPatternType,
			)
		}},
		//--- ttp.pattern.url ---
		"ttp.pattern.url": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.pattern.url",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyURL,
			)
		}},
		//--- ttp.pattern.version ---
		"ttp.pattern.version": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.pattern.version",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyVersion,
			)
		}},
		//--- ttp.pattern.platforms ---
		"ttp.pattern.platforms": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.pattern.platforms",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyPlatforms,
			)
		}},
		//--- ttp.pattern.permissionsRequired ---
		"ttp.pattern.permissionsRequired": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.pattern.permissionsRequired",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyPermissionsRequired,
			)
		}},
		//--- ttp.pattern.dataSources ---
		"ttp.pattern.dataSources": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.pattern.dataSources",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyDataSources,
			)
		}},
		//--- ttp.pattern.tactics ---
		"ttp.pattern.tactics": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.pattern.tactics",
				i,
				sttp.GetTtpTmp().ExtraData.Pattern.SetAnyTactics,
			)
		}},
		//****************** patternParent *******************
		//--- ttp.patternParent.remoteSupport ---
		"ttp.patternParent.remoteSupport": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.patternParent.remoteSupport",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyRemoteSupport,
			)
		}},
		//--- ttp.patternParent.revoked ---
		"ttp.patternParent.revoked": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.patternParent.revoked",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyRevoked,
			)
		}},
		//--- ttp.patternParent._createdAt ---
		"ttp.patternParent._createdAt": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.patternParent._createdAt",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyUnderliningCreatedAt,
			)
		}},
		//--- ttp.patternParent._createdBy ---
		"ttp.patternParent._createdBy": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.patternParent._createdBy",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyUnderliningCreatedBy,
			)
		}},
		//--- ttp.patternParent._id ---
		"ttp.patternParent._id": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.patternParent._id",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyUnderliningId,
			)
		}},
		//--- ttp.patternParent._type ---
		"ttp.patternParent._type": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.patternParent._type",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyUnderliningType,
			)
		}},
		//--- ttp.patternParent.detection ---
		"ttp.patternParent.detection": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.patternParent.detection",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyDetection,
			)
		}},
		//--- ttp.patternParent.description ---
		"ttp.patternParent.description": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.patternParent.description",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyDescription,
			)
		}},
		// --- ttp.patternParent.name ---
		"ttp.patternParent.name": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.patternParent.name",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyName,
			)
		}},
		// --- ttp.patternParent.patternId ---
		"ttp.patternParent.patternId": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.patternParent.patternId",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyPatternId,
			)
		}},
		// --- ttp.patternParent.patternType ---
		"ttp.patternParent.patternType": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.patternParent.patternType",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyPatternType,
			)
		}},
		//--- ttp.patternParent.url ---
		"ttp.patternParent.url": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.patternParent.url",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyURL,
			)
		}},
		//--- ttp.patternParent.version ---
		"ttp.patternParent.version": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.patternParent.version",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyVersion,
			)
		}},
		//--- ttp.patternParent.platforms ---
		"ttp.patternParent.platforms": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.patternParent.platforms",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyPlatforms,
			)
		}},
		//--- ttp.patternParent.permissionsRequired ---
		"ttp.patternParent.permissionsRequired": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.patternParent.permissionsRequired",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyPermissionsRequired,
			)
		}},
		//--- ttp.patternParent.dataSources ---
		"ttp.patternParent.dataSources": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.patternParent.dataSources",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyDataSources,
			)
		}},
		//--- ttp.patternParent.tactics ---
		"ttp.patternParent.tactics": {func(i interface{}) {
			sttp.HandlerValue(
				"ttp.patternParent.tactics",
				i,
				sttp.GetTtpTmp().ExtraData.PatternParent.SetAnyTactics,
			)
		}},
	}
}
