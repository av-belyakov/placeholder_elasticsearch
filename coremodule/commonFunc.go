package coremodule

import (
	"fmt"
	"strings"

	"placeholder_elasticsearch/datamodels"
)

func joinRawFieldsToString(list map[string]string, tag, id string) string {
	var str strings.Builder = strings.Builder{}

	for k, v := range list {
		str.WriteString(fmt.Sprintf("\n\t%s %s field: '%s', value: '%s'", tag, id, k, v))
	}

	return str.String()
}

// searchEventSource выполняет поиск источника события
func searchEventSource(fieldBranch string, value interface{}) (string, bool) {
	if fieldBranch != "source" {
		return "", false
	}

	if v, ok := value.(string); ok {
		return v, true
	}

	return "", false
}

type UserTypeGetter interface {
	GetData() string
	GetDataType() string
	SetValueData(string)
	SetValueSensorId(string)
	SetValueSnortSid(string)
}

// PostProcessingUserType выполняет постобработку некоторых пользовательских типов
/*func PostProcessingUserType[T UserTypeGetter](ut T) (T, bool) {
	handlers := map[string]func(utg UserTypeGetter){
		"snort_sid": func(utg UserTypeGetter) {
			if !strings.Contains((utg).GetData(), ",") {
				return
			}

			tmp := strings.Split((utg).GetData(), ",")
			for _, v := range tmp {
				(utg).SetValueSnortSid(strings.TrimSpace(v))
			}
		},
		"ip_home": func(utg UserTypeGetter) {
			if !strings.Contains((utg).GetData(), ":") {
				return
			}

			tmp := strings.Split((utg).GetData(), ":")
			if len(tmp) != 2 {
				return
			}

			(utg).SetValueSensorId(tmp[0])
			(utg).SetValueData(tmp[1])
		},
	}

	f, ok := handlers[ut.GetDataType()]
	if !ok {
		return ut, false
	}

	f(ut)

	return ut, true
}*/

// PostProcessingListUserType выполняет постобработку некоторых пользовательских типов
func PostProcessingListUserType[T UserTypeGetter](list map[string][]T) map[string][]T {
	handlers := map[string]func(utg UserTypeGetter){
		"snort_sid": func(utg UserTypeGetter) {
			if !strings.Contains((utg).GetData(), ",") {
				return
			}

			tmp := strings.Split((utg).GetData(), ",")
			for _, v := range tmp {
				(utg).SetValueSnortSid(strings.TrimSpace(v))
			}
		},
		"ip_home": func(utg UserTypeGetter) {
			if !strings.Contains((utg).GetData(), ":") {
				return
			}

			tmp := strings.Split((utg).GetData(), ":")
			if len(tmp) != 2 {
				return
			}

			(utg).SetValueSensorId(tmp[0])
			(utg).SetValueData(tmp[1])
		},
	}

	for k, v := range list {
		for key, value := range v {
			f, ok := handlers[value.GetDataType()]
			if !ok {
				continue
			}

			f(list[k][key])
		}
	}

	return list
}

// PostProcessingListArtifacts выполняет постобработку объектов Artifacts
func PostProcessingListArtifacts(list map[string][]datamodels.ArtifactForEsAlert) map[string][]datamodels.ArtifactForEsAlert {
	handlers := map[string]func(a *datamodels.ArtifactForEsAlert){
		"snort_sid": func(a *datamodels.ArtifactForEsAlert) {
			if strings.Contains(a.Data, ",") {
				tmp := strings.Split(a.Data, ",")
				sids := make([]string, 0, len(tmp))
				for _, v := range tmp {
					sids = append(sids, strings.TrimSpace(v))
				}

				a.SnortSid = sids

				return
			}

			a.SnortSid = append(a.SnortSid, a.Data)
		},
		"ip_home": func(a *datamodels.ArtifactForEsAlert) {
			tmp := strings.Split(a.Data, ":")

			if len(tmp) != 2 {
				return
			}

			a.SensorId = tmp[0]
			a.Data = tmp[1]
		},
	}

	for k, v := range list {
		for key, value := range v {
			f, ok := handlers[value.DataType]
			if !ok {
				continue
			}

			f(&list[k][key])
		}
	}

	return list
}

// PostProcessingListObservables выполняет постобработку объектов Observables
func PostProcessingListObservables(list map[string][]datamodels.ObservableMessageEs) map[string][]datamodels.ObservableMessageEs {
	handlers := map[string]func(a *datamodels.ObservableMessageEs){
		"snort_sid": func(a *datamodels.ObservableMessageEs) {
			if strings.Contains(a.Data, ",") {
				sid := strings.Split(a.Data, ",")
				a.SnortSid = sid

				return
			}

			a.SnortSid = append(a.SnortSid, a.Data)
		},
		"ip_home": func(a *datamodels.ObservableMessageEs) {
			tmp := strings.Split(a.Data, ":")

			if len(tmp) != 2 {
				return
			}

			a.SensorId = tmp[0]
			a.Data = tmp[1]
		},
	}

	for k, v := range list {
		for key, value := range v {
			f, ok := handlers[value.DataType]
			if !ok {
				continue
			}

			f(&list[k][key])
		}
	}

	return list
}

func checkDatetimeFieldsEventObject(e *datamodels.EventMessageTheHiveCase) {
	if e.GetStartDate() == "" {
		e.SetValueStartDate("1970-01-01T00:00:00+00:00")
	}

	if e.Details.GetEndDate() == "" {
		e.Details.SetValueEndDate("1970-01-01T00:00:00+00:00")
	}

	if e.Object.GetStartDate() == "" {
		e.Object.SetValueStartDate("1970-01-01T00:00:00+00:00")
	}

	if e.Object.GetEndDate() == "" {
		e.Object.SetValueEndDate("1970-01-01T00:00:00+00:00")
	}

	if e.Object.GetCreatedAt() == "" {
		e.Object.SetValueCreatedAt("1970-01-01T00:00:00+00:00")
	}

	if e.Object.GetUpdatedAt() == "" {
		e.Object.SetValueUpdatedAt("1970-01-01T00:00:00+00:00")
	}
}
