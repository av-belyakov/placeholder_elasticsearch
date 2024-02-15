package listhandlerthehivejson

import "placeholder_elasticsearch/datamodels"

// SupportiveAlertArtifacts вспомогательный тип для для обработки alert.artifacts
type SupportiveAlertArtifacts struct {
	listAcceptedFields []string
	artifactTmp        datamodels.AlertArtifact
	artifacts          []datamodels.AlertArtifact
}

// NewSupportiveObservables формирует вспомогательный объект для обработки
// thehive объектов типа alert.artifacts
func NewSupportiveAlertArtifacts() *SupportiveAlertArtifacts {
	return &SupportiveAlertArtifacts{
		listAcceptedFields: []string(nil),
		artifactTmp:        *datamodels.NewAlertArtifact(),
		artifacts:          []datamodels.AlertArtifact(nil),
	}
}

// GetArtifacts возвращает []datamodels.AlertArtifact, однако, метод
// выполняет еще очень важное действие, перемещает содержимое из a.artifactTmp в
// список a.artifacts, так как artifacts автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются artifactTmp
func (a *SupportiveAlertArtifacts) GetArtifacts() []datamodels.AlertArtifact {
	a.listAcceptedFields = []string(nil)
	a.artifacts = append(a.artifacts, a.artifactTmp)

	return a.artifacts
}

// GetArtifactTmp возвращает временный объект artifact
func (a *SupportiveAlertArtifacts) GetArtifactTmp() *datamodels.AlertArtifact {
	return &a.artifactTmp
}

func (a *SupportiveAlertArtifacts) HandlerValue(fieldBranch string, i interface{}, f func(interface{})) {
	//если поле повторяется то считается что это уже новый объект
	if fieldBranch != "alert.artifacts.tags" && a.isExistFieldBranch(fieldBranch) {
		a.listAcceptedFields = []string(nil)
		a.artifacts = append(a.artifacts, a.artifactTmp)
		a.artifactTmp = datamodels.AlertArtifact{
			CreatedAt: "1970-01-01T03:00:00+03:00",
			UpdatedAt: "1970-01-01T03:00:00+03:00",
			StartDate: "1970-01-01T03:00:00+03:00",
		}
	}

	a.listAcceptedFields = append(a.listAcceptedFields, fieldBranch)

	f(i)
}

func (a *SupportiveAlertArtifacts) isExistFieldBranch(value string) bool {
	for _, v := range a.listAcceptedFields {
		if v == value {
			return true
		}
	}

	return false
}
