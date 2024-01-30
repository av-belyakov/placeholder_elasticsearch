package testmongodbconnect_test

import (
	"context"
	"fmt"

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
		Host           = "192.168.9.208"
		Port           = 27117
		NameDB         = "placeholder_elasticsearch"
		User           = "module_placeholder_elasticsearch"
		Passwd         = "gDbv5cf7*F2"
		CollectionName = "case_test"
	)

	BeforeAll(func() {
		clientOption := options.Client()
		clientOption.SetAuth(options.Credential{
			AuthMechanism: "SCRAM-SHA-256",
			AuthSource:    NameDB,
			Username:      User,
			Password:      Passwd,
		})

		confPath := fmt.Sprintf("mongodb://%s:%d/%s", Host, Port, NameDB)

		fmt.Println("NewConnection: ", confPath)

		conn, err = mongo.Connect(context.Background(), options.Client().ApplyURI(confPath))
	})

	Context("Тест 1. Проверка соединения с MongoDB", func() {
		It("Соединение должно быть успешно установлено", func() {
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("При отправки запроса Ping не должно быть ошибок", func() {
			err := conn.Ping(context.Background(), readpref.Primary())
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

			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
