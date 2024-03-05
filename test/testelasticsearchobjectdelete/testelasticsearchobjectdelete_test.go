package testelasticsearchobjectdelete_test

import (
	"context"
	"encoding/json"
	"fmt"
	"placeholder_elasticsearch/datamodels"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testelasticsearchobjectdelete", func() {
	var (
		Host   string = "datahook.cloud.gcm"
		Port   int    = 9200
		User   string = "writer"
		Passwd string = "XxZqesYXuk8C"
	)

	Context("Тест 0. Проверка поиска объекта", func() {
		It("При поиске объекта ошибок быть не должно", func() {
			queryCurrent := strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}", "rcmmsk", "~335884384"))
			es, err := elasticsearch.NewClient(elasticsearch.Config{
				Addresses: []string{fmt.Sprintf("http://%s:%d", Host, Port)},
				Username:  User,
				Password:  Passwd,
			})
			Expect(err).ShouldNot(HaveOccurred())

			res, err := es.Search(
				es.Search.WithContext(context.Background()),
				es.Search.WithIndex([]string{"module_placeholder_new_alert_rcmmsk_2024_3"}...),
				es.Search.WithBody(queryCurrent),
			)
			Expect(err).ShouldNot(HaveOccurred())

			fmt.Println("SEARCH STATUS:", res.Status())

			decEs := datamodels.ElasticsearchResponseCase{}
			err = json.NewDecoder(res.Body).Decode(&decEs)
			Expect(err).ShouldNot(HaveOccurred())

			fmt.Println("func 'FOUND DATA' found data:", decEs.Options.Total.Value)
			fmt.Println("found data:", decEs.Options)

			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	/*Context("Тест 1. Проверка удаления объекта", func() {
		It("При удалении объекта ошибок быть не должно", func() {
			es, err := elasticsearch.NewClient(elasticsearch.Config{
				Addresses: []string{fmt.Sprintf("http://%s:%d", Host, Port)},
				Username:  User,
				Password:  Passwd,
			})
			Expect(err).ShouldNot(HaveOccurred())

			res, err := es.Delete("module_placeholder_new_alert__2024_3", "30VyCY4BNzcaI2GK6Oul")
			fmt.Println(res)

			Expect(err).ShouldNot(HaveOccurred())
		})
	})*/
})
