package testhandlerttp_test

import (
	"fmt"
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/listhandlerthehivejson"
	testing "placeholder_elasticsearch/test"
)

var _ = Describe("Testttp", Ordered, func() {
	var (
		sttp           *listhandlerthehivejson.SupportiveTtp
		listHandlerTtp map[string][]func(interface{})
	)

	BeforeAll(func() {
		sttp = listhandlerthehivejson.NewSupportiveTtp()
		listHandlerTtp = listhandlerthehivejson.NewListHandlerTtpElement(sttp)
	})

	Context("Тест 1. Проверка заполнения  объекта для хранения ttp", func() {
		It("Вспомогательный объект должен быть успешно заполнен", func() {
			// ---- ttpOne ----
			for _, v := range testing.GetTtpOne() {
				if lf, ok := listHandlerTtp[v.ElemName]; ok {
					for _, f := range lf {
						r := reflect.TypeOf(v.ElemValue)
						switch r.Kind() {
						case reflect.Slice:
							fmt.Println("v.ElemName:", v.ElemName, "00000000000000000000 v.ElemValue:", v.ElemValue)

							if s, ok := v.ElemValue.([]interface{}); ok {
								for _, value := range s {
									fmt.Println("fffffffffffffff =", value)
									f(value)
								}
							}
						default:
							f(v.ElemValue)

						}
					}
				}
			}

			// ---- ttpTwo ----
			/*for _, v := range testing.GetTtpTwo() {
				if lf, ok := listHandlerTtp[v.ElemName]; ok {
					for _, f := range lf {
						r := reflect.TypeOf(v.ElemValue)
						switch r.Kind() {
						case reflect.Slice:
							//fmt.Println("v.ElemName:", v.ElemName, "00000000000000000000 v.ElemValue:", v.ElemValue)

							if s, ok := v.ElemValue.([]interface{}); ok {
								for _, value := range s {
									f(value)
								}
							}
						default:
							f(v.ElemValue)

						}
					}
				}
			}*/

			ttps := sttp.GetTtps()

			fmt.Println("-------- sttp.GetTtps() --------")
			for k, v := range ttps {
				fmt.Printf("%d. \n, %v\n", k, v)
			}

			Expect(len(ttps)).Should(Equal(2))
		})
	})
})
