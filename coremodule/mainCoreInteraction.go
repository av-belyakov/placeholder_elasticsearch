package coremodule

import (
	"context"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"
	"placeholder_elasticsearch/memorytemporarystorage"
	"placeholder_elasticsearch/mongodbinteractions"
	"placeholder_elasticsearch/natsinteractions"
	rules "placeholder_elasticsearch/rulesinteraction"
	"placeholder_elasticsearch/supportingfunctions"
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
	ctx context.Context,
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
		select {
		case <-ctx.Done():
			return

		case data := <-natsChanReception:
			switch data.SubjectType {
			case "subject_case":
				chanOutputDecodeJson, chanDecodeJsonDone := decodeJsonCase.HandlerJsonMessage(data.Data, data.MsgId, data.SubjectType)
				go NewVerifiedTheHiveFormatCase(chanOutputDecodeJson, chanDecodeJsonDone, esModule, mdbModule, settings.logging)

			case "subject_alert":
				chanOutputDecodeJson, chanDecodeJsonDone := decodeJsonAlert.HandlerJsonMessage(data.Data, data.MsgId, data.SubjectType)

				chansOut := supportingfunctions.CreateChannelDuplication[datamodels.ChanOutputDecodeJSON](chanOutputDecodeJson, 2)
				chansDone := supportingfunctions.CreateChannelDuplication[bool](chanDecodeJsonDone, 2)

				go NewVerifiedTheHiveFormatAlert(chansOut[0], chansDone[0], mdbModule, settings.logging)
				go NewVerifiedElasticsearchFormatAlert(chansOut[1], chansDone[1], esModule, settings.logging)
			}
		}
	}
}
