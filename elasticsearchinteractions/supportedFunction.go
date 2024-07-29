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
	var (
		query string = `{
			"index": {
				"mapping": {
					"total_fields": {
						"limit": 2000
						}
					}
				}
			}`

		indexSettings = map[string]struct {
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
	)

	if len(indexes) == 0 {
		return fmt.Errorf("an empty list of indexes was received")
	}

	indexForTotalFieldsLimit := []string(nil)
	for _, v := range indexes {
		res, err := hsd.GetIndexSetting(v, "")
		if err != nil {
			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("'%s' %s:%d", err, f, l-1),
				MsgType: "error",
			}

			continue
		}

		if res.StatusCode == 200 {
			err = json.NewDecoder(res.Body).Decode(&indexSettings)
			if err != nil {
				_, f, l, _ := runtime.Caller(0)
				logging <- datamodels.MessageLogging{
					MsgData: fmt.Sprintf("'%s' %s:%d", err, f, l-1),
					MsgType: "error",
				}

				continue
			}

			if indexSettings[v].Settings.Index.Mapping.TotalFields.Limit == "2000" {
				continue
			}

			indexForTotalFieldsLimit = append(indexForTotalFieldsLimit, v)
		}
	}

	if len(indexForTotalFieldsLimit) == 0 {
		return nil
	}

	if _, err := hsd.SetIndexSetting(indexForTotalFieldsLimit, query); err != nil {
		return err
	}

	return nil
}

// SearchUnderlineIdCase выполняет поиск уникального идентификатора (_id) кейса
func SearchUnderlineIdCase(indexName, rootId, source string, hsd HandlerSendData) (string, error) {
	var caseId string

	//выполняем поиск _id индекса
	res, err := hsd.SearchDocument([]string{indexName}, strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}", source, rootId)))
	defer func() {
		errClose := res.Body.Close()
		if err == nil {
			err = errClose
		}
	}()
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		return caseId, fmt.Errorf("'rootId: '%s', %s' %s:%d", err.Error(), rootId, f, l-1)
	}

	if res.StatusCode != http.StatusOK {
		_, f, l, _ := runtime.Caller(0)
		return caseId, fmt.Errorf("'rootId: '%s', %d %s' %s:%d", rootId, res.StatusCode, res.Status(), f, l-2)
	}

	tmp := datamodels.ElasticsearchResponseCase{}
	if err = json.NewDecoder(res.Body).Decode(&tmp); err != nil {
		_, f, l, _ := runtime.Caller(0)
		return caseId, fmt.Errorf("'rootId: '%s', '%s' %s:%d", rootId, err.Error(), f, l-2)
	}

	fmt.Println("_________________________________________________")
	fmt.Printf("func 'SearchUnderlineIdCase', source:'%s' rootId:'%s', tmp.Options:%v\n", source, rootId, tmp)
	fmt.Println("_________________________________________________")

	for _, v := range tmp.Options.Hits {
		caseId = v.ID
	}

	return caseId, nil
}
