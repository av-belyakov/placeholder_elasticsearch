package elasticsearchinteractions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/elastic/go-elasticsearch/v8"

	"placeholder_elasticsearch/confighandler"
	"placeholder_elasticsearch/datamodels"
)

type SettingsInputChan struct {
	UUID           string
	Section        string
	Command        string
	Data           []byte
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
	client *elasticsearch.Client
	conf   confighandler.AppConfigElasticSearch
}

func (h *handlerSendData) New() error {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{fmt.Sprintf("http://%s:%d", h.conf.Host, h.conf.Port)},
		Username:  h.conf.User,
		Password:  h.conf.Passwd,
	})

	if err != nil {
		return err
	}

	h.client = es

	return nil
}

func (hsd handlerSendData) sendingData(
	data []byte,
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) {
	if !hsd.conf.Send {
		return
	}

	if hsd.client == nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'the client parameters for connecting to the Elasticsearch database are not set correctly' %s:%d", f, l-2),
			MsgType: "error",
		}

		return
	}

	t := time.Now()
	buf := bytes.NewReader(data)
	res, err := hsd.client.Index(fmt.Sprintf("%s%s_%d_%d", hsd.conf.Prefix, hsd.conf.Index, t.Year(), int(t.Month())), buf)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}

		return
	}

	if res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusOK {
		return
	}

	var errMsg string
	r := map[string]interface{}{}
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}
	}

	if e, ok := r["error"]; ok {
		errMsg = fmt.Sprintln(e)
	}

	logging <- datamodels.MessageLogging{
		MsgData: fmt.Sprintf("received from module Elsaticsearch: %s (%s)", res.Status(), errMsg),
		MsgType: "warning",
	}

	//счетчик
	counting <- datamodels.DataCounterSettings{
		DataType: "update count insert Elasticserach",
		Count:    1,
	}
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

	hsd := handlerSendData{conf: conf}

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
					b, err := json.Marshal(data.VerifiedObject)
					if err != nil {
						_, f, l, _ := runtime.Caller(0)
						logging <- datamodels.MessageLogging{
							MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
							MsgType: "error",
						}
					}

					go hsd.sendingData(b, logging, counting)
				}

			case "":
			}
		}
	}()

	return module, nil
}
