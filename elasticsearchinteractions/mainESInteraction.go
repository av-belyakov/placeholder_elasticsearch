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

var es ModuleElasticSearch

type SettingsInputChan struct {
	UUID string
	Data []byte
}

// ModuleElasticSearch инициализированный модуль
// chanInputElasticSearch - канал для отправки данных в модуль
// chanOutputElasticSearch - канал для принятия данных из модуля
type ModuleElasticSearch struct {
	chanInputElasticSearch  chan SettingsInputChan
	chanOutputElasticSearch chan interface{}
}

type handlerSendData struct {
	client  *elasticsearch.Client
	conf    confighandler.AppConfigElasticSearch
	logging chan<- datamodels.MessageLogging
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

func (hsd handlerSendData) sendingData(data []byte) {
	if !hsd.conf.Send {
		return
	}

	if hsd.client == nil {
		_, f, l, _ := runtime.Caller(0)
		hsd.logging <- datamodels.MessageLogging{
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
		hsd.logging <- datamodels.MessageLogging{
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
		hsd.logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}
	}

	if e, ok := r["error"]; ok {
		errMsg = fmt.Sprintln(e)
	}

	hsd.logging <- datamodels.MessageLogging{
		MsgData: fmt.Sprintf("received from module Elsaticsearch: %s (%s)", res.Status(), errMsg),
		MsgType: "warning",
	}
}

func init() {
	es = ModuleElasticSearch{
		chanInputElasticSearch:  make(chan SettingsInputChan),
		chanOutputElasticSearch: make(chan interface{}),
	}
}

func HandlerElasticSearch(
	conf confighandler.AppConfigElasticSearch,
	logging chan<- datamodels.MessageLogging) (*ModuleElasticSearch, error) {

	hsd := handlerSendData{
		conf:    conf,
		logging: logging,
	}
	if err := hsd.New(); err != nil {
		if err != nil {
			return &es, err
		}
	}

	go func() {
		for data := range es.chanInputElasticSearch {

			//получаем  data.Data и сохраняем ее в Redis list

			go hsd.sendingData(data.Data)
		}
	}()

	return &es, nil
}

func (es ModuleElasticSearch) HandlerData(data SettingsInputChan) {
	es.chanInputElasticSearch <- data
}
