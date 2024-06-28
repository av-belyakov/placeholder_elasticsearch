package testelasticsearchindexsettings_test

import (
	"fmt"

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

		hsd       elasticsearchinteractions.HandlerSendData
		indexName string = "module_placeholder_new_case_2024_4"
		errConn   error
	)

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
	})

	Context("Тест 1. Проверяем возможность изменения настроек уже существующего индекса", func() {
		It("При подключении не должно быть ошибок", func() {
			Expect(errConn).ShouldNot(HaveOccurred())
		})

		It("Должен существовать индекс с определенным именем", func() {
			indexes, err := hsd.GetExistingIndexes(indexName)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(indexes)).Should(Equal(1))
		})

		It("При изменении настроек выбранного индекса не должно быть ошибок", func() {
			query := `{
				"index": {
					"mapping": {
						"total_fields": {
							"limit": 2000
						}
					}
				}
			}`
			ok, err := hsd.SetIndexSetting([]string{indexName}, query)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ok).Should(BeTrue())
		})

		It("Должен быть найден индекс по заданному шаблону", func() {
			indexes, err := hsd.GetExistingIndexes("module_placeholder_new_case_2024")

			fmt.Println("index:", indexes)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(indexes)).ShouldNot(Equal(0))
		})
		/*It("При создании тестового индекса не должно быть ошибок", func() {
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

			res, err := hsd.InsertDocument("tag_name", indexName, document)
			Expect(err).ShouldNot(HaveOccurred())

			fmt.Println("Response Body:", res.Body)
			err = res.Body.Close()
			Expect(err).ShouldNot(HaveOccurred())
		})*/
	})
})
