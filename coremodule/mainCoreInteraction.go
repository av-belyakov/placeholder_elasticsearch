package coremodule

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"strings"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"
	"placeholder_elasticsearch/eventenrichmentmodule"
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
	mdbModule *mongodbinteractions.MongoDBModule,
	eeModule *eventenrichmentmodule.EventEnrichmentModule) {

	coreStorage := newStorage()

	natsReception := natsModule.GetDataReceptionChannel()
	decodeJsonCase := NewDecodeJsonMessageSettings(listRuleCase, settings.logging, settings.counting)
	decodeJsonAlert := NewDecodeJsonMessageSettings(listRuleAlert, settings.logging, settings.counting)

	for {
		select {
		case <-ctx.Done():
			return

		case info := <-eeModule.ChanOutputModule:
			//если в Zabbix было найденно хоть что то по сенсорам
			if len(info.GetSensorsId()) > 0 {
				// отправляем, найденную о сенсорах информацию, в MongoDB
				mdbModule.ChanInputModule <- mongodbinteractions.SettingsInputChan{
					Section: "handling eventenrichment",
					Command: "add sensor eventenrichment",
					Data:    info,
				}

				//отправляем найденную информацию в СУБД Elasticsearch
				esModule.ChanInputModule <- elasticsearchinteractions.SettingsInputChan{
					Section: "handling case",
					Command: "add eventenrichment information",
					Data:    info,
				}
			}

			if len(info.SensorsId) > 0 {
				settings.logging <- datamodels.MessageLogging{
					MsgData: fmt.Sprintf("'the following sensors '%s' were not found in Zabbix, a search will be performed for them in the database MongoDB'", strings.Join(info.SensorsId, ",")),
					MsgType: "info",
				}

				//перечень сенсоров по которым по каким то причинам ничего не найдено
				mdbModule.ChanInputModule <- mongodbinteractions.SettingsInputChan{
					Section: "handling eventenrichment",
					Command: "get sensor eventenrichment",
					RootId:  info.RootId,
					Source:  info.Source,
					Data:    info.SensorsId,
				}
			}

		case data := <-mdbModule.ChanOutputModule:
			if data.Section == "handling eventenrichment" && data.Command == "sensor eventenrichment response" {
				// отправляем найденную информацию в СУБД Elasticsearch
				esModule.ChanInputModule <- elasticsearchinteractions.SettingsInputChan{
					Section: "handling case",
					Command: "add eventenrichment information",
					Data:    data.Data,
				}
			}

		//канал для взаимодействия с NATS
		case data := <-natsReception:
			eventSettings := shortEventSettings{}

			if err := json.Unmarshal(data.Data, &eventSettings); err != nil {
				_, f, l, _ := runtime.Caller(0)
				settings.logging <- datamodels.MessageLogging{
					MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l+1),
					MsgType: "error",
				}

				continue
			}

			switch eventSettings.Event.ObjectType {
			case "case":
				chanOutputDecodeJson, chanDecodeJsonDone := decodeJsonCase.HandlerJsonMessage(data.Data, data.MsgId, "subject_case")

				chansOut := supportingfunctions.CreateChannelDuplication[datamodels.ChanOutputDecodeJSON](chanOutputDecodeJson, 2)
				chansDone := supportingfunctions.CreateChannelDuplication[bool](chanDecodeJsonDone, 2)

				//используется для хранения в MongoDB
				go NewVerifiedTheHiveFormatCase(chansOut[0], chansDone[0], mdbModule, settings.logging)
				//используется для хранения в Elasticsearch
				go NewVerifiedElasticsearchFormatCase(VerifiedElasticsearchFormatCaseOptions{
					msgId:    data.MsgId,
					cs:       coreStorage,
					input:    chansOut[1],
					done:     chansDone[1],
					natsChan: natsModule.GetDataDeliveryChannel(),
					esmChan:  esModule.ChanInputModule,
					eemChan:  eeModule.ChanInputModule,
					logging:  settings.logging})

			case "alert":
				chanOutputDecodeJson, chanDecodeJsonDone := decodeJsonAlert.HandlerJsonMessage(data.Data, data.MsgId, "subject_alert")

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
