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
				Expect(0).Should(Equal(0))
			} else {
				Expect(err).ShouldNot(HaveOccurred())
			}
		})
	})

	Context("Тест 5. Проверка обновления типа VerifiedTheHiveAlert", func() {
		verifiedObject := datamodels.VerifiedTheHiveAlert{
			ID:              "6b3be7fe-94a5-2133-a923-70b2a445f0ab",
			CreateTimestamp: "2024-01-31T16:17:22+03:00",
			Source:          "GCM",
			Event: datamodels.EventMessageTheHiveAlert{
				Base:           true,
				StartDate:      "2024-02-06T15:20:30+03:00",
				RootId:         "~84625227848",
				ObjectId:       "~4192",
				ObjectType:     "alert",
				Organisation:   "GCM",
				OrganisationId: "~4192",
				Operation:      "UPDATE___UPDATE",
				RequestId:      "55994d44f3b276c1:6483e5ec:18d786c2f83:-8000:138497",
				Details: datamodels.EventAlertDetails{
					SourceRef:   "TSK-8MSK-6-ZPM-240206-1215999",
					Title:       "title test replacing",
					Description: "description test replacing",
					Tags: []string{
						"ATs:reason=\"INFO Controlled FGS\"",
						"Sensor:id=\"8030066\"",
						"geoip:new=VB",
						"geoip:country=RU",
						"ATs:reason=\"Full stack\"",
					},
				},
				Object: datamodels.EventAlertObject{
					Severity:        2,
					Tlp:             3,
					Pap:             4,
					UnderliningId:   "~85455464790",
					Id:              "~85455464790",
					CreatedBy:       "test_replacing@example.net",
					CreatedAt:       "1970-01-01T03:00:00+03:00",
					UpdatedAt:       "1970-01-01T03:00:00+03:00",
					UnderliningType: "ALERT test replacing",
					Title:           "title test replacing",
					Description:     "test replacing description",
					Status:          "Update",
					Date:            "2024-02-06T15:37:52+03:00",
					Type:            "snort_alert",
					ObjectType:      "SNORT test replacing 111111",
					Source:          "zsiеmSystems",
					SourceRef:       "TSK-8MSK-6-ZPM-240206-1215999",
					Case:            "alert",
					CaseTemplate:    "alert_snort",
					Tags: []string{
						"'Sensor:id=\"8030012\"'",
						"'Webhook:send=ES'",
						"Ast:FFF",
						"AST:BBIN",
						"ATs:report=LLOmmn",
					},
					CustomFields: datamodels.CustomFields{
						"first-time": &datamodels.CustomFieldDateType{
							Order: 0,
							Date:  "2024-02-06T15:20:30+03:00",
						},
						"last-time": &datamodels.CustomFieldDateType{
							Order: 1,
							Date:  "2024-01-11T11:11:31+03:00",
						},
					},
				},
			},
			Alert: datamodels.AlertMessageTheHiveAlert{
				Follow:          false,
				Tlp:             2,
				Severity:        2,
				Date:            "1970-01-01T03:00:00+03:00",
				CreatedAt:       "2024-02-22T22:22:22+03:00",
				UpdatedAt:       "2024-02-23T22:23:24+03:00",
				UpdatedBy:       "webhook@cloud.gcm",
				UnderliningId:   "~88026357960",
				Status:          "NewUpdate",
				Type:            "SNORT",
				UnderliningType: "ALERT_Snort",
				Description:     "free alerts test replacing",
				CaseTemplate:    "sNOrt",
				SourceRef:       "TSK-8MSK-6-ZPM-240206-1216137",
				Tags: []string{
					"Sensor:id=\"8030105\"",
					"ATs:reason=\"Редко встречающиеся признаки ВПО\"",
					"'Webhook:send=ES'",
					"Ats:test=replacing",
				},
				CustomFields: datamodels.CustomFields{
					"first-time": &datamodels.CustomFieldDateType{
						Order: 0,
						Date:  "2024-01-03T03:33:33+03:00",
					},
					"last-time": &datamodels.CustomFieldDateType{
						Order: 0,
						Date:  "2024-01-17T00:18:13+03:00",
					},
				},
				Artifacts: []datamodels.AlertArtifact{
					{
						Ioc:              true,
						Sighted:          false,
						IgnoreSimilarity: true,
						Tlp:              3,
						UnderliningId:    "~84302220012",
						Id:               "~84302220012",
						CreatedAt:        "2024-01-26T13:02:01+03:00",
						//UpdatedAt: ,
						StartDate: "2024-01-26T13:02:01+03:00",
						CreatedBy: "friman@email.net",
						UpdatedBy: "friman@email.net",
						Data:      "63.5656 89.12",
						DataType:  "coordinates test replacing",
						Message:   "Any message test replacing",
						Tags: []string{
							"Sensor:id=\"1111111\"",
							"geoip:country=CH",
							"'Webhook:send=ES'",
							"'Webhook:SEND=RT'",
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
						Data:      "115.115.115.115",
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
			currentData := datamodels.NewVerifiedTheHiveAlert()
			err = collection.FindOne(
				context.TODO(),
				bson.D{
					{Key: "source", Value: verifiedObject.GetSource()},
					{Key: "event.rootId", Value: verifiedObject.GetEvent().GetRootId()},
				},
				opts,
			).Decode(currentData)

			fmt.Println("ERROR decode:", err)

			//fmt.Println("_________________ Result Test 5. __________________")
			//fmt.Println(currentData.ToStringBeautiful(0))

			if errors.Is(err, mongo.ErrNoDocuments) {
				//если похожего документа нет в БД
				currentData = &verifiedObject
			} else {
				//если похожий документ есть, выполняем замену старых значений новыми
				_, err = currentData.GetEvent().ReplacingOldValues(*verifiedObject.GetEvent())
				_, err = currentData.GetAlert().ReplacingOldValues(*verifiedObject.GetAlert())

				_, err = collection.DeleteMany(
					context.TODO(),
					bson.D{{
						Key:   "@id",
						Value: bson.D{{Key: "$in", Value: []string{currentData.GetID()}}}}},
					options.Delete())
			}

			_, err = collection.InsertMany(context.TODO(), []interface{}{currentData})

			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	/*Context("Тест 6. Проверка обновления типа VerifiedTheHiveAlert", func() {
		verifiedObject := datamodels.VerifiedTheHiveCase{
			ID:              "6b3be7fe-94a5-2133-a923-70b2a445f0ab",
			CreateTimestamp: "2024-01-31T16:17:22+03:00",
			Source:          "GCM",
			Event: datamodels.EventMessageTheHiveCase{
				Base: false,
				//StartDate      string           `json:"startDate" bson:"startDate"` //в формате RFC3339
				RootId:         "~2778345496",
				Organisation:   "RCM_DFO",
				OrganisationId: "~4136",
				ObjectId:       "~2778345496",
				ObjectType:     "case",
				Operation:      "New",
				RequestId:      "219560ec6156d5fb:7296c75a:18cf0a8f2c3:-8000:459828",
				Details: datamodels.EventCaseDetails{
					EndDate:          "2024-01-31T16:17:15+03:00",
					ResolutionStatus: "TruePositive",
					Summary:          "Сведения о компьютерной атаке доведены до ответственных лиц.",
					Status:           "Resolved",
					ImpactStatus:     "NoImpact",
					CustomFields: datamodels.CustomFields{
						"first-time": &datamodels.CustomFieldDateType{
							Order: 0,
							Date:  "2024-01-12T12:41:05+03:00",
						},
						"last-time": &datamodels.CustomFieldDateType{
							Order: 10,
							Date:  "2024-01-12T12:41:19+03:00",
						},
						"geoip": &datamodels.CustomFieldStringType{
							Order:  0,
							String: "Нидерланды",
						},
						"id-soa": &datamodels.CustomFieldStringType{
							Order:  2,
							String: "220041",
						},
					},
				},
				Object: datamodels.EventCaseObject{
					Flag:      false,
					CaseId:    15795,
					Severity:  3,
					Tlp:       1,
					Pap:       2,
					StartDate: "2024-01-31T16:16:42+03:00",
					//EndDate          string       `json:"endDate" bson:"endDate"`     //формат RFC3339
					CreatedAt:        "2024-01-31T16:16:42+03:00",
					UpdatedAt:        "2024-01-31T16:17:15+03:00",
					UnderliningId:    "~2778345496",
					Id:               "~2778345496",
					CreatedBy:        "zsiem@rcm.lcl",
					UpdatedBy:        "zsiem@rcm.lcl",
					UnderliningType:  "case",
					Title:            "Атаки типа Local File Inclusion",
					Description:      "системы: Заслон-Пост",
					ImpactStatus:     "NoImpact",
					ResolutionStatus: "TruePositive",
					Status:           "Resolved",
					Summary:          "Сведения о компьютерной атаке доведены до ответственных лиц.",
					Owner:            "zsiem@rcm.lcl",
					Tags: []string{
						"NCIRCC:attack=\"Попытки эксплуатации уязвимости\"",
						"Sensor:name=\"Правительство Республики Саха (Якутия)\"",
						"Sensor:id=\"310078\"",
					},
					CustomFields: datamodels.CustomFields{
						"geoip": &datamodels.CustomFieldStringType{
							Order:  0,
							String: "RU",
						},
						"id-soa": &datamodels.CustomFieldStringType{
							Order:  2,
							String: "33374",
						},
					},
				},
			},
		}
		verifiedObject.Observables = []datamodels.ObservableMessage{
			{
				Ioc:                  false,
				Sighted:              true,
				IgnoreSimilarity:     false,
				Tlp:                  2,
				UnderliningCreatedAt: "2023-04-16T17:03:39+03:00",
				UnderliningUpdatedAt: "2023-04-16T17:03:39+03:00",
				StartDate:            "2023-04-16T17:03:39+03:00",
				UnderliningCreatedBy: "a.feoktistov@cloud.rcm",
				//UnderliningUpdatedBy: "",
				UnderliningId:   "~89444416",
				UnderliningType: "Observable",
				Data:            "690006:192.168.128.46",
				DataType:        "ip_home",
				Message:         "Download a piece of traffic",
				Tags: []string{
					"snort",
				},
				Attachment: datamodels.AttachmentData{
					Size: 59963,
					Id:   "geee2333",
				},
				Reports: map[string]datamodels.ReportTaxonomies{
					"one": datamodels.ReportTaxonomies{
						Taxonomies: []datamodels.Taxonomy{
							{
								Level:     "red",
								Namespace: "all host",
								Value:     "anymore",
							},
						},
					},
				},
			},
		}
		verifiedObject.Ttp = []datamodels.TtpMessage{
			{
				OccurDate:            "2023-04-20T18:22:00+03:00",
				UnderliningCreatedAt: "2023-07-11T17:02:00+03:00",
				UnderliningId:        "~212127944",
				UnderliningCreatedBy: "a.feoktistov@cloud.rcm",
				PatternId:            "T1190",
				Tactic:               "initial-access",
				ExtraData: datamodels.ExtraDataTtpMessage{
					Pattern: datamodels.PatternExtraData{
						RemoteSupport:        false,
						Revoked:              true,
						UnderliningCreatedAt: "2023-03-27T12:46:20+03:00",
						UnderliningCreatedBy: "admin@thehive.local",
						UnderliningId:        "~118880",
						UnderliningType:      "Pattern",
						Description:          "Adversaries may attempt to exploit",
						Detection:            "logs for abnormal",
						Name:                 "Exploit Public-Facing Application",
						PatternId:            "T1190",
						PatternType:          "attack-pattern",
						URL:                  "https://attack.mitre.org/techniques/T1190",
						Version:              "2.4",
						Platforms: []string{
							"Windows",
							"Linux",
						},
						PermissionsRequired: []string{},
						DataSources: []string{
							"Network Traffic: Network Traffic Content",
						},
						Tactics: []string{
							"initial-access",
						},
					},
				},
			},
		}

		It("Должно быть выполнено успешное добавление или обновление", func() {
			var err error

			collection := conn.Database(NameDB).Collection("case_test_collection")
			opts := options.FindOne()
			currentData := datamodels.NewVerifiedTheHiveCase()
			err = collection.FindOne(
				context.TODO(),
				bson.D{
					{Key: "source", Value: verifiedObject.GetSource()},
					{Key: "event.rootId", Value: verifiedObject.GetEvent().GetRootId()},
				},
				opts,
			).Decode(currentData)

			fmt.Println("ERROR decode:", err)

			fmt.Println("_________________ Result Test 6. __________________")
			fmt.Println(currentData.ToStringBeautiful(0))

			if errors.Is(err, mongo.ErrNoDocuments) {
				//если похожего документа нет в БД
				currentData = &verifiedObject
			} else {
				//***************************************
				//*
				//* Для того чтобы продолжить тестирование
				//* необходимо сделать метод ReplacingOldValues
				//* для datamodels.EventMessageTheHiveCase,
				//* datamodels.ObservableMessage и datamodels.TtpMessage
				//*
				//***************************************
				//если похожий документ есть, выполняем замену старых значений новыми
				_, err = currentData.GetEvent().ReplacingOldValues(*verifiedObject.GetEvent())
				_, err = currentData.GetAlert().ReplacingOldValues(*verifiedObject.GetAlert())

				_, err = collection.DeleteMany(
					context.TODO(),
					bson.D{{
						Key:   "@id",
						Value: bson.D{{Key: "$in", Value: []string{currentData.GetID()}}}}},
					options.Delete())
			}

			_, err = collection.InsertMany(context.TODO(), []interface{}{currentData})

			Expect(err).ShouldNot(HaveOccurred())
		})
	})*/
})
