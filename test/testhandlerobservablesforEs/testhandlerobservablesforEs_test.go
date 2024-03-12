package testhandlerobservablesforEs_test

import (
	"fmt"
	"reflect"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/listhandlerforesjson"
	testing "placeholder_elasticsearch/test"
)

var _ = Describe("TesthandlerobservablesforEs", Ordered, func() {
	var (
		so                     *listhandlerforesjson.SupportiveObservables
		listHandlerObservables map[string][]func(interface{})
	)

	BeforeAll(func() {
		so = listhandlerforesjson.NewSupportiveObservables()
		listHandlerObservables = listhandlerforesjson.NewListHandlerObservablesElement(so)
	})

	Context("Тест 1. Проверка заполнения ES объекта для хранения observables", func() {
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

			for _, v := range testing.GetObservableThree() {
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

			observablesMessageEs := *datamodels.NewObservablesMessageEs()

			fmt.Println("-----------------------------------------------")
			for k, v := range so.GetObservables() {
				fmt.Println(k)
				for key, value := range v {
					fmt.Printf("%d.\n%s\n", key, value.ToStringBeautiful(2))
				}
			}
			fmt.Println("===============================================")

			observablesMessageEs.SetValueObservables(so.GetObservables())
			obser := observablesMessageEs.GetObservables()

			fmt.Println("-------- GetObservables() --------")
			fmt.Println(observablesMessageEs.ToStringBeautiful(0))

			Expect(len(obser)).Should(Equal(2))

			listSnortSid, ok := observablesMessageEs.GetKeyObservables("snort_sid")
			Expect(ok).Should(BeTrue())
			Expect(len(listSnortSid)).Should(Equal(2))

			Expect(listSnortSid[0].GetUnderliningId()).Should(Equal("~3460985064"))
			Expect(listSnortSid[0].GetData()).Should(Equal("9608643"))
			Expect(listSnortSid[0].GetIoc()).Should(BeTrue())
			Expect(listSnortSid[0].GetUnderliningType()).Should(Equal("Observable"))
			Expect(listSnortSid[0].GetTlp()).Should(Equal(uint64(2)))
			Expect(listSnortSid[0].GetUnderliningCreatedBy()).Should(Equal("uds@crimea-rcm"))
			Expect(len(listSnortSid[0].GetTagsAll())).Should(Equal(2))
			Expect(len(listSnortSid[0].GetSnortSid())).Should(Equal(1))
			Expect(listSnortSid[0].GetAttachment().GetContentType()).Should(Equal("text/plain"))
			Expect(listSnortSid[0].GetAttachment().GetSize()).Should(Equal(uint64(817)))
			Expect(len(listSnortSid[0].GetAttachment().GetHashes())).Should(Equal(3))

			cct, ok := listSnortSid[0].GetTaxonomies("CyberCrime-Tracker_1_0")
			Expect(ok).Should(BeTrue())
			Expect(len(cct.Taxonomies)).Should(Equal(2))
			Expect(cct.Taxonomies[0].GetLevel()).Should(Equal("info"))
			Expect(cct.Taxonomies[0].GetNamespace()).Should(Equal("CCT"))
			Expect(cct.Taxonomies[0].GetPredicate()).Should(Equal("C2 Search"))
			Expect(cct.Taxonomies[0].GetValue()).Should(Equal("0 hits"))
		})
	})
})
