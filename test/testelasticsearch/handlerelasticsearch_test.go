package testelasticsearch_test

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	//"placeholder_elasticsearch/test/testelasticsearch"
)

var _ = Describe("Handlerelasticsearch", Ordered, func() {
	var (
		errES    error
		esClient *elasticsearch.Client
	)

	BeforeAll(func() {
		//инициализируем клиента Elasticsearch
		esClient, errES = elasticsearch.NewClient(elasticsearch.Config{
			Addresses: []string{fmt.Sprintf("http://%s:%d", "datahook.cloud.gcm", 9200)},
			Username:  "writer",
			Password:  "XxZqesYXuk8C",
		})
	})

	Context("Тест 1. Проверка подключения к СУБД Elasticsearch", func() {
		It("При подключени к СУБД ошибок быть не должно", func() {
			Expect(errES).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Выполнение запросов к СУБД", func() {
		It("Запрос на получения документа по id не должен вызывать ошибок, должен быть получен результат", func() {
			//"test_module_placeholder_elasticsearch_case_2024_1"
			res, err := esClient.Get("test_module_placeholder_elasticsearch_case_2024_1", "hEUzWo0BNzcaI2GK2cLu")
			Expect(err).ShouldNot(HaveOccurred())

			r := map[string]interface{}{}
			errJsonDecode := json.NewDecoder(res.Body).Decode(&r)
			Expect(errJsonDecode).ShouldNot(HaveOccurred())

			var isExist bool
			fmt.Println("Read response message:")
			for k, v := range r {
				if k == "found" {
					if tmp, ok := v.(bool); ok {
						isExist = tmp
					}
				}

				fmt.Println(k, ":", v)
			}

			Expect(isExist).Should(BeTrue())
		})

		/*It("Запрос на поиск документа не должен вызывать ошибок", func ()  {
			res, err := esClient.Search()
			Expect(err).ShouldNot(HaveOccurred())
		})*/
	})

	/*
		Context("", func ()  {
			It("", func ()  {

			})
		})
	*/
})
