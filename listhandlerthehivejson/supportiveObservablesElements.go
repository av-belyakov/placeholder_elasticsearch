package listhandlerthehivejson

import "placeholder_elasticsearch/datamodels"

type SupportiveObservables struct {
	currentNum         int
	listAcceptedFields []string
	observableTmp      datamodels.ObservableMessage
	observables        []datamodels.ObservableMessage
}

func NewSupportiveObservables() *SupportiveObservables {
	return &SupportiveObservables{
		listAcceptedFields: make([]string, 0),
		observableTmp:      datamodels.ObservableMessage{},
		observables:        make([]datamodels.ObservableMessage, 0)}
}

func (o *SupportiveObservables) GetCurrentNum() int {
	return o.currentNum
}

// GetObservables возвращает []datamodels.ObservableMessage, однако, метод
// выполняет еще очень важное действие, перемещает содержимое из o.observableTmp в
// список o.observables, так как observables автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются observableTmp
func (o *SupportiveObservables) GetObservables() []datamodels.ObservableMessage {
	o.observables = append(o.observables, o.observableTmp)

	return o.observables
}

// GetObservableTmp возвращает временный объект observable
func (o *SupportiveObservables) GetObservableTmp() *datamodels.ObservableMessage {
	return &o.observableTmp
}

func (o *SupportiveObservables) HandlerValue(fieldBranch string, i interface{}, f func(interface{})) {
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

func (o *SupportiveObservables) isExistFieldBranch(value string) bool {
	for _, v := range o.listAcceptedFields {
		if v == value {
			return true
		}
	}

	return false
}
