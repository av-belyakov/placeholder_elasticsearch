package elasticsearchinteractions

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"

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
		//счетчик
		counting <- datamodels.DataCounterSettings{
			DataType: "update count insert Elasticserach",
			Count:    1,
		}

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
}

// insertDocument добавляет новый документ в заданный индекс
func (hsd handlerSendData) insertDocument(index string, b []byte) (*esapi.Response, error) {
	var res *esapi.Response

	if hsd.client == nil {
		_, f, l, _ := runtime.Caller(0)
		return res, fmt.Errorf("'the client parameters for connecting to the Elasticsearch database are not set correctly' %s:%d", f, l-1)
	}

	buf := bytes.NewReader(b)
	res, err := hsd.client.Index(index, buf)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		return res, fmt.Errorf("'%s' %s:%d", err.Error(), f, l-1)
	}

	if res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusOK {
		return res, nil
	}

	r := map[string]interface{}{}
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		_, f, l, _ := runtime.Caller(0)
		return res, fmt.Errorf("'%s' %s:%d", err.Error(), f, l-1)
	}

	if e, ok := r["error"]; ok {
		return res, fmt.Errorf("received from module Elsaticsearch: %s (%s)", res.Status(), e)
	}

	return res, nil
}

// deleteDocument выполняет поиск и удаление документов соответствующих
// параметрам заданным в запросе
func (hsd handlerSendData) deleteDocument(index []string, query *strings.Reader) (int, error) {
	var (
		err      error
		countDoc int
		res      *esapi.Response
	)

	res, err = hsd.client.Search(
		hsd.client.Search.WithContext(context.Background()),
		hsd.client.Search.WithIndex(index...),
		hsd.client.Search.WithBody(query),
	)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		return countDoc, fmt.Errorf("'%s' %s:%d", err.Error(), f, l-1)
	}

	decEs := datamodels.ElasticsearchResponseCase{}
	if err = json.NewDecoder(res.Body).Decode(&decEs); err != nil {
		_, f, l, _ := runtime.Caller(0)
		return countDoc, fmt.Errorf("'%s' %s:%d", err.Error(), f, l-1)
	}

	if decEs.Options.Total.Value > 0 {
		countDoc = decEs.Options.Total.Value
		for _, v := range decEs.Options.Hits {
			if _, errDel := hsd.client.Delete(v.Index, v.ID); errDel != nil {
				err = fmt.Errorf("%s, %s", err.Error(), errDel.Error())
			}
		}
	}

	return countDoc, err
}

// replacementDocument выполняет замену документа, но только в рамках одного индекса
func (hsd handlerSendData) replacementDocument(
	data *datamodels.VerifiedTheHiveCase,
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) {
	//нужно ли вообще добавлять документ
	if !hsd.conf.Send {
		return
	}

	t := time.Now()
	index := fmt.Sprintf("%s%s_%d_%d", hsd.conf.Prefix, hsd.conf.Index, t.Year(), int(t.Month()))
	queryDelete := strings.NewReader(
		fmt.Sprintf(
			"{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}",
			data.GetSource(),
			data.GetEvent().GetRootId(),
		))

	countDel, err := hsd.deleteDocument([]string{index}, queryDelete)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-1),
			MsgType: "error",
		}
	}
	if countDel > 0 {
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("a total of '%d' data has been deleted that corresponds to the parameters: ", countDel),
			MsgType: "warning",
		}
	}

	b, err := json.Marshal(data)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}
	}

	_, err = hsd.insertDocument(index, b)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}
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
					//*********************************
					//ЭТОТ новый метод должен заменить весь закоментированный
					//код который идет ниже
					//Однако тесты нового метода я не успел провести
					//************************************************
					go hsd.replacementDocument(data.VerifiedObject, logging, counting)

					/*
						Это старый код который не умеет удалять дубли документов

						b, err := json.Marshal(data.VerifiedObject)
						if err != nil {
							_, f, l, _ := runtime.Caller(0)
							logging <- datamodels.MessageLogging{
								MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
								MsgType: "error",
							}
						}

						go hsd.sendingData(b, logging, counting)
					*/
				}

			case "":
			}
		}
	}()

	return module, nil
}
