package testconfighandler_test

import (
	"os"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/confighandler"
)

var _ = Describe("Confighandler", func() {
	var (
		err  error
		conf confighandler.ConfigApp
	)

	Context("Тест 1. Чтение конфигурационного файла (по умолчанию config_prod.yaml)", func() {
		conf, err = confighandler.NewConfig("placeholder_elasticsearch")

		It("При чтении конфигурационного файла ошибок быть не должно", func() {
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("Все пораметры конфигрурационного файла для NATS должны быть успешно получены", func() {
			cn := conf.GetAppNATS()

			Expect(cn.Host).Should(Equal("nats.cloud.gcm"))
			Expect(cn.Port).Should(Equal(4222))
			Expect(cn.SubjectCase).Should(Equal("main_caseupdate"))
			Expect(cn.SubjectAlert).Should(Equal("main_alertupdate"))
		})
		It("Все пораметры конфигрурационного файла для Elasticsearch должны быть успешно получены", func() {
			ces := conf.GetAppES()

			Expect(ces.Host).Should(Equal("datahook.cloud.gcm"))
			Expect(ces.Port).Should(Equal(9200))
			Expect(ces.User).Should(Equal("writer"))
			Expect(ces.Passwd).Should(Equal("XxZqesYXuk8C"))
			Expect(ces.PrefixCase).Should(Equal(""))
			Expect(ces.IndexCase).Should(Equal("module_placeholder_case"))
			Expect(ces.PrefixAlert).Should(Equal(""))
			Expect(ces.IndexAlert).Should(Equal("module_placeholder_alert"))
		})
		It("Все пораметры конфигрурационного файла для MongoDB должны быть успешно получены", func() {
			cmdb := conf.GetAppMongoDB()

			Expect(cmdb.Host).Should(Equal("172.10.11.2"))
			Expect(cmdb.Port).Should(Equal(27017))
			Expect(cmdb.User).Should(Equal("module_placeholder_elasticsearch"))
			Expect(cmdb.Passwd).Should(Equal("gDbv5cf7*F2"))
			Expect(cmdb.NameDB).Should(Equal("placeholder_elasticsearch"))
		})
		It("Все пораметры конфигрурационного файла для загрузки правил должны быть успешно получены", func() {
			cr := conf.GetAppRulesProcMsg()

			Expect(cr.Directory).Should(Equal("rules"))
			Expect(cr.FileCase).Should(Equal("msgrule_case.yaml"))
			Expect(cr.FileAlert).Should(Equal("msgrule_alert.yaml"))
		})
	})

	Context("Тест 2. Проверяем установленные значения переменных окружения", func() {
		const (
			NATS_HOST         = "nats.cloud.gcm.test.test"
			NATS_PORT         = 4545
			NATS_SUBJECTCASE  = "main_CASEupdate"
			NATS_SUBJECTALERT = "main_ALERTupdate"

			ES_HOST         = "datahook.cloud.gcm.test.test"
			ES_PORT         = 99999
			ES_USER         = "writer.test.test"
			ES_PASSWD       = "XxZqesYXuk8C.test.test"
			ES_PREFIX_CASE  = "any more"
			ES_INDEX_CASE   = "module_placeholder_elasticsearch.test.test"
			ES_PREFIX_ALERT = "more alert"
			ES_INDEX_ALERT  = "module_placeholder_elasticsearch.alert.test"

			MDB_HOST   = "199.166.199.166"
			MDB_PORT   = 11111
			MDB_USER   = "module_placeholder_elasticsearch.test.test"
			MDB_PASSWD = "gDbv5cf7*F2.test.test"
			MDB_NAMEDB = "placeholder_elasticsearch.test.test"

			R_DIR  = "rules_dir_test"
			R_FILE = "filre_reules.test.test"
		)

		os.Setenv("GO_PHELASTIC_NHOST", NATS_HOST)
		os.Setenv("GO_PHELASTIC_NPORT", strconv.Itoa(NATS_PORT))
		os.Setenv("GO_PHELASTIC_SUBJECTCASE", NATS_SUBJECTCASE)
		os.Setenv("GO_PHELASTIC_SUBJECTALERT", NATS_SUBJECTALERT)

		os.Setenv("GO_PHELASTIC_ESHOST", ES_HOST)
		os.Setenv("GO_PHELASTIC_ESPORT", strconv.Itoa(ES_PORT))
		os.Setenv("GO_PHELASTIC_ESUSER", ES_USER)
		os.Setenv("GO_PHELASTIC_ESPASSWD", ES_PASSWD)
		os.Setenv("GO_PHELASTIC_ESPREFIXCASE", ES_PREFIX_CASE)
		os.Setenv("GO_PHELASTIC_ESINDEXCASE", ES_INDEX_CASE)
		os.Setenv("GO_PHELASTIC_ESPREFIXALERT", ES_PREFIX_ALERT)
		os.Setenv("GO_PHELASTIC_ESINDEXALERT", ES_INDEX_ALERT)

		os.Setenv("GO_PHELASTIC_MONGOHOST", MDB_HOST)
		os.Setenv("GO_PHELASTIC_MONGOPORT", strconv.Itoa(MDB_PORT))
		os.Setenv("GO_PHELASTIC_MONGOUSER", MDB_USER)
		os.Setenv("GO_PHELASTIC_MONGOPASSWD", MDB_PASSWD)
		os.Setenv("GO_PHELASTIC_MONGONAMEDB", MDB_NAMEDB)

		os.Setenv("GO_PHELASTIC_RULES_DIR", R_DIR)
		os.Setenv("GO_PHELASTIC_RULES_FILECASE", R_FILE)

		confEnv, err := confighandler.NewConfig("placeholder_elasticsearch")

		It("При чтении конфигурационного файла ошибок быть не должно", func() {
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("Все параметры конфигурации для NATS должны быть успешно установлены через соответствующие переменные окружения", func() {
			cn := confEnv.GetAppNATS()

			Expect(cn.Host).Should(Equal(NATS_HOST))
			Expect(cn.Port).Should(Equal(NATS_PORT))
			Expect(cn.SubjectCase).Should(Equal("main_CASEupdate"))
			Expect(cn.SubjectAlert).Should(Equal("main_ALERTupdate"))
		})
		It("Все параметры конфигурации для ELASTICSEARCH должны быть успешно установлены через соответствующие переменные окружения", func() {
			ces := confEnv.GetAppES()

			Expect(ces.Host).Should(Equal(ES_HOST))
			Expect(ces.Port).Should(Equal(ES_PORT))
			Expect(ces.User).Should(Equal(ES_USER))
			Expect(ces.Passwd).Should(Equal(ES_PASSWD))
			Expect(ces.PrefixCase).Should(Equal(ES_PREFIX_CASE))
			Expect(ces.IndexCase).Should(Equal(ES_INDEX_CASE))
			Expect(ces.PrefixAlert).Should(Equal(ES_PREFIX_ALERT))
			Expect(ces.IndexAlert).Should(Equal(ES_INDEX_ALERT))
		})
		It("Все параметры конфигурации для MONGODB должны быть успешно установлены через соответствующие переменные окружения", func() {
			cmdb := confEnv.GetAppMongoDB()

			Expect(cmdb.Host).Should(Equal(MDB_HOST))
			Expect(cmdb.Port).Should(Equal(MDB_PORT))
			Expect(cmdb.User).Should(Equal(MDB_USER))
			Expect(cmdb.Passwd).Should(Equal(MDB_PASSWD))
			Expect(cmdb.NameDB).Should(Equal(MDB_NAMEDB))
		})
	})

	Context("Тест 3. Проверяем работу функции NewConfig с разными значениями переменной окружения GO_PHMISP_MAIN", func() {
		It("Должно быть получено содержимое общего файла 'config.yaml'", func() {
			conf, err := confighandler.NewConfig("placeholder_elasticsearch")

			//fmt.Println("conf = ", conf)

			//for k, v := range conf.GetListOrganization() {
			//	fmt.Printf("%d. OrgName: %s, SourceName: %s\n", k, v.OrgName, v.SourceName)
			//}

			commonApp := conf.GetCommonApp()

			//fmt.Println("------------------------ ZABBIX -------------------------")
			//fmt.Println("NetworkHost:", conf.GetCommonApp().Zabbix.NetworkHost)
			//fmt.Println("NetworkPort:", conf.GetCommonApp().Zabbix.NetworkPort)

			Expect(commonApp.Zabbix.NetworkHost).Should(Equal("192.168.9.45"))
			Expect(commonApp.Zabbix.NetworkPort).Should(Equal(10051))
			Expect(commonApp.Zabbix.ZabbixHost).Should(Equal("test-uchet-db.cloud.gcm"))
			Expect(len(commonApp.Zabbix.EventTypes)).Should(Equal(3))
			Expect(commonApp.Zabbix.EventTypes[0].EventType).Should(Equal("error"))
			Expect(commonApp.Zabbix.EventTypes[0].ZabbixKey).Should(Equal("placeholder_elasticsearch.error"))
			Expect(commonApp.Zabbix.EventTypes[0].IsTransmit).Should(BeTrue())
			Expect(commonApp.Zabbix.EventTypes[0].Handshake.TimeInterval).Should(Equal(0))
			Expect(commonApp.Zabbix.EventTypes[0].Handshake.Message).Should(Equal(""))
			Expect(commonApp.Zabbix.EventTypes[1].EventType).Should(Equal("info"))
			Expect(commonApp.Zabbix.EventTypes[1].ZabbixKey).Should(Equal("placeholder_elasticsearch.info"))
			Expect(commonApp.Zabbix.EventTypes[1].IsTransmit).Should(BeTrue())
			Expect(commonApp.Zabbix.EventTypes[2].EventType).Should(Equal("handshake"))
			Expect(commonApp.Zabbix.EventTypes[2].ZabbixKey).Should(Equal("placeholder_elasticsearch.handshake"))
			Expect(commonApp.Zabbix.EventTypes[2].IsTransmit).Should(BeTrue())

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(conf.GetListLogs())).Should(Equal(8))
		})
	})
})
