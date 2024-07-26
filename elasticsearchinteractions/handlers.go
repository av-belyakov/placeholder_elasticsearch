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

func (hsd HandlerSendData) AddEventenrichment(
	data interface{},
	indexName string,
	logging chan<- datamodels.MessageLogging) {

	//приводим к интерфейсу позволяющему получить доступ к информации о сенсорах
	infoEvent, ok := data.(datamodels.InformationFromEventEnricher)

	/*
		Здесь надо будет сделать следующее:
		1. Сделать поисковый запрос на поиск _id кейса
		querySearch *strings.Reader = strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}", source, rootId))
		res, err := hsd.SearchDocument([]string{indexCurrent}, querySearch)

		2. Сделать добавление дополнительной информации в уже существующий кейс
			bodyUpdate := strings.NewReader(fmt.Sprintf("{\"doc\": {\"@geoCode\": \"%s\", \"@objectArea\": \"%s\", \"@subjectRF\": \"%s\", \"@inn\": \"%s\", \"@homeNet\": \"%s\", \"@orgName\": \"%s\", \"@fullOrgName\": \"%s\"}}}", GeoCode, ObjectArea, SubjectRF, INN, HomeNet, OrgName, FullOrgName))
			res, err = hsd.Client.Update(indexCurrent, caseid, bodyUpdate)

	*/
}

// ReplacementDocumentCase выполняет замену документа, но только в рамках одного индекса
func (hsd HandlerSendData) ReplacementDocumentCase(
	data interface{},
	indexName string,
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {
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

	indexes, err := hsd.GetExistingIndexes(indexName)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}

		return
	}

	//будет выполнятся поиск по индексам только в текущем году
	//так как при накоплении большого количества индексов, поиск
	//по всем серьезно замедлит работу
	indexesOnlyCurrentYear := []string(nil)
	for _, v := range indexes {
		if strings.Contains(v, fmt.Sprint(t.Year())) {
			indexesOnlyCurrentYear = append(indexesOnlyCurrentYear, v)
		}
	}
	res, err := hsd.SearchDocument(indexesOnlyCurrentYear, queryCurrent)
	defer func() {
		errClose := res.Body.Close()
		if err == nil {
			err = errClose
		}
	}()
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

		//устанавливаем максимальный лимит количества полей для всех индексов которые
		//содержат значение по умолчанию в 1000 полей
		if err := SetMaxTotalFieldsLimit(hsd, indexes, logging); err != nil {
			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
				MsgType: "error",
			}
		}

		return
	}

	//устанавливаем максимальный лимит количества полей для всех индексов которые
	//содержат значение по умолчанию в 1000 полей
	if err := SetMaxTotalFieldsLimit(hsd, indexes, logging); err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}
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
		countReplacingFields += updateVerified.TtpsMessageTheHive.ReplacingOldValues(v.Source.TtpsMessageTheHive)

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
	countReplacingFields += updateVerified.TtpsMessageTheHive.ReplacingOldValues(*newDocument.GetTtps())

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
	defer func() {
		errClose := res.Body.Close()
		if err == nil {
			err = errClose
		}
	}()
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
			DataMsg:  "subject_case",
			Count:    1,
		}

		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("count delete: '%d', count replacing fields '%d' for alert with rootId: '%s'", countDel, countReplacingFields, newDocument.GetEvent().GetRootId()),
			MsgType: "warning",
		}
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

	var (
		countReplacingFields int
		tag                  string = fmt.Sprintf("alert rootId: '%s'", newDocument.GetEvent().GetRootId())
	)

	t := time.Now()
	month := int(t.Month())
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

	indexes, err := hsd.GetExistingIndexes(indexName)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}

		return
	}

	//будет выполнятся поиск по индексам только в текущем году
	//так как при накоплении большого количества индексов, поиск
	//по всем серьезно замедлит работу
	indexesOnlyCurrentYear := []string(nil)
	for _, v := range indexes {
		if strings.Contains(v, fmt.Sprint(t.Year())) {
			indexesOnlyCurrentYear = append(indexesOnlyCurrentYear, v)
		}
	}

	res, err := hsd.SearchDocument(indexesOnlyCurrentYear, queryCurrent)
	defer func() {
		if res.Body == nil {
			return
		}

		errClose := res.Body.Close() //здесь бывает паника !!!!
		if err == nil {
			err = errClose
		}
	}()
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

		//устанавливаем максимальный лимит количества полей для всех индексов которые
		//содержат значение по умолчанию в 1000 полей
		if err := SetMaxTotalFieldsLimit(hsd, indexes, logging); err != nil {
			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
				MsgType: "error",
			}
		}

		return
	}

	//устанавливаем максимальный лимит количества полей для всех индексов которые
	//содержат значение по умолчанию в 1000 полей
	if err := SetMaxTotalFieldsLimit(hsd, indexes, logging); err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}
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
		MsgType: "test_object_replaced",
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
	defer func() {
		if res.Body == nil {
			return
		}

		errClose := res.Body.Close()
		if err == nil {
			err = errClose
		}
	}()
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
