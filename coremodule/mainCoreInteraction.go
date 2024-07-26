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

type SettingsCommonChanInput struct {
	Section      string
	Command      string
	MsgType      string
	SourceModule string
	Description  string
	Data         interface{}
	SomeData     interface{}
}

type SettingsCommonChanOutput struct{}

type CoreHandlerSettings struct {
	storageApp       *memorytemporarystorage.CommonStorageTemporary
	commonChanInput  chan SettingsCommonChanInput
	commonChanOutput chan SettingsCommonChanOutput
	logging          chan<- datamodels.MessageLogging
	counting         chan<- datamodels.DataCounterSettings
}

func NewCoreHandler(
	storage *memorytemporarystorage.CommonStorageTemporary,
	log chan<- datamodels.MessageLogging,
	count chan<- datamodels.DataCounterSettings) *CoreHandlerSettings {

	return &CoreHandlerSettings{
		storageApp:       storage,
		commonChanInput:  make(chan SettingsCommonChanInput),
		commonChanOutput: make(chan SettingsCommonChanOutput),
		logging:          log,
		counting:         count,
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

	natsChanReception := natsModule.GetDataReceptionChannel()
	decodeJsonCase := NewDecodeJsonMessageSettings(listRuleCase, settings.logging, settings.counting)
	decodeJsonAlert := NewDecodeJsonMessageSettings(listRuleAlert, settings.logging, settings.counting)

	for {
		select {
		case <-ctx.Done():
			return

		//канал для взаимодействия модулей с ядром
		case data := <-settings.commonChanInput:
			switch data.SourceModule {
			case "case_for_elasticsearch":
				//готовый кейс отправляется в Elasticsearch
				esModule.ChanInputModule <- elasticsearchinteractions.SettingsInputChan{
					Section: data.Section,
					Command: data.Command,
					Data:    data.Data,
				}

				/*tmp := strings.Split(data.Description, ":")
				if len(tmp) < 3 {
					var rootId string
					if len(tmp) != 0 {
						rootId = tmp[1]
					}

					_, f, l, _ := runtime.Caller(0)
					settings.logging <- datamodels.MessageLogging{
						MsgData: fmt.Sprintf("'there is not enough data to attempt to enrich with additional information (root_id:'%s')' %s:%d", rootId, f, l-7),
						MsgType: "error",
					}

					break
				}*/
				if someData, ok := data.SomeData.(struct {
					rootId    string
					source    string
					sensorsId []string
				}); ok {
					//информация отправляется в модуль обогащения доп. информацией
					eeModule.ChanInputModule <- eventenrichmentmodule.SettingsChanInputEEM{
						RootId:    someData.rootId,
						Source:    someData.source,
						SensorsId: someData.sensorsId,
					}
				}
			}

		//канал для взаимодействия с модулем обогащения доп. информацией об организацией
		case data := <-eeModule.ChanOutputModule:
			if len(data.Sensors) == 0 {
				//делаем запрос модулю MongoDB
				//
				// надо сделать
				//

				continue
			}

			//отправляем данные для записи в Elasticsearch
			esModule.ChanInputModule <- elasticsearchinteractions.SettingsInputChan{
				Section: "handling case",
				Command: "add eventenrichment information",
				Data:    data,
			}

			//отправляем данные для записи в MongoDB
			//
			// надо сделать
			//

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
				go NewVerifiedElasticsearchFormatCase(chansOut[1], chansDone[1], settings.logging, settings.commonChanInput)

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
