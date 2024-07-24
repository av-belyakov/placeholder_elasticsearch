package testncirccrequest_test

import (
	"context"
	"fmt"
	"placeholder_elasticsearch/eventenrichmentmodule"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testncirccrequest", func() {
	Context("Тест 1. Проверка возможности получения подробной информации по ИНН", func() {
		It("При выполнении запроса не должно быть ошибок", func() {
			var (
				url   string = "https://lk.cert.local/api/v2/companies"
				token string = "fdd2c5e743960ec9ea80d1ff8868cc6d8439b02f4d61075efd69a46eaa52ff0e"
				inn   string = "7722377866"
			)

			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

			settingsFullOrgNameByINN, err := eventenrichmentmodule.NewSettingsFuncFullNameOrganizationByINN(ctx, url, token, (5 * time.Second))
			Expect(err).ShouldNot(HaveOccurred())

			rd, err := settingsFullOrgNameByINN.GetFullNameOrganizationByINN(inn)
			Expect(err).ShouldNot(HaveOccurred())

			fmt.Println("RESULT:")
			for _, v := range rd.Data {
				fmt.Println("Organization name:", v.Name)
				fmt.Println("Full name organization:", v.Sname)
			}

			Expect(rd.Count).ShouldNot(Equal(0))
		})
	})
})
