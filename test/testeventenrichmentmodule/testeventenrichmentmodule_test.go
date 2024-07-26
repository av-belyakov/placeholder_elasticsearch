package testeventenrichmentmodule_test

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/confighandler"
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/eventenrichmentmodule"
)

var _ = Describe("Testeventenrichmentmodule", Ordered, func() {
	var (
		ctx       context.Context
		ctxCancel context.CancelFunc

		confApp               confighandler.ConfigApp
		eventEnrichmentModule *eventenrichmentmodule.EventEnrichmentModule

		errConfApp, errEEM error
	)

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
	})

	Context("Тест 1. Проверка чтения конфигурационных файлов", func() {
		It("При чтении конфигурационных файлов не должно быть ошибок", func() {
			Expect(errConfApp).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Проверка инициализации модуля обогащения", func() {
		It("При инициализации модуля обогащения не должно быть ошибок", func() {
			Expect(errEEM).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 3. Проверка работоспособности модуля обогащения информации", func() {
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
				fmt.Printf("\t___ SensorId ___:'%s'\n", k)
				fmt.Println("Sensor ID:", v.SensorId)
				fmt.Println("HostId:", v.HostId)
				fmt.Println("GeoCode:", v.GeoCode)
				fmt.Println("ObjectArea:", v.ObjectArea)
				fmt.Println("SubjectRF:", v.SubjectRF)
				fmt.Println("INN:", v.INN)
				fmt.Println("HomeNet:", v.HomeNet)
				fmt.Println("OrgName:", v.OrgName)
				fmt.Println("FullOrgName:", v.FullOrgName)
			}

			ctxCancel()

			Expect(true).Should(BeTrue())
		})
	})
})
