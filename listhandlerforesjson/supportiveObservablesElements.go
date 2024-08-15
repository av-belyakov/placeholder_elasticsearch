package listhandlerforesjson

import (
	"fmt"

	"placeholder_elasticsearch/datamodels"
)

// SupportiveObservables вспомогательный тип для для обработки observables
type SupportiveObservables struct {
	currentKey                string
	listAcceptedFields        []string
	listAcceptedFieldsReports []string
	observableTmp             datamodels.ObservableMessageEs
	observables               map[string][]datamodels.ObservableMessageEs
}

// NewSupportiveObservables формирует вспомогательный объект для обработки
// thehive объектов типа observables
func NewSupportiveObservables() *SupportiveObservables {
	return &SupportiveObservables{
		listAcceptedFields:        []string(nil),
		listAcceptedFieldsReports: []string(nil),
		observableTmp:             *datamodels.NewObservableMessageEs(),
		observables:               make(map[string][]datamodels.ObservableMessageEs),
	}
}

// GetObservables возвращает map[string][]datamodels.ObservableMessage, однако, метод
// выполняет еще очень важное действие, перемещает содержимое из o.observableTmp в
// o.observables, так как observables автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются observableTmp
func (o *SupportiveObservables) GetObservables() map[string][]datamodels.ObservableMessageEs {
	o.listAcceptedFields = []string(nil)

	if o.currentKey != "" {
		_, _ = datamodels.PostProcessingUserType[*datamodels.ObservableMessageEs](&o.observableTmp)
		o.observables[o.currentKey] = append(o.observables[o.currentKey], o.observableTmp)
	}

	o.currentKey = ""
	o.observableTmp = *datamodels.NewObservableMessageEs()

	return o.observables
}

// GetObservableTmp возвращает временный объект observable
func (o *SupportiveObservables) GetObservableTmp() *datamodels.ObservableMessageEs {
	return &o.observableTmp
}

func (o *SupportiveObservables) HandlerValue(fieldBranch string, i interface{}, f func(interface{})) {
	if fieldBranch == "observables.dataType" {
		str := fmt.Sprint(i)
		if _, ok := o.observables[str]; !ok {
			o.observables[str] = []datamodels.ObservableMessageEs(nil)
		}

		if o.isExistFieldBranch(fieldBranch) {
			o.listAcceptedFields = []string(nil)
			_, _ = datamodels.PostProcessingUserType[*datamodels.ObservableMessageEs](&o.observableTmp)
			o.observables[o.currentKey] = append(o.observables[o.currentKey], o.observableTmp)

			o.observableTmp = *datamodels.NewObservableMessageEs()
		}

		o.currentKey = str
	}

	//если поле повторяется то считается что это уже новый объект
	if fieldBranch != "observables.tags" && fieldBranch != "observables.attachment.hashes" && o.isExistFieldBranch((fieldBranch)) {
		o.listAcceptedFields = []string(nil)

		if _, ok := o.observables[o.currentKey]; !ok {
			o.observables[o.currentKey] = []datamodels.ObservableMessageEs(nil)
		}

		_, _ = datamodels.PostProcessingUserType[*datamodels.ObservableMessageEs](&o.observableTmp)
		o.observables[o.currentKey] = append(o.observables[o.currentKey], o.observableTmp)

		o.observableTmp = *datamodels.NewObservableMessageEs()
	}

	o.listAcceptedFields = append(o.listAcceptedFields, fieldBranch)

	f(i)
}

func (o *SupportiveObservables) isExistFieldBranch(value string) bool {
	for _, v := range o.listAcceptedFields {
		if v == value {
			return true
		}
	}

	return false
}

/*
func (o *SupportiveObservables) HandlerReportValue(fieldBranch string, i interface{}) {
	fields := strings.Split(fieldBranch, ".")
	if len(fields) != 5 {
		return
	}

	//пока обрабатываем только taxonomies
	if fields[3] != "taxonomies" {
		return
	}

	if _, ok := o.observableTmp.Reports[fields[2]]; !ok {
		o.observableTmp.Reports[fields[2]] = datamodels.ReportTaxonomies{Taxonomies: make([]datamodels.Taxonomy, 1)}
		o.previousFieldReports = fields[2]
		o.listAcceptedFieldsReports = []string{}
	}

	//для того чтобы понять нужно ли создавать новый элемент среза
	//используем хранилище listAcceptedFieldsReports для временного хранения
	//наименований полей, создаем новый элемент среза, если попадается
	// повторяющееся свойство структуры Taxonomy
	if o.previousFieldReports == fields[2] && o.isExistFieldBranchReports(fields[4]) {
		tmpSlice := o.observableTmp.Reports[fields[2]]
		tmpSlice.Taxonomies = append(tmpSlice.Taxonomies, datamodels.Taxonomy{})
		o.observableTmp.Reports[fields[2]] = tmpSlice

		o.listAcceptedFieldsReports = []string{}
	}

	o.listAcceptedFieldsReports = append(o.listAcceptedFieldsReports, fields[4])
	lastNum := len(o.observableTmp.Reports[fields[2]].Taxonomies) - 1
	if lastNum < 0 {
		lastNum = 0
	}

	switch fields[4] {
	case "level":
		o.observableTmp.Reports[fields[2]].Taxonomies[lastNum].SetAnyLevel(i)

	case "namespace":
		o.observableTmp.Reports[fields[2]].Taxonomies[lastNum].SetAnyNamespace(i)

	case "predicate":
		o.observableTmp.Reports[fields[2]].Taxonomies[lastNum].SetAnyPredicate(i)

	case "value":
		o.observableTmp.Reports[fields[2]].Taxonomies[lastNum].SetAnyValue(i)
	}
}

func (o *SupportiveObservables) isExistFieldBranchReports(value string) bool {
	for _, v := range o.listAcceptedFieldsReports {
		if v == value {
			return true
		}
	}

	return false
}
*/
