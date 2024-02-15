package elasticsearchinteractions

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"placeholder_elasticsearch/datamodels"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

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
		return res, fmt.Errorf("'%v' %s:%d", err, f, l-1)
	}

	if res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusOK {
		return res, nil
	}

	r := map[string]interface{}{}
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		_, f, l, _ := runtime.Caller(0)
		return res, fmt.Errorf("'%v' %s:%d", err, f, l-1)
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
		return countDoc, fmt.Errorf("'%v' %s:%d", err, f, l-1)
	}

	decEs := datamodels.ElasticsearchResponseCase{}
	if err = json.NewDecoder(res.Body).Decode(&decEs); err != nil {
		_, f, l, _ := runtime.Caller(0)
		return countDoc, fmt.Errorf("'%v' %s:%d", err, f, l-1)
	}

	if decEs.Options.Total.Value > 0 {
		countDoc = decEs.Options.Total.Value
		for _, v := range decEs.Options.Hits {
			if _, errDel := hsd.client.Delete(v.Index, v.ID); errDel != nil {
				err = fmt.Errorf("%v, %v", err, errDel)
			}
		}
	}

	return countDoc, err
}

// replacementDocumentCase выполняет замену документа, но только в рамках одного индекса
func (hsd handlerSendData) replacementDocumentCase(
	data interface{},
	index string,
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) {
	obj, ok := data.(*datamodels.VerifiedTheHiveCase)
	if !ok {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'error converting the type to type *datamodels.VerifiedTheHiveCase' %s:%d", f, l-1),
			MsgType: "error",
		}

		return
	}

	queryDelete := strings.NewReader(
		fmt.Sprintf(
			"{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}",
			obj.GetSource(),
			obj.GetEvent().GetRootId(),
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
			MsgData: fmt.Sprintf("a total of '%d' data has been deleted that corresponds to the parameters: source = '%s' and event.rootId = '%s'", countDel, obj.GetSource(), obj.GetEvent().GetRootId()),
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
		DataMsg:  "subject_case",
		Count:    1,
	}
}
