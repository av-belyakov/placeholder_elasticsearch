package testreplacingobservables_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/datamodels/commonobservable"
)

var _ = Describe("Testreplacingobservables", func() {
	Context("Тест 1. Проверка замены старых значений ObservablesMessageEs объекта, новыми значениями, если они отличаются", func() {
		oldStruct := datamodels.ObservablesMessageEs{
			Observables: map[string][]datamodels.ObservableMessageEs{
				"domain": {
					{
						CommonObservableType: commonobservable.CommonObservableType{
							Ioc:                  true,
							Sighted:              false,
							IgnoreSimilarity:     true,
							Tlp:                  2,
							UnderliningCreatedAt: "2024-03-05T20:58:09+03:00",
							UnderliningUpdatedAt: "1970-01-01T00:00:00+00:00",
							StartDate:            "2024-02-21T21:47:16+03:00",
							UnderliningCreatedBy: "a.pinyaskin@cloud.gcm",
							UnderliningUpdatedBy: "",
							UnderliningId:        "~86227013760",
							UnderliningType:      "Observable",
							Data:                 "wxanalytics.ru",
							DataType:             "domain",
							Message:              "message doamin 111",
						},
						SensorId: "",
						SnortSid: []string(nil),
						Tags: map[string][]string{
							"ats:reason": {
								"\"Evidence Malware 100\"",
								"\"Зафиксирована подозрительная активность по признакам НКЦКИ\"",
								"\"INFO Controlled FGS\"",
								"\"Evidence Malware 50\"",
								"\"bmpol List 2\"",
								"\"bmpol WhiteList\"",
							},
							"ats:geoip": {"\"США\""},
						},
						TagsAll: []string{
							"'Webhook:send=ES'",
							"ATs:reason=\"Evidence Malware 100\"",
							"ATs:reason=\"Зафиксирована подозрительная активность по признакам НКЦКИ\"",
							"Job",
							"ATs:reason=\"INFO Controlled FGS\"",
							"ATs:reason=\"Evidence Malware 50\"",
							"ATs:reason=\"bmpol List 2\"",
							"ATs:reason=\"bmpol WhiteList\"",
							"Sensor:id=\"8030059\"",
							"ATs:geoip=\"США\"",
							"APPA:Direction=\"outbound\"",
						},
						Attachment: datamodels.AttachmentData{},
						//Reports:    make(map[string]datamodels.ReportTaxonomies),
					},
					{
						CommonObservableType: commonobservable.CommonObservableType{
							Ioc:                  true,
							Sighted:              false,
							IgnoreSimilarity:     true,
							Tlp:                  2,
							UnderliningCreatedAt: "2024-03-05T17:22:28+03:00",
							UnderliningUpdatedAt: "1970-01-01T00:00:00+00:00",
							StartDate:            "2024-03-05T17:22:28+03:00",
							UnderliningCreatedBy: "denis.a@mimo.sec",
							UnderliningUpdatedBy: "",
							UnderliningId:        "~616284224",
							UnderliningType:      "Observable",
							Data:                 "contabo.com",
							DataType:             "domain",
							Message:              "message frfrgr efgon",
						},
						SensorId: "",
						SnortSid: []string(nil),
						Tags: map[string][]string{
							"misp:External analysis": {"\"domain\""},
						},
						TagsAll: []string{
							"misp:External analysis=\"domain\"",
						},
						Attachment: datamodels.AttachmentData{},
						/*Reports: map[string]datamodels.ReportTaxonomies{
							"AbuseIPDB_1_0": {
								Taxonomies: []datamodels.Taxonomy{
									{
										Level:     "malicious",
										Namespace: "AbuseIPDB",
										Predicate: "Records",
										Value:     "1",
									},
								},
							},
							"MISP_2_1": {
								Taxonomies: []datamodels.Taxonomy{
									{
										Level:     "suspicious",
										Namespace: "MISP",
										Predicate: "Search",
										Value:     "1 event(s)",
									},
								},
							},
						},*/
					},
				},
				"ip_home": {
					{
						CommonObservableType: commonobservable.CommonObservableType{
							Ioc:                  false,
							Sighted:              false,
							IgnoreSimilarity:     false,
							Tlp:                  2,
							UnderliningCreatedAt: "2024-03-05T20:58:09+03:00",
							UnderliningUpdatedAt: "1970-01-01T00:00:00+00:00",
							StartDate:            "2024-03-05T20:58:09+03:00",
							UnderliningCreatedBy: "hiveapi@rcm.ekb",
							UnderliningUpdatedBy: "",
							UnderliningId:        "~902983904",
							UnderliningType:      "Observable",
							Data:                 "680036:212.23.84.226",
							DataType:             "ip_home",
							Message:              "message my description",
						},
						SensorId: "680036",
						SnortSid: []string(nil),
						Tags: map[string][]string{
							"misp:Network activity": {"\"domain\""},
						},
						TagsAll: []string{
							"Malware ",
							"misp:Network activity=\"domain\"",
						},
						Attachment: datamodels.AttachmentData{
							Size:        343648,
							Id:          "44b04611c2a5d728797c6f79805f37a89a6dd048b2ba751682040d63778c6b8e",
							Name:        "UFO-20240305-5_sample.txt",
							ContentType: "application/octet-stream",
							Hashes: []string{
								"44b04611c2a5d728797c6f79805f37a89a6dd048b2ba751682040d63778c6b8e",
								"4d05d488d9b887a9c4eb0f4abc6c6529ab04c5ea",
								"03f2aa58955a9f542642e7d7672cf360",
							},
						},
						//Reports: make(map[string]datamodels.ReportTaxonomies),
					},
				},
			},
		}

		newStruct := datamodels.ObservablesMessageEs{
			Observables: map[string][]datamodels.ObservableMessageEs{
				"domain": {
					{
						CommonObservableType: commonobservable.CommonObservableType{
							Ioc:                  true,
							Sighted:              true, //замена
							IgnoreSimilarity:     true,
							Tlp:                  2,
							UnderliningCreatedAt: "2024-03-05T20:58:09+03:00",
							UnderliningUpdatedAt: "2024-03-06T10:26:10+03:00", //замена
							StartDate:            "2024-02-21T21:47:16+03:00",
							UnderliningCreatedBy: "a.pinyaskin@cloud.gcm",
							UnderliningUpdatedBy: "",
							UnderliningId:        "~86227013760",
							UnderliningType:      "Observable",
							Data:                 "wxanalytics.ru",
							DataType:             "domain",
							Message:              "тут заменяем сообщение", //замена
						},
						SensorId: "",
						SnortSid: []string(nil),
						Tags: map[string][]string{
							"ats:reason": {
								"\"Evidence Malware 100\"",
								"\"Зафиксирована подозрительная активность по признакам НКЦКИ\"",
								"\"INFO Controlled FGS\"",
								"\"Evidence Malware 50\"",
								"\"bmpol List 2\"",
								"\"bmpol WhiteList\"",
							},
							"ats:geoip":             {"\"США\""},
							"misp:Network activity": {"\"snort\""}, //добавление
						},
						TagsAll: []string{
							"'Webhook:send=ES'",
							"ATs:reason=\"Evidence Malware 100\"",
							"ATs:reason=\"Зафиксирована подозрительная активность по признакам НКЦКИ\"",
							"Job",
							"ATs:reason=\"INFO Controlled FGS\"",
							"ATs:reason=\"Evidence Malware 50\"",
							"ATs:reason=\"bmpol List 2\"",
							"ATs:reason=\"bmpol WhiteList\"",
							"Sensor:id=\"8030059\"",
							"ATs:geoip=\"США\"",
							"APPA:Direction=\"outbound\"",
							"misp:Network activity=\"snort\"", //добавление
						},
						Attachment: datamodels.AttachmentData{},
						/*Reports: map[string]datamodels.ReportTaxonomies{
							"AbuseIPDB_1_0": {
								Taxonomies: []datamodels.Taxonomy{
									{
										Level:     "malicious",
										Namespace: "_______AbuseIPDB_____", //замена
										Predicate: "Records",
										Value:     "1",
									},
								},
							},
							//добавление
							"ZPMSig_1_0": {
								Taxonomies: []datamodels.Taxonomy{
									{
										Level:     "malicious",
										Namespace: "ZPM_Sensor",
										Predicate: "IP",
										Value:     "Found",
									},
								},
							},
						},*/
					},
					//добавление
					{
						CommonObservableType: commonobservable.CommonObservableType{
							Ioc:                  true,
							Sighted:              false,
							IgnoreSimilarity:     true,
							Tlp:                  2,
							UnderliningCreatedAt: "2024-03-05T17:22:28+03:00",
							UnderliningUpdatedAt: "1970-01-01T00:00:00+00:00",
							StartDate:            "2024-03-05T17:22:28+03:00",
							UnderliningCreatedBy: "denis.a@mimo.sec",
							UnderliningUpdatedBy: "",
							UnderliningId:        "~616284224",
							UnderliningType:      "Observable",
							Data:                 "contabo.com",
							DataType:             "domain",
							Message:              "message frfrgr efgon",
						},
						SensorId: "",
						SnortSid: []string(nil),
						Tags: map[string][]string{
							"misp:External analysis": {"\"domain\""},
						},
						TagsAll: []string{
							"misp:External analysis=\"domain\"",
						},
						Attachment: datamodels.AttachmentData{},
						/*Reports: map[string]datamodels.ReportTaxonomies{
							"AbuseIPDB_1_0": {
								Taxonomies: []datamodels.Taxonomy{
									{
										Level:     "malicious",
										Namespace: "AbuseIPDB",
										Predicate: "Records",
										Value:     "1",
									},
								},
							},
							"MISP_2_1": {
								Taxonomies: []datamodels.Taxonomy{
									{
										Level:     "suspicious",
										Namespace: "MISP",
										Predicate: "Search",
										Value:     "1 event(s)",
									},
								},
							},
						},*/
					},
				},
				//добавление
				"url_pcap": {
					{
						CommonObservableType: commonobservable.CommonObservableType{
							Ioc:                  false,
							Sighted:              false,
							IgnoreSimilarity:     true,
							Tlp:                  2,
							UnderliningCreatedAt: "2024-03-06T12:33:42+03:00",
							UnderliningUpdatedAt: "1970-01-01T00:00:00+00:00",
							StartDate:            "2024-03-06T12:33:42+03:00",
							UnderliningCreatedBy: "s.chinkov@cloud.gcm",
							UnderliningUpdatedBy: "denis.a@mimo.sec",
							UnderliningId:        "~86794760200",
							UnderliningType:      "Observable",
							Data:                 "ftp://ftp.cloud.gcm/traffic/8030139/1709680601_2024_03_06____02_16_41_898767.pcap",
							DataType:             "url_pcap",
							Message:              "Download a piece of traffic",
						},
						SensorId:   "",
						SnortSid:   []string(nil),
						Tags:       map[string][]string{},
						TagsAll:    []string{"45.8 KB"},
						Attachment: datamodels.AttachmentData{},
						/*Reports: map[string]datamodels.ReportTaxonomies{
							"Moloch_1_7": {
								Taxonomies: []datamodels.Taxonomy{
									{
										Level:     "info",
										Namespace: "Moloch",
										Predicate: "Uploading",
										Value:     "1 pcap(s)",
									},
								},
							},
						},*/
					},
				},
				"ip_home": {
					{
						CommonObservableType: commonobservable.CommonObservableType{
							Ioc:                  false,
							Sighted:              false,
							IgnoreSimilarity:     false,
							Tlp:                  2,
							UnderliningCreatedAt: "2024-03-05T20:58:09+03:00",
							UnderliningUpdatedAt: "1970-01-01T00:00:00+00:00",
							StartDate:            "2024-03-05T20:58:09+03:00",
							UnderliningCreatedBy: "hiveapi@rcm.ekb",
							UnderliningUpdatedBy: "",
							UnderliningId:        "~902983904",
							UnderliningType:      "Observable",
							Data:                 "680036:212.23.84.226",
							DataType:             "ip_home",
							Message:              "message my description",
						},
						SensorId: "680036",
						SnortSid: []string(nil),
						Tags: map[string][]string{
							"misp:Network activity": {"\"domain\""},
						},
						TagsAll: []string{
							"Malware ",
							"misp:Network activity=\"domain\"",
						},
						Attachment: datamodels.AttachmentData{
							Size:        343000648, //замена
							Id:          "44b04611c2a5d728797c6f79805f37a89a6dd048b2ba751682040d63778c6b8e",
							Name:        "замена UFO-20240305-5_sample.txt", //замена
							ContentType: "замена application/octet-stream",  //замена
							Hashes: []string{
								"44b04611c2a5d728797c6f79805f37a89a6dd048b2ba751682040d63778c6b8e",
								"4d05d488d9b887a9c4eb0f4abc6c6529ab04c5ea",
								"03f2aa58955a9f542642e7d7672cf360",
								//добавление
								"84584584a224af43454b4554776",
							},
						},
						//Reports: make(map[string]datamodels.ReportTaxonomies),
					},
				},
				//добавление
				"snort_sid": {
					{
						CommonObservableType: commonobservable.CommonObservableType{
							Ioc:                  true,
							Sighted:              false,
							IgnoreSimilarity:     false,
							Tlp:                  2,
							UnderliningCreatedAt: "2024-03-04T11:57:33+03:00",
							UnderliningUpdatedAt: "1970-01-01T00:00:00+00:00",
							StartDate:            "2024-03-04T11:57:33+03:00",
							UnderliningCreatedBy: "a.pinyaskin@cloud.gcm",
							UnderliningUpdatedBy: "",
							UnderliningId:        "~88166945016",
							UnderliningType:      "Observable",
							Data:                 "3005204, 2016683, 3023574",
							DataType:             "snort_sid",
							Message:              "many snort signature",
						},
						SensorId: "",
						SnortSid: []string{
							"3005204",
							"2016683",
							"3023574",
						},
						Tags:       map[string][]string{},
						TagsAll:    []string{"snort"},
						Attachment: datamodels.AttachmentData{},
						/*Reports: map[string]datamodels.ReportTaxonomies{
							"ARIADNA_GeoIP_1_0": {
								Taxonomies: []datamodels.Taxonomy{
									{
										Level:     "info",
										Namespace: "Ariadna_GeoIP",
										Predicate: "Country",
										Value:     "Великобритания",
									},
								},
							},
						},*/
					},
				},
			},
		}

		It("Ряд полей в ObservablesMessageEs должны быть успешно заменены", func() {
			num := oldStruct.ReplacingOldValues(newStruct)

			//кол-во замененных полей
			Expect(num).Should(Equal(16))

			fmt.Println("---=== VERIFED ObservablesMessageEs ===---")
			fmt.Println(oldStruct.ToStringBeautiful(0))

			//добавляется
			_, ok := oldStruct.Observables["url_pcap"]
			Expect(ok).Should(BeTrue())

			_, ok = oldStruct.Observables["snort_sid"]
			Expect(ok).Should(BeTrue())

			domains, ok := oldStruct.Observables["domain"]
			Expect(ok).Should(BeTrue())
			Expect(len(domains)).Should(Equal(2))

			ipHome, ok := oldStruct.Observables["ip_home"]
			Expect(ok).Should(BeTrue())
			Expect(ipHome[0].GetAttachment().GetSize()).Should(Equal(uint64(343000648)))
			Expect(len(ipHome[0].GetAttachment().GetHashes())).Should(Equal(4))

			Expect(true).Should(BeTrue())
		})
	})
})
