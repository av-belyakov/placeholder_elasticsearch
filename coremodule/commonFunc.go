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
