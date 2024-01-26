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

	"placeholder_elasticsearch/coremodule"
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"
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
		errReadFile, errRule, errES, errMarshal, errESSend error
		exampleByte                                        []byte
		listRule                                           *rules.ListRule
		esModule                                           *elasticsearchinteractions.ElasticSearchModule
		decodeJson                                         *coremodule.DecodeJsonMessageSettings
		warnings                                           []string
		esClient                                           *elasticsearch.Client
		b                                                  []byte
		res                                                *esapi.Response

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

		decodeJson = coremodule.NewDecodeJsonMessageSettings(listRule, logging, counting)
		chanOutputJsonDecode, chanDecodeDone = decodeJson.HandlerJsonMessage(exampleByte, uuid.NewString())
		go coremodule.NewVerifiedTheHiveFormat(chanOutputJsonDecode, chanDecodeDone, esModule, logging)
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

	Context("Тест 3. Проверка подключения к Elasticsearch", func() {
		It("При подключении не должно быть ошибок", func() {
			Expect(errES).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 4. Получаем валидированный финальный объект", func() {
		It("Должен быть получен валидный объект", func() {
			var esSettings elasticsearchinteractions.SettingsInputChan

			for esSettings = range esModule.ChanInputModule {
				verifiedObject := esSettings.VerifiedObject.Get()

				fmt.Println("func 'NewVerifiedTheHiveFormat' is STOPPED")
				fmt.Println("------------------ VerifiedObject RESULT ----------------")
				fmt.Println(verifiedObject.ToStringBeautiful(0))
				//fmt.Println("CreateTimestatmp:", verifiedObject.CreateTimestatmp)
				//fmt.Println("Source:", verifiedObject.Source)
				//fmt.Println("Event:", verifiedObject.Event)
				//fmt.Println("Observables:", verifiedObject.Observables)

				//************************************************
				//тестовая отправка данных в Elastisearch

				b, errMarshal = json.Marshal(verifiedObject)

				t := time.Now()
				buf := bytes.NewReader(b)
				res, errESSend = esClient.Index(fmt.Sprintf("%s%s_%d_%d", "", "test_module_placeholder_elasticsearch_case", t.Year(), int(t.Month())), buf)

				/*if res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusOK {
					return
				}

				var errMsg string
				r := map[string]interface{}{}
				if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
					_, f, l, _ := runtime.Caller(0)
					hsd.logging <- datamodels.MessageLogging{
						MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
						MsgType: "error",
					}
				}

				if e, ok := r["error"]; ok {
					errMsg = fmt.Sprintln(e)
				}*/
			}

			close(counting)

			Expect(res.StatusCode).Should(Equal(http.StatusCreated))

			Expect(errMarshal).ShouldNot(HaveOccurred())
			Expect(errESSend).ShouldNot(HaveOccurred())
			Expect(esSettings.Command).Should(Equal("add new object"))
		})
	})
	/*
		Context("", func ()  {
			It("", func ()  {

			})
		})
	*/
})
