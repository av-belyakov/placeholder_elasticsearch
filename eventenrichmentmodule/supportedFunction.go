package eventenrichmentmodule

import "net/http"

func responseClose(res *http.Response) {
	if res == nil || res.Body == nil {
		return
	}

	res.Body.Close()
}
