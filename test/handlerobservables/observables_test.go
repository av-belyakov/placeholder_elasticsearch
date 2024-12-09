package testhandlerobservables_test

import (
	"fmt"
	"reflect"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/listhandlerthehivejson"
	testing "placeholder_elasticsearch/test"
)

var _ = Describe("Testobservablesreports", Ordered, func() {
	var (
		so                     *listhandlerthehivejson.SupportiveObservables
		sor                    *listhandlerthehivejson.SupportiveObservablesReports
		listHandlerObservables map[string][]func(interface{})
	)

	BeforeAll(func() {
		so = listhandlerthehivejson.NewSupportiveObservables()
		sor = listhandlerthehivejson.NewSupportiveObservablesReports()
		listHandlerObservables = listhandlerthehivejson.NewListHandlerObservablesElement(so)
	})

	Context("Тест 1. Проверка заполнения вспомогательного объекта для хранения observables.reports", func() {
		It("Вспомогательный объект должен быть успешно заполнен", func() {
			//****************
			//Метод sor.HandlerReportValue используется только исключительно
			//для теста, в продуктовом приложении НЕ ИСПОЛЬЗУЕТСЯ
			//****************
			for _, v := range testing.GetReportsList() {
				sor.HandlerReportValue(v.ElemName, v.ElemValue)
			}

			fmt.Println(sor.GetReports())

			Expect(len(sor.GetReports())).Should(Equal(4))
		})
	})

	Context("Тест 2. Проверка заполнения  объекта для хранения observables", func() {
		It("Вспомогательный объект должен быть успешно заполнен", func() {
			observableExample := []struct {
				ElemName  string
				ElemValue interface{}
			}(nil)

			observableExample = append(observableExample, testing.GetObservableOne()...)
			observableExample = append(observableExample, testing.GetReportsList()...)

			// ------ OBSERVABLES ------
			for _, v := range observableExample {
				if lf, ok := listHandlerObservables[v.ElemName]; ok {
					//fmt.Printf("%d. v.ElemName: %s, v.ElemValue: '%v'\n", k, v.ElemName, v.ElemValue)

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

				if strings.Contains(v.ElemName, "observables.reports.") {
					so.HandlerReportValue(v.ElemName, v.ElemValue)
				}
			}

			for _, v := range testing.GetObservableTwo() {
				if lf, ok := listHandlerObservables[v.ElemName]; ok {
					//fmt.Printf("%d. v.ElemName: %s, v.ElemValue: '%v'\n", k, v.ElemName, v.ElemValue)

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
			}

			observables := so.GetObservables()

			fmt.Println("-------- soTmp.GetObservables() --------")
			for k, v := range observables {
				fmt.Printf("%d. \n, %v\n", k, v)
			}

			Expect(len(observables)).Should(Equal(2))

			//Тест ПЕРВОГО элемента из объекта observables
			Expect(observables[0].GetUnderliningId()).Should(Equal("~3460985064"))
			Expect(observables[0].GetData()).Should(Equal("9608643"))
			Expect(observables[0].GetDataType()).Should(Equal("snort_sid"))
			Expect(observables[0].GetIoc()).Should(BeTrue())
			Expect(observables[0].GetUnderliningType()).Should(Equal("Observable"))
			Expect(observables[0].GetTlp()).Should(Equal(uint64(2)))
			Expect(observables[0].GetUnderliningCreatedAt()).Should(Equal("2023-08-02T12:31:04+03:00"))
			Expect(observables[0].GetUnderliningCreatedBy()).Should(Equal("uds@crimea-rcm"))
			Expect(observables[0].GetUnderliningUpdatedAt()).Should(Equal("2024-01-11T15:45:51+03:00"))
			Expect(observables[0].GetSighted()).ShouldNot(BeTrue())
			Expect(observables[0].GetStartDate()).Should(Equal("2023-08-02T12:31:04+03:00"))
			Expect(len(observables[0].GetTags())).Should(Equal(2))

			Expect(observables[0].GetAttachment().GetContentType()).Should(Equal("text/plain"))
			Expect(observables[0].GetAttachment().GetId()).Should(Equal("c29438b04791184d3eba39bdb7cf99560ab62068fee9509d50cf59723c398ac1"))
			Expect(observables[0].GetAttachment().GetName()).Should(Equal("n[n.txt"))
			Expect(observables[0].GetAttachment().GetSize()).Should(Equal(uint64(817)))
			Expect(len(observables[0].GetAttachment().GetHashes())).Should(Equal(3))
			//Expect(observables.).Should(Equal(""))

			//Тест ПЕРВОГО элемента из объекта observables.reports
			Expect(len(observables[0].GetReports())).Should(Equal(4))
			for k, v := range observables[0].GetReports() {
				if k == "CyberCrime-Tracker_1_0" {
					Expect(len(v.GetReportTaxonomys().Taxonomies)).Should(Equal(2))

					for key, value := range v.GetReportTaxonomys().Taxonomies {
						fmt.Printf("%d.\n", key)
						fmt.Println(" - value.Level: ", value.Level)
						fmt.Println(" - value.Namespace: ", value.Namespace)
						fmt.Println(" - value.Predicate: ", value.Predicate)
						fmt.Println(" - value.Value: ", value.Value)

						if key == 0 {
							Expect(value.Level).Should(Equal("info"))
							Expect(value.Namespace).Should(Equal("CCT"))
							Expect(value.Predicate).Should(Equal("C2 Search"))
							Expect(value.Value).Should(Equal("0 hits"))
						}
					}
				}
			}
			//fmt.Println("Observables:", observables)

			//Тест ВТОРОГО элемента из объекта observables
			Expect(observables[1].GetUnderliningId()).Should(Equal("~542580736"))
			Expect(observables[1].GetData()).Should(Equal("/dbdata/dump/events/58964/B2M-58964.pcap"))
			Expect(observables[1].GetDataType()).Should(Equal("url_pcap"))
			Expect(observables[1].GetIoc()).Should(BeTrue())
			Expect(observables[1].GetUnderliningType()).Should(Equal("Observable"))
			Expect(observables[1].GetTlp()).Should(Equal(uint64(2)))
			Expect(observables[1].GetUnderliningCreatedAt()).Should(Equal("2024-01-12T11:50:49+03:00"))
			Expect(observables[1].GetUnderliningCreatedBy()).Should(Equal("zhmurchuk@mail.rcm"))
			Expect(observables[1].GetUnderliningUpdatedAt()).Should(Equal("2024-01-12T11:50:48+03:00"))
			Expect(observables[1].GetSighted()).ShouldNot(BeTrue())
			Expect(observables[1].GetStartDate()).Should(Equal("2024-01-12T11:50:49+03:00"))
			Expect(len(observables[1].GetTags())).Should(Equal(2))

			for k, v := range observables[1].GetTags() {
				if k == 0 {
					Expect(v).Should(Equal("misp:Network activity=\"ip-src\""))
				}

				if k == 1 {
					Expect(v).Should(Equal("b2m:ip_ext=206.189.15.25"))
				}
			}
		})
	})
})
