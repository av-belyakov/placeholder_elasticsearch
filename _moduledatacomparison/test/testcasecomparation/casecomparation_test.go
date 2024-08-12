package casecomparation_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"placeholder_elasticsearch/_moduledatacomparison/datamodel"
	"placeholder_elasticsearch/_moduledatacomparison/decoder"
	"placeholder_elasticsearch/confighandler"
	"placeholder_elasticsearch/datamodels"
)

func ConnectElasticsearch(host string, port int, user, passwd string) (*elasticsearch.Client, error) {
	return elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{fmt.Sprintf("http://%s:%d", host, port)},
		Username:  user,
		Password:  passwd,
		Transport: &http.Transport{
			MaxIdleConns:          10, //число открытых TCP-соединений, которые в данный момент не используются
			IdleConnTimeout:       5,  //время, через которое закрываются такие неактивные соединения
			MaxIdleConnsPerHost:   10, //число неактивных TCP-соединений, которые допускается устанавливать на один хост
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
		},
	})
}

func ConnectMongoDB(ctx context.Context, host string, port int, nameDB, user, passwd string) (*mongo.Client, error) {
	clientOption := options.Client().SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    nameDB,
		Username:      user,
		Password:      passwd,
	})

	confPath := fmt.Sprintf("mongodb://%s:%d/%s", host, port, nameDB)

	fmt.Println("NewConnection: ", confPath)

	return mongo.Connect(ctx, clientOption.ApplyURI(confPath))
}

func ConvertFormatCaseMongoDBToElasticsearch(
	casesFormatMongoDB []datamodels.VerifiedTheHiveCase,
	listCasesIdConvert []string) ([]datamodels.VerifiedEsCase, error) {
	countCasesIdConvert := len(listCasesIdConvert)
	if countCasesIdConvert == 0 {
		return []datamodels.VerifiedEsCase{}, fmt.Errorf("the list of case IDs to be converted should not be empty")
	}

	verifiedEsCases := make([]datamodels.VerifiedEsCase, countCasesIdConvert)

	for _, caseId := range listCasesIdConvert {
		for _, caseInfo := range casesFormatMongoDB {
			if caseId == caseInfo.GetID() {
				verifedCase := datamodels.NewVerifiedEsCase()
				verifedCase.SetID(caseInfo.GetID())
				verifedCase.SetSource(caseInfo.GetSource())
				verifedCase.SetCreateTimestamp(caseInfo.GetCreateTimestamp())

			}
		}
	}

	return verifiedEsCases, nil
}

var _ = Describe("Casecomparation", Ordered, func() {
	var (
		confApp confighandler.ConfigApp

		ctxMdb       context.Context
		ctxMdbCancel context.CancelFunc
		connMdb      *mongo.Client

		connEs *elasticsearch.Client

		logging chan datamodels.MessageLogging

		f                          *os.File
		err, errMdb, errEs, errApp error
	)

	BeforeAll(func() {
		logging := make(chan datamodels.MessageLogging)

		go func() {
			for msg := range logging {
				fmt.Println("LOG MSG:", msg)
			}
		}()

		go func() {
			f, err = os.Create("proff.out")
			if err != nil {
				log.Fatal("could not create CPU profile: ", err)
			}

			if err = pprof.StartCPUProfile(f); err != nil {
				log.Fatal("could not start CPU profile: ", err)
			}
		}()

		//конфигурационный файл
		os.Setenv("GO_PHELASTIC_MAIN", "development")
		confApp, errApp = confighandler.NewConfig("placeholder_elasticsearch")

		//инициализация соединения с MongoDB
		confMdb := confApp.GetAppMongoDB()
		ctxMdb, ctxMdbCancel = context.WithTimeout(context.Background(), 7*time.Second)
		connMdb, errMdb = ConnectMongoDB(ctxMdb, confMdb.Host, confMdb.Port, confMdb.NameDB, confMdb.User, confMdb.Passwd)

		//инициализация соединения с Elasticsearch
		confEs := confApp.GetAppES()
		connEs, errEs = ConnectElasticsearch(confEs.Host, confEs.Port, confEs.User, confEs.Passwd)
	})

	Context("Тест 1. Проверка инициализации модулей", func() {
		It("При инициализации модуля чтения конфигурационного файла не должно быть ошибок", func() {
			Expect(errApp).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Верификация данных", func() {
		It("Должны быть определены верные параметры конфигурации", func() {
			Expect(confApp.GetAppES().Host).Should(Equal("datahook.cloud.gcm"))
			Expect(confApp.GetAppMongoDB().Host).Should(Equal("192.168.9.208"))
		})
	})

	Context("Тест 3. Подключение к СУБД", func() {
		It("При подключении к MongoDB не должно быть ошибки", func() {
			Expect(errMdb).ShouldNot(HaveOccurred())
		})

		It("При подключении к Elasticsearch не должно быть ошибки", func() {
			Expect(errEs).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 4. Выполнение запросов к MongoDB", func() {
		It("При выполнении запросов не должно быть ошибки", func() {
			confMdb := confApp.GetAppMongoDB()

			//запрос к MongoDB
			collection := connMdb.Database(confMdb.NameDB).Collection("case_collection")
			opts := options.Count()
			count, err := collection.CountDocuments(ctxMdb, bson.D{}, opts)
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Println("Count document:", count)
			Expect(count).ShouldNot(Equal(0))

			newopts := options.Find().SetAllowDiskUse(true).SetSort(bson.D{{Key: "@timestamp", Value: 1}}).SetSkip(int64(10)).SetLimit(int64(15))
			cur, err := collection.Find(ctxMdb, bson.D{}, newopts)
			Expect(err).ShouldNot(HaveOccurred())

			num := 1
			casesId := []string(nil)
			verifedTheHiveCases := []datamodel.VerifiedTheHiveCase(nil)
			for cur.Next(context.Background()) {
				var verifedCase datamodel.VerifiedTheHiveCase
				if err := cur.Decode(&verifedCase); err != nil {
					Expect(err).ShouldNot(HaveOccurred())
				}

				//fmt.Printf("%d.\n", num)
				//fmt.Println("@id:", verifedCase.GetID())
				//fmt.Println("@timestamp:", verifedCase.Get().CreateTimestamp)
				//fmt.Println("RootId:", verifedCase.GetEvent().GetRootId())

				verifedTheHiveCases = append(verifedTheHiveCases, verifedCase)
				//casesId = append(casesId, fmt.Sprintf("{ \"span_term\" : { \"@id\" : \"%s\" } }", verifedCase.GetID()))
				casesId = append(casesId, verifedCase.GetID())

				num++
			}

			Expect(len(verifedTheHiveCases)).Should(Equal(15))
			Expect(len(casesId)).ShouldNot(Equal(0))
			/*var str string
			for i := 0; i < len(casesId); i++ {
				str += fmt.Sprintf("{\"span_term\": {\"@id\": \"%s\"}}", casesId[i])
				if i < (len(casesId) - 1) {
					str += ", "
				}
			}

			fmt.Println("Query string:", str)*/

			//запрос к Elasticsearch
			//SetSkip(int64(10)).SetLimit(int64(15) таких кейсов в эластике нет
			//queryCurrent := strings.NewReader(fmt.Sprintf("{\"query\": {\"span_or\": {\"clauses\": [%s]}}}", str))

			//есть похожие кейсы
			queryCurrent := strings.NewReader("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"@id\": \"e3c9aa7a-d4e8-46f3-90d4-dae71434cde2\"}}]}}}")
			//нет похожих кейсов
			//queryCurrent := strings.NewReader("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"@id\": \"2dbd8d6f-657e-4d61-af95-9945f81bc340\"}}]}}}")

			res, err := connEs.Search(
				connEs.Search.WithContext(context.Background()),
				connEs.Search.WithIndex("module_placeholder_new_case_2024_3"), //	1031
				//connEs.Search.WithIndex("module_placeholder_new_case_2024_8"),
				connEs.Search.WithBody(queryCurrent))
			Expect(err).ShouldNot(HaveOccurred())

			Expect(res.StatusCode).Should(Equal(http.StatusOK))

			msg := datamodels.ElasticsearchResponseCase{}
			err = json.NewDecoder(res.Body).Decode(&msg)

			fmt.Println("_________________________________________________")
			fmt.Println("____ RESULT Total:", &msg.Options.Total)

			foundCaseIdElasticsearch := []string(nil)
			for _, v := range msg.Options.Hits {
				foundCaseIdElasticsearch = append(foundCaseIdElasticsearch, v.Source.GetID())
			}

			fmt.Println("Found id:", strings.Join(foundCaseIdElasticsearch, ","))

			//конвертируем, полученные их MongoDB, кейсы, которых не хватает в
			//Elasticsearch в формат подходящий для этой БД формат
			newVerifedEsCases := []datamodels.VerifiedEsCase(nil)
			for _, vthc := range verifedTheHiveCases {
				var isExist bool
				for _, caseId := range foundCaseIdElasticsearch {
					if vthc.GetID() == caseId {
						isExist = true

						break
					}
				}

				if vthc.GetID() == "f52930a9-5caf-4299-b00b-781b3aba094f" {
					fmt.Println(vthc.ToStringBeautiful(0))

					//if b, err := json.MarshalIndent(vthc, "", " "); err == nil {
					//	fmt.Println("___________________________")
					//	fmt.Println(string(b))
					//	fmt.Println("___________________________")
					//}
				}

				if !isExist {
					if b, err := json.Marshal(vthc); err == nil {
						newVerifedCase := datamodels.VerifiedEsCase{}

						isOk := decoder.FormatCaseJsonMongoDBHandler(b, &newVerifedCase, logging)

						if isOk {
							newVerifedEsCases = append(newVerifedEsCases, newVerifedCase)
						}
					}
				}
			}

			fmt.Println("____ NEW VERIFIED CASE for Elasticsherach: _____")
			fmt.Println("Count:", len(newVerifedEsCases))
			for k, v := range newVerifedEsCases {
				fmt.Printf("%d.\n", k)
				if v.GetID() == "f52930a9-5caf-4299-b00b-781b3aba094f" {
					fmt.Println(v.ToStringBeautiful(0))
				}
			}

			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	AfterAll(func() {
		ctxMdbCancel()
		//close(logging)

		defer f.Close() // error handling omitted for example
		defer pprof.StopCPUProfile()
	})
})
