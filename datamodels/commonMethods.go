package datamodels

import (
	"fmt"
	"reflect"
	"strings"

	"placeholder_elasticsearch/supportingfunctions"
)

func replacingSlice(current, new reflect.Value) (list reflect.Value, ok bool) {
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
		str.WriteString(fmt.Sprintf("%s%s\n", supportingfunctions.GetWhitespace(num+1), k))
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
		strB.WriteString(fmt.Sprintf("%s%s:\n", supportingfunctions.GetWhitespace(num+1), k))

		nameOne, dataOne, nameTwo, dataTwo := v.Get()
		strB.WriteString(fmt.Sprintf("%s%s: %d\n", ws, nameOne, dataOne))
		strB.WriteString(fmt.Sprintf("%s%s: %s\n", ws, nameTwo, dataTwo))
	}

	return strB.String()
}
