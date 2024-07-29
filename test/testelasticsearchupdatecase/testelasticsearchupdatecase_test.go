package testelasticsearchupdatecase_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/confighandler"
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"
	"placeholder_elasticsearch/eventenrichmentmodule"
)

var _ = Describe("Testelasticsearchupdatecase", Ordered, func() {
	var (
		hsd          elasticsearchinteractions.HandlerSendData
		source       string = "rcmlnx"
		rootId       string = "~338731008"
		indexName    string = "test_module_placeholder_case"
		indexCurrent string = indexName + "_2024_7"

		querySearch *strings.Reader = strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}", source, rootId))

		ctxEEM       context.Context
		ctxCancelEEM context.CancelFunc

		confApp               confighandler.ConfigApp
		eventEnrichmentModule *eventenrichmentmodule.EventEnrichmentModule

		errConn, errConfApp, errEEM error
	)

	BeforeAll(func() {
		logging := make(chan datamodels.MessageLogging)

		go func() {
			for msg := range logging {
				fmt.Println("ERROR:", msg.MsgData)
			}
		}()

		confApp, errConfApp = confighandler.NewConfig("placeholder_elasticsearch")
		commonConf := confApp.GetCommonApp()
		confAppEs := confApp.GetAppES()

		hsd = elasticsearchinteractions.HandlerSendData{
			Settings: elasticsearchinteractions.SettingsHandler{
				Port:   confAppEs.Port,
				Host:   confAppEs.Host,
				User:   confAppEs.User,
				Passwd: confAppEs.Passwd,
			},
		}

		errConn = hsd.New()

		ctxEEM, ctxCancelEEM = context.WithCancel(context.Background())
		eventEnrichmentModule, errEEM = eventenrichmentmodule.NewEventEnrichmentModule(ctxEEM, commonConf.NCIRCC, commonConf.ZabbixJsonRPC, logging)

		//получаем информацию о сенсорах
		eventEnrichmentModule.ChanInputModule <- eventenrichmentmodule.SettingsChanInputEEM{
			RootId:    rootId,
			Source:    source,
			SensorsId: []string{"690017", "8030160", "690013", "570084"},
		}
	})

	Context("Тест 0. Проверяем наличие ошибок при инициализации модулей", func() {
		It("При чтении конфигурационного файла не должно быть ошибок", func() {
			Expect(errConfApp).ShouldNot(HaveOccurred())
		})

		It("При подключении не должно быть ошибок", func() {
			Expect(errConn).ShouldNot(HaveOccurred())
		})

		It("При получении дополнительной информации о сенсорах не должно быть ошибок", func() {
			Expect(errEEM).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 1. Проверяем возможность добавление данных в существующий кейс", func() {
		It("При добавлении новых данных к существующему кейсу не должно быть ошибок", func() {
			type tmpRequest struct {
				SensorAdditionalInformation []eventenrichmentmodule.FoundSensorInformation `json:"@sensorAdditionalInformation"`
			}

			//делаем запрос для поиска _id кейса
			//----------------------------------
			res, err := hsd.SearchDocument([]string{indexCurrent}, querySearch)
			Expect(err).ShouldNot(HaveOccurred())

			fmt.Println("Searach Status code:", res.StatusCode, ", Status:", res.Status())
			Expect(res.StatusCode).Should(Equal(http.StatusOK))

			tmp := datamodels.ElasticsearchResponseCase{}
			err = json.NewDecoder(res.Body).Decode(&tmp)
			Expect(err).ShouldNot(HaveOccurred())

			var caseid string
			for _, v := range tmp.Options.Hits {
				caseid = v.ID
			}

			fmt.Println("CASE ID =", caseid)

			Expect(caseid).ShouldNot(Equal(""))

			//получаем найденную, дополнительную информацию по сенсорам
			//---------------------------------------------------------
			additionalInformation := <-eventEnrichmentModule.ChanOutputModule

			func( /*i datamodels.InformationFromEventEnricher*/ i interface{}) {

				infoEvent, ok := i.(datamodels.InformationFromEventEnricher)

				fmt.Println("___ IS OK:", ok)
				fmt.Println("___ infoEvent:", infoEvent)

			}(additionalInformation)

			fmt.Println("RootId:", additionalInformation.GetRootId())
			fmt.Println("Source:", additionalInformation.GetSource())

			tmpReq := tmpRequest{SensorAdditionalInformation: additionalInformation.Sensors}
			request, err := json.MarshalIndent(tmpReq, "", " ")
			Expect(err).ShouldNot(HaveOccurred())

			bodyUpdate := strings.NewReader(fmt.Sprintf("{\"doc\": %s}", string(request)))
			res, err = hsd.Client.Update(indexCurrent, caseid, bodyUpdate)
			Expect(err).ShouldNot(HaveOccurred())

			fmt.Println("Update Status code:", res.StatusCode, ", Status:", res.Status())
			Expect(res.StatusCode).Should(Equal(http.StatusOK))

			t := map[string]interface{}{}
			err = json.NewDecoder(res.Body).Decode(&tmp)
			Expect(err).ShouldNot(HaveOccurred())

			for k, v := range t {
				fmt.Println(k, " - ", v)
			}

			ctxCancelEEM()

			Expect(len(tmp.Options.Hits)).ShouldNot(Equal(0))
		})
	})
})
