package elasticsearchinteractions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8/esapi"

	"placeholder_elasticsearch/datamodels"
)

func GetVerifiedForEsAlert(res *esapi.Response) (datamodels.ElasticsearchPatternVerifiedForEsAlert, error) {
	mp := datamodels.ElasticsearchPatternVerifiedForEsAlert{}
	err := json.NewDecoder(res.Body).Decode(&mp)
	if err != nil {
		if err != io.EOF {
			return mp, err
		}
	}

	return mp, nil
}

func (hsd HandlerSendData) InsertNewDocument(
	index string,
	document []byte,
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) {
	res, err := hsd.InsertDocument(index, document)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}
	}

	if res.StatusCode == http.StatusCreated {
		//счетчик
		counting <- datamodels.DataCounterSettings{
			DataType: "update count insert Elasticserach",
			DataMsg:  "subject_alert",
			Count:    1,
		}
	}
}

// ReplacementDocumentCase выполняет замену документа, но только в рамках одного индекса
func (hsd HandlerSendData) ReplacementDocumentCase(
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

	t := time.Now()
	index = fmt.Sprintf("%s_%d_%d", index, t.Year(), int(t.Month()))

	queryDelete := strings.NewReader(
		fmt.Sprintf(
			"{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}",
			obj.GetSource(),
			obj.GetEvent().GetRootId(),
		))

	countDel, err := hsd.DeleteDocument([]string{index}, queryDelete)
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

		return
	}

	_, err = hsd.InsertDocument(index, b)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}

		return
	}

	//счетчик
	counting <- datamodels.DataCounterSettings{
		DataType: "update count insert Elasticserach",
		DataMsg:  "subject_case",
		Count:    1,
	}
}

// ReplacementDocumentAlert выполняет замену документа, но только в рамках одного индекса
func (hsd HandlerSendData) ReplacementDocumentAlert(
	data interface{},
	indexName string,
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) {
	newDocument, ok := data.(*datamodels.VerifiedForEsAlert)
	if !ok {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'error converting the type to type *datamodels.VerifiedTheHiveAlert' %s:%d", f, l-1),
			MsgType: "error",
		}

		return
	}

	var countReplacingFields int

	t := time.Now()
	month := int(t.Month())
	indexPattern := fmt.Sprintf("%s_%s_%d", indexName, newDocument.GetSource(), t.Year())
	indexCurrent := fmt.Sprintf("%s_%s_%d_%d", indexName, newDocument.GetSource(), t.Year(), month)
	queryCurrent := strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}", newDocument.GetSource(), newDocument.GetEvent().GetRootId()))

	newDocumentBinary, err := json.Marshal(newDocument.Get())
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}

		return
	}

	indexes, err := hsd.GetExistingIndexes(indexPattern)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}

		return
	}

	if len(indexes) == 0 {
		hsd.InsertNewDocument(indexCurrent, newDocumentBinary, logging, counting)

		return
	}

	res, err := hsd.SearchDocument(indexes, queryCurrent)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}

		return
	}

	if res.StatusCode != http.StatusOK {
		//выполняется только когда не найден искомый документ
		hsd.InsertNewDocument(indexCurrent, newDocumentBinary, logging, counting)

		return
	}

	//*** при наличие искомого документа выполняем его замену ***
	//***********************************************************
	object, err := GetVerifiedForEsAlert(res)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}

		return
	}

	updateVerified := datamodels.NewVerifiedForEsAlert()
	for _, v := range object.Hits.Hits {
		var err error

		_, errTmp := updateVerified.Event.ReplacingOldValues(v.Source.Event)
		if errTmp != nil {
			err = fmt.Errorf("%w event replacing error '%w'", err, errTmp)
		}

		_, errTmp = updateVerified.Alert.ReplacingOldValues(v.Source.Alert)
		if errTmp != nil {
			err = fmt.Errorf("%w alert replacing error '%w'", err, errTmp)
		}

		if err != nil {
			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
				MsgType: "error",
			}
		}
	}

	//выполняем обновление объекта типа Event
	num, errTmp := updateVerified.Event.ReplacingOldValues(*newDocument.GetEvent())
	if errTmp != nil {
		err = fmt.Errorf("%w event replacing error '%w'", err, errTmp)
	}
	countReplacingFields += num

	//выполняем обновление объекта типа Alert
	num, errTmp = updateVerified.Alert.ReplacingOldValues(*newDocument.GetAlert())
	if errTmp != nil {
		err = fmt.Errorf("%w alert replacing error '%w'", err, errTmp)
	}
	countReplacingFields += num

	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}
	}

	//******** TEST ********
	//только в рамках тестирования, отправка обновленного объекта
	//в специальный файл
	infoUpdate, err := json.MarshalIndent(updateVerified, "", "  ")
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}
	}
	logging <- datamodels.MessageLogging{
		MsgData: string(infoUpdate),
		MsgType: "test_object_update",
	}
	//***********************

	nvbyte, err := json.Marshal(updateVerified)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}

		return
	}

	res, err = hsd.UpdateDocument(indexCurrent, indexPattern, queryCurrent, nvbyte)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}

		return
	}

	if res.StatusCode == http.StatusCreated {
		//счетчик
		counting <- datamodels.DataCounterSettings{
			DataType: "update count insert Elasticserach",
			DataMsg:  "subject_alert",
			Count:    1,
		}

		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("count replacing fields '%d' for alert with rootId: '%s'", countReplacingFields, newDocument.GetEvent().GetRootId()),
			MsgType: "warning",
		}
	}
}
