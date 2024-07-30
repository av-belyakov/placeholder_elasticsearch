package elasticsearchinteractions

import (
	"fmt"
	"runtime"

	"github.com/elastic/go-elasticsearch/v8"

	"placeholder_elasticsearch/confighandler"
	"placeholder_elasticsearch/datamodels"
)

func (h *HandlerSendData) New() error {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{fmt.Sprintf("http://%s:%d", h.Settings.Host, h.Settings.Port)},
		Username:  h.Settings.User,
		Password:  h.Settings.Passwd,
	})
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		return fmt.Errorf("'%v' %s:%d", err, f, l-1)
	}

	h.Client = es

	return nil
}

func HandlerElasticSearch(
	conf confighandler.AppConfigElasticSearch,
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) (*ElasticSearchModule, error) {

	module := &ElasticSearchModule{
		ChanInputModule:  make(chan SettingsInputChan),
		ChanOutputModule: make(chan SettingsOutputChan),
	}

	hsd := HandlerSendData{
		Settings: SettingsHandler{
			Port:   conf.Port,
			Host:   conf.Host,
			User:   conf.User,
			Passwd: conf.Passwd,
		},
	}

	if err := hsd.New(); err != nil {
		return module, err
	}

	/*
				Была ошибка
				06.05.2024 15:43:04
		error: 'dial tcp 192.168.9.10:9200: connect: cannot assign requested address' /go/src/elasticsearchinteractions/handlers.go:90
		06.05.2024 15:42:57
		error: 'dial tcp 192.168.9.10:9200: connect: cannot assign requested address' /go/src/elasticsearchinteractions/handlers.go:280

			И она так и осталась. Надо разбиратся с ней!!!!!!
	*/

	go func() {
		for data := range module.ChanInputModule {
			switch data.Section {
			case "handling case":
				index := fmt.Sprintf("%s%s", conf.PrefixCase, conf.IndexCase)

				if data.Command == "add new case" {
					go hsd.ReplacementDocumentCase(data.Data, index, module.ChanOutputModule, logging, counting)
				}

				//if data.Command == "add eventenrichment information" {
				//	go hsd.AddEventenrichmentCase(data.Data, index, logging)
				//}

			case "handling alert":
				index := fmt.Sprintf("%s%s", conf.PrefixAlert, conf.IndexAlert)

				if data.Command == "add new alert" {
					go hsd.ReplacementDocumentAlert(data.Data, index, module.ChanOutputModule, logging, counting)
				}

				//if data.Command == "add eventenrichment information" {
				//	go hsd.AddEventenrichmentAlert(data.Data, index, logging)
				//}
			}
		}
	}()

	return module, nil
}
