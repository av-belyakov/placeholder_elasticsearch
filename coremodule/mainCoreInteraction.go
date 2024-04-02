package coremodule

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"
	"placeholder_elasticsearch/memorytemporarystorage"
	"placeholder_elasticsearch/mongodbinteractions"
	"placeholder_elasticsearch/natsinteractions"
	rules "placeholder_elasticsearch/rulesinteraction"
	"placeholder_elasticsearch/supportingfunctions"
)

type shortEventSettings struct {
	Source string         `json:"source"`
	Event  nameObjectType `json:"event"`
}

type nameObjectType struct {
	ObjectType string `json:"objectType"`
}

type CoreHandlerSettings struct {
	storageApp *memorytemporarystorage.CommonStorageTemporary
	logging    chan<- datamodels.MessageLogging
	counting   chan<- datamodels.DataCounterSettings
}

func NewCoreHandler(
	storage *memorytemporarystorage.CommonStorageTemporary,
	log chan<- datamodels.MessageLogging,
	count chan<- datamodels.DataCounterSettings) *CoreHandlerSettings {
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
	mdbModule *mongodbinteractions.MongoDBModule) {
	natsChanReception := natsModule.GetDataReceptionChannel()
	decodeJsonCase := NewDecodeJsonMessageSettings(listRuleCase, settings.logging, settings.counting)
	decodeJsonAlert := NewDecodeJsonMessageSettings(listRuleAlert, settings.logging, settings.counting)

	for {
		select {
		case <-ctx.Done():
			return

		case data := <-natsChanReception:
			eventSettings := shortEventSettings{}

			if err := json.Unmarshal(data.Data, &eventSettings); err != nil {
				_, f, l, _ := runtime.Caller(0)
				settings.logging <- datamodels.MessageLogging{
					MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l+1),
					MsgType: "error",
				}

				return
			}

			switch eventSettings.Event.ObjectType {
			case "case":
				chanOutputDecodeJson, chanDecodeJsonDone := decodeJsonCase.HandlerJsonMessage(data.Data, data.MsgId, data.SubjectType)

				chansOut := supportingfunctions.CreateChannelDuplication[datamodels.ChanOutputDecodeJSON](chanOutputDecodeJson, 2)
				chansDone := supportingfunctions.CreateChannelDuplication[bool](chanDecodeJsonDone, 2)

				//используется для хранения в MongoDB
				go NewVerifiedTheHiveFormatCase(chansOut[0], chansDone[0], mdbModule, settings.logging)
				//используется для хранения в Elasticsearch
				go NewVerifiedElasticsearchFormatCase(chansOut[1], chansDone[1], esModule, settings.logging)

			case "alert":
				chanOutputDecodeJson, chanDecodeJsonDone := decodeJsonAlert.HandlerJsonMessage(data.Data, data.MsgId, data.SubjectType)

				chansOut := supportingfunctions.CreateChannelDuplication[datamodels.ChanOutputDecodeJSON](chanOutputDecodeJson, 2)
				chansDone := supportingfunctions.CreateChannelDuplication[bool](chanDecodeJsonDone, 2)

				//используется для хранения в MongoDB
				go NewVerifiedTheHiveFormatAlert(chansOut[0], chansDone[0], mdbModule, settings.logging)
				//используется для хранения в Elasticsearch
				go NewVerifiedElasticsearchFormatAlert(chansOut[1], chansDone[1], esModule, settings.logging)

			default:
				_, f, l, _ := runtime.Caller(0)
				settings.logging <- datamodels.MessageLogging{
					MsgData: fmt.Sprintf("'undefined type objectType' %s:%d", f, l+1),
					MsgType: "error",
				}
			}
		}
	}
}
