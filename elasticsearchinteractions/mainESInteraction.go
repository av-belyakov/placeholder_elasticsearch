package elasticsearchinteractions

import (
	"fmt"
	"time"

	"github.com/elastic/go-elasticsearch/v8"

	"placeholder_elasticsearch/confighandler"
	"placeholder_elasticsearch/datamodels"
)

type SettingsInputChan struct {
	UUID           string
	Section        string
	Command        string
	Data           interface{}
	VerifiedObject *datamodels.VerifiedTheHiveCase
}

// ElasticSearchModule инициализированный модуль
// ChanInputModule - канал для отправки данных В модуль
// ChanOutputModule - канал для принятия данных ИЗ модуля
type ElasticSearchModule struct {
	ChanInputModule  chan SettingsInputChan
	ChanOutputModule chan interface{}
}

type handlerSendData struct {
	client   *elasticsearch.Client
	settings settingsHandler
}

type settingsHandler struct {
	port   int
	host   string
	user   string
	passwd string
}

func (h *handlerSendData) New() error {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{fmt.Sprintf("http://%s:%d", h.settings.host, h.settings.port)},
		Username:  h.settings.user,
		Password:  h.settings.passwd,
	})

	if err != nil {
		return err
	}

	h.client = es

	return nil
}

func HandlerElasticSearch(
	conf confighandler.AppConfigElasticSearch,
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) (*ElasticSearchModule, error) {

	module := &ElasticSearchModule{
		ChanInputModule:  make(chan SettingsInputChan),
		ChanOutputModule: make(chan interface{}),
	}

	hsd := handlerSendData{
		settings: settingsHandler{
			port:   conf.Port,
			host:   conf.Host,
			user:   conf.User,
			passwd: conf.Passwd,
		},
	}

	if err := hsd.New(); err != nil {
		if err != nil {
			return module, err
		}
	}

	go func() {
		for data := range module.ChanInputModule {
			switch data.Section {
			case "handling case":
				if data.Command == "add new case" {
					t := time.Now()
					index := fmt.Sprintf("%s%s_%d_%d", conf.PrefixCase, conf.IndexCase, t.Year(), int(t.Month()))

					go hsd.replacementDocumentCase(data.Data, index, logging, counting)
				}

			case "handling alert":
				if data.Command == "add new alert" {
					fmt.Println(data.Data)

					/*
						Тут должна быть обработка полученных от модуля core
						verifiedAlert.Get(),
						при этом нужно выполнять update объекта в СУБД
						который совпадает с полученным объектом

						использовать для сравнения методы ReplacingOldValues
						для datamodels.NewVerifiedTheHiveAlert()
					*/

				}
			}
		}
	}()

	return module, nil
}
