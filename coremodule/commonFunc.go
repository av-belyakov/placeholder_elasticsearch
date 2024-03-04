package coremodule

import (
	"fmt"
	"strings"
)

func joinRawFieldsToString(list map[string]string, tag, id string) string {
	var str strings.Builder = strings.Builder{}

	for k, v := range list {
		str.WriteString(fmt.Sprintf("/t%s %s field: '%s', value: '%s'\n", tag, id, k, v))
	}

	return str.String()
}
