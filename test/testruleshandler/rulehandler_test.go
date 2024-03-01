package testruleshandler_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/confighandler"
	rules "placeholder_elasticsearch/rulesinteraction"
)

var _ = Describe("Rulehandler", Ordered, func() {
	var (
		lr               *rules.ListRule
		conf             confighandler.ConfigApp
		ruleWarnings     []string
		errConf, errRule error
	)

	BeforeAll(func() {
		conf, errConf = confighandler.NewConfig("placeholder_elasticsearch")

		// инициализируем модуль чтения правил обработки MISP сообщений
		lr, ruleWarnings, errRule = rules.NewListRule("placeholder_elasticsearch", conf.AppConfigRulesProcMsg.Directory, conf.AppConfigRulesProcMsg.FileAlert)
	})

	Context("Тест 1. Проверка чтения конфигурационного файла", func() {
		It("При чтении файла ошибки быть не должно", func() {
			Expect(errConf).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Проверка чтения файла с правилами", func() {
		It("При чтении файла с правилами ошибки быть не должно", func() {
			Expect(errRule).ShouldNot(HaveOccurred())
		})

		It("В файле не должно быть каких либо логических ошибок", func() {
			Expect(len(ruleWarnings)).Should(Equal(0))
		})

		It("Должно быть некоторое количество правил типа Pass", func() {

			fmt.Println("List rule type Pass:")
			fmt.Println(lr.GetRulePass())

			Expect(len(lr.GetRulePass())).ShouldNot(Equal(0))
		})

		It("Правило типа Passany должно быть в False", func() {
			Expect(lr.GetRulePassany()).ShouldNot(BeTrue())
		})

		It("Должно быть успешно обработано правило REPLACE", func() {
			fmt.Println("List rule type REPLACE:")
			lrr := lr.GetRuleReplace()
			fmt.Println(lrr)

			for k, v := range lrr {
				fmt.Println("key: ", k, " value: ", v)
			}

			_, _, err := lr.ReplacementRuleHandler("string", "objectType", "dfff")
			Expect(err).ShouldNot(HaveOccurred())

		})
	})
})
