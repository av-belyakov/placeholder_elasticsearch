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
		rootId       string = "~129876088"
		indexName    string = "test_module_placeholder_case"
		indexCurrent string = indexName + "_2024_7"

		querySearch *strings.Reader = strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}", source, rootId))

		/*
			GeoCode     string = "RU-IRK"
			ObjectArea  string = "Наука и образование"
			SubjectRF   string = "Иркутская область"
			INN         string = "3812011682"
			HomeNet     string = "[84.237.16.0/20]"
			OrgName     string = "ФЕДЕРАЛЬНОЕ ГОСУДАРСТВЕННОЕ БЮДЖЕТНОЕ УЧРЕЖДЕНИЕ НАУКИ ИНСТИТУТ ДИНАМИКИ СИСТЕМ И ТЕОРИИ УПРАВЛЕНИЯ ИМЕНИ В.М. МАТРОСОВА СИБИРСКОГО ОТДЕЛЕНИЯ РОССИЙСКОЙ АКАДЕМИИ НАУК"
			FullOrgName string = "ИДСТУ СО РАН"
		*/

		ctxEEM       context.Context
		ctxCancelEEM context.CancelFunc

		confApp               confighandler.ConfigApp
		eventEnrichmentModule *eventenrichmentmodule.EventEnrichmentModule

		errConfApp, errEEM error

		errConn error
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
			SensorsId: []string{"570084", "690017", "690013"},
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

			fmt.Println("RootId:", additionalInformation.GetRootId())
			fmt.Println("Source:", additionalInformation.GetSource())
			for k, v := range additionalInformation.Sensors {
				fmt.Printf("\tSensor:'%s'\n", k)
				fmt.Println("SensorId:", v.SensorId)
				fmt.Println("HostId:", v.HostId)
				fmt.Println("GeoCode:", v.GeoCode)
				fmt.Println("ObjectArea:", v.ObjectArea)
				fmt.Println("SubjectRF:", v.SubjectRF)
				fmt.Println("INN:", v.INN)
				fmt.Println("HomeNet:", v.HomeNet)
				fmt.Println("OrgName:", v.OrgName)
				fmt.Println("FullOrgName:", v.FullOrgName)
			}

			//
			//здесь надо сформировать запрос на добавление доп. данных
			// с учетом того, что необхлдимо формировать список объектов
			//

			//выполняем запрос на добавление в кейс дополнительных данных
			//-----------------------------------------------------------

			//этот запрос уже не подходит
			bodyUpdate := strings.NewReader(fmt.Sprintf("{\"doc\": {\"@geoCode\": \"%s\", \"@objectArea\": \"%s\", \"@subjectRF\": \"%s\", \"@inn\": \"%s\", \"@homeNet\": \"%s\", \"@orgName\": \"%s\", \"@fullOrgName\": \"%s\"}}}", GeoCode, ObjectArea, SubjectRF, INN, HomeNet, OrgName, FullOrgName))

			//bodyUpdate := strings.NewReader(fmt.Sprintf("{\"must\": [{\"script\": \"ctx._source.event.GeoCode='%s'\"}, {\"script\": \"ctx._source.event.ObjectArea='%s'\"}, {\"script\": \"ctx._source.event.SubjectRF='%s'\"}, {\"script\": \"ctx._source.event.INN='%s'\"}, {\"script\": \"ctx._source.event.HomeNet='%s'\"}, {\"script\": \"ctx._source.event.OrgName='%s'\"}, {\"script\": \"ctx._source.event.FullOrgName='%s'\"}]}", GeoCode, ObjectArea, SubjectRF, INN, HomeNet, OrgName, FullOrgName))
			//bodyUpdate := strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"script\": \"ctx._source.event.GeoCode='%s'\"}, {\"script\": \"ctx._source.event.ObjectArea='%s'\"}, {\"script\": \"ctx._source.event.SubjectRF='%s'\"}, {\"script\": \"ctx._source.event.INN='%s'\"}, {\"script\": \"ctx._source.event.HomeNet='%s'\"}, {\"script\": \"ctx._source.event.OrgName='%s'\"}, {\"script\": \"ctx._source.event.FullOrgName='%s'\"}]}}}", GeoCode, ObjectArea, SubjectRF, INN, HomeNet, OrgName, FullOrgName))
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

			/*

				здесь по тестам все проходит, однако сенсоров может быть много
				поэтому следуюе продумать как в СУБД добавлять информацию по нескольким
				сенсорам

			*/

			ctxCancelEEM()

			Expect(len(tmp.Options.Hits)).ShouldNot(Equal(0))
		})
	})
})
