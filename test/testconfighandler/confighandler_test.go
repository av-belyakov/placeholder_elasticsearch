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
		conf confighandler.ConfigApp
		err  error
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
		})
		It("Все пораметры конфигрурационного файла для Elasticsearch должны быть успешно получены", func() {
			ces := conf.GetAppES()

			Expect(ces.Send).ShouldNot(BeTrue())
			Expect(ces.Host).Should(Equal("datahook.cloud.gcm"))
			Expect(ces.Port).Should(Equal(9200))
			Expect(ces.Prefix).Should(Equal(""))
			Expect(ces.Index).Should(Equal("module_placeholder_elasticsearch"))
			Expect(ces.User).Should(Equal("writer"))
			Expect(ces.Passwd).Should(Equal("XxZqesYXuk8C"))
		})
		It("Все пораметры конфигрурационного файла для MongoDB должны быть успешно получены", func() {
			cmdb := conf.GetAppMongoDB()

			Expect(cmdb.Host).Should(Equal("192.168.9.208"))
			Expect(cmdb.Port).Should(Equal(37017))
			Expect(cmdb.User).Should(Equal("module_placeholder_elasticsearch"))
			Expect(cmdb.Passwd).Should(Equal("gDbv5cf7*F2"))
			Expect(cmdb.NameDB).Should(Equal("placeholder_elasticsearch"))
		})
	})

	Context("Тест 2. Проверяем установленные значения переменных окружения", func() {
		const (
			NATS_HOST = "nats.cloud.gcm.test.test"
			NATS_PORT = 4545

			ES_SEND   = true
			ES_HOST   = "datahook.cloud.gcm.test.test"
			ES_PORT   = 99999
			ES_PREFIX = "any more"
			ES_INDEX  = "module_placeholder_elasticsearch.test.test"
			ES_USER   = "writer.test.test"
			ES_PASSWD = "XxZqesYXuk8C.test.test"

			MDB_HOST   = "199.166.199.166"
			MDB_PORT   = 11111
			MDB_USER   = "module_placeholder_elasticsearch.test.test"
			MDB_PASSWD = "gDbv5cf7*F2.test.test"
			MDB_NAMEDB = "placeholder_elasticsearch.test.test"
		)

		os.Setenv("GO_PHELASTIC_NHOST", NATS_HOST)
		os.Setenv("GO_PHELASTIC_NPORT", strconv.Itoa(NATS_PORT))

		os.Setenv("GO_PHELASTIC_ESSEND", "true")
		os.Setenv("GO_PHELASTIC_ESHOST", ES_HOST)
		os.Setenv("GO_PHELASTIC_ESPORT", strconv.Itoa(ES_PORT))
		os.Setenv("GO_PHELASTIC_ESPREFIX", ES_PREFIX)
		os.Setenv("GO_PHELASTIC_ESINDEX", ES_INDEX)
		os.Setenv("GO_PHELASTIC_ESUSER", ES_USER)
		os.Setenv("GO_PHELASTIC_ESPASSWD", ES_PASSWD)

		os.Setenv("GO_PHELASTIC_MONGOHOST", MDB_HOST)
		os.Setenv("GO_PHELASTIC_MONGOPORT", strconv.Itoa(MDB_PORT))
		os.Setenv("GO_PHELASTIC_MONGOUSER", MDB_USER)
		os.Setenv("GO_PHELASTIC_MONGOPASSWD", MDB_PASSWD)
		os.Setenv("GO_PHELASTIC_MONGONAMEDB", MDB_NAMEDB)

		confEnv, err := confighandler.NewConfig("placeholder_elasticsearch")

		It("При чтении конфигурационного файла ошибок быть не должно", func() {
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("Все параметры конфигурации для NATS должны быть успешно установлены через соответствующие переменные окружения", func() {
			cn := confEnv.GetAppNATS()

			Expect(cn.Host).Should(Equal(NATS_HOST))
			Expect(cn.Port).Should(Equal(NATS_PORT))
		})
		It("Все параметры конфигурации для ELASTICSEARCH должны быть успешно установлены через соответствующие переменные окружения", func() {
			ces := confEnv.GetAppES()

			Expect(ces.Send).Should(Equal(ES_SEND))
			Expect(ces.Host).Should(Equal(ES_HOST))
			Expect(ces.Port).Should(Equal(ES_PORT))
			Expect(ces.Prefix).Should(Equal(ES_PREFIX))
			Expect(ces.Index).Should(Equal(ES_INDEX))
			Expect(ces.User).Should(Equal(ES_USER))
			Expect(ces.Passwd).Should(Equal(ES_PASSWD))
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
})
