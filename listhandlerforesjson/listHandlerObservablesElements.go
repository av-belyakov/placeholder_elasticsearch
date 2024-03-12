package listhandlerforesjson

import "placeholder_elasticsearch/listhandlercommon"

func NewListHandlerObservablesElement(so *SupportiveObservables) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		//--- ioc ---
		"observables.ioc": {func(i interface{}) {
			so.HandlerValue(
				"observables.ioc",
				i,
				so.GetObservableTmp().SetAnyIoc,
			)
		}},
		//--- sighted ---
		"observables.sighted": {func(i interface{}) {
			so.HandlerValue(
				"observables.sighted",
				i,
				so.GetObservableTmp().SetAnySighted,
			)
		}},
		//--- ignoreSimilarity ---
		"observables.ignoreSimilarity": {func(i interface{}) {
			so.HandlerValue(
				"observables.ignoreSimilarity",
				i,
				so.GetObservableTmp().SetAnyIgnoreSimilarity,
			)
		}},
		//--- tlp ---
		"observables.tlp": {func(i interface{}) {
			so.HandlerValue(
				"observables.tlp",
				i,
				so.GetObservableTmp().SetAnyTlp,
			)
		}},
		//--- _createdAt ---
		"observables._createdAt": {func(i interface{}) {
			so.HandlerValue(
				"observables._createdAt",
				i,
				so.GetObservableTmp().SetAnyUnderliningCreatedAt,
			)
		}},
		//--- _updatedAt ---
		"observables._updatedAt": {func(i interface{}) {
			so.HandlerValue(
				"observables._updatedAt",
				i,
				so.GetObservableTmp().SetAnyUnderliningUpdatedAt,
			)
		}},
		//--- startDate ---
		"observables.startDate": {func(i interface{}) {
			so.HandlerValue(
				"observables.startDate",
				i,
				so.GetObservableTmp().SetAnyStartDate,
			)
		}},
		//--- _createdBy ---
		"observables._createdBy": {func(i interface{}) {
			so.HandlerValue(
				"observables._createdBy",
				i,
				so.GetObservableTmp().SetAnyUnderliningCreatedBy,
			)
		}},
		//--- _updatedBy ---
		"observables._updatedBy": {func(i interface{}) {
			so.HandlerValue(
				"observables._updatedBy",
				i,
				so.GetObservableTmp().SetAnyUnderliningUpdatedBy,
			)
		}},
		//--- _id ---
		"observables._id": {func(i interface{}) {
			so.HandlerValue(
				"observables._id",
				i,
				so.GetObservableTmp().SetAnyUnderliningId,
			)
		}},
		//--- _type ---
		"observables._type": {func(i interface{}) {
			so.HandlerValue(
				"observables._type",
				i,
				so.GetObservableTmp().SetAnyUnderliningType,
			)
		}},
		//--- data ---
		"observables.data": {func(i interface{}) {
			so.HandlerValue(
				"observables.data",
				i,
				so.GetObservableTmp().SetAnyData,
			)
		}},
		//--- dataType ---
		"observables.dataType": {func(i interface{}) {
			so.HandlerValue(
				"observables.dataType",
				i,
				so.GetObservableTmp().SetAnyDataType,
			)
		}},
		//--- message ---
		"observables.message": {func(i interface{}) {
			so.HandlerValue(
				"observables.message",
				i,
				so.GetObservableTmp().SetAnyMessage,
			)
		}},

		//--- tags ---
		"observables.tags": {
			func(i interface{}) {
				so.HandlerValue(
					"observables.tags",
					i,
					func(i interface{}) {
						key, value := listhandlercommon.HandlerTag(i)
						if value == "" {
							return
						}

						so.GetObservableTmp().SetAnyTags(key, value)
					},
				)
			},
			so.GetObservableTmp().SetAnyTagsAll,
		},
		//--- attachment.id ---
		"observables.attachment.id": {func(i interface{}) {
			so.HandlerValue(
				"observables.attachment.id",
				i,
				so.GetObservableTmp().Attachment.SetAnyId,
			)
		}},
		//--- attachment.size ---
		"observables.attachment.size": {func(i interface{}) {
			so.HandlerValue(
				"observables.attachment.size",
				i,
				so.GetObservableTmp().Attachment.SetAnySize,
			)
		}},
		// --- attachment.name ---
		"observables.attachment.name": {func(i interface{}) {
			so.HandlerValue(
				"observables.attachment.name",
				i,
				so.GetObservableTmp().Attachment.SetAnyName,
			)
		}},
		// --- attachment.contentType ---
		"observables.attachment.contentType": {func(i interface{}) {
			so.HandlerValue(
				"observables.attachment.contentType",
				i,
				so.GetObservableTmp().Attachment.SetAnyContentType,
			)
		}},
		// --- attachment.hashes ---
		"observables.attachment.hashes": {func(i interface{}) {
			so.HandlerValue(
				"observables.attachment.hashes",
				i,
				so.GetObservableTmp().Attachment.SetAnyHashes,
			)
		}},
	}
}
