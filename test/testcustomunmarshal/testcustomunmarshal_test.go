package testcustomunmarshal_test

import (
	"encoding/json"
	"fmt"
	"placeholder_elasticsearch/datamodels"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testcustomunmarshal", func() {
	Context("Test 1.", func() {
		customFieldJson := `{
			"customFields": {		
				"first-time": {
				  "order": 0,
				  "date":  "2024-01-01T05:22:30+03:00"
				},
				"last-time": {
				   "order": 0,
				  "date":  "2024-01-17T00:18:13+03:00"
				}
			  }
		}`

		It("Do test", func() {
			var cf datamodels.CustomFields

			err := json.Unmarshal([]byte(customFieldJson), &cf)
			Expect(err).ShouldNot(HaveOccurred())

			b, err := json.MarshalIndent(cf, "", "  ")
			fmt.Println("customFields:", string(b))

			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
