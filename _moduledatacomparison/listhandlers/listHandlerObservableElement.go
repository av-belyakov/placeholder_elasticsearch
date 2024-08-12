package supportingfunctions

import (
	"strings"

	"placeholder_elasticsearch/_moduledatacomparison/datamodel"
)

func NewListHandlerObservablesElement(so *datamodel.SupportiveObservables) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		//--- ioc ---
		"observablesmessagethehive.observables.ioc": {func(i interface{}) {
			so.HandlerValue(
				"observablesmessagethehive.observables.ioc",
				i,
				so.GetObservableTmp().SetAnyIoc,
			)
		}},
		//--- sighted ---
		"observablesmessagethehive.observables.sighted": {func(i interface{}) {
			so.HandlerValue(
				"observablesmessagethehive.observables.sighted",
				i,
				so.GetObservableTmp().SetAnySighted,
			)
		}},
		//--- ignoreSimilarity ---
		"observablesmessagethehive.observables.ignoreSimilarity": {func(i interface{}) {
			so.HandlerValue(
				"observablesmessagethehive.observables.ignoreSimilarity",
				i,
				so.GetObservableTmp().SetAnyIgnoreSimilarity,
			)
		}},
		//--- tlp ---
		"observablesmessagethehive.observables.tlp": {func(i interface{}) {
			so.HandlerValue(
				"observablesmessagethehive.observables.tlp",
				i,
				so.GetObservableTmp().SetAnyTlp,
			)
		}},
		//--- _createdAt ---
		"observablesmessagethehive.observables._createdAt": {func(i interface{}) {
			so.HandlerValue(
				"observablesmessagethehive.observables._createdAt",
				i,
				so.GetObservableTmp().SetAnyUnderliningCreatedAt,
			)
		}},
		//--- _updatedAt ---
		"observablesmessagethehive.observables._updatedAt": {func(i interface{}) {
			so.HandlerValue(
				"observablesmessagethehive.observables._updatedAt",
				i,
				so.GetObservableTmp().SetAnyUnderliningUpdatedAt,
			)
		}},
		//--- startDate ---
		"observablesmessagethehive.observables.startDate": {func(i interface{}) {
			so.HandlerValue(
				"observablesmessagethehive.observables.startDate",
				i,
				so.GetObservableTmp().SetAnyStartDate,
			)
		}},
		//--- _createdBy ---
		"observablesmessagethehive.observables._createdBy": {func(i interface{}) {
			so.HandlerValue(
				"observablesmessagethehive.observables._createdBy",
				i,
				so.GetObservableTmp().SetAnyUnderliningCreatedBy,
			)
		}},
		//--- _updatedBy ---
		"observablesmessagethehive.observables._updatedBy": {func(i interface{}) {
			so.HandlerValue(
				"observablesmessagethehive.observables._updatedBy",
				i,
				so.GetObservableTmp().SetAnyUnderliningUpdatedBy,
			)
		}},
		//--- _id ---
		"observablesmessagethehive.observables._id": {func(i interface{}) {
			so.HandlerValue(
				"observablesmessagethehive.observables._id",
				i,
				so.GetObservableTmp().SetAnyUnderliningId,
			)
		}},
		//--- _type ---
		"observablesmessagethehive.observables._type": {func(i interface{}) {
			so.HandlerValue(
				"observablesmessagethehive.observables._type",
				i,
				so.GetObservableTmp().SetAnyUnderliningType,
			)
		}},
		//--- data ---
		"observablesmessagethehive.observables.data": {func(i interface{}) {
			so.HandlerValue(
				"observablesmessagethehive.observables.data",
				i,
				so.GetObservableTmp().SetAnyData,
			)
		}},
		//--- dataType ---
		"observablesmessagethehive.observables.dataType": {func(i interface{}) {
			so.HandlerValue(
				"observablesmessagethehive.observables.dataType",
				i,
				so.GetObservableTmp().SetAnyDataType,
			)
		}},
		//--- message ---
		"observablesmessagethehive.observables.message": {func(i interface{}) {
			so.HandlerValue(
				"observablesmessagethehive.observables.message",
				i,
				so.GetObservableTmp().SetAnyMessage,
			)
		}},

		//--- tags ---
		"observablesmessagethehive.observables.tags": {
			func(i interface{}) {
				so.HandlerValue(
					"observablesmessagethehive.observables.tags",
					i,
					func(i interface{}) {
						key, value := HandlerTag(i)
						if value == "" {
							return
						}

						value = strings.TrimSpace(value)
						value = strings.Trim(value, "\"")
						so.GetObservableTmp().SetAnyTags(key, value)
					},
				)
			},
			so.GetObservableTmp().SetAnyTagsAll,
		},
		//--- attachment.id ---
		"observablesmessagethehive.observables.attachment.id": {func(i interface{}) {
			so.HandlerValue(
				"observablesmessagethehive.observables.attachment.id",
				i,
				so.GetObservableTmp().Attachment.SetAnyId,
			)
		}},
		//--- attachment.size ---
		"observablesmessagethehive.observables.attachment.size": {func(i interface{}) {
			so.HandlerValue(
				"observablesmessagethehive.observables.attachment.size",
				i,
				so.GetObservableTmp().Attachment.SetAnySize,
			)
		}},
		// --- attachment.name ---
		"observablesmessagethehive.observables.attachment.name": {func(i interface{}) {
			so.HandlerValue(
				"observablesmessagethehive.observables.attachment.name",
				i,
				so.GetObservableTmp().Attachment.SetAnyName,
			)
		}},
		// --- attachment.contentType ---
		"observablesmessagethehive.observables.attachment.contentType": {func(i interface{}) {
			so.HandlerValue(
				"observablesmessagethehive.observables.attachment.contentType",
				i,
				so.GetObservableTmp().Attachment.SetAnyContentType,
			)
		}},
		// --- attachment.hashes ---
		"observablesmessagethehive.observables.attachment.hashes": {func(i interface{}) {
			so.HandlerValue(
				"observablesmessagethehive.observables.attachment.hashes",
				i,
				so.GetObservableTmp().Attachment.SetAnyHashes,
			)
		}},
	}
}
