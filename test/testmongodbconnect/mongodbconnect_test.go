package testmongodbconnect_test

import (
	"context"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	//"placeholder_elasticsearch/test/testmongodbconnect"
)

var _ = Describe("Mongodbconnect", Ordered, func() {
	var (
		conn *mongo.Client
		err  error
	)

	const (
		Host = "192.168.9.208"
		Port = 27117
		//Port = 37017
		NameDB = "placeholder_elasticsearch"
		//NameDB = "isems-mrsict"
		User = "module_placeholder_elasticsearch"
		//User = "module-isems-mrsict"
		Passwd = "gDbv5cf7*F2"
		//Passwd = "vkL6Znj$Pmt1e1"
		CollectionName = "case_collection"
		//CollectionName = "stix_object_collection"
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
		ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
		defer cancel()

		conn, err = mongo.Connect(ctx, clientOption.ApplyURI(confPath))
	})

	Context("Тест 1. Проверка соединения с MongoDB", func() {
		It("Соединение должно быть успешно установлено", func() {
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("При отправки запроса Ping не должно быть ошибок", func() {
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			err := conn.Ping(ctx, readpref.Primary())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Отправка тестового запроса", func() {
		It("Не должно быть ошибки", func() {
			collection := conn.Database(NameDB).Collection(CollectionName)

			opts := options.FindOne()
			var result bson.M
			err := collection.FindOne(
				context.TODO(),
				bson.D{{Key: "@id", Value: "6b3be7fe-94a5-4143-a923-70b2a44720cc"}},
				opts,
			).Decode(&result)

			for _, v := range result {
				fmt.Println(v)
			}

			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 3. Поиск объектов для дальнейшего удаления", func() {
		It("Некоторое количество объектов должно быть найдено", func() {
			collection := conn.Database(NameDB).Collection(CollectionName)
			opts := options.Find().SetAllowDiskUse(true).SetSort(bson.D{{Key: "@timestamp", Value: 1}})
			cur, err := collection.Find(
				context.TODO(),
				bson.D{
					{Key: "source", Value: "gcm"},
					{Key: "event.rootId", Value: "~85917298816"},
				},
				opts,
			)
			//"source": "gcm", "event.rootId": "~85917298816"
			Expect(err).ShouldNot(HaveOccurred())

			//****************************************************
			// Я не мог взять тип datamodels.VerifiedTheHiveCase
			// так как при дикодировании документа из MongoDB
			// возникают какие то трудности с полями event.object.customFields и
			// event.object.customFields возможно из-за определения этих
			// полей как interfece CustomerFields
			// Для текущего типа хватит и такой реализации
			// однако в далбнейшем этот вопрос придется решать
			//****************************************************

			list := []string(nil)

			for cur.Next(context.Background()) {
				var modelType struct {
					ID     string `bson:"@id"`
					Source string `bson:"source"`
				}
				if err := cur.Decode(&modelType); err != nil {

					fmt.Println("Decode ERROR:", err)

					continue
				}

				list = append(list, modelType.ID)
			}

			for _, v := range list {
				fmt.Println(v)
			}

			//{"source": "gcm", "event.rootId": "~85917298816"}
			Expect(len(list)).Should(Equal(2))
		})
	})
})
