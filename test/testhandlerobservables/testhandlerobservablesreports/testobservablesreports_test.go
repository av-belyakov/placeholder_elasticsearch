package testhandlerobservablesreports_test

import (
	"fmt"
	"reflect"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	//"placeholder_elasticsearch/test/testhandlerobservables/testhandlerobservablesreports"
	"placeholder_elasticsearch/datamodels"
	testing "placeholder_elasticsearch/test"
)

//----------- Observables Reports -----------
//----------------- START -------------------

// supportiveObservablesReports вспомогательный тип для формирования объекта типа reports
type supportiveObservablesReports struct {
	currentNum         int
	previousField      string
	listAcceptedFields []string
	reports            map[string]datamodels.ReportTaxonomies
}

func NewSupportiveObservablesReports() *supportiveObservablesReports {
	return &supportiveObservablesReports{
		reports: make(map[string]datamodels.ReportTaxonomies),
	}
}

func (sor *supportiveObservablesReports) HandlerReportValue(fieldBranch string, i interface{}) {
	fields := strings.Split(fieldBranch, ".")
	if len(fields) != 4 {
		return
	}

	//пока обрабатываем только taxonomies
	if fields[2] != "taxonomies" {
		return
	}

	if _, ok := sor.reports[fields[1]]; !ok {
		sor.reports[fields[1]] = datamodels.ReportTaxonomies{Taxonomies: make([]datamodels.Taxonomy, 1)}
		sor.previousField = fields[1]
		sor.listAcceptedFields = []string{}
	}

	//для того чтобы понять нужно ли создавать новый элемент среза
	//используем хранилище listAcceptedFields для временного хранения
	//наименований полей, создаем новый элемент среза, если попадается
	// повторяющееся свойство структуры Taxonomy
	if sor.previousField == fields[1] && sor.isExistFieldBranch(fields[3]) {
		tmpSlice := sor.reports[fields[1]]
		tmpSlice.Taxonomies = append(tmpSlice.Taxonomies, datamodels.Taxonomy{})
		sor.reports[fields[1]] = tmpSlice

		sor.listAcceptedFields = []string{}
	}

	sor.listAcceptedFields = append(sor.listAcceptedFields, fields[3])
	lastNum := len(sor.reports[fields[1]].Taxonomies) - 1
	if lastNum < 0 {
		lastNum = 0
	}

	switch fields[3] {
	case "level":
		sor.reports[fields[1]].Taxonomies[lastNum].SetAnyLevel(i)

	case "namespace":
		sor.reports[fields[1]].Taxonomies[lastNum].SetAnyNamespace(i)

	case "predicate":
		sor.reports[fields[1]].Taxonomies[lastNum].SetAnyPredicate(i)

	case "value":
		sor.reports[fields[1]].Taxonomies[lastNum].SetAnyValue(i)
	}

}

func (sor *supportiveObservablesReports) GetReports() map[string]datamodels.ReportTaxonomies {
	return sor.reports
}

func (sor *supportiveObservablesReports) isExistFieldBranch(value string) bool {
	for _, v := range sor.listAcceptedFields {
		if v == value {
			return true
		}
	}

	return false
}

//---------- Observables Reports ----------
//----------------- END -------------------

//-------------- Observables ----------------
//----------------- START -------------------

type supportiveObservables struct {
	currentNum         int
	listAcceptedFields []string
	observableTmp      datamodels.ObservableMessage
	observables        []datamodels.ObservableMessage
}

func NewSupportiveObservables() *supportiveObservables {
	return &supportiveObservables{
		listAcceptedFields: make([]string, 0),
		observableTmp:      datamodels.ObservableMessage{},
		observables:        make([]datamodels.ObservableMessage, 0)}
}

func (o *supportiveObservables) GetCurrentNum() int {
	return o.currentNum
}

func (o *supportiveObservables) HandlerValue(fieldBranch string, i interface{}, f func(interface{})) {
	//если поле повторяется то считается что это уже новый объект
	if fieldBranch != "observables.tags" && fieldBranch != "observables.attachment.hashes" && o.isExistFieldBranch(fieldBranch) {
		o.currentNum += o.currentNum
		o.listAcceptedFields = make([]string, 0)
		o.observables = append(o.observables, o.observableTmp)
		o.observableTmp = datamodels.ObservableMessage{}
	}

	o.listAcceptedFields = append(o.listAcceptedFields, fieldBranch)
	f(i)
}

func (o *supportiveObservables) isExistFieldBranch(value string) bool {
	for _, v := range o.listAcceptedFields {
		if v == value {
			return true
		}
	}

	return false
}

// GetObservables возвращает []datamodels.ObservableMessage, однако, метод
// выполняет еще очень важное действие, перемещает содержимое из o.observableTmp в
// список o.observables, так как observables автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются observableTmp
func (o *supportiveObservables) GetObservables() []datamodels.ObservableMessage {
	o.observables = append(o.observables, o.observableTmp)

	return o.observables
}

//------------- Observables ---------------
//----------------- END -------------------

func newCustomFieldsElement(elem, objType string, customFields *map[string]datamodels.CustomerFields) {
	if _, ok := (*customFields)[elem]; !ok {
		switch objType {
		case "string":
			(*customFields)[elem] = &datamodels.CustomFieldStringType{}
		case "date":
			(*customFields)[elem] = &datamodels.CustomFieldDateType{}
		case "integer":
			(*customFields)[elem] = &datamodels.CustomFieldIntegerType{}
		}
	}
}

var _ = Describe("Testobservablesreports", Ordered, func() {
	var (
		event        datamodels.EventMessageTheHive = datamodels.EventMessageTheHive{}
		eventObject  datamodels.EventObject         = datamodels.EventObject{}
		eventDetails datamodels.EventDetails        = datamodels.EventDetails{}

		eventObjectCustomFields  map[string]datamodels.CustomerFields = make(map[string]datamodels.CustomerFields)
		eventDetailsCustomFields map[string]datamodels.CustomerFields = make(map[string]datamodels.CustomerFields)

		listHandlerEvent                    map[string][]func(interface{})
		listHandlerEventDetails             map[string][]func(interface{})
		listHandlerEventDetailsCustomFields map[string][]func(interface{})

		listHandlerEventObject             map[string][]func(interface{})
		listHandlerEventObjectCustomFields map[string][]func(interface{})

		so                     *supportiveObservables
		sor                    *supportiveObservablesReports
		listHandlerObservables map[string][]func(interface{})
	)

	BeforeAll(func() {
		so = NewSupportiveObservables()
		sor = NewSupportiveObservablesReports()

		listHandlerObservables = map[string][]func(interface{}){
			//--- ioc ---
			"observables.ioc": {func(i interface{}) {
				so.HandlerValue(
					"observables.ioc",
					i,
					so.observableTmp.SetAnyIoc,
				)
			}},
			//--- sighted ---
			"observables.sighted": {func(i interface{}) {
				so.HandlerValue(
					"observables.sighted",
					i,
					so.observableTmp.SetAnySighted,
				)
			}},
			//--- ignoreSimilarity ---
			"observables.ignoreSimilarity": {func(i interface{}) {
				so.HandlerValue(
					"observables.ignoreSimilarity",
					i,
					so.observableTmp.SetAnyIgnoreSimilarity,
				)
			}},
			//--- tlp ---
			"observables.tlp": {func(i interface{}) {
				so.HandlerValue(
					"observables.tlp",
					i,
					so.observableTmp.SetAnyTlp,
				)
			}},
			//--- _createdAt ---
			"observables._createdAt": {func(i interface{}) {
				fmt.Println("-------- _createdAt", i)

				so.HandlerValue(
					"observables._createdAt",
					i,
					so.observableTmp.SetAnyUnderliningCreatedAt,
				)
			}},
			//--- _updatedAt ---
			"observables._updatedAt": {func(i interface{}) {
				fmt.Println("-------- _createdAt", i)

				so.HandlerValue(
					"observables._updatedAt",
					i,
					so.observableTmp.SetAnyUnderliningUpdatedAt,
				)
			}},
			//--- startDate ---
			"observables.startDate": {func(i interface{}) {
				fmt.Println("-------- startDate", i)

				so.HandlerValue(
					"observables.startDate",
					i,
					so.observableTmp.SetAnyStartDate,
				)
			}},
			//--- _createdBy ---
			"observables._createdBy": {func(i interface{}) {
				so.HandlerValue(
					"observables._createdBy",
					i,
					so.observableTmp.SetAnyUnderliningCreatedBy,
				)
			}},
			//--- _updatedBy ---
			"observables._updatedBy": {func(i interface{}) {
				fmt.Println("-------- _updatedBy", i)

				so.HandlerValue(
					"observables._updatedBy",
					i,
					so.observableTmp.SetAnyUnderliningUpdatedBy,
				)
			}},
			//--- _id ---
			"observables._id": {func(i interface{}) {
				so.HandlerValue(
					"observables._id",
					i,
					so.observableTmp.SetAnyUnderliningId,
				)
			}},
			//--- _type ---
			"observables._type": {func(i interface{}) {
				so.HandlerValue(
					"observables._type",
					i,
					so.observableTmp.SetAnyUnderliningType,
				)
			}},
			//--- data ---
			"observables.data": {func(i interface{}) {
				so.HandlerValue(
					"observables.data",
					i,
					so.observableTmp.SetAnyData,
				)
			}},
			//--- dataType ---
			"observables.dataType": {func(i interface{}) {
				so.HandlerValue(
					"observables.dataType",
					i,
					so.observableTmp.SetAnyDataType,
				)
			}},
			//--- message ---
			"observables.message": {func(i interface{}) {
				so.HandlerValue(
					"observables.message",
					i,
					so.observableTmp.SetAnyMessage,
				)
			}},
			//--- tags ---
			"observables.tags": {func(i interface{}) {
				so.HandlerValue(
					"observables.tags",
					i,
					so.observableTmp.SetAnyTags,
				)
			}},

			//--- attachment.id ---
			"observables.attachment.id": {func(i interface{}) {
				so.HandlerValue(
					"observables.attachment.id",
					i,
					so.observableTmp.Attachment.SetAnyId,
				)
			}},
			//--- attachment.size ---
			"observables.attachment.size": {func(i interface{}) {
				so.HandlerValue(
					"observables.attachment.size",
					i,
					so.observableTmp.Attachment.SetAnySize,
				)
			}},
			// --- attachment.name ---
			"observables.attachment.name": {func(i interface{}) {
				so.HandlerValue(
					"observables.attachment.name",
					i,
					so.observableTmp.Attachment.SetAnyName,
				)
			}},
			// --- attachment.contentType ---
			"observables.attachment.contentType": {func(i interface{}) {
				so.HandlerValue(
					"observables.attachment.contentType",
					i,
					so.observableTmp.Attachment.SetAnyContentType,
				)
			}},
			// --- attachment.hashes ---
			"observables.attachment.hashes": {func(i interface{}) {
				so.HandlerValue(
					"observables.attachment.hashes",
					i,
					so.observableTmp.Attachment.SetAnyHashes,
				)
			}},
		}

		// ------ EVENT ------
		listHandlerEvent = map[string][]func(interface{}){
			"event.rootId":         {event.SetAnyRootId},
			"event.objectId":       {event.SetAnyObjectId},
			"event.objectType":     {event.SetAnyObjectType},
			"event.base":           {event.SetAnyBase},
			"event.startDate":      {event.SetAnyStartDate},
			"event.requestId":      {event.SetAnyRequestId},
			"event.organisation":   {event.SetAnyOperation},
			"event.organisationId": {event.SetAnyOrganisationId},
			"event.operation":      {event.SetAnyOperation},
		}

		// ------ EVENT DETAILS ------
		listHandlerEventDetails = map[string][]func(interface{}){
			"event.details.endDate":          {eventDetails.SetAnyEndDate},
			"event.details.resolutionStatus": {eventDetails.SetAnyResolutionStatus},
			"event.details.summary":          {eventDetails.SetAnySummary},
			"event.details.status":           {eventDetails.SetAnyStatus},
			"event.details.impactStatus":     {eventDetails.SetAnyImpactStatus},
		}

		// ------ EVENT DETAILS CUSTOMFIELDS ------
		listHandlerEventDetailsCustomFields = map[string][]func(interface{}){
			//--- attack-type ---
			"event.details.customFields.attack-type.order": {func(i interface{}) {
				//создаем элемент "attack-type" если его нет
				newCustomFieldsElement("attack-type", "string", &eventDetailsCustomFields)
				_, _, _, str := eventDetailsCustomFields["attack-type"].Get()
				eventDetailsCustomFields["attack-type"].Set(i, str)
			}},
			"event.details.customFields.attack-type.string": {func(i interface{}) {
				newCustomFieldsElement("attack-type", "string", &eventDetailsCustomFields)
				_, order, _, _ := eventDetailsCustomFields["attack-type"].Get()
				eventDetailsCustomFields["attack-type"].Set(order, i)
			}},
			//--- class-attack ---
			"event.details.customFields.class-attack.order": {func(i interface{}) {
				newCustomFieldsElement("class-attack", "string", &eventDetailsCustomFields)
				_, _, _, str := eventDetailsCustomFields["class-attack"].Get()
				eventDetailsCustomFields["class-attack"].Set(i, str)
			}},
			"event.details.customFields.class-attack.string": {func(i interface{}) {
				newCustomFieldsElement("class-attack", "string", &eventDetailsCustomFields)
				_, order, _, _ := eventDetailsCustomFields["class-attack"].Get()
				eventDetailsCustomFields["class-attack"].Set(order, i)
			}},
			//--- event-source ---
			"event.details.customFields.event-source.order": {func(i interface{}) {
				newCustomFieldsElement("event-source", "string", &eventDetailsCustomFields)
				_, _, _, str := eventDetailsCustomFields["event-source"].Get()
				eventDetailsCustomFields["event-source"].Set(i, str)
			}},
			"event.details.customFields.event-source.string": {func(i interface{}) {
				newCustomFieldsElement("event-source", "string", &eventDetailsCustomFields)
				_, order, _, _ := eventDetailsCustomFields["event-source"].Get()
				eventDetailsCustomFields["event-source"].Set(order, i)
			}},
			//--- ncircc-class-attack ---
			"event.details.customFields.ncircc-class-attack.order": {func(i interface{}) {
				newCustomFieldsElement("ncircc-class-attack", "string", &eventDetailsCustomFields)
				_, _, _, str := eventDetailsCustomFields["ncircc-class-attack"].Get()
				eventDetailsCustomFields["ncircc-class-attack"].Set(i, str)
			}},
			"event.details.customFields.ncircc-class-attack.string": {func(i interface{}) {
				newCustomFieldsElement("ncircc-class-attack", "string", &eventDetailsCustomFields)
				_, order, _, _ := eventDetailsCustomFields["ncircc-class-attack"].Get()
				eventDetailsCustomFields["ncircc-class-attack"].Set(order, i)
			}},
			//--- ncircc-bulletin-id ---
			"event.details.customFields.ncircc-bulletin-id.order": {func(i interface{}) {
				newCustomFieldsElement("ncircc-bulletin-id", "string", &eventDetailsCustomFields)
				_, _, _, str := eventDetailsCustomFields["ncircc-bulletin-id"].Get()
				eventDetailsCustomFields["ncircc-bulletin-id"].Set(i, str)
			}},
			"event.details.customFields.ncircc-bulletin-id.string": {func(i interface{}) {
				newCustomFieldsElement("ncircc-bulletin-id", "string", &eventDetailsCustomFields)
				_, order, _, _ := eventDetailsCustomFields["ncircc-bulletin-id"].Get()
				eventDetailsCustomFields["ncircc-bulletin-id"].Set(order, i)
			}},
			//--- ir-name ---
			"event.details.customFields.ir-name.order": {func(i interface{}) {
				newCustomFieldsElement("ir-name", "string", &eventDetailsCustomFields)
				_, _, _, str := eventDetailsCustomFields[" "].Get()
				eventDetailsCustomFields[" "].Set(i, str)
			}},
			"event.details.customFields.ir-name.string": {func(i interface{}) {
				newCustomFieldsElement("ir-name", "string", &eventDetailsCustomFields)
				_, order, _, _ := eventDetailsCustomFields["ir-name"].Get()
				eventDetailsCustomFields["ir-name"].Set(order, i)
			}},
			//--- sphere ---
			"event.details.customFields.sphere.order": {func(i interface{}) {
				newCustomFieldsElement("sphere", "string", &eventDetailsCustomFields)
				_, _, _, str := eventDetailsCustomFields["sphere"].Get()
				eventDetailsCustomFields["sphere"].Set(i, str)
			}},
			"event.details.customFields.sphere.string": {func(i interface{}) {
				newCustomFieldsElement("sphere", "string", &eventDetailsCustomFields)
				_, order, _, _ := eventDetailsCustomFields["sphere"].Get()
				eventDetailsCustomFields["sphere"].Set(order, i)
			}},
			//--- id-soa ---
			"event.details.customFields.id-soa.order": {func(i interface{}) {
				newCustomFieldsElement("id-soa", "string", &eventDetailsCustomFields)
				_, _, _, str := eventDetailsCustomFields["id-soa"].Get()
				eventDetailsCustomFields["id-soa"].Set(i, str)
			}},
			"event.details.customFields.id-soa.string": {func(i interface{}) {
				newCustomFieldsElement("id-soa", "string", &eventDetailsCustomFields)
				_, order, _, _ := eventDetailsCustomFields["id-soa"].Get()
				eventDetailsCustomFields["id-soa"].Set(order, i)
			}},
			//--- state ---
			"event.details.customFields.state.order": {func(i interface{}) {
				newCustomFieldsElement("state", "string", &eventDetailsCustomFields)
				_, _, _, str := eventDetailsCustomFields["state"].Get()
				eventDetailsCustomFields["state"].Set(i, str)
			}},
			"event.details.customFields.state.string": {func(i interface{}) {
				newCustomFieldsElement("state", "string", &eventDetailsCustomFields)
				_, order, _, _ := eventDetailsCustomFields["state"].Get()
				eventDetailsCustomFields["state"].Set(order, i)
			}},
			//--- external-letter ---
			"event.details.customFields.external-letter.order": {func(i interface{}) {
				newCustomFieldsElement("external-letter", "string", &eventDetailsCustomFields)
				_, _, _, str := eventDetailsCustomFields["external-letter"].Get()
				eventDetailsCustomFields["external-letter"].Set(i, str)
			}},
			//--- inbox1 ---
			"event.details.customFields.inbox1.order": {func(i interface{}) {
				newCustomFieldsElement("inbox1", "string", &eventDetailsCustomFields)
				_, _, _, str := eventDetailsCustomFields["inbox1"].Get()
				eventDetailsCustomFields["inbox1"].Set(i, str)
			}},
			//--- inner-letter ---
			"event.details.customFields.inner-letter.order": {func(i interface{}) {
				newCustomFieldsElement("inner-letter", "string", &eventDetailsCustomFields)
				_, _, _, str := eventDetailsCustomFields["inner-letter"].Get()
				eventDetailsCustomFields["inner-letter"].Set(i, str)
			}},
			//--- notification ---
			"event.details.customFields.notification.order": {func(i interface{}) {
				newCustomFieldsElement("notification", "string", &eventDetailsCustomFields)
				_, _, _, str := eventDetailsCustomFields["notification"].Get()
				eventDetailsCustomFields["notification"].Set(i, str)
			}},
			//--- report ---
			"event.details.customFields.report.order": {func(i interface{}) {
				newCustomFieldsElement("report", "string", &eventDetailsCustomFields)
				_, _, _, str := eventDetailsCustomFields["report"].Get()
				eventDetailsCustomFields["report"].Set(i, str)
			}},
			//--- first-time ---
			"event.details.customFields.first-time.order": {func(i interface{}) {
				newCustomFieldsElement("first-time", "string", &eventDetailsCustomFields)
				_, _, _, str := eventDetailsCustomFields["first-time"].Get()
				eventDetailsCustomFields["first-time"].Set(i, str)
			}},
			"event.details.customFields.first-time.date": {func(i interface{}) {
				newCustomFieldsElement("first-time", "date", &eventDetailsCustomFields)
				_, order, _, _ := eventDetailsCustomFields["first-time"].Get()
				eventDetailsCustomFields["first-time"].Set(order, i)
			}},
			//--- last-time ---
			"event.details.customFields.last-time.order": {func(i interface{}) {
				newCustomFieldsElement("last-time", "string", &eventDetailsCustomFields)
				_, _, _, str := eventDetailsCustomFields["last-time"].Get()
				eventDetailsCustomFields["last-time"].Set(i, str)
			}},
			"event.details.customFields.last-time.date": {func(i interface{}) {
				newCustomFieldsElement("last-time", "date", &eventDetailsCustomFields)
				_, order, _, _ := eventDetailsCustomFields["last-time"].Get()
				eventDetailsCustomFields["last-time"].Set(order, i)
			}},
			//--- b2mid ---
			"event.details.customFields.b2mid.order": {func(i interface{}) {
				newCustomFieldsElement("b2mid", "integer", &eventDetailsCustomFields)
				_, _, _, str := eventDetailsCustomFields["b2mid"].Get()
				eventDetailsCustomFields["b2mid"].Set(i, str)
			}},
			"event.details.customFields.b2mid.integer": {func(i interface{}) {
				newCustomFieldsElement("b2mid", "integer", &eventDetailsCustomFields)
				_, order, _, _ := eventDetailsCustomFields["b2mid"].Get()
				eventDetailsCustomFields["b2mid"].Set(order, i)
			}},
		}

		// ------ EVENT OBJECT ------
		listHandlerEventObject = map[string][]func(interface{}){
			"event.object.flag":             {eventObject.SetAnyFlag},
			"event.object.caseId":           {eventObject.SetAnyCaseId},
			"event.object.severity":         {eventObject.SetAnySeverity},
			"event.object.tlp":              {eventObject.SetAnyTlp},
			"event.object.pap":              {eventObject.SetAnyPap},
			"event.object.startDate":        {eventObject.SetAnyStartDate},
			"event.object.endDate":          {eventObject.SetAnyEndDate},
			"event.object.createdAt":        {eventObject.SetAnyCreatedAt},
			"event.object.updatedAt":        {eventObject.SetAnyUpdatedAt},
			"event.object._id":              {eventObject.SetAnyUnderliningId},
			"event.object.id":               {eventObject.SetAnyId},
			"event.object.createdBy":        {eventObject.SetAnyCreatedBy},
			"event.object.updatedBy":        {eventObject.SetAnyUpdatedBy},
			"event.object._type":            {eventObject.SetAnyUnderliningType},
			"event.object.title":            {eventObject.SetAnyTitle},
			"event.object.description":      {eventObject.SetAnyDescription},
			"event.object.impactStatus":     {eventObject.SetAnyImpactStatus},
			"event.object.resolutionStatus": {eventObject.SetAnyResolutionStatus},
			"event.object.status":           {eventObject.SetAnyStatus},
			"event.object.summary":          {eventObject.SetAnySummary},
			"event.object.owner":            {eventObject.SetAnyOwner},
			"event.object.tags":             {eventObject.SetAnyTags},

			//ниже следующие поля редко используются, думаю пока они не требуют реализации
			//"event.object.stats.impactStatus":    {},
			//"event.object.permissions.id":        {},
			//"event.object.permissions.createdAt": {},
			//"event.object.permissions.pap":       {},
		}

		// ------ EVENT OBJECT CUSTOMFIELDS ------
		listHandlerEventObjectCustomFields = map[string][]func(interface{}){
			//--- attack-type ---
			"event.object.customFields.attack-type.order": {func(i interface{}) {
				//создаем элемент "attack-type" если его нет
				newCustomFieldsElement("attack-type", "string", &eventObjectCustomFields)
				_, _, _, str := eventObjectCustomFields["attack-type"].Get()
				eventObjectCustomFields["attack-type"].Set(i, str)
			}},
			"event.object.customFields.attack-type.string": {func(i interface{}) {
				newCustomFieldsElement("attack-type", "string", &eventObjectCustomFields)
				_, order, _, _ := eventObjectCustomFields["attack-type"].Get()
				eventObjectCustomFields["attack-type"].Set(order, i)
			}},
			//--- class-attack ---
			"event.object.customFields.class-attack.order": {func(i interface{}) {
				newCustomFieldsElement("class-attack", "string", &eventObjectCustomFields)
				_, _, _, str := eventObjectCustomFields["class-attack"].Get()
				eventObjectCustomFields["class-attack"].Set(i, str)
			}},
			"event.object.customFields.class-attack.string": {func(i interface{}) {
				newCustomFieldsElement("class-attack", "string", &eventObjectCustomFields)
				_, order, _, _ := eventObjectCustomFields["class-attack"].Get()
				eventObjectCustomFields["class-attack"].Set(order, i)
			}},
			//--- ncircc-class-attack ---
			"event.object.customFields.ncircc-class-attack.order": {func(i interface{}) {
				newCustomFieldsElement("ncircc-class-attack", "string", &eventObjectCustomFields)
				_, _, _, str := eventObjectCustomFields["ncircc-class-attack"].Get()
				eventObjectCustomFields["ncircc-class-attack"].Set(i, str)
			}},
			"event.object.customFields.ncircc-class-attack.string": {func(i interface{}) {
				newCustomFieldsElement("ncircc-class-attack", "string", &eventObjectCustomFields)
				_, order, _, _ := eventObjectCustomFields["ncircc-class-attack"].Get()
				eventObjectCustomFields["ncircc-class-attack"].Set(order, i)
			}},
			//--- inbox1 ---
			"event.object.customFields.inbox1.order": {func(i interface{}) {
				newCustomFieldsElement("inbox1", "string", &eventObjectCustomFields)
				_, _, _, str := eventObjectCustomFields["inbox1"].Get()
				eventObjectCustomFields["inbox1"].Set(i, str)
			}},
			//--- inner-letter ---
			"event.object.customFields.inner-letter.order": {func(i interface{}) {
				newCustomFieldsElement("inner-letter", "string", &eventObjectCustomFields)
				_, _, _, str := eventObjectCustomFields["inner-letter"].Get()
				eventObjectCustomFields["inner-letter"].Set(i, str)
			}},
			//--- notification ---
			"event.object.customFields.notification.order": {func(i interface{}) {
				newCustomFieldsElement("notification", "string", &eventObjectCustomFields)
				_, _, _, str := eventObjectCustomFields["notification"].Get()
				eventObjectCustomFields["notification"].Set(i, str)
			}},
			//--- report ---
			"event.object.customFields.report.order": {func(i interface{}) {
				newCustomFieldsElement("report", "string", &eventObjectCustomFields)
				_, _, _, str := eventObjectCustomFields["report"].Get()
				eventObjectCustomFields["report"].Set(i, str)
			}},
		}

	})

	Context("Тест 1. Проверка заполнения вспомогательного объекта для хранения observables reports", func() {
		It("Вспомогательный объект должен быть успешно заполнен", func() {
			for _, v := range testing.GetReportsList() {
				//fmt.Printf("%d. %s: %s\n", k, v.ElemName, v.ElemValue)

				sor.HandlerReportValue(v.ElemName, v.ElemValue)
			}

			/**********************
			***********************
			***********************

				Вопрос с заполнением вспомогательного объекта решен.
			Однако, теперь надо продумать как этот вспомогательный объект
			ассоциировать с НОМЕРОМ текущего observables

			***********************
			***********************
			***********************/

			fmt.Println(sor.reports)

			Expect(len(sor.reports)).Should(Equal(4))
		})
	})

	Context("Тест 2. Проверка заполнения  объекта для хранения observables", func() {
		It("Вспомогательный объект должен быть успешно заполнен", func() {
			// ------ OBSERVABLES ------
			for _, v := range testing.GetObservableOne() {
				if lf, ok := listHandlerObservables[v.ElemName]; ok {
					//fmt.Printf("%d. v.ElemName: %s, v.ElemValue: '%v'\n", k, v.ElemName, v.ElemValue)

					for _, f := range lf {
						r := reflect.TypeOf(v.ElemValue)
						switch r.Kind() {
						case reflect.Slice:
							//fmt.Println("v.ElemName:", v.ElemName, "00000000000000000000 v.ElemValue:", v.ElemValue)

							if s, ok := v.ElemValue.([]interface{}); ok {
								for _, value := range s {
									f(value)
								}
							}
						default:
							f(v.ElemValue)

						}
					}
				}
			}

			observables := so.GetObservables()[0]

			//fmt.Println("Observables:", observables)

			Expect(observables.GetUnderliningId()).Should(Equal("~3460985064"))
			Expect(observables.GetData()).Should(Equal("9608643"))
			Expect(observables.GetDataType()).Should(Equal("snort_sid"))
			Expect(observables.GetIoc()).Should(BeTrue())
			Expect(observables.GetUnderliningType()).Should(Equal("Observable"))
			Expect(observables.GetTlp()).Should(Equal(uint64(2)))
			Expect(observables.GetUnderliningCreatedAt()).Should(Equal(uint64(1690968664227)))
			Expect(observables.GetUnderliningCreatedBy()).Should(Equal("uds@crimea-rcm"))
			Expect(observables.GetUnderliningUpdatedAt()).Should(Equal(uint64(1704977151860)))
			Expect(observables.GetSighted()).ShouldNot(BeTrue())
			Expect(observables.GetStartDate()).Should(Equal(uint64(1690968664227)))
			Expect(len(observables.GetTags())).Should(Equal(2))

			Expect(observables.GetAttachment().GetContentType()).Should(Equal("text/plain"))
			Expect(observables.GetAttachment().GetId()).Should(Equal("c29438b04791184d3eba39bdb7cf99560ab62068fee9509d50cf59723c398ac1"))
			Expect(observables.GetAttachment().GetName()).Should(Equal("n[n.txt"))
			Expect(observables.GetAttachment().GetSize()).Should(Equal(uint64(817)))
			Expect(len(observables.GetAttachment().GetHashes())).Should(Equal(3))
			//Expect(observables.).Should(Equal(""))
		})
	})

	Context("Тест 3. Проверка заполнения объекта для хранения events", func() {
		It("Вспомогательный объект должен быть успешно заполнен", func() {
			for k, v := range testing.GetEventOne() {
				//Simple elements
				if lf, ok := listHandlerEvent[v.ElemName]; ok {
					fmt.Printf("%d. (event) v.ElemName: %s, v.ElemValue: '%v'\n", k, v.ElemName, v.ElemValue)

					for _, f := range lf {
						f(v.ElemValue)
					}
				}

				//В первом тесте надо доделать

				/*
					listHandlerEventDetails             map[string][]func(interface{})
					listHandlerEventDetailsCustomFields map[string][]func(interface{})

					listHandlerEventObject             map[string][]func(interface{})
					listHandlerEventObjectCustomFields map[string][]func(interface{})
				*/

				//event.object

				//event.object.customFields

				//event.details

				//event.details.customFields
			}
		})
	})
})
