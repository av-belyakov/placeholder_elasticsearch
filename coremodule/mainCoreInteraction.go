package coremodule

import (
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"
	"placeholder_elasticsearch/memorytemporarystorage"
	"placeholder_elasticsearch/mongodbinteractions"
	"placeholder_elasticsearch/natsinteractions"
	rules "placeholder_elasticsearch/rulesinteraction"
)

func CoreHandler(natsModule *natsinteractions.ModuleNATS,
	storageApp *memorytemporarystorage.CommonStorageTemporary,
	listRule *rules.ListRule,
	esModule *elasticsearchinteractions.ElasticSearchModule,
	mdbModule *mongodbinteractions.MongoDBModule,
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) {
	natsChanReception := natsModule.GetDataReceptionChannel()
	decodeJson := NewDecodeJsonMessageSettings(listRule, logging, counting)

	for {
		data := <-natsChanReception

		//обработчик сообщений из TheHive (выполняется разбор сообщения и его разбор на основе правил)
		chanOutputJsonDecode, chanDecodeDone := decodeJson.HandlerJsonMessage(data.Data, data.MsgId)
		go NewVerifiedTheHiveFormat(chanOutputJsonDecode, chanDecodeDone, esModule, mdbModule, logging)
	}
}
