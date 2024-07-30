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

// GetIndexSetting получает натройки выбранного индекса
func (hsd HandlerSendData) GetIndexSetting(index, query string) (*esapi.Response, error) {
	var (
		res *esapi.Response
		err error
	)

	req := esapi.IndicesGetSettingsRequest{
		Index:  []string{index},
		Pretty: true,
		Human:  true,
	}

	res, err = req.Do(context.Background(), hsd.Client.Transport)
	if err != nil {
		return res, err
	}

	return res, nil
}

// SetIndexSetting устанавливает новые настройки индекса
func (hsd HandlerSendData) SetIndexSetting(indexes []string, query string) (bool, error) {
	indicesSettings := esapi.IndicesPutSettingsRequest{
		Index: indexes,
		Body:  strings.NewReader(query),
	}

	res, err := indicesSettings.Do(context.Background(), hsd.Client.Transport)
	defer func() {
		errClose := res.Body.Close()
		if err == nil {
			err = errClose
		}
	}()
	if err != nil {
		return false, err
	}

	if res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusOK {
		return true, nil
	}

	r := map[string]interface{}{}
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		_, f, l, _ := runtime.Caller(0)
		return true, fmt.Errorf("'%v' %s:%d", err, f, l-1)
	}

	if e, ok := r["error"]; ok {
		return true, fmt.Errorf("received from module Elsaticsearch: %s (%s)", res.Status(), e)
	}

	return false, nil
}

// DelIndexSetting удаляет выбранные индексы
func (hsd HandlerSendData) DelIndexSetting(indexes []string) (*esapi.Response, error) {
	var (
		res *esapi.Response
		err error
	)

	req := esapi.IndicesDeleteRequest{
		Index: indexes,
	}

	res, err = req.Do(context.Background(), hsd.Client.Transport)
	if err != nil {
		return res, err
	}

	return res, err
}

// InsertDocument добавляет новый документ в заданный индекс
func (hsd HandlerSendData) InsertDocument(tag, index string, b []byte) (*esapi.Response, error) {
	var res *esapi.Response

	if hsd.Client == nil {
		_, f, l, _ := runtime.Caller(0)
		return res, fmt.Errorf("'the client parameters for connecting to the Elasticsearch database are not set correctly' %s:%d", f, l-1)
	}

	buf := bytes.NewReader(b)
	res, err := hsd.Client.Index(index, buf)
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
		return res, fmt.Errorf("%s received from module Elsaticsearch: %s (%s)", tag, res.Status(), e)
	}

	return res, nil
}

// DeleteDocument выполняет поиск и удаление документов соответствующих
// параметрам заданным в запросе
func (hsd HandlerSendData) DeleteDocument(index []string, query *strings.Reader) (int, error) {
	var (
		err      error
		countDoc int
		res      *esapi.Response
	)

	res, err = hsd.Client.Search(
		hsd.Client.Search.WithContext(context.Background()),
		hsd.Client.Search.WithIndex(index...),
		hsd.Client.Search.WithBody(query),
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
			_, errDel := hsd.Client.Delete(v.Index, v.ID)

			if errDel != nil {
				err = fmt.Errorf("%v, %v", err, errDel)
			}
		}
	}

	return countDoc, err
}

// SearchDocument выполняет поиск документов соответствующих параметрам заданным в запросе
func (hsd HandlerSendData) SearchDocument(index []string, query *strings.Reader) (*esapi.Response, error) {
	res, err := hsd.Client.Search(
		hsd.Client.Search.WithContext(context.Background()),
		hsd.Client.Search.WithIndex(index...),
		hsd.Client.Search.WithBody(query),
	)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		return nil, fmt.Errorf("'%v' %s:%d", err, f, l-1)
	}

	return res, err
}

// UpdateDocument выполняет поиск и обновление документов соответствующих
// параметрам заданным в запросе
func (hsd HandlerSendData) UpdateDocument(tag, currentIndex string, list []datamodels.ServiseOption, document []byte) (res *esapi.Response, countDel int, err error) {
	for _, v := range list {
		_, errDel := hsd.Client.Delete(v.Index, v.ID)
		if errDel != nil {
			err = fmt.Errorf("%v, %v", err, errDel)
		}

		countDel++
	}

	res, err = hsd.InsertDocument(tag, currentIndex, document)

	return res, countDel, err
}

// GetExistingIndexes выполняет проверку наличия индексов соответствующих
// определенному шаблону и возвращает список наименований тндексов
// подходящих под заданный шаблон
func (hsd HandlerSendData) GetExistingIndexes(pattern string) ([]string, error) {
	listIndexes := []string(nil)
	msg := []struct {
		Index string `json:"index"`
	}(nil)

	res, err := hsd.Client.Cat.Indices(
		hsd.Client.Cat.Indices.WithContext(context.TODO()),
		hsd.Client.Cat.Indices.WithFormat("json"),
	)
	if err != nil {
		return nil, err
	}

	defer func() {
		errClose := res.Body.Close() //здесь бывает паника !!!!
		if err == nil {
			err = errClose
		}
	}()

	if err = json.NewDecoder(res.Body).Decode(&msg); err != nil {
		return nil, err
	}

	for _, v := range msg {
		if !strings.Contains(v.Index, pattern) {
			continue
		}

		listIndexes = append(listIndexes, v.Index)
	}

	return listIndexes, err
}

// Get возвращает список идентификаторов сенсоров
func (e *listSensorId) Get() []string {
	return e.sensors
}

// AddElem добавляет только уникальные элементы
func (e *listSensorId) AddElem(sensorId string) {
	for _, v := range e.sensors {
		if v == sensorId {
			return
		}
	}

	e.sensors = append(e.sensors, sensorId)
}
