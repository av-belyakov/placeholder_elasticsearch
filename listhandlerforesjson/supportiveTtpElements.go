package listhandlerforesjson

import (
	"fmt"
	"placeholder_elasticsearch/datamodels"
)

var fieldsRepresentedAsList []string = []string{
	"ttp.extraData.pattern.tactics",
	"ttp.extraData.pattern.platforms",
	"ttp.extraData.pattern.dataSources",
	"ttp.extraData.pattern.systemRequirements",
	"ttp.extraData.pattern.permissionsRequired",
	"ttp.extraData.patternParent.tactics",
	"ttp.extraData.patternParent.platforms",
	"ttp.extraData.patternParent.dataSources",
	"ttp.extraData.patternParent.systemRequirements",
	"ttp.extraData.patternParent.permissionsRequired",
}

// SupportiveTtp вспомогательный тип для для обработки ttp
type SupportiveTtp struct {
	currentKey         string
	listAcceptedFields []string
	ttpTmp             datamodels.TtpMessage
	ttps               map[string][]datamodels.TtpMessage
}

// NewSupportiveTtp формирует вспомогательный объект для обработки
// thehive объектов типа ttp
func NewSupportiveTtp() *SupportiveTtp {
	return &SupportiveTtp{
		listAcceptedFields: []string(nil),
		ttpTmp:             *datamodels.NewTtpMessage(),
		ttps:               make(map[string][]datamodels.TtpMessage),
	}
}

// GetTtps возвращает []datamodels.TtpMessage, однако, метод
// выполняет еще очень важное действие, перемещает содержимое из sttp.ttpTmp в
// список sttp.ttps, так как ttps автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются sttp.ttpTmp
func (sttp *SupportiveTtp) GetTtps() map[string][]datamodels.TtpMessage {
	sttp.listAcceptedFields = []string(nil)

	if sttp.currentKey != "" {
		sttp.ttps[sttp.currentKey] = append(sttp.ttps[sttp.currentKey], sttp.ttpTmp)
	}

	sttp.currentKey = ""
	sttp.ttpTmp = *datamodels.NewTtpMessage()

	return sttp.ttps
}

// GetTtpTmp возвращает временный объект ttpTmp
func (sttp *SupportiveTtp) GetTtpTmp() *datamodels.TtpMessage {
	return &sttp.ttpTmp
}

func (sttp *SupportiveTtp) HandlerValue(fieldBranch string, i interface{}, f func(interface{})) {
	if fieldBranch == "ttp.tactic" {
		str := fmt.Sprint(i)
		if _, ok := sttp.ttps[str]; !ok {
			sttp.ttps[str] = []datamodels.TtpMessage(nil)
		}

		if sttp.isExistFieldBranch(fieldBranch) {
			sttp.listAcceptedFields = []string(nil)
			//_, _ = datamodels.PostProcessingUserType[*datamodels.ObservableMessageEs](&o.observableTmp)
			sttp.ttps[sttp.currentKey] = append(sttp.ttps[sttp.currentKey], sttp.ttpTmp)

			sttp.ttpTmp = *datamodels.NewTtpMessage()
		}

		sttp.currentKey = str
	}

	//если поле повторяется то считается что это уже новый объект
	isExist := isExistFieldsRepresentedAsList(fieldBranch, fieldsRepresentedAsList)
	if !isExist && sttp.isExistFieldBranch(fieldBranch) {
		sttp.listAcceptedFields = []string(nil)

		if _, ok := sttp.ttps[sttp.currentKey]; !ok {
			sttp.ttps[sttp.currentKey] = []datamodels.TtpMessage(nil)
		}

		sttp.ttps[sttp.currentKey] = append(sttp.ttps[sttp.currentKey], sttp.ttpTmp)

		sttp.ttpTmp = *datamodels.NewTtpMessage()
	}

	sttp.listAcceptedFields = append(sttp.listAcceptedFields, fieldBranch)

	f(i)
}

func (sttp SupportiveTtp) isExistFieldBranch(value string) bool {
	for _, v := range sttp.listAcceptedFields {
		if v == value {
			return true
		}
	}

	return false
}

func isExistFieldsRepresentedAsList(field string, list []string) bool {
	for _, v := range list {
		if field == v {
			return true
		}
	}

	return false
}
