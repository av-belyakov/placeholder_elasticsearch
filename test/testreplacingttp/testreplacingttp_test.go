package testreplacingttp_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/datamodels"
)

var _ = Describe("Testreplacingttp", func() {
	Context("Тест 1. Проверка замены старых значений TtpMessageEs объекта, новыми значениями, если они отличаются", func() {
		oldStruct := datamodels.TtpsMessageEs{}

		newStruct := datamodels.TtpsMessageEs{}

		It("Ряд полей в TtpMessageEs должны быть успешно заменены", func() {
			num := oldStruct.ReplacingOldValues(newStruct)

			//кол-во замененных полей
			Expect(num).Should(Equal(16))

			fmt.Println("---=== VERIFED TtpMessageEs ===---")
			fmt.Println(oldStruct.ToStringBeautiful(0))

			/*
				//добавляется
				_, ok := oldStruct.Observables["url_pcap"]
				Expect(ok).Should(BeTrue())

				_, ok = oldStruct.Observables["snort_sid"]
				Expect(ok).Should(BeTrue())

				domains, ok := oldStruct.Observables["domain"]
				Expect(ok).Should(BeTrue())
				Expect(len(domains)).Should(Equal(2))

				ipHome, ok := oldStruct.Observables["ip_home"]
				Expect(ok).Should(BeTrue())
				Expect(ipHome[0].GetAttachment().GetSize()).Should(Equal(uint64(343000648)))
				Expect(len(ipHome[0].GetAttachment().GetHashes())).Should(Equal(4))
			*/

			Expect(true).Should(BeTrue())
		})
	})
})
