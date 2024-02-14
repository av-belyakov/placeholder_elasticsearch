package testmongodbconnect_test

import (
	"context"
	"errors"
	"fmt"
	"placeholder_elasticsearch/datamodels"
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

			/*for _, v := range result {
				fmt.Println(v)
			}*/

			if errors.Is(err, mongo.ErrNoDocuments) {
				Expect(len(result)).Should(Equal(0))
			} else {
				Expect(err).ShouldNot(HaveOccurred())
			}
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
			// так как при декодировании документа из MongoDB
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
			Expect(len(list)).Should(Equal(1))
		})
	})

	Context("Тест 4. Поиск не существующего значения", func() {
		It("Должна быть успешно обработана ошибка", func() {
			collection := conn.Database(NameDB).Collection("alert_collection")
			opts := options.FindOne()
			var result datamodels.VerifiedTheHiveAlert
			err := collection.FindOne(
				context.TODO(),
				bson.D{{Key: "@id", Value: "6b3be7fe-94a5-4143-a923-70b2a44720cc"}},
				opts,
			).Decode(&result)

			if errors.Is(err, mongo.ErrNoDocuments) {
				fmt.Println("Document not found")
				Expect(0).Should(Equal(0))
			} else {
				Expect(err).ShouldNot(HaveOccurred())
			}
		})
	})

	Context("Тест 5. Проверка обновления типа VerifiedTheHiveAlert", func() {
		verifiedObject := datamodels.VerifiedTheHiveAlert{
			ID:               "6b3be7fe-94a5-2133-a923-70b2a445f0ab",
			CreateTimestatmp: "2024-01-31T16:17:22+03:00",
			Source:           "GCM",
			Event: datamodels.EventMessageTheHiveAlert{
				Base:           false,
				StartDate:      "2024-02-06T15:20:30+03:00",
				RootId:         "~84625227848",
				ObjectId:       "~4192",
				ObjectType:     "alert",
				Organisation:   "GCM",
				OrganisationId: "~4192",
				Operation:      "new",
				RequestId:      "55994d44f3b276c1:6483e5ec:18d786c2f83:-8000:138497",
				Details: datamodels.EventAlertDetails{
					SourceRef:   "TSK-8MSK-6-ZPM-240206-1215999",
					Title:       "222",
					Description: "111",
					Tags: []string{
						"ATs:reason=\"INFO Controlled FGS\"",
						"Sensor:id=\"8030066\"",
					},
				},
				Object: datamodels.EventAlertObject{
					Severity:        1,
					Tlp:             1,
					Pap:             1,
					UnderliningId:   "~85455464790",
					Id:              "~85455464790",
					CreatedBy:       "ddddd",
					CreatedAt:       "1970-01-01T03:00:00+03:00",
					UpdatedAt:       "1970-01-01T03:00:00+03:00",
					UnderliningType: "aalllert",
					Title:           "vbb er3",
					Description:     "any more",
					Status:          "None",
					Date:            "2024-02-06T15:37:52+03:00",
					Type:            "snort_alert",
					ObjectType:      "",
					Source:          "zsiеmSystems",
					SourceRef:       "TSK-8MSK-6-ZPM-240206-1215999",
					Case:            "alert",
					CaseTemplate:    "alert_snort",
					Tags: []string{
						"'Sensor:id=\"8030012\"'",
						"'Webhook:send=ES'",
					},
					CustomFields: datamodels.CustomFields{
						CustomFields: map[string]datamodels.CustomerFields{
							"first-time": &datamodels.CustomFieldDateType{
								Order: 0,
								Date:  "2024-02-06T15:20:30+03:00",
							},
							"last-time": &datamodels.CustomFieldDateType{
								Order: 0,
								Date:  "2024-02-06T15:20:30+03:00",
							},
						},
					},
					/*map[string]datamodels.CustomerFields{
						"first-time": &datamodels.CustomFieldDateType{
							Order: 0,
							Date:  "2024-02-06T15:20:30+03:00",
						},
						"last-time": &datamodels.CustomFieldDateType{
							Order: 0,
							Date:  "2024-02-06T15:20:30+03:00",
						},
					},*/
				},
			},
			Alert: datamodels.AlertMessageTheHiveAlert{
				Follow:    false,
				Tlp:       2,
				Severity:  2,
				Date:      "1970-01-01T03:00:00+03:00",
				CreatedAt: "2024-02-07T11:11:11+03:00",
				// UpdatedAt: ,
				UpdatedBy:       "webhook@cloud.gcm",
				UnderliningId:   "~88026357960",
				Status:          "New",
				Type:            "snort",
				UnderliningType: "__Snort",
				Description:     "free alerts",
				CaseTemplate:    "sonr",
				SourceRef:       "TSK-8MSK-6-ZPM-240206-1216137",
				Tags: []string{
					"Sensor:id=\"8030105\"",
					"ATs:reason=\"Редко встречающиеся признаки ВПО\"",
					"'Webhook:send=ES'",
				},
				CustomFields: datamodels.CustomFields{
					CustomFields: map[string]datamodels.CustomerFields{
						"first-time": &datamodels.CustomFieldDateType{
							Order: 0,
							Date:  "2024-01-01T05:22:30+03:00",
						},
						"last-time": &datamodels.CustomFieldDateType{
							Order: 0,
							Date:  "2024-01-17T00:18:13+03:00",
						},
					},
				},
				/*map[string]datamodels.CustomerFields{
					"first-time": &datamodels.CustomFieldDateType{
						Order: 0,
						Date:  "2024-01-01T05:22:30+03:00",
					},
					"last-time": &datamodels.CustomFieldDateType{
						Order: 0,
						Date:  "2024-01-17T00:18:13+03:00",
					},
				},*/
				Artifacts: []datamodels.AlertArtifact{
					{
						Ioc:              false,
						Sighted:          false,
						IgnoreSimilarity: true,
						Tlp:              1,
						UnderliningId:    "~84302220012",
						Id:               "~84302220012",
						CreatedAt:        "2024-01-26T13:02:01+03:00",
						//UpdatedAt: ,
						StartDate: "2024-01-26T13:02:01+03:00",
						CreatedBy: "friman@email.net",
						UpdatedBy: "friman@email.net",
						Data:      "63.5656 89.12",
						DataType:  "coordinates",
						Message:   "Any message",
						Tags: []string{
							"Sensor:id=\"1111111\"",
							"geoip:country=CH",
							"'Webhook:send=ES'",
						},
					},
					{
						Ioc:              true,
						Sighted:          false,
						IgnoreSimilarity: true,
						Tlp:              2,
						UnderliningId:    "~306522241",
						Id:               "~306522241",
						CreatedAt:        "2024-01-16T03:32:01+03:00",
						//UpdatedAt: ,
						StartDate: "2024-01-04T19:32:01+03:00",
						CreatedBy: "example@email.net",
						UpdatedBy: "example@email.net",
						Data:      "5.63.123.99",
						DataType:  "ipaddr",
						Message:   "ffdffd fdg",
						Tags: []string{
							"Sensor:id=\"3411\"",
							"geoip:country=RU",
							"'Webhook:send=ES'",
						},
					},
				},
			},
		}

		It("Должно быть выполнено успешное добавление или обновление", func() {
			var err error

			collection := conn.Database(NameDB).Collection("alert_collection")
			opts := options.FindOne()
			//currentData := datamodels.VerifiedTheHiveAlert{}
			currentData := *datamodels.NewVerifiedTheHiveAlert()
			err = collection.FindOne(
				context.TODO(),
				bson.D{
					{Key: "source", Value: verifiedObject.GetSource()},
					{Key: "event.rootId", Value: verifiedObject.GetEvent().GetRootId()},
				},
				opts,
			).Decode(&currentData)

			fmt.Println("_________________ Result Test 5. __________________")
			fmt.Println(currentData.ToStringBeautiful(0))

			if errors.Is(err, mongo.ErrNoDocuments) {
				//если похожего документа нет в БД
				currentData = verifiedObject
			} else {
				//если похожий документ есть, выполняем замену старых значений новыми
				_, err = currentData.GetEvent().ReplacingOldValues(*verifiedObject.GetEvent())
				_, err = currentData.GetAlert().ReplacingOldValues(*verifiedObject.GetAlert())

			}

			_, err = collection.InsertMany(context.TODO(), []interface{}{currentData}) /*[]mongo.IndexModel{
				{
					Keys: bson.D{
						{Key: "@id", Value: 1},
					},
					Options: &options.IndexOptions{},
				}, {
					Keys: bson.D{
						{Key: "source", Value: 1},
						{Key: "event.rootId", Value: 1},
					},
					Options: &options.IndexOptions{},
				},
			}*/

			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
