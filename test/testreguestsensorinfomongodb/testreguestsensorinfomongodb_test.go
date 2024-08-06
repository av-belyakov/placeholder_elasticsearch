package testreguestsensorinfomongodb_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"placeholder_elasticsearch/datamodels"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ = Describe("Testreguestsensorinfomongodb", Ordered, func() {
	const (
		Host           = "192.168.9.208"
		Port           = 27117
		NameDB         = "placeholder_elasticsearch"
		User           = "module_placeholder_elasticsearch"
		Passwd         = "gDbv5cf7*F2"
		CollectionName = "collection_sensor_information"
	)

	var (
		ctx       context.Context
		ctxCancel context.CancelFunc

		connMdb *mongo.Client

		f *os.File

		errMdb, errFile error
	)

	BeforeAll(func() {
		clientOption := options.Client().SetAuth(options.Credential{
			AuthMechanism: "SCRAM-SHA-256",
			AuthSource:    NameDB,
			Username:      User,
			Password:      Passwd,
		})

		confPath := fmt.Sprintf("mongodb://%s:%d/%s", Host, Port, NameDB)

		fmt.Println("NewConnection: ", confPath)
		ctx, ctxCancel = context.WithTimeout(context.Background(), 7*time.Second)

		connMdb, errMdb = mongo.Connect(ctx, clientOption.ApplyURI(confPath))

		f, errFile = os.OpenFile("sensors_information.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		//bufio.NewWriter()
	})

	Context("Тест 1. Подключение к СУБД", func() {
		It("При подключении к СУБД не должно быть ошибки", func() {
			Expect(errMdb).ShouldNot(HaveOccurred())
		})

		It("При создании дескриптора файла не должно быть ошибки", func() {
			Expect(errFile).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Поиск информации по сенсорам", func() {
		It("При выполнении поиска не должно быть ошибок", func() {
			collection := connMdb.Database(NameDB).Collection(CollectionName)
			opts := options.Find().SetAllowDiskUse(true).SetSort(bson.D{{Key: "sensorId", Value: 1}})
			cur, err := collection.Find(
				context.TODO(),
				bson.D{
					//{Key: "source", Value: "gcm"},
					//{Key: "event.rootId", Value: "~85917298816"},
				},
				opts,
			)
			Expect(err).ShouldNot(HaveOccurred())

			sensorsInformation := []datamodels.SensorInformation(nil)
			for cur.Next(context.Background()) {
				var sensorInfo datamodels.SensorInformation
				if err := cur.Decode(&sensorInfo); err != nil {
					Expect(err).ShouldNot(HaveOccurred())
				}

				fmt.Println("SensorId: ", sensorInfo.SensorId)

				sensorsInformation = append(sensorsInformation, sensorInfo)
			}

			b, err := json.MarshalIndent(sensorsInformation, "", " ")
			Expect(err).ShouldNot(HaveOccurred())

			_, err = f.Write(b)
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	AfterAll(func() {
		connMdb.Disconnect(context.Background())
		ctxCancel()
		f.Close()
	})
})
