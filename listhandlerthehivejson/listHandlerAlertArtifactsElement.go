package listhandlerthehivejson

func NewListHandlerAlertArtifactsElement(saa *SupportiveAlertArtifacts) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		//--- ioc ---
		"alert.artifacts.ioc": {func(i interface{}) {
			saa.HandlerValue(
				"alert.artifacts.ioc",
				i,
				saa.GetArtifactTmp().SetAnyIoc,
			)
		}},
		//--- sighted ---
		"alert.artifacts.sighted": {func(i interface{}) {
			saa.HandlerValue(
				"alert.artifacts.sighted",
				i,
				saa.GetArtifactTmp().SetAnySighted,
			)
		}},
		//--- ignoreSimilarity ---
		"alert.artifacts.ignoreSimilarity": {func(i interface{}) {
			saa.HandlerValue(
				"alert.artifacts.ignoreSimilarity",
				i,
				saa.GetArtifactTmp().SetAnyIgnoreSimilarity,
			)
		}},
		//--- tlp ---
		"alert.artifacts.tlp": {func(i interface{}) {
			saa.HandlerValue(
				"alert.artifacts.tlp",
				i,
				saa.GetArtifactTmp().SetAnyTlp,
			)
		}},
		//--- _id ---
		"alert.artifacts._id": {func(i interface{}) {
			saa.HandlerValue(
				"alert.artifacts._id",
				i,
				saa.GetArtifactTmp().SetAnyUnderliningId,
			)
		}},
		//--- id ---
		"alert.artifacts.id": {func(i interface{}) {
			saa.HandlerValue(
				"alert.artifacts.id",
				i,
				saa.GetArtifactTmp().SetAnyId,
			)
		}},
		//--- _type ---
		"alert.artifacts._type": {func(i interface{}) {
			saa.HandlerValue(
				"alert.artifacts._type",
				i,
				saa.GetArtifactTmp().SetAnyUnderliningType,
			)
		}},
		//--- createdAt ---
		"alert.artifacts.createdAt": {func(i interface{}) {
			saa.HandlerValue(
				"alert.artifacts.createdAt",
				i,
				saa.GetArtifactTmp().SetAnyCreatedAt,
			)
		}},
		//--- updatedAt ---
		"alert.artifacts.updatedAt": {func(i interface{}) {
			saa.HandlerValue(
				"alert.artifacts.updatedAt",
				i,
				saa.GetArtifactTmp().SetAnyUpdatedAt,
			)
		}},
		//--- startDate ---
		"alert.artifacts.startDate": {func(i interface{}) {
			saa.HandlerValue(
				"alert.artifacts.startDate",
				i,
				saa.GetArtifactTmp().SetAnyStartDate,
			)
		}},
		//--- createdBy ---
		"alert.artifacts.createdBy": {func(i interface{}) {
			saa.HandlerValue(
				"alert.artifacts.createdBy",
				i,
				saa.GetArtifactTmp().SetAnyCreatedBy,
			)
		}},
		//--- updatedBy ---
		"alert.artifacts.updatedBy": {func(i interface{}) {
			saa.HandlerValue(
				"alert.artifacts.updatedBy",
				i,
				saa.GetArtifactTmp().SetAnyUpdatedBy,
			)
		}},
		//--- data ---
		"alert.artifacts.data": {func(i interface{}) {
			saa.HandlerValue(
				"alert.artifacts.data",
				i,
				saa.GetArtifactTmp().SetAnyData,
			)
		}},
		//--- dataType ---
		"alert.artifacts.dataType": {func(i interface{}) {
			saa.HandlerValue(
				"alert.artifacts.dataType",
				i,
				saa.GetArtifactTmp().SetAnyDataType,
			)
		}},
		//--- message ---
		"alert.artifacts.message": {func(i interface{}) {
			saa.HandlerValue(
				"alert.artifacts.message",
				i,
				saa.GetArtifactTmp().SetAnyMessage,
			)
		}},
		//--- tags ---
		"alert.artifacts.tags": {func(i interface{}) {
			saa.HandlerValue(
				"alert.artifacts.tags",
				i,
				saa.GetArtifactTmp().SetAnyTags,
			)
		}},
	}
}
