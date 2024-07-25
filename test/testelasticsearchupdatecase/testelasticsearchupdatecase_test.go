package testelasticsearchupdatecase_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"
)

var _ = Describe("Testelasticsearchupdatecase", Ordered, func() {
	var (
		Host   string = "datahook.cloud.gcm"
		Port   int    = 9200
		User   string = "writer"
		Passwd string = "XxZqesYXuk8C"

		hsd          elasticsearchinteractions.HandlerSendData
		source       string = "rcmlnx"
		rootId       string = "~129876088"
		indexName    string = "test_module_placeholder_case"
		indexCurrent string = indexName + "_2024_7"

		querySearch *strings.Reader = strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}", source, rootId))

		GeoCode     string = "RU-IRK"
		ObjectArea  string = "Наука и образование"
		SubjectRF   string = "Иркутская область"
		INN         string = "3812011682"
		HomeNet     string = "[84.237.16.0/20]"
		OrgName     string = "ФЕДЕРАЛЬНОЕ ГОСУДАРСТВЕННОЕ БЮДЖЕТНОЕ УЧРЕЖДЕНИЕ НАУКИ ИНСТИТУТ ДИНАМИКИ СИСТЕМ И ТЕОРИИ УПРАВЛЕНИЯ ИМЕНИ В.М. МАТРОСОВА СИБИРСКОГО ОТДЕЛЕНИЯ РОССИЙСКОЙ АКАДЕМИИ НАУК"
		FullOrgName string = "ИДСТУ СО РАН"

		errConn error
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

		/*
			i, err := hsd.DelIndexSetting([]string{indexTest})
			if err != nil {
				fmt.Println("ERROR:", err)
			}

			fmt.Println("INT:", i)
		*/
	})

	Context("Тест 1. Проверяем возможность добавление данных в существующий кейс", func() {
		It("При подключении не должно быть ошибок", func() {
			Expect(errConn).ShouldNot(HaveOccurred())
		})

		It("При добавлении новых данных к существующему кейсу не должно быть ошибок", func() {
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

			//bodyUpdate := strings.NewReader(fmt.Sprintf("{\"doc\": {\"event.GeoCode\": \"%s\", \"event.ObjectArea\": \"%s\", \"event.SubjectRF\": \"%s\", \"event.INN\": \"%s\", \"event.HomeNet\": \"%s\", \"event.OrgName\": \"%s\", \"event.FullOrgName\": \"%s\"}}}", GeoCode, ObjectArea, SubjectRF, INN, HomeNet, OrgName, FullOrgName))
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

			Expect(len(tmp.Options.Hits)).ShouldNot(Equal(0))
		})
	})
})
