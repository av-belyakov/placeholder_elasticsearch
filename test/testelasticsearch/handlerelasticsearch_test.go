package testelasticsearch_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	//"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"
	"placeholder_elasticsearch/test/testelasticsearch"
)

func GetVerifiedForEsAlert(res *esapi.Response) (datamodels.ElasticsearchPatternVerifiedForEsAlert, error) {
	mp := datamodels.ElasticsearchPatternVerifiedForEsAlert{}
	err := json.NewDecoder(res.Body).Decode(&mp)
	if err != nil {
		if err != io.EOF {
			return mp, err
		}
	}

	return mp, nil
}

var _ = Describe("Handlerelasticsearch", Ordered, func() {
	var (
		Host   string = "datahook.cloud.gcm"
		Port   int    = 9200
		User   string = "writer"
		Passwd string = "XxZqesYXuk8C"
	)

	/*Context("Тест 1. Поиск, обновление и удаление индексов с Events", func() {
		var (
			index                      string = "module_placeholder_case"
			indexCurrent, indexPattern string
			source                     string = "gcm"
			rootId                     string = "~84625227848"
			queryCurrent               *strings.Reader
			hsd                        elasticsearchinteractions.HandlerSendData

			newVerifiedForEsAlert *datamodels.VerifiedForEsAlert = datamodels.NewVerifiedForEsAlert()
			verifiedForEsAlert    *datamodels.VerifiedTheHiveCase //= datamodels.NewVerifiedFor

			errConn error
		)
	})*/

	Context("Тест 2. Поиск, обновление и удаление индексов с Alerts", func() {
		var (
			index                      string = "module_placeholder_new_alert"
			indexCurrent, indexPattern string
			source                     string = "gcm"
			rootId                     string = "~86079815800" //только это objectId
			queryCurrent               *strings.Reader
			hsd                        elasticsearchinteractions.HandlerSendData

			newVerifiedForEsAlert *datamodels.VerifiedForEsAlert = datamodels.NewVerifiedForEsAlert()
			verifiedForEsAlert    *datamodels.VerifiedForEsAlert = datamodels.NewVerifiedForEsAlert()

			errConn error
		)

		verifiedForEsAlert.SetSource(source)
		verifiedForEsAlert.SetID("jf99r3u9rtt045059y9h49yh9fef93")
		verifiedForEsAlert.SetCreateTimestatmp("2024-02-06T15:37:52+03:00")
		verifiedForEsAlert.SetEvent(testelasticsearch.EventForEsAlertTestOne)
		verifiedForEsAlert.SetAlert(testelasticsearch.AlertForEsAlertTestOne)

		newVerifiedForEsAlert.SetSource(source)
		newVerifiedForEsAlert.SetID("jf99r3u9rtt045059y9h49yh9fef93")
		newVerifiedForEsAlert.SetCreateTimestatmp("2024-02-06T15:37:52+03:00")
		newVerifiedForEsAlert.SetEvent(testelasticsearch.EventForEsAlertTestTwo)
		newVerifiedForEsAlert.SetAlert(testelasticsearch.AlertForEsAlertTestTwo)

		BeforeAll(func() {
			t := time.Now()
			month := int(t.Month())

			indexPattern = fmt.Sprintf("module_placeholder_new_alert_%s_%d", source, t.Year())
			indexCurrent = fmt.Sprintf("%s_%s_%d_%d", index, source, t.Year(), month)

			hsd = elasticsearchinteractions.HandlerSendData{
				Settings: elasticsearchinteractions.SettingsHandler{
					Port:   Port,
					Host:   Host,
					User:   User,
					Passwd: Passwd,
				},
			}

			errConn = hsd.New()

			queryCurrent = strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.objectId\": \"%s\"}}]}}}", source, rootId))
		})

		It("При подключении не должно быть ошибок", func() {
			//fmt.Println("indexCurrent:", indexCurrent)
			//fmt.Println("indexBefore:", indexBefore)
			//fmt.Println("queryCurrent:", queryCurrent)

			Expect(errConn).ShouldNot(HaveOccurred())
		})

		It("Запросы должны быть обработаны без ошибок", func() {
			indexes, err := hsd.GetExistingIndexes(indexPattern)
			Expect(err).ShouldNot(HaveOccurred())

			b, err := json.Marshal(verifiedForEsAlert.Get())
			Expect(err).ShouldNot(HaveOccurred())

			fmt.Println("INDEXES:", indexes)

			if len(indexes) == 0 {
				//ЭТО ВЫПОЛЯЕТСЯ ТОЛЬКО КОГДА ПОХОЖИЙ ИНДЕКС НЕ НАЙДЕН

				res, err := hsd.InsertDocument("my_tag", indexCurrent, b)
				Expect(err).ShouldNot(HaveOccurred())

				fmt.Println("Status Code:", res.Status())
				Expect(res.StatusCode).Should(Equal(http.StatusCreated))
			} else {
				res, errSearch := hsd.SearchDocument(indexes, queryCurrent)
				Expect(errSearch).ShouldNot(HaveOccurred())

				fmt.Println("======================= Response ======================")
				fmt.Println("||| Status:", res.Status())
				//fmt.Println("||| Body:", res.Body)

				decEs := datamodels.ElasticsearchResponseCase{}
				err = json.NewDecoder(res.Body).Decode(&decEs)
				Expect(err).ShouldNot(HaveOccurred())

				fmt.Println("==== decEs.Options.Total.Value ===")
				fmt.Println(decEs.Options.Total.Value)

				if decEs.Options.Total.Value == 0 {
					//ВЫПОЛНЯЕТСЯ ТОГДА КОГДА ДОКУМЕНТ НЕ НАЙДЕН

					res, err = hsd.InsertDocument("my_tag", indexCurrent, b)
					Expect(err).ShouldNot(HaveOccurred())

					fmt.Println("Status Code:", res.Status())
					Expect(res.StatusCode).Should(Equal(http.StatusCreated))
				} else {
					/*
						здесь проблемма так как res был УЖЕ обработан ранее

						datamodels.ElasticsearchPatternVerifiedForEsAlert{}
					*/

					//при наличие похожего индекса его замена
					object, err := GetVerifiedForEsAlert(res)
					Expect(err).ShouldNot(HaveOccurred())

					updateVerified := datamodels.NewVerifiedForEsAlert()
					for _, v := range object.Hits.Hits {
						fmt.Println("****************** v.Source.Event *********************")
						fmt.Println(v.Source.Event)

						num, err := updateVerified.Event.ReplacingOldValues(v.Source.Event)
						Expect(err).ShouldNot(HaveOccurred())
						fmt.Println("EVENT Replacing:", num)

						num, err = updateVerified.Alert.ReplacingOldValues(v.Source.Alert)
						Expect(err).ShouldNot(HaveOccurred())
						fmt.Println("ALERT Replacing:", num)
					}

					fmt.Println("_____ OLD DATA ______")
					fmt.Println(updateVerified.ToStringBeautiful(0))
					/*
						num, err := updateVerified.Event.ReplacingOldValues(newVerifiedForEsAlert.Event)
						Expect(err).ShouldNot(HaveOccurred())
						fmt.Println("EVENT Replacing:", num)

						num, err = updateVerified.Alert.ReplacingOldValues(newVerifiedForEsAlert.Alert)
						Expect(err).ShouldNot(HaveOccurred())
						fmt.Println("ALERT Replacing:", num)

						fmt.Println("_____ UPDATE OLD DATA ______")
						fmt.Println(updateVerified.ToStringBeautiful(0))

						/*nvbyte*/
					_, err = json.Marshal(updateVerified)
					Expect(err).ShouldNot(HaveOccurred())

					//Пока временно выключаем замену в БД
					/*
						res, countDel, err := hsd.UpdateDocument(
							"my_tag",
							indexCurrent,
							decEs.Options.Hits,
							nvbyte,
						)

						fmt.Println("Status Code:", res.Status(), " countDel:", countDel)

						Expect(err).ShouldNot(HaveOccurred())
						Expect(res.StatusCode).Should(Equal(http.StatusCreated))
					*/
				}
			}

			Expect(true).Should(BeTrue())
		})
	})
})
