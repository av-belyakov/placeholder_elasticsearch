package listhandlerthehivejson

import (
	"strings"

	"placeholder_elasticsearch/datamodels"
)

// SupportiveObservablesReports вспомогательный тип для формирования объекта типа reports
type SupportiveObservablesReports struct {
	previousField      string
	listAcceptedFields []string
	reports            map[string]datamodels.ReportTaxonomies
}

func NewSupportiveObservablesReports() *SupportiveObservablesReports {
	return &SupportiveObservablesReports{
		reports: make(map[string]datamodels.ReportTaxonomies),
	}
}

func (sor *SupportiveObservablesReports) HandlerReportValue(fieldBranch string, i interface{}) {
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

func (sor *SupportiveObservablesReports) GetReports() map[string]datamodels.ReportTaxonomies {
	return sor.reports
}

func (sor *SupportiveObservablesReports) isExistFieldBranch(value string) bool {
	for _, v := range sor.listAcceptedFields {
		if v == value {
			return true
		}
	}

	return false
}
