package testmappingobjectarea_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/confighandler"
	"placeholder_elasticsearch/eventenrichmentmodule"
)

var _ = Describe("Testmappingobjectarea", Ordered, func() {
	var (
		confApp confighandler.ConfigApp
		err     error
	)

	BeforeAll(func() {
		confApp, err = confighandler.NewConfig("placeholder_elasticsearch")
	})

	Context("Тест 1. Чтение конфигурационного файла", func() {
		It("При чтении конфигурационного файла не дложно быть ошибок", func() {
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Сопоставление и замена значений наименования сферы деятельности", func() {
		It("Сопоставление и замена значений должна быть успешно выполнена", func() {
			Expect("Оборонная промышленность").Should(Equal(eventenrichmentmodule.MappingObjectArea("Оборонная", confApp.GetMappingParameters().AreaActivity)))
			Expect("Государственная/муниципальная власть").Should(Equal(eventenrichmentmodule.MappingObjectArea("Государственный орган", confApp.GetMappingParameters().AreaActivity)))
			Expect("Иная").Should(Equal(eventenrichmentmodule.MappingObjectArea("Другие отрасли", confApp.GetMappingParameters().AreaActivity)))
		})

		It("Сопоставление и замена значений не выполняется", func() {
			Expect("Транспорт").Should(Equal(eventenrichmentmodule.MappingObjectArea("Транспорт", confApp.GetMappingParameters().AreaActivity)))
			Expect("Связьи").Should(Equal(eventenrichmentmodule.MappingObjectArea("Связьи", confApp.GetMappingParameters().AreaActivity)))
			Expect("Атомная пром.").Should(Equal(eventenrichmentmodule.MappingObjectArea("Атомная пром.", confApp.GetMappingParameters().AreaActivity)))
		})
	})
})
