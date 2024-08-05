package testsearchsensorinfomongodb_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/confighandler"
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/mongodbinteractions"
)

var _ = Describe("Testsearchsensorinfomongodb", Ordered, func() {
	var (
		done chan struct{}

		confApp   confighandler.ConfigApp
		mdbModule *mongodbinteractions.MongoDBModule

		errApp, errMdb error

		rootId    string   = "~90677472"
		source    string   = "rcmniz"
		sensorsId []string = []string{
			"500065",
			"310052",
			"570062",
			"570097",
			"8030070",
			"8030090",
			"8030013",
		}
	)

	BeforeAll(func() {
		logging := make(chan datamodels.MessageLogging)
		counting := make(chan datamodels.DataCounterSettings)
		done = make(chan struct{})

		go func() {
			for {
				select {
				case <-done:
					return

				case msg := <-logging:
					fmt.Println("ERROR:", msg.MsgData)

				case count := <-counting:
					fmt.Println("Count:", count.DataMsg)

				}
			}
		}()

		confApp, errApp = confighandler.NewConfig("placeholder_elasticsearch")

		// инициализация модуля для взаимодействия с СУБД MongoDB
		mdbModule, errMdb = mongodbinteractions.HandlerMongoDB(*confApp.GetAppMongoDB(), logging, counting)

	})

	Context("Тест 1. Инициализация модулей", func() {
		It("Не должно быть ошибки при инициализации модуля конфигурационного файла", func() {
			Expect(errApp).ShouldNot(HaveOccurred())
		})

		It("Не должно быть ошибки при инициализации модуля взаимодействия с MongoDB", func() {
			Expect(errMdb).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Поиск информации по сенсорам", func() {
		It("Не должно быть ошибки, должно быть успешно найдено определенное количество сенсоров", func() {
			mdbModule.ChanInputModule <- mongodbinteractions.SettingsInputChan{
				Section: "handling eventenrichment",
				Command: "get sensor eventenrichment",
				RootId:  rootId,
				Source:  source,
				Data:    sensorsId,
			}

			data := <-mdbModule.ChanOutputModule

			Expect(data.Section).Should(Equal("handling eventenrichment"))
			Expect(data.Command).Should(Equal("sensor eventenrichment response"))

			info, ok := data.Data.(datamodels.InformationFromEventEnricher)

			Expect(ok).Should(BeTrue())
			Expect(info.GetRootId()).Should(Equal(rootId))
			Expect(info.GetSource()).Should(Equal(source))
			Expect(len(info.GetSensorsId())).Should(Equal(7))

			for k, sensorId := range info.GetSensorsId() {
				fmt.Printf("\t%d. Sensor id: %s\n", k, sensorId)
				fmt.Println("HostId:", info.GetHostId(sensorId))
				fmt.Println("GeoCode:", info.GetGeoCode(sensorId))
				fmt.Println("ObjectArea:", info.GetObjectArea(sensorId))
				fmt.Println("SubjectRF:", info.GetSubjectRF(sensorId))
				fmt.Println("INN:", info.GetINN(sensorId))
				fmt.Println("HomeNet:", info.GetHomeNet(sensorId))
				fmt.Println("OrgName:", info.GetOrgName(sensorId))
				fmt.Println("FullOrgName:", info.GetFullOrgName(sensorId))

			}

			Expect(true).Should(BeTrue())
		})
	})

	AfterAll(func() {
		done <- struct{}{}
	})
})
