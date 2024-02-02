package testelasticsearch_test

import (
	"context"
	"encoding/json"
	"fmt"
	"placeholder_elasticsearch/datamodels"
	"strings"

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
		It("Запрос на получения документа по source и event.rootId не должен вызывать ошибок, должен быть получен результат", func() {
			//"test_module_placeholder_elasticsearch_case_2024_1"

			source := "rcmsr"
			rootId := "~1682026696"

			query := strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}", source, rootId))
			res, err := esClient.Search(
				esClient.Search.WithContext(context.Background()),
				esClient.Search.WithIndex("module_placeholder_elasticsearch_thehive_case_2024_2"),
				esClient.Search.WithBody(query),
			)

			Expect(err).ShouldNot(HaveOccurred())

			decEs := datamodels.ElasticsearchResponseCase{}
			errJsonDecode := json.NewDecoder(res.Body).Decode(&decEs)

			fmt.Println("Total:", decEs.Options.Total.Value)

			Expect(errJsonDecode).ShouldNot(HaveOccurred())
			Expect(decEs.Options.Total.Value).Should(Equal(2))
		})

		It("Запрос на получения документа только по event.rootId не должен вызывать ошибок, должен быть получен результат", func() {
			//"test_module_placeholder_elasticsearch_case_2024_1"

			query := strings.NewReader(`
			{"query": {
				"bool": {
					"must": [
        				{ "match": { "event.rootId": "~1682026696" }}]
					}
				}
			}`)
			res, err := esClient.Search(
				esClient.Search.WithContext(context.Background()),
				esClient.Search.WithIndex("module_placeholder_elasticsearch_thehive_case_2024_2"),
				esClient.Search.WithBody(query),
			)

			Expect(err).ShouldNot(HaveOccurred())

			r := map[string]interface{}{}
			errJsonDecode := json.NewDecoder(res.Body).Decode(&r)
			Expect(errJsonDecode).ShouldNot(HaveOccurred())

			//var isExist bool
			//fmt.Println("Read response message:")
			//fmt.Println(r)
			//for k, v := range r {
			//	if k == "found" {
			//		if tmp, ok := v.(bool); ok {
			//			isExist = tmp
			//		}
			//	}
			//				fmt.Println(k, ":", v)
			//}
			//for k, v := range r {
			//	if k == "hits" {
			//		fmt.Printf("%s: %v\n", k, v)
			//	}
			//}

			Expect(len(r)).Should(Equal(4))
		})

		It("При запросе по всем индексам не должно быть ошибок", func() {
			//"test_module_placeholder_elasticsearch_case_2024_1"

			query := strings.NewReader(`
			{"query": {
				"bool": {
					"must": [
						{ "match": { "source": "rcmsr" }},
        				{ "match": { "event.rootId": "~1682026696" }}]
					}
				}
			}`)
			res, err := esClient.Search(
				esClient.Search.WithContext(context.Background()),
				//esClient.Search.WithIndex("_all"),
				esClient.Search.WithIndex("module_placeholder_elasticsearch_thehive_case_2024_1"),
				esClient.Search.WithBody(query),
			)

			Expect(err).ShouldNot(HaveOccurred())

			decEs := datamodels.ElasticsearchResponseCase{}
			errJsonDecode := json.NewDecoder(res.Body).Decode(&decEs)

			fmt.Println("Total value:", decEs.Options.Total.Value)
			fmt.Println("MaxScore:", decEs.Options.MaxScore)

			if decEs.Options.Total.Value > 0 {
				fmt.Println("Hits:")
				for k, v := range decEs.Options.Hits {
					fmt.Printf("%d.\n", k)
					fmt.Println("_id:", v.ID)
					fmt.Println("index:", v.Index)
					//fmt.Println("source:", v.Source)

					if _, err := esClient.Delete("module_placeholder_elasticsearch_thehive_case_2024_1", v.ID); err != nil {
						fmt.Println("DELETE error:", err)
					}
				}
			}

			Expect(errJsonDecode).ShouldNot(HaveOccurred())
			//Expect(len(des)).Should(Equal(4))
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
