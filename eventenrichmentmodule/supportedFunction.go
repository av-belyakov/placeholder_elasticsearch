package eventenrichmentmodule

import (
	"net/http"
	"placeholder_elasticsearch/confighandler"
)

func responseClose(res *http.Response) {
	if res == nil || res.Body == nil {
		return
	}

	res.Body.Close()
}

func MappingObjectArea(objectArea string, listObjectArea []confighandler.ObjectAreaActivity) string {
	result := objectArea

	for _, v := range listObjectArea {
		if len(v.VariationsName) == 0 {
			continue
		}

		for _, name := range v.VariationsName {
			if objectArea == name {
				result = v.ApprovedName

				return result
			}
		}
	}

	return result
}
