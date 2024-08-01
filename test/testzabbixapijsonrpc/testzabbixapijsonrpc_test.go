package testzabbixapijsonrpc_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/zabbixinteractions"
)

var _ = Describe("Testzabbixapijsonrpc", Ordered, func() {
	var (
		settings zabbixinteractions.SettingsZabbixConnectionJsonRPC

		zabbixConnHandler *zabbixinteractions.ZabbixConnectionJsonRPC
		zabbixConnErr     error
	)

	BeforeAll(func() {
		connTimeout := time.Duration(3 * time.Second)

		settings = zabbixinteractions.SettingsZabbixConnectionJsonRPC{
			Host:   "192.168.9.45", //правильный
			Login:  "Cherry",
			Passwd: "v-2ymX!aVg3eS*hC",
			//Host:              "192.168.9.145", //не правильный
			ConnectionTimeout: &connTimeout,
		}

		zabbixConnHandler, zabbixConnErr = zabbixinteractions.NewZabbixConnectionJsonRPC(settings)
	})

	Context("Тест 0. Выполняем создание нового JSON RPC соединения", func() {
		It("При инициализации нового JSON RPS соединения ошибок быть не должно", func() {
			fmt.Println("RESULT:", zabbixConnHandler)

			Expect(zabbixConnErr).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 1. Выполняем POST запрос к Zabbix", func() {
		It("Не должно быть ошибок при запросе", func() {
			//hostName := "690028"
			hostName := "8030010"
			//hostName := "570084"
			//hostName := 530043 (содержит некорректную запись в Zabbix)

			fullInfo, err := zabbixinteractions.GetFullSensorInformationFromZabbixAPI(hostName, zabbixConnHandler)
			Expect(err).ShouldNot(HaveOccurred())

			fmt.Println("_____ SensorId _____:", fullInfo.SensorId)
			fmt.Println("_____ HostId _____:", fullInfo.HostId)
			fmt.Println("_____ GeoCode _____:", fullInfo.GeoCode)
			fmt.Println("_____ ObjectArea _____:", fullInfo.ObjectArea)
			fmt.Println("_____ SubjectRF _____:", fullInfo.SubjectRF)
			fmt.Println("_____ INN _____:", fullInfo.INN)
			fmt.Println("_____ HomeNet _____:", fullInfo.HomeNet)

			Expect(true).Should(BeTrue())
		})
	})
})
