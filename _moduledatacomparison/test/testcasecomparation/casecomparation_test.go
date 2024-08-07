package casecomparation_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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

var _ = Describe("Casecomparation", Ordered, func() {
	var (
		confApp confighandler.ConfigApp

		ctxMdb       context.Context
		ctxMdbCancel context.CancelFunc
		connMdb      *mongo.Client

		connEs *elasticsearch.Client

		errMdb, errEs, errApp error
	)

	BeforeAll(func() {
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
			verifedCases := []datamodels.VerifiedTheHiveCase(nil)
			for cur.Next(context.Background()) {
				var verifedCase datamodels.VerifiedTheHiveCase
				if err := cur.Decode(&verifedCase); err != nil {
					Expect(err).ShouldNot(HaveOccurred())
				}

				fmt.Printf("%d.\n", num)
				fmt.Println("@id:", verifedCase.GetID())
				fmt.Println("@timestamp:", verifedCase.Get().CreateTimestamp)
				fmt.Println("RootId:", verifedCase.GetEvent().GetRootId())

				verifedCases = append(verifedCases, verifedCase)
				casesId = append(casesId, fmt.Sprintf("{ \"span_term\" : { \"@id\" : \"%s\" } }", verifedCase.GetID()))

				num++
			}

			Expect(len(verifedCases)).Should(Equal(15))

			var str string
			for i := 0; i < len(casesId); i++ {
				if i == (len(casesId) - 1) {
					str += casesId[i]
				} else {
					str += casesId[i] + ","
				}
			}

			fmt.Println("Query string:", str)

			//запрос к Elasticsearch
			//queryCurrent := strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"@id\": \"%s\"}}]}}}", ))
			queryCurrent := strings.NewReader(fmt.Sprintf("{\"query\": {\"span_or\": {\"clauses\": [%s]}}}", str))

			res, err := connEs.Search(
				connEs.Search.WithContext(context.Background()),
				connEs.Search.WithIndex("module_placeholder_new_case_2024_3"),
				connEs.Search.WithBody(queryCurrent))
			Expect(err).ShouldNot(HaveOccurred())

			Expect(res.StatusCode).Should(Equal(http.StatusOK))

			/*msg := []struct {
				id        string `json:"@id"`
				timestamp string `json:"@timestamp"`
				source    string `json:"source"`
			}{}*/
			msg := datamodels.ElasticsearchResponseCase{}
			err = json.NewDecoder(res.Body).Decode(&msg)

			fmt.Println("RESULT:", &msg)

			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	AfterAll(func() {
		ctxMdbCancel()
	})
})
