package elasticsearchinteractions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"

	"placeholder_elasticsearch/datamodels"
)

func (hsd HandlerSendData) InsertNewDocument(
	tag string,
	index string,
	document []byte,
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) {
	res, err := hsd.InsertDocument(tag, index, document)
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
	indexName string,
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) {
	newDocument, ok := data.(*datamodels.VerifiedEsCase)
	if !ok {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'error converting the type to type *datamodels.VerifiedTheHiveCase' %s:%d", f, l-1),
			MsgType: "error",
		}

		return
	}

	var (
		countReplacingFields int
		tag                  string = fmt.Sprintf("case rootId: '%s'", newDocument.GetEvent().GetRootId())
	)

	t := time.Now()
	month := int(t.Month())
	indexPattern := fmt.Sprintf("%s_%d", indexName, t.Year())
	indexCurrent := fmt.Sprintf("%s_%d_%d", indexName, t.Year(), month)
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
		hsd.InsertNewDocument(tag, indexCurrent, newDocumentBinary, logging, counting)

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

	decEs := datamodels.ElasticsearchResponseCase{}
	err = json.NewDecoder(res.Body).Decode(&decEs)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}

		return
	}

	if decEs.Options.Total.Value == 0 {
		//выполняется только когда не найден искомый документ
		hsd.InsertNewDocument(tag, indexCurrent, newDocumentBinary, logging, counting)

		return
	}

	//*** при наличие искомого документа выполняем его замену ***
	//***********************************************************
	listDeleting := []datamodels.ServiseOption(nil)
	updateVerified := datamodels.NewVerifiedEsCase()
	for _, v := range decEs.Options.Hits {
		count, err := updateVerified.Event.ReplacingOldValues(*v.Source.GetEvent())
		if err != nil {
			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
				MsgType: "error",
			}
		} else {
			countReplacingFields += count
		}

		countReplacingFields += updateVerified.ObservablesMessageEs.ReplacingOldValues(v.Source.ObservablesMessageEs)
		countReplacingFields += updateVerified.TtpsMessageEs.ReplacingOldValues(v.Source.TtpsMessageEs)

		listDeleting = append(listDeleting, datamodels.ServiseOption{
			ID:    v.ID,
			Index: v.Index,
		})
	}

	//выполняем обновление объекта типа Event
	updateVerified.SetSource(newDocument.GetSource())
	num, err := updateVerified.Event.ReplacingOldValues(*newDocument.GetEvent())
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}
	} else {
		countReplacingFields += num
	}

	countReplacingFields += updateVerified.ObservablesMessageEs.ReplacingOldValues(*newDocument.GetObservables())
	countReplacingFields += updateVerified.TtpsMessageEs.ReplacingOldValues(*newDocument.GetTtps())

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

	res, countDel, err := hsd.UpdateDocument(tag, indexCurrent, listDeleting, nvbyte)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("rootId '%s' '%s' %s:%d", newDocument.GetEvent().GetRootId(), err.Error(), f, l-2),
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
			MsgData: fmt.Sprintf("count delete: '%d', count replacing fields '%d' for alert with rootId: '%s'", countDel, countReplacingFields, newDocument.GetEvent().GetRootId()),
			MsgType: "warning",
		}
	}

	/*

		tag := fmt.Sprintf("case rootId: '%s' ", obj.GetEvent().GetRootId())
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
				MsgData: fmt.Sprintf("DocumentCase - a total of '%d' data has been deleted that corresponds to the parameters: source = '%s' and event.rootId = '%s'", countDel, obj.GetSource(), obj.GetEvent().GetRootId()),
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

		_, err = hsd.InsertDocument(tag, index, b)
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
	*/
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

	var (
		countReplacingFields int
		tag                  string = fmt.Sprintf("alert rootId: '%s'", newDocument.GetEvent().GetRootId())
	)

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
		hsd.InsertNewDocument(tag, indexCurrent, newDocumentBinary, logging, counting)

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

	decEs := datamodels.ElasticsearchResponseAlert{}
	err = json.NewDecoder(res.Body).Decode(&decEs)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}

		return
	}

	if decEs.Options.Total.Value == 0 {
		//выполняется только когда не найден искомый документ
		hsd.InsertNewDocument(tag, indexCurrent, newDocumentBinary, logging, counting)

		return
	}

	//*** при наличие искомого документа выполняем его замену ***
	//***********************************************************
	listDeleting := []datamodels.ServiseOption(nil)
	updateVerified := datamodels.NewVerifiedForEsAlert()
	for _, v := range decEs.Options.Hits {
		var err error
		_, errTmp := updateVerified.Event.ReplacingOldValues(*v.Source.GetEvent())
		if errTmp != nil {
			err = fmt.Errorf("%w event replacing error '%w'", err, errTmp)
		}

		_, errTmp = updateVerified.Alert.ReplacingOldValues(*v.Source.GetAlert())
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

		listDeleting = append(listDeleting, datamodels.ServiseOption{
			ID:    v.ID,
			Index: v.Index,
		})
	}

	//выполняем обновление объекта типа Event
	updateVerified.SetSource(newDocument.GetSource())
	num, errTmp := updateVerified.Event.ReplacingOldValues(*newDocument.GetEvent())
	if errTmp != nil {
		err = fmt.Errorf("%w event replacing error '%w'", err, errTmp)
	} else {
		countReplacingFields += num
	}

	//выполняем обновление объекта типа Alert
	num, errTmp = updateVerified.Alert.ReplacingOldValues(*newDocument.GetAlert())
	if errTmp != nil {
		err = fmt.Errorf("%w alert replacing error '%w'", err, errTmp)
	} else {
		countReplacingFields += num
	}

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

	res, countDel, err := hsd.UpdateDocument(tag, indexCurrent, listDeleting, nvbyte)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("rootId '%s' '%s' %s:%d", newDocument.GetEvent().GetRootId(), err.Error(), f, l-2),
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
			MsgData: fmt.Sprintf("count delete: '%d', count replacing fields '%d' for alert with rootId: '%s'", countDel, countReplacingFields, newDocument.GetEvent().GetRootId()),
			MsgType: "warning",
		}
	}
}
