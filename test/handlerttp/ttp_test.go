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

			// ---- ttpTwo ----
			for _, v := range testing.GetTtpTwo() {
				if lf, ok := listHandlerTtp[v.ElemName]; ok {
					for _, f := range lf {
						r := reflect.TypeOf(v.ElemValue)
						switch r.Kind() {
						case reflect.Slice:
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

			ttps := sttp.GetTtps()

			fmt.Println("-------- sttp.GetTtps() --------")
			for k, v := range ttps {
				fmt.Printf("%d. \n, %v\n", k, v.ToStringBeautiful(0))
			}

			Expect(len(ttps)).Should(Equal(2))

			//Тест ПЕРВОГО элемента из объекта ttp
			Expect(ttps[0].GetUnderliningCreatedAt()).Should(Equal("2024-01-12T09:37:09+03:00"))
			Expect(ttps[0].GetUnderliningCreatedBy()).Should(Equal("prs@rcm"))
			Expect(ttps[0].GetUnderliningId()).Should(Equal("~185286688"))
			Expect(ttps[0].GetOccurDate()).Should(Equal("2024-01-12T09:37:00+03:00"))
			Expect(ttps[0].GetPatternId()).Should(Equal("T1110.001"))
			Expect(ttps[0].GetTactic()).Should(Equal("credential-access"))

			Expect(ttps[0].GetPattern().GetRemoteSupport()).Should(BeTrue())
			Expect(ttps[0].GetPattern().GetRevoked()).Should(BeTrue())
			Expect(ttps[0].GetPattern().GetUnderliningCreatedAt()).Should(Equal("2024-01-12T09:37:09+03:00"))
			Expect(ttps[0].GetPattern().GetUnderliningCreatedBy()).Should(Equal("admin@thehive.local"))
			Expect(ttps[0].GetPattern().GetUnderliningId()).Should(Equal("~164016"))
			Expect(ttps[0].GetPattern().GetUnderliningType()).Should(Equal("Pattern"))
			Expect(ttps[0].GetPattern().GetDescription()).Should(Equal("exploit edge network"))
			Expect(ttps[0].GetPattern().GetDetection()).Should(Equal("Monitor authentication logs for system"))
			Expect(ttps[0].GetPattern().GetName()).Should(Equal("Password Guessing"))
			Expect(ttps[0].GetPattern().GetPatternId()).Should(Equal("T1110.001"))
			Expect(ttps[0].GetPattern().GetPatternType()).Should(Equal("attack-pattern"))
			Expect(ttps[0].GetPattern().GetURL()).Should(Equal("https://attack.mitre.org/techniques/T1110/001"))
			Expect(ttps[0].GetPattern().GetVersion()).Should(Equal("1.4"))
			Expect(len(ttps[0].GetPattern().GetPlatforms())).Should(Equal(3))
			Expect(len(ttps[0].GetPattern().GetPermissionsRequired())).Should(Equal(2))
			Expect(len(ttps[0].GetPattern().GetDataSources())).Should(Equal(2))
			Expect(len(ttps[0].GetPattern().GetTactics())).Should(Equal(1))

			Expect(ttps[0].GetPatternParent().GetRemoteSupport()).Should(BeTrue())
			Expect(ttps[0].GetPatternParent().GetRevoked()).Should(BeTrue())
			Expect(ttps[0].GetPatternParent().GetUnderliningCreatedAt()).Should(Equal("2024-01-12T09:21:21+03:00"))
			Expect(ttps[0].GetPatternParent().GetUnderliningCreatedBy()).Should(Equal("root@thehive.gcm"))
			Expect(ttps[0].GetPatternParent().GetUnderliningId()).Should(Equal("~236516"))
			Expect(ttps[0].GetPatternParent().GetUnderliningType()).Should(Equal("Pattern Parent"))
			Expect(ttps[0].GetPatternParent().GetDescription()).Should(Equal("Adversaries may"))
			Expect(ttps[0].GetPatternParent().GetDetection()).Should(Equal("Adversaries with no prior knowledge"))
			Expect(ttps[0].GetPatternParent().GetName()).Should(Equal("Password Guessing"))
			Expect(ttps[0].GetPatternParent().GetPatternId()).Should(Equal("T1110"))
			Expect(ttps[0].GetPatternParent().GetPatternType()).Should(Equal("attack-pattern"))
			Expect(ttps[0].GetPatternParent().GetURL()).Should(Equal("https://attack.mitre.org/techniques/T1110"))
			Expect(ttps[0].GetPatternParent().GetVersion()).Should(Equal("1.14"))
			Expect(len(ttps[0].GetPatternParent().GetPlatforms())).Should(Equal(6))
			Expect(len(ttps[0].GetPatternParent().GetPermissionsRequired())).Should(Equal(3))
			Expect(len(ttps[0].GetPatternParent().GetDataSources())).Should(Equal(3))
			Expect(len(ttps[0].GetPatternParent().GetTactics())).Should(Equal(1))

			// Тест ВТОРОГО элемента из объекта ttp
			Expect(ttps[1].GetUnderliningCreatedAt()).Should(Equal("2024-01-12T07:13:49+03:00"))
			Expect(ttps[1].GetUnderliningCreatedBy()).Should(Equal("dv-cloud@gcm"))
			Expect(ttps[1].GetUnderliningId()).Should(Equal("~104177744"))
			Expect(ttps[1].GetOccurDate()).Should(Equal("2024-01-12T07:13:00+03:00"))
			Expect(ttps[1].GetPatternId()).Should(Equal("T1190"))
			Expect(ttps[1].GetTactic()).Should(Equal("initial-access"))

			Expect(ttps[1].GetPattern().GetRemoteSupport()).Should(BeTrue())
			Expect(ttps[1].GetPattern().GetRevoked()).ShouldNot(BeTrue())
			Expect(ttps[1].GetPattern().GetUnderliningCreatedAt()).Should(Equal("2024-01-12T07:13:49+03:00"))
			Expect(ttps[1].GetPattern().GetUnderliningCreatedBy()).Should(Equal("prs@rcm"))
			Expect(ttps[1].GetPattern().GetUnderliningId()).Should(Equal("~104177744"))
			Expect(ttps[1].GetPattern().GetUnderliningType()).Should(Equal("Pattern"))
			Expect(ttps[1].GetPattern().GetDescription()).Should(Equal("It`s very important message"))
			Expect(ttps[1].GetPattern().GetDetection()).Should(Equal("Monitor application logs"))
			Expect(ttps[1].GetPattern().GetName()).Should(Equal("Password Guessing"))
			Expect(ttps[1].GetPattern().GetPatternId()).Should(Equal("T1190"))
			Expect(ttps[1].GetPattern().GetPatternType()).Should(Equal("attack-pattern"))
			Expect(ttps[1].GetPattern().GetURL()).Should(Equal("https://attack.mitre.org/techniques/T1110/001"))
			Expect(ttps[1].GetPattern().GetVersion()).Should(Equal("2.4"))
			Expect(len(ttps[1].GetPattern().GetPlatforms())).Should(Equal(4))
			Expect(len(ttps[1].GetPattern().GetPermissionsRequired())).Should(Equal(4))
			Expect(len(ttps[1].GetPattern().GetDataSources())).Should(Equal(2))
			Expect(len(ttps[1].GetPattern().GetTactics())).Should(Equal(1))

			Expect(ttps[1].GetPatternParent().GetRemoteSupport()).Should(BeTrue())
			Expect(ttps[1].GetPatternParent().GetRevoked()).Should(BeTrue())
			Expect(ttps[1].GetPatternParent().GetUnderliningCreatedAt()).Should(Equal("2024-01-12T09:21:21+03:00"))
			Expect(ttps[1].GetPatternParent().GetUnderliningCreatedBy()).Should(Equal("root@example.gcm"))
			Expect(ttps[1].GetPatternParent().GetUnderliningId()).Should(Equal("~16400016"))
			Expect(ttps[1].GetPatternParent().GetUnderliningType()).Should(Equal("Pattern Parent"))
			Expect(ttps[1].GetPatternParent().GetDescription()).Should(Equal(""))
			Expect(ttps[1].GetPatternParent().GetDetection()).Should(Equal("Adversaries with no prior knowledge"))
			Expect(ttps[1].GetPatternParent().GetName()).Should(Equal("Password Guessing"))
			Expect(ttps[1].GetPatternParent().GetPatternId()).Should(Equal("T1220"))
			Expect(ttps[1].GetPatternParent().GetPatternType()).Should(Equal("attack-pattern"))
			Expect(ttps[1].GetPatternParent().GetURL()).Should(Equal("https://attack.mitre.org/techniques/T1220"))
			Expect(ttps[1].GetPatternParent().GetVersion()).Should(Equal("1.1411"))
			Expect(len(ttps[1].GetPatternParent().GetPlatforms())).Should(Equal(6))
			Expect(len(ttps[1].GetPatternParent().GetPermissionsRequired())).Should(Equal(3))
			Expect(len(ttps[1].GetPatternParent().GetDataSources())).Should(Equal(3))
			Expect(len(ttps[1].GetPatternParent().GetTactics())).Should(Equal(1))

			//Expect(ttps[0]).Should(Equal())
			//Expect(ttps[0]).Should(Equal())
			//Expect(ttps[0]).Should(Equal())
		})
	})
})
