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
		/*customFieldJson := `{
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
		}`*/
		/*testJson := `{
				"severity":      1,
				"tlp":             1,
				"pap":             1,
				"_id":   "~85455464790",
				"id":              "~85455464790",
				"createdBy":       "ddddd",
				"createdAt":       "1970-01-01T03:00:00+03:00",
				"updatedAt":       "1970-01-01T03:00:00+03:00",
				"_Type": "aalllert",
				"title":           "vbb er3",
				"description":     "any more",
				"status":          "None",
				"date":            "2024-02-06T15:37:52+03:00",
				"type":            "snort_alert",
				"objectType":      "",
				"source":          "zsiеmSystems",
				"sourceRef":       "TSK-8MSK-6-ZPM-240206-1215999",
				"case":            "alert",
				"caseTemplate":    "alert_snort",
				"tags": [
					"'Sensor:id=\"8030012\"'",
					"'Webhook:send=ES'"
				],
				"customFields": {
					"first-time": {
						"order": 0,
						"date":  "2024-02-06T15:20:30+03:00"
					},
					"last-time": {
						"order": 0,
						"date":  "2024-02-06T15:20:30+03:00"
					}
				}
		}`*/
		testJson := `{
			"@id":         "6b3be7fe-94a5-2133-a923-70b2a445f0ab",
			"@timestamp": "2024-01-31T16:17:22+03:00",
			"source":      "GCM",
			"event": {
				"base":           false,
				"startDate":      "2024-02-06T15:20:30+03:00",
				"rootId":         "~84625227848",
				"objectId":       "~4192",
				"objectType":     "alert",
				"organisation":   "GCM",
				"organisationId": "~4192",
				"operation":      "new",
				"requestId":      "55994d44f3b276c1:6483e5ec:18d786c2f83:-8000:138497",
				"details": {
					"sourceRef":   "TSK-8MSK-6-ZPM-240206-1215999",
					"title":       "222",
					"description": "111",
					"tags": [
						"ATs:reason=\"INFO Controlled FGS\"",
						"Sensor:id=\"8030066\""
					]
				},
				"object": {
					"severity":      1,
					"tlp":             1,
					"pap":             1,
					"_id":             "~85455464790",
					"id":              "~85455464790",
					"createdBy":       "ddddd",
					"createdAt":       "1970-01-01T03:00:00+03:00",
					"updatedAt":       "1970-01-01T03:00:00+03:00",
					"_Type":           "aalllert",
					"title":           "vbb er3",
					"description":     "any more",
					"status":          "None",
					"date":            "2024-02-06T15:37:52+03:00",
					"type":            "snort_alert",
					"objectType":      "",
					"source":          "zsiеmSystems",
					"sourceRef":       "TSK-8MSK-6-ZPM-240206-1215999",
					"case":            "alert",
					"caseTemplate":    "alert_snort",
					"tags": [
						"'Sensor:id=\"8030012\"'",
						"'Webhook:send=ES'"
					],
					"customFields": {
						"first-time": {
							"order": 0,
							"date":  "2024-02-06T15:20:30+03:00"
						},
						"last-time": {
							"order": 0,
							"date":  "2024-02-06T15:20:30+03:00"
						}
					}
				}
			}
		}`

		It("Do test", func() {
			/*type CustomFieldType struct {
				CustomFields datamodels.CustomFields `json:"customFields"`
			}*/

			cf := datamodels.NewVerifiedTheHiveAlert()
			//cf := datamodels.EventAlertObject{}
			//cf := CustomFieldType{}

			err := json.Unmarshal([]byte(testJson), cf)
			Expect(err).ShouldNot(HaveOccurred())

			b, err := json.MarshalIndent(cf, "", "  ")
			fmt.Println("Finaly object:", string(b))

			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
