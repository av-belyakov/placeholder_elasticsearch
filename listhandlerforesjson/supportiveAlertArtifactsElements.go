package listhandlerforesjson

import (
	"fmt"

	"placeholder_elasticsearch/datamodels"
)

// SupportiveAlertArtifacts вспомогательный тип для для обработки alert.artifacts
type SupportiveAlertArtifacts struct {
	currentKey         string
	listAcceptedFields []string
	artifactTmp        datamodels.ArtifactForEsAlert
	artifacts          map[string][]datamodels.ArtifactForEsAlert
}

// NewSupportiveObservables формирует вспомогательный объект для обработки
// thehive объектов типа alert.artifacts
func NewSupportiveAlertArtifacts() *SupportiveAlertArtifacts {
	return &SupportiveAlertArtifacts{
		listAcceptedFields: []string(nil),
		artifactTmp:        *datamodels.NewArtifactForEsAlert(),
		artifacts:          make(map[string][]datamodels.ArtifactForEsAlert),
	}
}

// GetArtifacts возвращает map[string][]datamodels.AlertArtifact, однако, метод
// выполняет еще очень важное действие, перемещает содержимое из a.artifactTmp в
// a.artifacts, так как artifacts автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются artifactTmp
func (a *SupportiveAlertArtifacts) GetArtifacts() map[string][]datamodels.ArtifactForEsAlert {
	a.listAcceptedFields = []string(nil)

	if a.currentKey != "" {
		a.artifacts[a.currentKey] = append(a.artifacts[a.currentKey], a.artifactTmp)
	}

	a.currentKey = ""
	a.artifactTmp = *datamodels.NewArtifactForEsAlert()

	return a.artifacts
}

// GetArtifactTmp возвращает временный объект artifact
func (a *SupportiveAlertArtifacts) GetArtifactTmp() *datamodels.ArtifactForEsAlert {
	return &a.artifactTmp
}

func (a *SupportiveAlertArtifacts) HandlerValue(fieldBranch string, i interface{}, f func(interface{})) {
	if fieldBranch == "alert.artifacts.dataType" {
		str := fmt.Sprint(i)
		if _, ok := a.artifacts[str]; !ok {
			a.artifacts[str] = []datamodels.ArtifactForEsAlert(nil)
		}

		if a.isExistFieldBranch(fieldBranch) {
			a.listAcceptedFields = []string(nil)

			if _, ok := a.artifacts[a.currentKey]; !ok {
				a.artifacts[a.currentKey] = []datamodels.ArtifactForEsAlert(nil)
			}

			a.artifacts[a.currentKey] = append(a.artifacts[a.currentKey], a.artifactTmp)
			a.artifactTmp = *datamodels.NewArtifactForEsAlert()
		}

		a.currentKey = str
	}

	//если поле повторяется то считается что это уже новый объект
	if fieldBranch != "alert.artifacts.tags" && a.isExistFieldBranch(fieldBranch) {
		a.listAcceptedFields = []string(nil)

		if _, ok := a.artifacts[a.currentKey]; !ok {
			a.artifacts[a.currentKey] = []datamodels.ArtifactForEsAlert(nil)
		}

		a.artifacts[a.currentKey] = append(a.artifacts[a.currentKey], a.artifactTmp)
		a.artifactTmp = *datamodels.NewArtifactForEsAlert()
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
