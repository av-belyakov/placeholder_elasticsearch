package datamodels

import (
	"fmt"
	"reflect"
	"strings"

	"placeholder_elasticsearch/supportingfunctions"
)

type UserTypeGetter interface {
	GetData() string
	GetDataType() string
	SetValueData(string)
	SetValueSensorId(string)
	SetValueSnortSid(string)
}

// PostProcessingUserType выполняет постобработку некоторых пользовательских типов
func PostProcessingUserType[T UserTypeGetter](ut T) (T, bool) {
	handlers := map[string]func(utg UserTypeGetter){
		"snort_sid": func(utg UserTypeGetter) {
			if !strings.Contains(utg.GetData(), ",") {
				if utg.GetData() != "" {
					utg.SetValueSnortSid(utg.GetData())
				}

				return
			}

			tmp := strings.Split(utg.GetData(), ",")
			for _, v := range tmp {
				utg.SetValueSnortSid(strings.TrimSpace(v))
			}
		},
		"ip_home": func(utg UserTypeGetter) {
			if !strings.Contains(utg.GetData(), ":") {
				return
			}

			tmp := strings.Split(utg.GetData(), ":")
			if len(tmp) != 2 {
				return
			}

			utg.SetValueSensorId(tmp[0])
			utg.SetValueData(tmp[1])
		},
	}

	f, ok := handlers[ut.GetDataType()]
	if !ok {
		return ut, false
	}

	f(ut)

	return ut, true
}

func replacingSliceString(current, new reflect.Value) (list reflect.Value, ok bool) {
	if reflect.DeepEqual(current, new) {
		return list, false
	}

	currentTags, okCurr := current.Interface().([]string)
	newTags, okNew := new.Interface().([]string)
	if !okCurr || !okNew {
		return list, false
	}

	if !current.CanSet() {
		return list, false
	}

	list = reflect.ValueOf(supportingfunctions.SliceJoinUniq[string](currentTags, newTags))

	return list, true
}

func ToStringBeautifulSlice(num int, l []string) string {
	str := strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num + 1)

	for k, v := range l {
		str.WriteString(fmt.Sprintf("%s%d. '%s'\n", ws, k+1, v))
	}

	return str.String()
}

func ToStringBeautifulMapSlice(num int, m map[string][]string) string {
	str := strings.Builder{}

	for k, v := range m {
		str.WriteString(fmt.Sprintf("%s'%s'\n", supportingfunctions.GetWhitespace(num+1), k))
		for key, value := range v {
			str.WriteString(fmt.Sprintf("%s%d. %s\n", supportingfunctions.GetWhitespace(num+2), key+1, value))
		}
	}

	return str.String()
}

func CustomFieldsToStringBeautiful(l CustomFields, num int) string {
	strB := strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num + 2)

	for k, v := range l {
		strB.WriteString(fmt.Sprintf("%s'%s':\n", supportingfunctions.GetWhitespace(num+1), k))

		nameOne, dataOne, nameTwo, dataTwo := v.Get()
		strB.WriteString(fmt.Sprintf("%s'%s': %d\n", ws, nameOne, dataOne))
		strB.WriteString(fmt.Sprintf("%s'%s': %s\n", ws, nameTwo, dataTwo))
	}

	return strB.String()
}
