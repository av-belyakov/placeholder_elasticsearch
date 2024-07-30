package coremodule

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"

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

	//coreStorage := newStorage()

	natsChanReception := natsModule.GetDataReceptionChannel()
	decodeJsonCase := NewDecodeJsonMessageSettings(listRuleCase, settings.logging, settings.counting)
	decodeJsonAlert := NewDecodeJsonMessageSettings(listRuleAlert, settings.logging, settings.counting)

	for {
		select {
		case <-ctx.Done():
			return

		/*case data := <-esModule.ChanOutputModule:
		switch data.Section {
		//получаем запрос от модуля Elasticsearch на обогащение кейса
		case "eventenrichment case":
		if data.Command == "get eventenrichment information" {
			if settings, ok := data.Data.(struct {
				Source    string
				RootId    string
				SensorsId []string
			}); ok {
				coreStorage.setRequest(fmt.Sprintf("%s:%s", settings.RootId, settings.Source), "handling case")

				//информация отправляется в модуль обогащения доп. информацией
				eeModule.ChanInputModule <- eventenrichmentmodule.SettingsChanInputEEM{
					RootId:    settings.RootId,
					Source:    settings.Source,
					SensorsId: settings.SensorsId,
				}
			}
		}


		//получаем запрос от модуля Elasticsearch на обогащение алерта
		case "eventenrichment alert":
			if data.Command == "get eventenrichment information" {
				if settings, ok := data.Data.(struct {
					Source    string
					RootId    string
					SensorsId []string
				}); ok {
					coreStorage.setRequest(fmt.Sprintf("%s:%s", settings.RootId, settings.Source), "handling alert")

					//fmt.Printf("===== eventenrichment alert %s:%s -> send to eventenrichmentmodule", settings.RootId, "handling alert")

					//информация отправляется в модуль обогащения доп. информацией
					eeModule.ChanInputModule <- eventenrichmentmodule.SettingsChanInputEEM{
						RootId:    settings.RootId,
						Source:    settings.Source,
						SensorsId: settings.SensorsId,
					}
				}
			}
		}*/

		//канал для взаимодействия с модулем обогащения доп. информацией об организацией
		/*case data := <-eeModule.ChanOutputModule:
		if len(data.Sensors) == 0 {
			//делаем запрос модулю MongoDB
			//
			// надо сделать
			//

			continue
		}

		id := fmt.Sprintf("%s:%s", data.GetRootId(), data.GetSource())
		if section, ok := coreStorage.getRequest(id); ok {
			//отправляем данные для записи в Elasticsearch
			esModule.ChanInputModule <- elasticsearchinteractions.SettingsInputChan{
				Section: section,
				Command: "add eventenrichment information",
				Data:    data,
			}

			coreStorage.deleteElement(id)
		}

		//отправляем данные для записи в MongoDB
		//
		// надо сделать
		//
		*/
		//канал для взаимодействия с NATS
		case data := <-natsChanReception:
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
				go NewVerifiedElasticsearchFormatCase(chansOut[1], chansDone[1], esModule, eeModule, settings.logging)

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
