package testelasticsearchindexsettings_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/elasticsearchinteractions"
)

var _ = Describe("Testelasticsearchindexsettings", Ordered, func() {
	var (
		Host   string = "datahook.cloud.gcm"
		Port   int    = 9200
		User   string = "writer"
		Passwd string = "XxZqesYXuk8C"

		hsd          elasticsearchinteractions.HandlerSendData
		source       string = "rcmmsk"
		indexName    string = "module_placeholder_new_alert"
		indexPattern string = indexName + "_" + source
		indexCurrent string = indexName + "_" + source + "_" + "2024" + "_" + "4"
		indexTest    string = indexPattern + "_2025_1"

		query string = `{
			"index": {
				"mapping": {
					"total_fields": {
						"limit": 2000
					}
				}
			}
		}`

		errConn error
	)

	//при создании нового индекса вносим в его настройку дополнительный
	//параметр позволяющий увеличить лимит количества создаваемых полей
	//в текущем индексе, что, позволяет убрать ошибку Elasticsearch типа
	//Limit of total fields [1000] has been exceeded while adding new fields

	BeforeAll(func() {
		hsd = elasticsearchinteractions.HandlerSendData{
			Settings: elasticsearchinteractions.SettingsHandler{
				Port:   Port,
				Host:   Host,
				User:   User,
				Passwd: Passwd,
			},
		}

		errConn = hsd.New()

		i, err := hsd.DelIndexSetting([]string{indexTest})
		if err != nil {
			fmt.Println("ERROR:", err)
		}

		fmt.Println("INT:", i)

	})

	Context("Тест 1. Проверяем возможность изменения настроек уже существующего индекса", func() {
		It("При подключении не должно быть ошибок", func() {
			Expect(errConn).ShouldNot(HaveOccurred())
		})

		It("Должен существовать индекс с определенным именем", func() {
			indexes, err := hsd.GetExistingIndexes(indexPattern)
			Expect(err).ShouldNot(HaveOccurred())

			var isExist bool
			for k, v := range indexes {
				fmt.Printf("%d. search index:%s\n", k, v)

				if v == indexCurrent {
					isExist = true

					break
				}
			}

			Expect(len(indexes)).Should(Equal(5))
			Expect(isExist).Should(BeTrue())
		})

		It("При изменении настроек выбранного индекса не должно быть ошибок", func() {
			ok, err := hsd.SetIndexSetting([]string{indexCurrent}, query)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ok).Should(BeTrue())
		})

		It("Должен быть найден индекс по заданному шаблону", func() {
			indexes, err := hsd.GetExistingIndexes(indexPattern)

			for k, v := range indexes {
				fmt.Printf("%d. index:%s\n", k, v)
			}

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(indexes)).ShouldNot(Equal(0))
		})
		It("При создании тестового индекса не должно быть ошибок", func() {
			document, err := json.Marshal(struct {
				Id   string
				Name string
				Age  int
			}{
				Id:   "1334356",
				Name: "Onmsddf",
				Age:  24,
			})
			Expect(err).ShouldNot(HaveOccurred())

			_, err = hsd.SetIndexSetting([]string{indexTest}, query)
			Expect(err).Should(HaveOccurred())
			if err != nil {
				res, err := hsd.InsertDocument("tag_name", indexTest, document)
				Expect(err).ShouldNot(HaveOccurred())

				fmt.Println("Response Body:", res.Body)
				err = res.Body.Close()
				Expect(err).ShouldNot(HaveOccurred())
			}

			ok, err := hsd.SetIndexSetting([]string{indexTest}, query)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ok).Should(BeTrue())
		})
	})
	Context("Тест 2. Получить список настроек индекса", func() {
		It("При получении настроек не должно быть ошибок", func() {
			indexSettings := map[string]struct {
				Settings struct {
					Index struct {
						Mapping struct {
							TotalFields struct {
								Limit string `json:"limit"`
							} `json:"total_fields"`
						} `json:"mapping"`
					} `json:"index"`
				} `json:"settings"`
			}{}

			t := time.Now()

			indexes, err := hsd.GetExistingIndexes(indexPattern)
			Expect(err).ShouldNot(HaveOccurred())

			indexForTotalFieldsLimit := []string(nil)
			if len(indexes) > 0 {
				for _, v := range indexes {
					res, err := hsd.GetIndexSetting(v, "")
					Expect(err).ShouldNot(HaveOccurred())

					if strings.Contains(v, fmt.Sprint(t.Year())) {
						fmt.Println("Index", v, " contains", t.Year(), " year!!!")
					}

					if res.StatusCode == 200 {
						Expect(res.StatusCode).Should(Equal(200))

						err = json.NewDecoder(res.Body).Decode(&indexSettings)
						Expect(err).ShouldNot(HaveOccurred())

						fmt.Println("______________________________")
						fmt.Printf("Index %s: mapping.total_fields.limit: %s\n", v, indexSettings[v].Settings.Index.Mapping.TotalFields.Limit)

						if indexSettings[v].Settings.Index.Mapping.TotalFields.Limit == "2000" {
							continue
						}

						fmt.Printf("Set total fields limit for index %s\n", v)
						indexForTotalFieldsLimit = append(indexForTotalFieldsLimit, v)
					}
				}
			}

			if len(indexForTotalFieldsLimit) > 0 {
				ok, err := hsd.SetIndexSetting(indexForTotalFieldsLimit, query)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(ok).Should(BeTrue())
			}

			Expect(true).Should(BeTrue())

		})
	})
})
