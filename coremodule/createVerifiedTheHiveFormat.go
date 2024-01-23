package coremodule

import (
	"fmt"

	"placeholder_elasticsearch/datamodels"
)

func NewVerifiedTheHiveFormat(
	input <-chan datamodels.ChanOutputDecodeJSON,
	done <-chan bool,
	//esm *elasticsearchinteractions.ModuleElasticSearch,
	logging chan<- datamodels.MessageLogging,
) {
	//var taskId string

	/*listHandler := joinListHandler([]map[string][]func(interface{}){
		listHandlerEvent,
		listHandlerEventDetails,
		listHandlerEventDetailsCustomFields,
	})*/

	for {
		select {
		case <-input:
			//fmt.Println("func 'NewVerifiedTheHiveFormat' DATA.Value")

		case <-done:
			/*bytes, err := json.Marshal(struct {
				datamodels.SourceMessageTheHive
				datamodels.ObservablesMessageTheHive
				event datamodels.EventMessageTheHive
			}{
				source,
				observables,
				event,
			})
			if err != nil {
				_, f, l, _ := runtime.Caller(0)
				logging <- datamodels.MessageLogging{
					MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-10),
					MsgType: "error",
				}

				return
			}*/

			//Тут надо отправить в модули Elasticsearch и MongoDB

			//esm.HandlerData(elasticsearchinteractions.SettingsInputChan{
			//	UUID: taskId,
			//	Data: bytes,
			//})

			fmt.Println("func 'NewVerifiedTheHiveFormat' PROCESSING IS STOPED")

			return
		}
	}
}
