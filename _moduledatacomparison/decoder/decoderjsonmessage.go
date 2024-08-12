package decoder

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"placeholder_elasticsearch/_moduledatacomparison/datamodel"
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/listhandlerforesjson"
	"placeholder_elasticsearch/listhandlerthehivejson"
)

func FormatCaseJsonMongoDBHandler(formatCaseJson []byte, resultObject *datamodels.VerifiedEsCase, logging chan<- datamodels.MessageLogging) bool {
	chanOutputDecodeJson, chanDone := HandlerDecodeJsonMessage(formatCaseJson, logging)

	var (
		event *datamodels.EventMessageForEsCase = datamodels.NewEventMessageForEsCase()

		eventObject  *datamodels.EventForEsCaseObject = datamodels.NewEventForEsCaseObject()
		eventDetails *datamodels.EventCaseDetails     = datamodels.NewEventCaseDetails()

		eventObjectCustomFields  datamodels.CustomFields = datamodels.CustomFields{}
		eventDetailsCustomFields datamodels.CustomFields = datamodels.CustomFields{}
	)

	//******************* Основные обработчики для Event **********************
	// ------ EVENT ------
	listHandlerEvent := listhandlerforesjson.NewListHandlerEventCaseElement(event)
	// ------ EVENT OBJECT ------
	listHandlerEventObject := listhandlerforesjson.NewListHandlerEventCaseObjectElement(eventObject)
	// ------ EVENT OBJECT CUSTOMFIELDS ------
	listHandlerEventObjectCustomFields := listhandlerthehivejson.NewListHandlerEventObjectCustomFieldsElement(eventObjectCustomFields)
	// ------ EVENT DETAILS ------
	listHandlerEventDetails := listhandlerforesjson.NewListHandlerEventCaseDetailsElement(eventDetails)
	// ------ EVENT DETAILS CUSTOMFIELDS ------
	listHandlerEventDetailsCustomFields := listhandlerthehivejson.NewListHandlerEventDetailsCustomFieldsElement(eventDetailsCustomFields)

	//******************* Вспомогательный объект для Observables **********************
	so := listhandlerforesjson.NewSupportiveObservables()
	listHandlerObservables := listhandlerforesjson.NewListHandlerObservablesElement(so)

	//******************* Вспомогательный объект для Ttp **********************
	sttp := listhandlerthehivejson.NewSupportiveTtp()
	listHandlerTtp := listhandlerthehivejson.NewListHandlerTtpElement(sttp)

	for data := range chanOutputDecodeJson {
		if data.FieldBranch == "source" {
			resultObject.SetSource(fmt.Sprint(data.Value))
		}

		if data.FieldBranch == "@id" {
			resultObject.SetID(fmt.Sprint(data.Value))
		}

		if data.FieldBranch == "@timestamp" {
			resultObject.SetCreateTimestamp(fmt.Sprint(data.Value))
		}

		//******************************************************************
		//********** Сбор всех объектов относящихся к полю Event  **********
		// event element
		if lf, ok := listHandlerEvent[data.FieldBranch]; ok {
			for _, f := range lf {
				f(data.Value)
			}

			continue
		}

		// event.object element
		if lf, ok := listHandlerEventObject[data.FieldBranch]; ok {
			for _, f := range lf {
				f(data.Value)
			}

			continue
		}

		// event.object.customFields element
		if lf, ok := listHandlerEventObjectCustomFields[data.FieldBranch]; ok {
			for _, f := range lf {
				f(data.Value)
			}

			continue
		}

		// event.details element
		if lf, ok := listHandlerEventDetails[data.FieldBranch]; ok {
			for _, f := range lf {
				f(data.Value)
			}

			continue
		}

		// event.details.customFields element
		if lf, ok := listHandlerEventDetailsCustomFields[data.FieldBranch]; ok {
			for _, f := range lf {
				f(data.Value)
			}

			continue
		}

		if strings.Contains(data.FieldBranch, "ignoreSimilarity") {
			fmt.Println("__FIND: ", data.FieldBranch)
		}

		//************************************************************************
		//********** Сбор всех объектов относящихся к полю Observables  **********
		// для всех полей входящих в observables, кроме содержимого
		//поля reports
		if lf, ok := listHandlerObservables[data.FieldBranch]; ok {
			for _, f := range lf {
				r := reflect.TypeOf(data.Value)
				switch r.Kind() {
				case reflect.Slice:
					if s, ok := data.Value.([]interface{}); ok {
						for _, value := range s {
							f(value)
						}
					}
				default:
					f(data.Value)

				}
			}

			continue
		}

		//убрал обработку observables.reports так как тип TtpsMessageEs
		//способствует росту черезмерно большого количества полей которое
		//влечет за собой превышения лимита маппинга в Elsticsearch), что
		//выражается в ошибке от СУБД типа "Limit of total fields [2000]
		//has been exceeded while adding new fields"
		//
		//для всех полей входящих в состав observables.reports
		//if strings.Contains(data.FieldBranch, "observables.reports.") {
		//		handlerIsExist = true
		//		so.HandlerReportValue(data.FieldBranch, data.Value)
		//}

		//*********************************************************************
		//********** Сбор всех объектов относящихся к полю Ttp  ***************
		if lf, ok := listHandlerTtp[data.FieldBranch]; ok {
			for _, f := range lf {
				r := reflect.TypeOf(data.Value)
				switch r.Kind() {
				case reflect.Slice:
					if s, ok := data.Value.([]interface{}); ok {
						for _, value := range s {
							f(value)
						}
					}

				default:
					f(data.Value)

				}
			}

			continue
		}
	}

	// Собираем объект Event
	eventObject.SetValueCustomFields(eventObjectCustomFields)
	eventDetails.SetValueCustomFields(eventDetailsCustomFields)
	event.SetValueObject(*eventObject)
	event.SetValueDetails(*eventDetails)

	// собираем объект observables
	observables := datamodels.NewObservablesMessageEs()
	observables.SetValueObservables(so.GetObservables())

	// собираем объект ttp
	ttps := datamodels.NewTtpsMessageTheHive()
	ttps.SetTtps(sttp.GetTtps())

	resultObject.SetEvent(*event)
	resultObject.SetObservables(*observables)
	resultObject.SetTtps(*ttps)

	return <-chanDone
}

func HandlerDecodeJsonMessage(b []byte, logging chan<- datamodels.MessageLogging) (chan datamodel.ChanOutputDecodeJSON, chan bool) {
	chanOutputJsonData := make(chan datamodel.ChanOutputDecodeJSON)
	chanDone := make(chan bool)

	isSuccess := true
	go func() {
		var (
			f   string
			l   int
			err error
		)

		//для карт
		_, f, l, _ = runtime.Caller(0)
		listMap := map[string]interface{}{}
		if err = json.Unmarshal(b, &listMap); err == nil {
			if len(listMap) == 0 {
				isSuccess = false
				logging <- datamodels.MessageLogging{
					MsgData: fmt.Sprintf("'error decoding the json message, it may be empty' %s:%d", f, l+2),
					MsgType: "error",
				}

				return
			}

			_ = reflectMap(logging, chanOutputJsonData, listMap, "")
		} else {
			// для срезов
			_, f, l, _ = runtime.Caller(0)
			listSlice := []interface{}{}
			if err = json.Unmarshal(b, &listSlice); err != nil {
				isSuccess = false
				logging <- datamodels.MessageLogging{
					MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l+2),
					MsgType: "error",
				}

				return
			}

			if len(listSlice) == 0 {
				isSuccess = false
				logging <- datamodels.MessageLogging{
					MsgData: fmt.Sprintf("'error decoding the json message, it may be empty' %s:%d", f, l+2),
					MsgType: "error",
				}

				return
			}

			_ = reflectSlice(logging, chanOutputJsonData, listSlice, "")
		}

		close(chanOutputJsonData)

		//останавливаем обработчик формирующий верифицированный объект
		chanDone <- isSuccess

		close(chanDone)
	}()

	return chanOutputJsonData, chanDone
}

func reflectAnySimpleType(
	logging chan<- datamodels.MessageLogging,
	chanOutMispFormat chan<- datamodel.ChanOutputDecodeJSON,
	name interface{},
	anyType interface{},
	fieldBranch string) interface{} {

	var nameStr string
	r := reflect.TypeOf(anyType)

	if n, ok := name.(int); ok {
		nameStr = fmt.Sprint(n)
	} else if n, ok := name.(string); ok {
		nameStr = n
	}

	if r == nil {
		return anyType
	}

	switch r.Kind() {
	case reflect.String:
		result := reflect.ValueOf(anyType).String()

		chanOutMispFormat <- datamodel.ChanOutputDecodeJSON{
			FieldName:   nameStr,
			ValueType:   "string",
			Value:       result,
			FieldBranch: fieldBranch,
		}

		return result
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		result := reflect.ValueOf(anyType).Int()

		chanOutMispFormat <- datamodel.ChanOutputDecodeJSON{
			FieldName:   nameStr,
			ValueType:   "int",
			Value:       result,
			FieldBranch: fieldBranch,
		}

		return result
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result := reflect.ValueOf(anyType).Uint()

		chanOutMispFormat <- datamodel.ChanOutputDecodeJSON{
			FieldName:   nameStr,
			ValueType:   "uint",
			Value:       result,
			FieldBranch: fieldBranch,
		}

		return result
	case reflect.Float32, reflect.Float64:
		result := reflect.ValueOf(anyType).Float()

		chanOutMispFormat <- datamodel.ChanOutputDecodeJSON{
			FieldName:   nameStr,
			ValueType:   "float",
			Value:       result,
			FieldBranch: fieldBranch,
		}

		return result
	case reflect.Bool:
		result := reflect.ValueOf(anyType).Bool()

		chanOutMispFormat <- datamodel.ChanOutputDecodeJSON{
			FieldName:   nameStr,
			ValueType:   "bool",
			Value:       result,
			FieldBranch: fieldBranch,
		}

		return result
	}

	return anyType
}

func reflectMap(
	logging chan<- datamodels.MessageLogging,
	chanOutMispFormat chan<- datamodel.ChanOutputDecodeJSON,
	l map[string]interface{},
	fieldBranch string) map[string]interface{} {

	var (
		newMap  map[string]interface{}
		newList []interface{}
	)
	nl := map[string]interface{}{}

	for k, v := range l {
		var fbTmp string
		r := reflect.TypeOf(v)

		if r == nil {
			return nl
		}

		fbTmp = fieldBranch
		if fbTmp == "" {
			fbTmp += k
		} else {
			fbTmp += "." + k
		}

		switch r.Kind() {
		case reflect.Map:
			if v, ok := v.(map[string]interface{}); ok {
				newMap = reflectMap(logging, chanOutMispFormat, v, fbTmp)
				nl[k] = newMap
			}

		case reflect.Slice:
			if v, ok := v.([]interface{}); ok {
				newList = reflectSlice(logging, chanOutMispFormat, v, fbTmp)
				nl[k] = newList
			}

		default:
			nl[k] = reflectAnySimpleType(logging, chanOutMispFormat, k, v, fbTmp)
		}
	}

	return nl
}

func reflectSlice(
	logging chan<- datamodels.MessageLogging,
	chanOutMispFormat chan<- datamodel.ChanOutputDecodeJSON,
	l []interface{},
	fieldBranch string) []interface{} {

	var (
		newMap  map[string]interface{}
		newList []interface{}
	)
	nl := make([]interface{}, 0, len(l))

	for k, v := range l {
		r := reflect.TypeOf(v)

		if r == nil {
			return nl
		}

		switch r.Kind() {
		case reflect.Map:
			if v, ok := v.(map[string]interface{}); ok {
				newMap = reflectMap(logging, chanOutMispFormat, v, fieldBranch)

				nl = append(nl, newMap)
			}

		case reflect.Slice:
			if v, ok := v.([]interface{}); ok {
				newList = reflectSlice(logging, chanOutMispFormat, v, fieldBranch)

				nl = append(nl, newList...)
			}

		default:
			nl = append(nl, reflectAnySimpleType(logging, chanOutMispFormat, k, v, fieldBranch))
		}
	}

	return nl
}
