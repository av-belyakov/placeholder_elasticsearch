package testgetsensorseventenrichmentmpngodb_test

import (
	"context"
	"encoding/json"
	"fmt"
	"placeholder_elasticsearch/confighandler"
	"placeholder_elasticsearch/datamodels"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ = Describe("Testgetsensorseventenrichmentmpngodb", Ordered, func() {
	var (
		confApp confighandler.ConfigApp
		confMdb *confighandler.AppConfigMongoDB
		connMdb *mongo.Client

		errApp, errMdb error

		ctxTimeout context.CancelFunc
	)

	BeforeAll(func() {
		confApp, errApp = confighandler.NewConfig("placeholder_elasticsearch")

		confMdb = confApp.GetAppMongoDB()

		clientOption := options.Client().SetAuth(options.Credential{
			AuthMechanism: "SCRAM-SHA-256",
			AuthSource:    confMdb.NameDB,
			Username:      confMdb.User,
			Password:      confMdb.Passwd,
		})

		confPath := fmt.Sprintf("mongodb://%s:%d/%s", confMdb.Host, confMdb.Port, confMdb.NameDB)

		fmt.Println("NewConnection: ", confPath)

		var ctx context.Context
		ctx, ctxTimeout = context.WithTimeout(context.Background(), 7*time.Second)

		connMdb, errMdb = mongo.Connect(ctx, clientOption.ApplyURI(confPath))
	})

	Context("Тест 1. Проверка на наличие ошибок при инициализации модулей", func() {
		It("При инициализации модуля конфигурации не должно быть ошибок", func() {
			Expect(errApp).ShouldNot(HaveOccurred())
		})

		It("При инициализации модуля взаимодействия с MongoDB не должно быть ошибок", func() {
			Expect(errMdb).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Поиск информации о сенсоре", func() {
		It("При поиске информации не должно быть ошибки, информация должна быть успешно найдена", func() {
			collection := connMdb.Database(confMdb.NameDB).Collection("collection_sensor_information")

			opts := &options.FindOneOptions{}
			result := datamodels.SensorInformation{}
			err := collection.FindOne(
				context.TODO(),
				bson.D{{Key: "sensorId", Value: "310073"}},
				opts,
			).Decode(&result)
			Expect(err).ShouldNot(HaveOccurred())

			b, err := json.MarshalIndent(result, "", " ")
			Expect(err).ShouldNot(HaveOccurred())

			fmt.Println("Result:")
			fmt.Println(string(b))

			Expect(result.GetINN()).Should(Equal("1435073060"))
		})
	})

	Context("Тест 3. Поиск информации о списке сенсоров", func() {
		It("При поиске ошибок быть не должно, должно быть найдено определенное количество сенсоров", func() {
			sensorsId := []string{
				"500065",
				"310052",
				"570062",
				"570097",
				"8030070",
				"8030090",
				"8030013",
			}

			collection := connMdb.Database(confMdb.NameDB).Collection("collection_sensor_information")
			options := options.Find().SetAllowDiskUse(true)
			cur, err := collection.Find(
				context.Background(),
				(bson.D{
					{Key: "sensorId", Value: bson.D{
						{Key: "$in", Value: sensorsId}},
					},
				}), options)
			Expect(err).ShouldNot(HaveOccurred())

			var sensorsInfo []datamodels.SensorInformation
			for cur.Next(context.Background()) {
				var sensorInfo datamodels.SensorInformation
				err := cur.Decode(&sensorInfo)
				Expect(err).ShouldNot(HaveOccurred())

				sensorsInfo = append(sensorsInfo, sensorInfo)
			}

			Expect(len(sensorsInfo)).Should(Equal(7))
		})
	})

	AfterAll(func() {
		ctxTimeout()
	})
})
