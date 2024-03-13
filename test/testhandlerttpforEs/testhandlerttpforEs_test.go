package testhandlerttpforEs_test

import (
	"fmt"
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/listhandlerforesjson"
	testing "placeholder_elasticsearch/test"
)

var _ = Describe("TesthandlerttpforEs", Ordered, func() {
	var (
		sttp           *listhandlerforesjson.SupportiveTtp
		listHandlerTtp map[string][]func(interface{})
	)

	BeforeAll(func() {
		sttp = listhandlerforesjson.NewSupportiveTtp()
		listHandlerTtp = listhandlerforesjson.NewListHandlerTtpElement(sttp)
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

			// ---- ttpThree ----
			for _, v := range testing.GetTtpThree() {
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

			ttpsMessaheEs := *datamodels.NewTtpsMessageEs()
			ttpsMessaheEs.SetValueTtp(sttp.GetTtps())

			fmt.Println("-------- Ttps() --------")
			fmt.Println(ttpsMessaheEs.ToStringBeautiful(0))

			Expect(len(ttpsMessaheEs.GetValueTtp())).Should(Equal(2))

			T1110_001, ok := ttpsMessaheEs.GetKeyTtp("T1110.001")
			Expect(ok).Should(BeTrue())

			//Тест T1110_001 элемента из объекта ttp
			Expect(T1110_001[0].GetUnderliningCreatedAt()).Should(Equal("2024-01-12T09:37:09+03:00"))
			Expect(T1110_001[0].GetUnderliningCreatedBy()).Should(Equal("prs@rcm"))
			Expect(T1110_001[0].GetUnderliningId()).Should(Equal("~185286688"))
			Expect(T1110_001[0].GetOccurDate()).Should(Equal("2024-01-12T09:37:00+03:00"))
			Expect(T1110_001[0].GetPatternId()).Should(Equal("T1110.001"))
			Expect(T1110_001[0].GetTactic()).Should(Equal("credential-access"))

			Expect(T1110_001[0].GetPattern().GetRemoteSupport()).Should(BeTrue())
			Expect(T1110_001[0].GetPattern().GetRevoked()).Should(BeTrue())
			Expect(T1110_001[0].GetPattern().GetUnderliningCreatedAt()).Should(Equal("2024-01-12T09:37:09+03:00"))
			Expect(T1110_001[0].GetPattern().GetUnderliningCreatedBy()).Should(Equal("admin@thehive.local"))
			Expect(T1110_001[0].GetPattern().GetUnderliningId()).Should(Equal("~164016"))
			Expect(T1110_001[0].GetPattern().GetUnderliningType()).Should(Equal("Pattern"))
			Expect(T1110_001[0].GetPattern().GetDescription()).Should(Equal("exploit edge network"))
			Expect(T1110_001[0].GetPattern().GetDetection()).Should(Equal("Monitor authentication logs for system"))
			Expect(T1110_001[0].GetPattern().GetName()).Should(Equal("Password Guessing"))
			Expect(T1110_001[0].GetPattern().GetPatternId()).Should(Equal("T1110.001"))
			Expect(T1110_001[0].GetPattern().GetPatternType()).Should(Equal("attack-pattern"))
			Expect(T1110_001[0].GetPattern().GetURL()).Should(Equal("https://attack.mitre.org/techniques/T1110/001"))
			Expect(T1110_001[0].GetPattern().GetVersion()).Should(Equal("1.4"))
			Expect(len(T1110_001[0].GetPattern().GetPlatforms())).Should(Equal(3))
			Expect(len(T1110_001[0].GetPattern().GetPermissionsRequired())).Should(Equal(2))
			Expect(len(T1110_001[0].GetPattern().GetDataSources())).Should(Equal(2))
			Expect(len(T1110_001[0].GetPattern().GetTactics())).Should(Equal(1))

			Expect(T1110_001[0].GetPatternParent().GetRemoteSupport()).Should(BeTrue())
			Expect(T1110_001[0].GetPatternParent().GetRevoked()).Should(BeTrue())
			Expect(T1110_001[0].GetPatternParent().GetUnderliningCreatedAt()).Should(Equal("2024-01-12T09:21:21+03:00"))
			Expect(T1110_001[0].GetPatternParent().GetUnderliningCreatedBy()).Should(Equal("root@thehive.gcm"))
			Expect(T1110_001[0].GetPatternParent().GetUnderliningId()).Should(Equal("~236516"))
			Expect(T1110_001[0].GetPatternParent().GetUnderliningType()).Should(Equal("Pattern Parent"))
			Expect(T1110_001[0].GetPatternParent().GetDescription()).Should(Equal("Adversaries may"))
			Expect(T1110_001[0].GetPatternParent().GetDetection()).Should(Equal("Adversaries with no prior knowledge"))
			Expect(T1110_001[0].GetPatternParent().GetName()).Should(Equal("Password Guessing"))
			Expect(T1110_001[0].GetPatternParent().GetPatternId()).Should(Equal("T1110"))
			Expect(T1110_001[0].GetPatternParent().GetPatternType()).Should(Equal("attack-pattern"))
			Expect(T1110_001[0].GetPatternParent().GetURL()).Should(Equal("https://attack.mitre.org/techniques/T1110"))
			Expect(T1110_001[0].GetPatternParent().GetVersion()).Should(Equal("1.14"))
			Expect(len(T1110_001[0].GetPatternParent().GetPlatforms())).Should(Equal(6))
			Expect(len(T1110_001[0].GetPatternParent().GetPermissionsRequired())).Should(Equal(3))
			Expect(len(T1110_001[0].GetPatternParent().GetDataSources())).Should(Equal(3))
			Expect(len(T1110_001[0].GetPatternParent().GetTactics())).Should(Equal(1))
			//----------------------------
			Expect(T1110_001[1].GetUnderliningCreatedAt()).Should(Equal("2024-01-12T09:37:09+03:00"))
			Expect(T1110_001[1].GetUnderliningCreatedBy()).Should(Equal("ag@rcm.ekb"))
			Expect(T1110_001[1].GetUnderliningId()).Should(Equal("~698159200"))
			Expect(T1110_001[1].GetOccurDate()).Should(Equal("2024-01-12T09:37:00+03:00"))
			Expect(T1110_001[1].GetPatternId()).Should(Equal("T1110.001"))
			Expect(T1110_001[1].GetTactic()).Should(Equal("initial-access"))

			Expect(T1110_001[1].GetPattern().GetRemoteSupport()).Should(BeTrue())
			Expect(T1110_001[1].GetPattern().GetRevoked()).Should(BeTrue())
			Expect(T1110_001[1].GetPattern().GetUnderliningCreatedAt()).Should(Equal("2024-01-12T09:37:09+03:00"))
			Expect(T1110_001[1].GetPattern().GetUnderliningCreatedBy()).Should(Equal("adminxxx@thehive.local"))
			Expect(T1110_001[1].GetPattern().GetUnderliningId()).Should(Equal("~61546640"))
			Expect(T1110_001[1].GetPattern().GetUnderliningType()).Should(Equal("Pattern"))
			Expect(T1110_001[1].GetPattern().GetDescription()).Should(Equal("Spearphishing attachment"))
			Expect(T1110_001[1].GetPattern().GetDetection()).Should(Equal("ACSC Email Spoofing"))
			Expect(T1110_001[1].GetPattern().GetName()).Should(Equal("Spearphishing Attachment"))
			Expect(T1110_001[1].GetPattern().GetPatternId()).Should(Equal("T1110.001"))
			Expect(T1110_001[1].GetPattern().GetPatternType()).Should(Equal("attack-pattern"))
			Expect(T1110_001[1].GetPattern().GetURL()).Should(Equal("https://attack.mitre.org/techniques/T1110/001"))
			Expect(T1110_001[1].GetPattern().GetVersion()).Should(Equal("2.2"))
			Expect(len(T1110_001[1].GetPattern().GetPlatforms())).Should(Equal(3))
			Expect(len(T1110_001[1].GetPattern().GetPermissionsRequired())).Should(Equal(1))
			Expect(len(T1110_001[1].GetPattern().GetDataSources())).Should(Equal(3))
			Expect(len(T1110_001[1].GetPattern().GetTactics())).Should(Equal(1))

			Expect(T1110_001[1].GetPatternParent().GetRemoteSupport()).Should(BeTrue())
			Expect(T1110_001[1].GetPatternParent().GetRevoked()).Should(BeTrue())
			Expect(T1110_001[1].GetPatternParent().GetUnderliningCreatedAt()).Should(Equal("2024-01-12T09:21:21+03:00"))
			Expect(T1110_001[1].GetPatternParent().GetUnderliningCreatedBy()).Should(Equal("rootxx@thehive.gcm"))
			Expect(T1110_001[1].GetPatternParent().GetUnderliningId()).Should(Equal("~346894448"))
			Expect(T1110_001[1].GetPatternParent().GetUnderliningType()).Should(Equal("Pattern cxv Parent"))
			Expect(T1110_001[1].GetPatternParent().GetDescription()).Should(Equal("Phishing may also"))
			Expect(T1110_001[1].GetPatternParent().GetDetection()).Should(Equal("SSL/TLS inspection"))
			Expect(T1110_001[1].GetPatternParent().GetName()).Should(Equal("Password Guessing"))
			Expect(T1110_001[1].GetPatternParent().GetPatternId()).Should(Equal("T1110"))
			Expect(T1110_001[1].GetPatternParent().GetPatternType()).Should(Equal("attack-pattern"))
			Expect(T1110_001[1].GetPatternParent().GetURL()).Should(Equal("https://attack.mitre.org/techniques/T1110"))
			Expect(T1110_001[1].GetPatternParent().GetVersion()).Should(Equal("2.14"))
			Expect(len(T1110_001[1].GetPatternParent().GetPlatforms())).Should(Equal(4))
			Expect(len(T1110_001[1].GetPatternParent().GetPermissionsRequired())).Should(Equal(3))
			Expect(len(T1110_001[1].GetPatternParent().GetDataSources())).Should(Equal(3))
			Expect(len(T1110_001[1].GetPatternParent().GetTactics())).Should(Equal(1))

			T1190, ok := ttpsMessaheEs.GetKeyTtp("T1190")
			Expect(ok).Should(BeTrue())

			// Тест T1190 элемента из объекта ttp
			Expect(T1190[0].GetUnderliningCreatedAt()).Should(Equal("2024-01-12T07:13:49+03:00"))
			Expect(T1190[0].GetUnderliningCreatedBy()).Should(Equal("dv-cloud@gcm"))
			Expect(T1190[0].GetUnderliningId()).Should(Equal("~104177744"))
			Expect(T1190[0].GetOccurDate()).Should(Equal("2024-01-12T07:13:00+03:00"))
			Expect(T1190[0].GetPatternId()).Should(Equal("T1190"))
			Expect(T1190[0].GetTactic()).Should(Equal("initial-access"))

			Expect(T1190[0].GetPattern().GetRemoteSupport()).Should(BeTrue())
			Expect(T1190[0].GetPattern().GetRevoked()).ShouldNot(BeTrue())
			Expect(T1190[0].GetPattern().GetUnderliningCreatedAt()).Should(Equal("2024-01-12T07:13:49+03:00"))
			Expect(T1190[0].GetPattern().GetUnderliningCreatedBy()).Should(Equal("prs@rcm"))
			Expect(T1190[0].GetPattern().GetUnderliningId()).Should(Equal("~104177744"))
			Expect(T1190[0].GetPattern().GetUnderliningType()).Should(Equal("Pattern"))
			Expect(T1190[0].GetPattern().GetDescription()).Should(Equal("It`s very important message"))
			Expect(T1190[0].GetPattern().GetDetection()).Should(Equal("Monitor application logs"))
			Expect(T1190[0].GetPattern().GetName()).Should(Equal("Password Guessing"))
			Expect(T1190[0].GetPattern().GetPatternId()).Should(Equal("T1190"))
			Expect(T1190[0].GetPattern().GetPatternType()).Should(Equal("attack-pattern"))
			Expect(T1190[0].GetPattern().GetURL()).Should(Equal("https://attack.mitre.org/techniques/T1110/001"))
			Expect(T1190[0].GetPattern().GetVersion()).Should(Equal("2.4"))
			Expect(len(T1190[0].GetPattern().GetPlatforms())).Should(Equal(4))
			Expect(len(T1190[0].GetPattern().GetPermissionsRequired())).Should(Equal(4))
			Expect(len(T1190[0].GetPattern().GetDataSources())).Should(Equal(2))
			Expect(len(T1190[0].GetPattern().GetTactics())).Should(Equal(1))

			Expect(T1190[0].GetPatternParent().GetRemoteSupport()).Should(BeTrue())
			Expect(T1190[0].GetPatternParent().GetRevoked()).Should(BeTrue())
			Expect(T1190[0].GetPatternParent().GetUnderliningCreatedAt()).Should(Equal("2024-01-12T09:21:21+03:00"))
			Expect(T1190[0].GetPatternParent().GetUnderliningCreatedBy()).Should(Equal("root@example.gcm"))
			Expect(T1190[0].GetPatternParent().GetUnderliningId()).Should(Equal("~16400016"))
			Expect(T1190[0].GetPatternParent().GetUnderliningType()).Should(Equal("Pattern Parent"))
			Expect(T1190[0].GetPatternParent().GetDescription()).Should(Equal(""))
			Expect(T1190[0].GetPatternParent().GetDetection()).Should(Equal("Adversaries with no prior knowledge"))
			Expect(T1190[0].GetPatternParent().GetName()).Should(Equal("Password Guessing"))
			Expect(T1190[0].GetPatternParent().GetPatternId()).Should(Equal("T1220"))
			Expect(T1190[0].GetPatternParent().GetPatternType()).Should(Equal("attack-pattern"))
			Expect(T1190[0].GetPatternParent().GetURL()).Should(Equal("https://attack.mitre.org/techniques/T1220"))
			Expect(T1190[0].GetPatternParent().GetVersion()).Should(Equal("1.1411"))
			Expect(len(T1190[0].GetPatternParent().GetPlatforms())).Should(Equal(6))
			Expect(len(T1190[0].GetPatternParent().GetPermissionsRequired())).Should(Equal(3))
			Expect(len(T1190[0].GetPatternParent().GetDataSources())).Should(Equal(3))
			Expect(len(T1190[0].GetPatternParent().GetTactics())).Should(Equal(1))

			Expect(true).Should(BeTrue())

			//Expect(ttps[0]).Should(Equal())
			//Expect(ttps[0]).Should(Equal())
			//Expect(ttps[0]).Should(Equal())
		})
	})
})
