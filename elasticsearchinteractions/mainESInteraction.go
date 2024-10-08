package elasticsearchinteractions

import (
	"fmt"
	"net"
	"net/http"
	"runtime"
	"time"

	"github.com/elastic/go-elasticsearch/v8"

	"placeholder_elasticsearch/confighandler"
	"placeholder_elasticsearch/datamodels"
)

func (h *HandlerSendData) New() error {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{fmt.Sprintf("http://%s:%d", h.Settings.Host, h.Settings.Port)},
		Username:  h.Settings.User,
		Password:  h.Settings.Passwd,
		Transport: &http.Transport{
			MaxIdleConns:          10,              //число открытых TCP-соединений, которые в данный момент не используются
			IdleConnTimeout:       1 * time.Second, //время, через которое закрываются такие неактивные соединения
			MaxIdleConnsPerHost:   10,              //число неактивных TCP-соединений, которые допускается устанавливать на один хост
			ResponseHeaderTimeout: 2 * time.Second, //время в течении которого сервер ожидает получение ответа после записи заголовка запроса
			DialContext: (&net.Dialer{
				Timeout: 3 * time.Second,
				//KeepAlive: 1 * time.Second,
			}).DialContext,
		},
		//RetryOnError: ,
		//RetryOnStatus: ,
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

	go func() {
		for data := range module.ChanInputModule {
			switch data.Section {
			case "handling case":
				index := fmt.Sprintf("%s%s", conf.PrefixCase, conf.IndexCase)

				if data.Command == "add new case" {
					go hsd.ReplacementDocumentCase(data.Data, index, module.ChanOutputModule, logging, counting)
				}

				if data.Command == "add eventenrichment information" {
					go hsd.AddEventenrichmentCase(data.Data, index, logging)
				}

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
