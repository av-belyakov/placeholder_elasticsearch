package elasticsearchinteractions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"placeholder_elasticsearch/datamodels"
)

// SetMaxTotalFieldsLimit устанавливает максимальный лимит полей для
// переданного списка индексов в 2000, если такой лимит не был установлен ранее.
// Данная функция позволяет убрать ошибку Elasticsearch типа Limit of total
// fields [1000] has been exceeded while adding new fields которая
// возникает при установленном максимальном количестве полей в 1000.
func SetMaxTotalFieldsLimit(hsd HandlerSendData, indexes []string, logging chan<- datamodels.MessageLogging) error {
	if len(indexes) == 0 {
		return fmt.Errorf("an empty list of indexes was received")
	}

	getIndexLimit := func(indexName string) (string, bool, error) {
		res, err := hsd.GetIndexSetting(indexName, "")
		defer func() {
			if res == nil || res.Body == nil {
				return
			}

			res.Body.Close()
		}()
		if err != nil {
			return "", false, err
		}

		if res.StatusCode != http.StatusOK {
			return "", false, fmt.Errorf("the server response when executing an index search query is equal to '%s'", res.Status())
		}

		indexSettings := map[string]struct {
			Settings struct {
				Index struct {
					Mapping struct {
						TotalFields struct {
							Limit string `json:"limit"`
						} `json:"total_fields"`
					} `json:"mapping"`
				} `json:"index"`
			} `json:"settings"`
		}{}

		err = json.NewDecoder(res.Body).Decode(&indexSettings)
		if err != nil {
			return "", false, err
		}

		info, ok := indexSettings[indexName]

		return info.Settings.Index.Mapping.TotalFields.Limit, ok, nil
	}

	indexForTotalFieldsLimit := []string(nil)
	for _, v := range indexes {
		limit, ok, err := getIndexLimit(v)
		if err != nil {
			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("'%s' %s:%d", err, f, l-2),
				MsgType: "error",
			}
		}

		if !ok || limit == "2000" {
			continue
		}

		indexForTotalFieldsLimit = append(indexForTotalFieldsLimit, v)
	}

	if len(indexForTotalFieldsLimit) == 0 {
		return nil
	}

	var query string = `{
		"index": {
			"mapping": {
				"total_fields": {
					"limit": 2000
					}
				}
			}
		}`
	if _, err := hsd.SetIndexSetting(indexForTotalFieldsLimit, query); err != nil {
		return err
	}

	return nil
}

// SearchUnderlineIdAlert выполняет поиск уникального идентификатора (_id)
func SearchUnderlineIdAlert(indexName, rootId, source string, hsd HandlerSendData) (string, error) {
	var alertId string

	//выполняем поиск _id индекса
	res, err := hsd.SearchDocument([]string{indexName}, strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}", source, rootId)))
	defer func() {
		if res == nil || res.Body == nil {
			return
		}

		errClose := res.Body.Close()
		if err == nil {
			err = errClose
		}
	}()
	if err != nil {
		return alertId, err
	}

	if res.StatusCode != http.StatusOK {
		return alertId, fmt.Errorf("%s", res.Status())
	}

	tmp := datamodels.ElasticsearchResponseAlert{}
	if err = json.NewDecoder(res.Body).Decode(&tmp); err != nil {
		return alertId, err
	}

	for _, v := range tmp.Options.Hits {
		alertId = v.ID
	}

	return alertId, nil
}

// SearchUnderlineIdCase выполняет поиск уникального идентификатора (_id) кейса
func SearchUnderlineIdCase(indexName, rootId, source string, hsd HandlerSendData) (string, error) {
	var caseId string

	//выполняем поиск _id индекса
	res, err := hsd.SearchDocument([]string{indexName}, strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}", source, rootId)))
	defer func() {
		if res == nil || res.Body == nil {
			return
		}

		errClose := res.Body.Close()
		if err == nil {
			err = errClose
		}
	}()
	if err != nil {
		return caseId, err
	}

	if res.StatusCode != http.StatusOK {
		return caseId, fmt.Errorf("%s", res.Status())
	}

	tmp := datamodels.ElasticsearchResponseCase{}
	if err = json.NewDecoder(res.Body).Decode(&tmp); err != nil {
		return caseId, err
	}

	for _, v := range tmp.Options.Hits {
		caseId = v.ID
	}

	return caseId, nil
}
