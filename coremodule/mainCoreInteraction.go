package coremodule

import (
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"
	"placeholder_elasticsearch/memorytemporarystorage"
	"placeholder_elasticsearch/mongodbinteractions"
	"placeholder_elasticsearch/natsinteractions"
	rules "placeholder_elasticsearch/rulesinteraction"
)

type CoreHandlerSettings struct {
	storageApp *memorytemporarystorage.CommonStorageTemporary
	logging    chan<- datamodels.MessageLogging
	counting   chan<- datamodels.DataCounterSettings
}

func NewCoreHandler(
	storage *memorytemporarystorage.CommonStorageTemporary,
	log chan<- datamodels.MessageLogging,
	count chan<- datamodels.DataCounterSettings,
) *CoreHandlerSettings {
	return &CoreHandlerSettings{
		storageApp: storage,
		logging:    log,
		counting:   count,
	}
}

func (settings *CoreHandlerSettings) CoreHandler(
	listRuleCase *rules.ListRule,
	listRuleAlert *rules.ListRule,
	natsModule *natsinteractions.ModuleNATS,
	esModule *elasticsearchinteractions.ElasticSearchModule,
	mdbModule *mongodbinteractions.MongoDBModule,
) {
	natsChanReception := natsModule.GetDataReceptionChannel()
	decodeJsonCase := NewDecodeJsonMessageSettings(listRuleCase, settings.logging, settings.counting)
	decodeJsonAlert := NewDecodeJsonMessageSettings(listRuleAlert, settings.logging, settings.counting)

	for {
		data := <-natsChanReception

		switch data.SubjectType {
		case "subject_case":
			chanOutputDecodeJson, chanDecodeJsonDone := decodeJsonCase.HandlerJsonMessage(data.Data, data.MsgId, data.SubjectType)
			go NewVerifiedTheHiveFormatCase(chanOutputDecodeJson, chanDecodeJsonDone, esModule, mdbModule, settings.logging)

		case "subject_alert":
			chanOutputDecodeJson, chanDecodeJsonDone := decodeJsonAlert.HandlerJsonMessage(data.Data, data.MsgId, data.SubjectType)
			go NewVerifiedTheHiveFormatAlert(chanOutputDecodeJson, chanDecodeJsonDone, esModule, mdbModule, settings.logging)
		}
	}
}
