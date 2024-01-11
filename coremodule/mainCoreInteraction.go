package coremodule

import (
	"fmt"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/memorytemporarystorage"
	"placeholder_elasticsearch/natsinteractions"
	rules "placeholder_elasticsearch/rulesinteraction"
)

func CoreHandler(natsModule *natsinteractions.ModuleNATS,
	storageApp *memorytemporarystorage.CommonStorageTemporary,
	listRule *rules.ListRule,
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) {
	natsChanReception := natsModule.GetDataReceptionChannel()
	decodeJson := NewDecodeJsonMessageSettings(listRule, logging, counting)

	for {
		data := <-natsChanReception

		fmt.Println("func 'CoreHandler'", data)

		//обработчик сообщений из TheHive (выполняется разбор сообщения и его разбор на основе правил)
		chanOutputJsonDecode, chanDecodeDone := decodeJson.HandlerJsonMessage(data.Data, data.MsgId)
		go NewVerifiedTheHiveFormat(chanOutputJsonDecode, chanDecodeDone, logging)
	}
}
