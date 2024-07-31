package testeventenrichmentmodule_test

import (
	"context"
	"errors"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"placeholder_elasticsearch/confighandler"
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/eventenrichmentmodule"
)

var _ = Describe("Testeventenrichmentmodule", Ordered, func() {
	var (
		collectionName string = "collection_sensor_information"

		ctx       context.Context
		ctxCancel context.CancelFunc

		confApp               confighandler.ConfigApp
		eventEnrichmentModule *eventenrichmentmodule.EventEnrichmentModule

		connMDB *mongo.Client

		errConfApp, errEEM, errMDB error
	)

	requestMongoDB := func(sensorId string, sensorSettings datamodels.InformationFromEventEnricher) error {
		collection := connMDB.Database(confApp.GetAppMongoDB().NameDB).Collection(collectionName)
		opts := options.FindOne()
		currentData := datamodels.NewSensorInformation()
		err := collection.FindOne(
			context.TODO(),
			bson.D{
				{Key: "sensorId", Value: sensorId},
			},
			opts,
		).Decode(currentData)
		if !errors.Is(err, mongo.ErrNoDocuments) {
			if _, err := collection.DeleteMany(
				context.TODO(),
				bson.D{{
					Key:   "sensorId",
					Value: bson.D{{Key: "$in", Value: []string{currentData.SensorId}}}}},
				options.Delete()); err != nil {

				return err
			}
		}

		newData := datamodels.NewSensorInformation()
		newData.SetSensorId(sensorId)
		newData.SetHostId(sensorSettings.GetHostId(sensorId))
		newData.SetGeoCode(sensorSettings.GetGeoCode(sensorId))
		newData.SetObjectArea(sensorSettings.GetObjectArea(sensorId))
		newData.SetSubjectRF(sensorSettings.GetSubjectRF(sensorId))
		newData.SetINN(sensorSettings.GetINN(sensorId))
		newData.SetHomeNet(sensorSettings.GetHomeNet(sensorId))
		newData.SetOrgName(sensorSettings.GetOrgName(sensorId))
		newData.SetFullOrgName(sensorSettings.GetFullOrgName(sensorId))

		//если похожего документа нет в БД
		if _, err := collection.InsertMany(context.TODO(), []interface{}{newData}); err != nil {
			return err
		}

		return nil
	}

	BeforeAll(func() {
		logging := make(chan datamodels.MessageLogging)

		go func() {
			for msg := range logging {
				fmt.Println("ERROR:", msg.MsgData)
			}
		}()

		confApp, errConfApp = confighandler.NewConfig("placeholder_elasticsearch")
		commonConf := confApp.GetCommonApp()

		ctx, ctxCancel = context.WithCancel(context.Background())
		eventEnrichmentModule, errEEM = eventenrichmentmodule.NewEventEnrichmentModule(ctx, commonConf.NCIRCC, commonConf.ZabbixJsonRPC, logging)

		mongoDBConf := confApp.GetAppMongoDB()

		clientOption := options.Client().SetAuth(options.Credential{
			AuthMechanism: "SCRAM-SHA-256",
			AuthSource:    mongoDBConf.NameDB,
			Username:      mongoDBConf.User,
			Password:      mongoDBConf.Passwd,
		})

		confPath := fmt.Sprintf("mongodb://%s:%d/%s", mongoDBConf.Host, mongoDBConf.Port, mongoDBConf.NameDB)

		fmt.Println("NewConnection: ", confPath)
		ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
		defer cancel()

		connMDB, errMDB = mongo.Connect(ctx, clientOption.ApplyURI(confPath))
	})

	Context("Тест 1. Проверка чтения конфигурационных файлов", func() {
		It("При чтении конфигурационных файлов не должно быть ошибок", func() {
			Expect(errConfApp).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Проверка подключения к MongoDB", func() {
		It("При подключении к СУБД MongoDB не должно быть ошибок", func() {
			Expect(errMDB).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 3. Проверка инициализации модуля обогащения", func() {
		It("При инициализации модуля обогащения не должно быть ошибок", func() {
			Expect(errEEM).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 4. Проверка работоспособности модуля обогащения информации", func() {
		It("Из модуля обогащения должны быть получены корректные данные", func() {
			rootId := "74374y82"
			source := "anysource"

			eventEnrichmentModule.ChanInputModule <- eventenrichmentmodule.SettingsChanInputEEM{
				RootId:    rootId,
				Source:    source,
				SensorsId: []string{"570084", "690017", "690013"},
			}

			response := <-eventEnrichmentModule.ChanOutputModule
			Expect(response.RootId).Should(Equal(rootId))
			Expect(response.Source).Should(Equal(source))

			for k, v := range response.Sensors {
				fmt.Printf("%d.\n", k)
				fmt.Println("Sensor ID:", v.SensorId)
				fmt.Println("HostId:", v.HostId)
				fmt.Println("GeoCode:", v.GeoCode)
				fmt.Println("ObjectArea:", v.ObjectArea)
				fmt.Println("SubjectRF:", v.SubjectRF)
				fmt.Println("INN:", v.INN)
				fmt.Println("HomeNet:", v.HomeNet)
				fmt.Println("OrgName:", v.OrgName)
				fmt.Println("FullOrgName:", v.FullOrgName)

				err := requestMongoDB(v.SensorId, response)
				Expect(err).ShouldNot(HaveOccurred())
			}

			ctxCancel()

			Expect(true).Should(BeTrue())
		})
	})
})
