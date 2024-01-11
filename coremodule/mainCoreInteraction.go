package coremodule

import (
	"fmt"
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/memorytemporarystorage"
	"placeholder_elasticsearch/natsinteractions"
)

func CoreHandler(natsModule *natsinteractions.ModuleNATS,
	storageApp *memorytemporarystorage.CommonStorageTemporary,
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) {
	natsChanReception := natsModule.GetDataReceptionChannel()

	for {
		data := <-natsChanReception

		fmt.Println("func 'CoreHandler'", data)

	}
}
