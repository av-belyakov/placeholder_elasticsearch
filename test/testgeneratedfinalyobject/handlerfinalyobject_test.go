package testgeneratedfinalyobject_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/confighandler"
	"placeholder_elasticsearch/coremodule"
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"
	"placeholder_elasticsearch/mongodbinteractions"
	rules "placeholder_elasticsearch/rulesinteraction"
	"placeholder_elasticsearch/supportingfunctions"
)

func readFileJson(fpath, fname string) ([]byte, error) {
	var newResult []byte

	rootPath, err := supportingfunctions.GetRootPath("placeholder_misp")
	if err != nil {
		return newResult, err
	}

	tmp := strings.Split(rootPath, "/")

	fmt.Println("func 'readFileJson', path = ", path.Join(path.Join(tmp[:6]...), fpath, fname))

	f, err := os.OpenFile("/"+path.Join(path.Join(tmp[:6]...), fpath, fname), os.O_RDONLY, os.ModePerm)
	if err != nil {
		return newResult, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		newResult = append(newResult, sc.Bytes()...)
	}

	return newResult, nil
}

var _ = Describe("Handlerfinalyobject", Ordered, func() {
	var (
		errReadFile, errRule, errES       error
		errMongoDB, errMarshal, errESSend error
		exampleByte                       []byte
		listRule                          *rules.ListRule
		esModule                          *elasticsearchinteractions.ElasticSearchModule
		mongodbModule                     *mongodbinteractions.MongoDBModule
		decodeJson                        *coremodule.DecodeJsonMessageSettings
		warnings                          []string
		esClient                          *elasticsearch.Client
		b                                 []byte
		res                               *esapi.Response

		logging              chan datamodels.MessageLogging
		counting             chan datamodels.DataCounterSettings
		chanOutputJsonDecode chan datamodels.ChanOutputDecodeJSON
		chanDecodeDone       chan bool
		chanStoppedCounting  chan struct{}
	)

	BeforeAll(func() {
		logging = make(chan datamodels.MessageLogging)
		counting = make(chan datamodels.DataCounterSettings)
		chanStoppedCounting = make(chan struct{})

		esModule = &elasticsearchinteractions.ElasticSearchModule{
			ChanInputModule:  make(chan elasticsearchinteractions.SettingsInputChan),
			ChanOutputModule: make(chan interface{}),
		}

		go func() {
			fmt.Println("___ Logging START")
			defer fmt.Println("___ Logging STOP")

			for log := range logging {
				if log.MsgType == "error" {
					fmt.Println("----", log, "----")
				}

				if log.MsgType == "STOP TEST" {
					chanStoppedCounting <- struct{}{}

					return
				}
			}
		}()

		//вывод данных счетчика
		go func() {
			for {
				select {
				case d := <-counting:
					fmt.Printf("\tСчетчик %v\n", d.DataType)

				case <-chanStoppedCounting:
					close(esModule.ChanInputModule)
					close(logging)

					return
				}
			}
		}()

		// инициализируем модуль чтения правил обработки сообщений поступающих через NATS
		listRule, warnings, errRule = rules.NewListRule("placeholder_elasticsearch", "rules", "mispmsgrule.yaml")
		exampleByte, errReadFile = readFileJson("test/test_json", "example_caseId_33705.json")

		//инициализируем клиента Elasticsearch
		esClient, errES = elasticsearch.NewClient(elasticsearch.Config{
			Addresses: []string{fmt.Sprintf("http://%s:%d", "datahook.cloud.gcm", 9200)},
			Username:  "writer",
			Password:  "XxZqesYXuk8C",
		})

		// инициализация модуля для взаимодействия с СУБД MongoDB
		mongodbModule, errMongoDB = mongodbinteractions.HandlerMongoDB(confighandler.AppConfigMongoDB{
			Host:   "192.168.9.208",
			Port:   27117,
			User:   "module_placeholder_elasticsearch",
			Passwd: "gDbv5cf7*F2",
			NameDB: "placeholder_elasticsearch",
		}, logging, counting)

		decodeJson = coremodule.NewDecodeJsonMessageSettings(listRule, logging, counting)
		chanOutputJsonDecode, chanDecodeDone = decodeJson.HandlerJsonMessage(exampleByte, uuid.NewString(), "subject_case")
		go coremodule.NewVerifiedTheHiveFormatCase(chanOutputJsonDecode, chanDecodeDone, mongodbModule, logging)
	})

	Context("Тест 1. Проверка чтения тестового JSON файла", func() {
		It("При чтении файла не должно быть ошибок", func() {
			Expect(errReadFile).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Проверка четния файла с правилами", func() {
		It("При чтении файла не должно быть ошибок", func() {
			Expect(errRule).ShouldNot(HaveOccurred())
		})

		It("Ошибок в написании правил быть не должно", func() {
			Expect(len(warnings)).Should(Equal(0))
		})
	})

	Context("Тест 3. Проверка подключения к СУБД", func() {
		It("При подключении к Elasticsearch не должно быть ошибок", func() {
			Expect(errES).ShouldNot(HaveOccurred())
		})
		It("При подключении к MongoDB не должно быть ошибок", func() {
			Expect(errMongoDB).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 4. Получаем валидированный финальный объект", func() {
		It("Должен быть получен валидный объект", func() {
			esSettings := <-esModule.ChanInputModule

			data, ok := esSettings.Data.(datamodels.VerifiedTheHiveCase)
			Expect(ok).Should(BeTrue())

			verifiedObject := data.Get()

			//fmt.Println("func 'NewVerifiedTheHiveFormat' is STOPPED")
			//fmt.Println("------------------ VerifiedObject RESULT ----------------")
			//fmt.Println(verifiedObject.ToStringBeautiful(0))

			//************************************************
			//тестовая отправка данных в Elastisearch
			b, errMarshal = json.Marshal(verifiedObject)

			t := time.Now()
			buf := bytes.NewReader(b)
			res, errESSend = esClient.Index(fmt.Sprintf("%s%s_%d_%d", "", "test_module_placeholder_elasticsearch_case", t.Year(), int(t.Month())), buf)

			r := map[string]interface{}{}
			errJsonDecode := json.NewDecoder(res.Body).Decode(&r)

			fmt.Println("Read response message:")
			for k, v := range r {
				fmt.Println(k, ":", v)
			}

			close(counting)

			Expect(errJsonDecode).ShouldNot(HaveOccurred())
			Expect(res.StatusCode).Should(Equal(http.StatusCreated))

			Expect(errMarshal).ShouldNot(HaveOccurred())
			Expect(errESSend).ShouldNot(HaveOccurred())
			Expect(esSettings.Command).Should(Equal("add new case"))
		})
	})
	/*
		Context("", func ()  {
			It("", func ()  {

			})
		})
	*/
})
