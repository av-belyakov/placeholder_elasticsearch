package supportingfunctions

import (
	"fmt"
	"strings"
)

// GetWhitespace возвращает необходимое количество пробелов
func GetWhitespace(num int) string {
	var str string

	if num == 0 {
		return str
	}

	for i := 0; i < num; i++ {
		str += "  "
	}

	return str
}

// HandlerTag выполняет обработку тегов, разделяя тег на его тип и значение
func HandlerTag(i interface{}) (key, value string) {
	isExistValidTag := func(item string) bool {
		validListTags := []string{
			"geo",
			"geoip",
			"reason",
			"sensor",
			"misp",
			"ioc",
		}

		for _, v := range validListTags {
			if strings.Contains(item, v) {
				return true
			}
		}

		return false
	}

	tag := strings.ToLower(fmt.Sprint(i))

	if isExistValidTag(tag) && strings.Contains(tag, "=") {
		elements := strings.Split(tag, "=")
		if len(elements) > 1 {
			if strings.Contains(elements[0], "geo") {
				return elements[0], strings.ToUpper(elements[1])
			}

			return elements[0], elements[1]
		}
	}

	return tag, ""
}

func ToStringBeautifulSlice(num int, l []string) string {
	str := strings.Builder{}
	ws := GetWhitespace(num + 1)

	for k, v := range l {
		str.WriteString(fmt.Sprintf("%s%d. '%s'\n", ws, k+1, v))
	}

	return str.String()
}

func ToStringBeautifulMapSlice(num int, m map[string][]string) string {
	str := strings.Builder{}

	for k, v := range m {
		str.WriteString(fmt.Sprintf("%s'%s'\n", GetWhitespace(num+1), k))
		for key, value := range v {
			str.WriteString(fmt.Sprintf("%s%d. %s\n", GetWhitespace(num+2), key+1, value))
		}
	}

	return str.String()
}
