package elasticsearchinteractions

import (
	"fmt"
	"runtime"

	"github.com/elastic/go-elasticsearch/v8"

	"placeholder_elasticsearch/confighandler"
	"placeholder_elasticsearch/datamodels"
)

type SettingsInputChan struct {
	UUID    string
	Section string
	Command string
	RootId  string
	Source  string
	Data    interface{}
}

// ElasticSearchModule инициализированный модуль
// ChanInputModule - канал для отправки данных В модуль
// ChanOutputModule - канал для принятия данных ИЗ модуля
type ElasticSearchModule struct {
	ChanInputModule  chan SettingsInputChan
	ChanOutputModule chan interface{}
}

type HandlerSendData struct {
	Client   *elasticsearch.Client
	Settings SettingsHandler
}

type SettingsHandler struct {
	Port   int
	Host   string
	User   string
	Passwd string
}

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
		ChanOutputModule: make(chan interface{}),
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
				if data.Command == "add new case" {
					index := fmt.Sprintf("%s%s", conf.PrefixCase, conf.IndexCase)

					go hsd.ReplacementDocumentCase(data.Data, index, logging, counting)
				}

			case "handling alert":
				if data.Command == "add new alert" {
					index := fmt.Sprintf("%s%s", conf.PrefixAlert, conf.IndexAlert)

					go hsd.ReplacementDocumentAlert(data.Data, index, logging, counting)
				}
			}
		}
	}()

	return module, nil
}
