package listhandlerthehivejson

import (
	"placeholder_elasticsearch/datamodels"
)

var fieldsRepresentedAsList []string

func isExistFieldsRepresentedAsList(field string, list []string) bool {
	for _, v := range list {
		if field == v {
			return true
		}
	}

	return false
}

func init() {
	fieldsRepresentedAsList = []string{
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
}

// SupportiveTtp вспомогательный тип для для обработки ttp
type SupportiveTtp struct {
	listAcceptedFields []string
	ttpTmp             datamodels.TtpMessage
	ttps               []datamodels.TtpMessage
}

// NewSupportiveTtp формирует вспомогательный объект для обработки
// thehive объектов типа ttp
func NewSupportiveTtp() *SupportiveTtp {
	return &SupportiveTtp{
		listAcceptedFields: []string(nil),
		ttpTmp: datamodels.TtpMessage{
			OccurDate:            "1970-01-01T03:00:00+03:00",
			UnderliningCreatedAt: "1970-01-01T03:00:00+03:00",
			ExtraData: datamodels.ExtraDataTtpMessage{
				Pattern: datamodels.PatternExtraData{
					UnderliningCreatedAt: "1970-01-01T03:00:00+03:00",
				},
				PatternParent: datamodels.PatternExtraData{
					UnderliningCreatedAt: "1970-01-01T03:00:00+03:00",
				},
			},
		},
		ttps: []datamodels.TtpMessage(nil),
	}
}

// GetTtps возвращает []datamodels.TtpMessage, однако, метод
// выполняет еще очень важное действие, перемещает содержимое из sttp.ttpTmp в
// список sttp.ttps, так как ttps автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются sttp.ttpTmp
func (sttp *SupportiveTtp) GetTtps() []datamodels.TtpMessage {
	sttp.listAcceptedFields = []string(nil)
	sttp.ttps = append(sttp.ttps, sttp.ttpTmp)

	return sttp.ttps
}

// GetTtpTmp возвращает временный объект ttpTmp
func (sttp *SupportiveTtp) GetTtpTmp() *datamodels.TtpMessage {
	return &sttp.ttpTmp
}

func (sttp *SupportiveTtp) HandlerValue(fieldBranch string, i interface{}, f func(interface{})) {
	//если поле повторяется то считается что это уже новый объект
	isExist := isExistFieldsRepresentedAsList(fieldBranch, fieldsRepresentedAsList)

	if !isExist && sttp.isExistFieldBranch(fieldBranch) {
		sttp.listAcceptedFields = []string(nil)
		sttp.ttps = append(sttp.ttps, sttp.ttpTmp)
		sttp.ttpTmp = datamodels.TtpMessage{
			OccurDate:            "1970-01-01T03:00:00+03:00",
			UnderliningCreatedAt: "1970-01-01T03:00:00+03:00",
			ExtraData: datamodels.ExtraDataTtpMessage{
				Pattern: datamodels.PatternExtraData{
					UnderliningCreatedAt: "1970-01-01T03:00:00+03:00",
					Platforms:            []string(nil),
					PermissionsRequired:  []string(nil),
					DataSources:          []string(nil),
					Tactics:              []string(nil),
				},
				PatternParent: datamodels.PatternExtraData{
					UnderliningCreatedAt: "1970-01-01T03:00:00+03:00",
					Platforms:            []string(nil),
					PermissionsRequired:  []string(nil),
					DataSources:          []string(nil),
					Tactics:              []string(nil),
				},
			},
		}
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
