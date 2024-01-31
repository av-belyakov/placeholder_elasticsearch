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
		CollectionName = "case_test"
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
				bson.D{{}},
				opts,
			).Decode(&result)

			for _, v := range result {
				fmt.Println(v)
			}

			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
