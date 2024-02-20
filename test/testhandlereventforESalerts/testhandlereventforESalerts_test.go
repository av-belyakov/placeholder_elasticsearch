package testhandlereventforESalerts_test

import (
	"fmt"
	"reflect"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/listhandlercommon"
	"placeholder_elasticsearch/listhandlerforesjson"
	testing "placeholder_elasticsearch/test"
)

var _ = Describe("TesthandlereventforESalerts", Ordered, func() {
	var (
		event        *datamodels.EventMessageForEsAlert        = datamodels.NewEventMessageForEsAlert()
		eventObject  *datamodels.EventMessageForEsAlertObject  = datamodels.NewEventMessageForEsAlertObject()
		eventDetails *datamodels.EventMessageForEsAlertDetails = datamodels.NewEventMessageForEsAlertDetails()

		alert *datamodels.AlertMessageForEsAlert = datamodels.NewAlertMessageForEsAlert()

		sa listhandlerforesjson.SupportiveAlertArtifacts = *listhandlerforesjson.NewSupportiveAlertArtifacts()

		eventObjectCustomFields datamodels.CustomFields = datamodels.CustomFields{}
		alertObjectCustomFields datamodels.CustomFields = datamodels.CustomFields{}

		listHandlerEvent        map[string][]func(interface{})
		listHandlerEventDetails map[string][]func(interface{})

		listHandlerEventObject             map[string][]func(interface{})
		listHandlerEventObjectCustomFields map[string][]func(interface{})

		listHandlerAlert             map[string][]func(interface{})
		listHandlerAlertCustomFields map[string][]func(interface{})
		listHandlerAlertArtifacts    map[string][]func(interface{})

		verifiedTheHiveAlert *datamodels.VerifiedForEsAlert = datamodels.NewVerifiedForEsAlert()
	)

	BeforeAll(func() {
		// ------ EVENT ------
		listHandlerEvent = listhandlerforesjson.NewListHandlerEventAlertElement(event)

		// ------ EVENT OBJECT ------
		listHandlerEventObject = listhandlerforesjson.NewListHandlerEventAlertObjectElement(eventObject)

		// ------ EVENT OBJECT CUSTOMFIELDS ------
		listHandlerEventObjectCustomFields = listhandlercommon.NewListHandlerAlertCustomFieldsElement(eventObjectCustomFields)

		// ------ EVENT DETAILS ------
		listHandlerEventDetails = listhandlerforesjson.NewListHandlerEventAlertDetailsElement(eventDetails)

		// ------ ALERT ------
		listHandlerAlert = listhandlerforesjson.NewListHandlerAlertElement(alert)

		// ------ ALERT CUSTOMFIELDS ------
		listHandlerAlertCustomFields = listhandlercommon.NewListHandlerAlertCustomFieldsElement(alertObjectCustomFields)

		// ------ ALERT ARTIFACTS ------
		listHandlerAlertArtifacts = listhandlerforesjson.NewListHandlerAlertArtifactsElement(&sa)
	})

	Context("Тест 1. Проверка заполнения объекта для хранения events", func() {
		BeforeAll(func() {
			sendData := make(chan struct {
				num   int
				name  string
				value interface{}
			})

			go func() {
				for data := range sendData {
					//************ Обработчики для Event ************
					//event element
					if lf, ok := listHandlerEvent[data.name]; ok {
						//fmt.Printf("%d. (event) v.ElemName: %s, v.ElemValue: '%v'\n", data.num, data.name, data.value)
						for _, f := range lf {
							f(data.value)
						}
					}

					//event.object element
					if lf, ok := listHandlerEventObject[data.name]; ok {
						//fmt.Printf("%d. (event.object) v.ElemName: %s, v.ElemValue: '%v'\n", data.num, data.name, data.value)
						for _, f := range lf {
							f(data.value)
						}
					}

					//event.object.customFields element
					if lf, ok := listHandlerEventObjectCustomFields[data.name]; ok {
						//fmt.Printf("%d. (event.object.customFields) v.ElemName: %s, v.ElemValue: '%v'\n", data.num, data.name, data.value)
						for _, f := range lf {
							f(data.value)
						}
					}

					//event.details element
					if lf, ok := listHandlerEventDetails[data.name]; ok {
						//fmt.Printf("%d. (event.details) v.ElemName: %s, v.ElemValue: '%v'\n", data.num, data.name, data.value)
						for _, f := range lf {
							f(data.value)
						}
					}

					//************ Обработчики для Alert ************
					//alert element
					if lf, ok := listHandlerAlert[data.name]; ok {
						for _, f := range lf {
							f(data.value)
						}
					}

					//alert.customFields
					if lf, ok := listHandlerAlertCustomFields[data.name]; ok {
						for _, f := range lf {
							f(data.value)
						}
					}

					//alert.artifacts
					if strings.Contains(data.name, "alert.artifacts.") {
						if lf, ok := listHandlerAlertArtifacts[data.name]; ok {
							for _, f := range lf {
								f(data.value)
							}
						}
					}
				}
			}()

			//тестовые Event для типа Alert
			for k, v := range testing.GetEventForAlertOne() {
				r := reflect.TypeOf(v.ElemValue)
				switch r.Kind() {
				case reflect.Slice:
					if s, ok := v.ElemValue.([]interface{}); ok {
						for _, value := range s {
							sendData <- struct {
								num   int
								name  string
								value interface{}
							}{
								num:   k,
								name:  v.ElemName,
								value: value,
							}
						}
					}
				default:
					sendData <- struct {
						num   int
						name  string
						value interface{}
					}{
						num:   k,
						name:  v.ElemName,
						value: v.ElemValue,
					}
				}
			}

			// тестовые Alert для типа Alert
			for k, v := range testing.GetAlertForAlertTwo() {
				r := reflect.TypeOf(v.ElemValue)
				switch r.Kind() {
				case reflect.Slice:
					if s, ok := v.ElemValue.([]interface{}); ok {
						for _, value := range s {
							sendData <- struct {
								num   int
								name  string
								value interface{}
							}{
								num:   k,
								name:  v.ElemName,
								value: value,
							}
						}
					}
				default:
					sendData <- struct {
						num   int
						name  string
						value interface{}
					}{
						num:   k,
						name:  v.ElemName,
						value: v.ElemValue,
					}
				}
			}

			close(sendData)

			eventObject.SetValueCustomFields(eventObjectCustomFields)

			event.SetValueObject(*eventObject)
			event.SetValueDetails(*eventDetails)

			alert.SetValueCustomFields(alertObjectCustomFields)
			alert.SetValueArtifacts(sa.GetArtifacts())

			verifiedTheHiveAlert.SetID("fhe78f838f88fg488398f8e3")
			verifiedTheHiveAlert.SetElasticsearchID("3883f8f9-839r983hf848g8h-f84")
			verifiedTheHiveAlert.SetSource("GCM")
			verifiedTheHiveAlert.SetCreateTimestatmp("2024-02-06T15:37:52+03:00")
			verifiedTheHiveAlert.SetEvent(*event)
			verifiedTheHiveAlert.SetAlert(*alert)
		})

		It("Объект Event должен быть успешно заполнен", func() {
			//event
			anyEvent := event.Get()

			Expect(anyEvent.GetBase()).Should(BeTrue())
			Expect(anyEvent.GetStartDate()).Should(Equal("2024-02-06T15:37:52+03:00"))
			Expect(anyEvent.GetRootId()).Should(Equal("~84625227848"))
			Expect(anyEvent.GetObjectId()).Should(Equal("~84625227848"))
			Expect(anyEvent.GetObjectType()).Should(Equal("alert"))
			Expect(anyEvent.GetOrganisation()).Should(Equal("GCM"))
			Expect(anyEvent.GetOrganisationId()).Should(Equal("~4192"))
			Expect(anyEvent.GetOperation()).Should(Equal("update"))
			Expect(anyEvent.GetRequestId()).Should(Equal("55994d44f3b276c1:6483e5ec:18d786c2f83:-8000:138497"))
		})

		It("Объект Event.details.customFields должен быть успешно заполнен", func() {
			//event.details
			anyEventDetails := eventDetails.Get()

			Expect(anyEventDetails.GetSourceRef()).Should(Equal("TSK-8MSK-6-ZPM-240206-1215999"))
			Expect(anyEventDetails.GetTitle()).Should(Equal("Зафиксированна КА"))
			Expect(anyEventDetails.GetDescription()).Should(Equal("**Задача переданная из смежной системы"))
			Expect(len(anyEventDetails.GetTagsAll())).Should(Equal(5))
		})

		It("Объект Event.object должен быть успешно заполнен", func() {
			//event.object
			anyEventObject := eventObject.Get()

			Expect(anyEventObject.GetTlp()).Should(Equal(uint64(3)))
			Expect(anyEventObject.GetUnderliningId()).Should(Equal("~85455464790"))
			Expect(anyEventObject.GetId()).Should(Equal("~85771464712"))
			Expect(anyEventObject.GetCreatedBy()).Should(Equal("v.kovanko@cloud.gcm"))
			Expect(anyEventObject.GetUpdatedBy()).Should(Equal("y.kovalenko@cloud.gcm"))
			Expect(anyEventObject.GetCreatedAt()).Should(Equal("2024-02-06T15:37:52+03:00"))
			Expect(anyEventObject.GetUpdatedAt()).Should(Equal("1970-01-01T03:00:00+03:00"))
			Expect(anyEventObject.GetUnderliningType()).Should(Equal("alert"))
			Expect(anyEventObject.GetTitle()).Should(Equal("Редко встречающиеся признаки ВПО"))
			Expect(anyEventObject.GetDescription()).Should(Equal("из смежной системы: Заслон-Пост-Модерн**"))
			Expect(anyEventObject.GetStatus()).Should(Equal("Ignored"))
			Expect(anyEventObject.GetDate()).Should(Equal("2024-02-06T15:36:57+03:00"))
			Expect(anyEventObject.GetType()).Should(Equal("snort_alert"))
			Expect(anyEventObject.GetObjectType()).Should(Equal("alert"))
			Expect(anyEventObject.GetSource()).Should(Equal("zsiеmSystems"))
			Expect(anyEventObject.GetSourceRef()).Should(Equal("TSK-8MSK-6-ZPM-240206-1215999"))
			Expect(anyEventObject.GetCase()).Should(Equal("aallert"))
			Expect(anyEventObject.GetCaseTemplate()).Should(Equal("snort_alert"))
			Expect(len(anyEventObject.GetTags())).Should(Equal(2))
		})

		It("Объект Event.object.customFields должен быть успешно заполнен", func() {
			anyEventCustomFields := eventObject.GetCustomFields()

			Expect(len(anyEventCustomFields)).Should(Equal(2))

			_, _, _, firstTime := anyEventCustomFields["first-time"].Get()
			Expect(firstTime).Should(Equal("2024-02-06T15:20:30+03:00"))

			_, _, _, lastTime := anyEventCustomFields["last-time"].Get()
			Expect(lastTime).Should(Equal("2024-02-06T15:20:30+03:00"))
		})

		It("Объект Alert должен быть успешно заполнен", func() {
			anyAlert := alert.Get()

			Expect(anyAlert.GetTlp()).Should(Equal(uint64(2)))
			Expect(anyAlert.GetDate()).Should(Equal("2024-02-06T07:59:05+03:00"))
			Expect(anyAlert.GetCreatedAt()).Should(Equal("2024-02-06T07:59:05+03:00"))
			Expect(anyAlert.GetUpdatedAt()).Should(Equal("2024-02-06T15:45:00+03:00"))
			Expect(anyAlert.GetUpdatedBy()).Should(Equal("webhook@cloud.gcm"))
			Expect(anyAlert.GetUnderliningId()).Should(Equal("~88026357960"))
			Expect(anyAlert.GetStatus()).Should(Equal("New"))
			Expect(anyAlert.GetType()).Should(Equal("snort_alert"))
			Expect(anyAlert.GetUnderliningType()).Should(Equal("alert"))
			Expect(anyAlert.GetDescription()).Should(Equal("Атака направлена **внутрь**, использует протоколы: **smtp/tcp**"))
			Expect(anyAlert.GetCaseTemplate()).Should(Equal("snort_alert"))
			Expect(anyAlert.GetSourceRef()).Should(Equal("TSK-8MSK-6-ZPM-240206-1216137"))
			Expect(len(anyAlert.GetTags())).Should(Equal(2))
		})

		It("Объект Alert.customFields должен быть успешно заполнен", func() {
			anyCustomField := alert.Get().GetCustomFields()

			Expect(len(anyCustomField)).Should(Equal(2))

			_, _, _, firstTime := anyCustomField["first-time"].Get()
			Expect(firstTime).Should(Equal("2024-02-06T15:20:30+03:00"))

			_, _, _, lastTime := anyCustomField["last-time"].Get()
			Expect(lastTime).Should(Equal("2024-02-06T15:20:30+03:00"))
		})

		It("Объект Artifacts должен быть успешно заполнен", func() {
			anyArtifacts := alert.Get().GetArtifacts()

			Expect(len(anyArtifacts)).Should(Equal(2))

			urlArkime, ok := anyArtifacts["url_arkime"]
			Expect(ok).Should(BeTrue())

			Expect(urlArkime[0].GetIoc()).Should(BeTrue())
			Expect(urlArkime[0].GetTlp()).Should(Equal(uint64(2)))
			Expect(urlArkime[0].GetUnderliningId()).Should(Equal("~84302291040"))
			Expect(urlArkime[0].GetId()).Should(Equal("~84302291040"))
			Expect(urlArkime[0].GetUnderliningType()).Should(Equal("case_artifact"))
			Expect(urlArkime[0].GetCreatedAt()).Should(Equal("2024-02-06T15:05:15+03:00"))
			Expect(urlArkime[0].GetStartDate()).Should(Equal("2024-02-06T15:05:15+03:00"))
			Expect(urlArkime[0].GetCreatedBy()).Should(Equal("e.anisimova@cloud.gcm"))
			Expect(urlArkime[0].GetData()).Should(Equal("http://anisimova.cloud.gcm:8005/sessions?expression=file%20%3D%3D%20%2Fopt%2Farkime%2Fraw%2F2024_02_06_15_03_01_522219____1707221109_2024_02_06____15_05_09_364449.pcap&date=-1"))
			Expect(urlArkime[0].GetDataType()).Should(Equal("url_arkime"))
			Expect(urlArkime[0].GetMessage()).Should(Equal("Ссылка на Arkime"))
			Expect(len(urlArkime[0].GetTags())).Should(Equal(0))
			Expect(len(urlArkime[0].GetTagsAll())).Should(Equal(1))

			urlPcap, ok := anyArtifacts["url_pcap"]
			Expect(ok).Should(BeTrue())

			Expect(urlPcap[0].GetIoc()).Should(BeTrue())
			Expect(urlPcap[0].GetTlp()).Should(Equal(uint64(1)))
			Expect(urlPcap[0].GetUnderliningId()).Should(Equal(""))
			Expect(urlPcap[0].GetId()).Should(Equal("~84998705312"))
			Expect(urlPcap[0].GetUnderliningType()).Should(Equal("case_artifact"))
			Expect(urlPcap[0].GetCreatedAt()).Should(Equal("2024-02-06T15:05:14+03:00"))
			Expect(urlPcap[0].GetStartDate()).Should(Equal("2024-02-06T15:05:14+03:00"))
			Expect(urlPcap[0].GetCreatedBy()).Should(Equal(""))
			Expect(urlPcap[0].GetData()).Should(Equal("ftp://ftp.cloud.gcm/traffic/8030030/1707221109_2024_02_06____15_05_09_364449.pcap"))
			Expect(urlPcap[0].GetDataType()).Should(Equal("url_pcap"))
			Expect(urlPcap[0].GetMessage()).Should(Equal("Download a piece of traffic"))
			Expect(len(urlPcap[0].GetTags())).Should(Equal(4))
			Expect(len(urlPcap[0].GetTagsAll())).Should(Equal(5))
		})

		It("Должен быть полностью заполнен объект verifiedTheHiveAlert", func() {
			verified := verifiedTheHiveAlert.Get()

			fmt.Println("---=== VERIFEDFORESALERT ===---")
			fmt.Println(verified.ToStringBeautiful(0))

			Expect(verified.GetID()).Should(Equal("fhe78f838f88fg488398f8e3"))
			Expect(verified.GetElasticsearchID()).Should(Equal("3883f8f9-839r983hf848g8h-f84"))
			Expect(verified.GetSource()).Should(Equal("GCM"))
			Expect(verified.GetCreateTimestatmp()).Should(Equal("2024-02-06T15:37:52+03:00"))
		})
	})

	Context("Тест 2. Проверка замены старых значений EventMessageForEsAlert объекта, новыми значениями, если они отличаются", func() {
		oldStruct := datamodels.EventMessageForEsAlert{
			Base:           false,
			StartDate:      "2024-02-06T15:20:30+03:00",
			RootId:         "~84625227848",
			ObjectId:       "~4192",
			ObjectType:     "alert",
			Organisation:   "GCM",
			OrganisationId: "~4192",
			Operation:      "new",
			RequestId:      "55994d44f3b276c1:6483e5ec:18d786c2f83:-8000:138497",
			Details: datamodels.EventMessageForEsAlertDetails{
				SourceRef:   "TSK-8MSK-6-ZPM-240206-1215999",
				Title:       "222",
				Description: "111",
				Tags: map[string][]string{
					"ats:reason": {"INFO Controlled FGS"},
					"sensor:id":  {"8030066"},
				},
				TagsAll: []string{
					"ATs:reason=\"INFO Controlled FGS\"",
					"Sensor:id=\"8030066\"",
				},
			},
			Object: datamodels.EventMessageForEsAlertObject{
				Tlp:             1,
				UnderliningId:   "~85455464790",
				Id:              "~85455464790",
				CreatedBy:       "ddddd",
				CreatedAt:       "1970-01-01T03:00:00+03:00",
				UpdatedAt:       "1970-01-01T03:00:00+03:00",
				UnderliningType: "aalllert",
				Title:           "vbb er3",
				Description:     "any more",
				Status:          "None",
				Date:            "2024-02-06T15:37:52+03:00",
				Type:            "snort_alert",
				ObjectType:      "",
				Source:          "zsiеmSystems",
				SourceRef:       "TSK-8MSK-6-ZPM-240206-1215999",
				Case:            "alert",
				CaseTemplate:    "alert_snort",
				Tags: map[string][]string{
					"sensor:id": {"8030012"},
				},
				TagsAll: []string{
					"'Sensor:id=\"8030012\"'",
					"'Webhook:send=ES'",
				},
				CustomFields: datamodels.CustomFields{
					"first-time": &datamodels.CustomFieldDateType{
						Order: 0,
						Date:  "2024-02-06T15:20:30+03:00",
					},
					"last-time": &datamodels.CustomFieldDateType{
						Order: 0,
						Date:  "2024-02-06T15:20:30+03:00",
					},
				},
			},
		}

		newStruct := datamodels.EventMessageForEsAlert{
			Base:           true,                       //замена
			StartDate:      "2024-02-13T5:12:24+03:00", //замена
			RootId:         "~84625227848",
			ObjectId:       "~4192",
			ObjectType:     "ALERT",   //замена
			Organisation:   "GCM-MSK", //замена
			OrganisationId: "~419211", //замена
			Operation:      "update",  //замена
			RequestId:      "55994d44f3b276c1:6483e5ec:18d786c2f83:-8000:138497",
			Details: datamodels.EventMessageForEsAlertDetails{
				SourceRef:   "TSK-8MSK-6-ZPM-240206-1215999",
				Title:       "протоколы: **smtp/tcp**",            //замена
				Description: "использует протоколы: **smtp/tcp**", //замена
				//замена
				Tags: map[string][]string{
					"ats:reason": {
						"INFO Controlled FGS",
						"Редко встречающиеся признаки ВПО",
					},
					"sensor:id":             {"8030066"},
					"ats:geoip":             {"Китай"},
					"misp:Network activity": {"snort"},
				},
				TagsAll: []string{
					"ATs:reason=\"INFO Controlled FGS\"",
					"Sensor:id=\"8030066\"",
					"'APPA:Direction=\"inbound\"'",
					"ATs:geoip=\"Китай\"",
					"misp:Network activity=\"snort\"",
					"ATs:reason=\"Редко встречающиеся признаки ВПО\"",
				},
			},
			Object: datamodels.EventMessageForEsAlertObject{
				Tlp:             1,
				UnderliningId:   "~85455464790",
				Id:              "~85455464790",
				CreatedBy:       "d.zablotsky@cloud.gcm",       //замена
				CreatedAt:       "2024-02-10T23:25:14+03:00",   //замена
				UpdatedAt:       "2024-02-06T15:15:14+03:00",   //замена
				UnderliningType: "ALERT",                       //замена
				Title:           "Атака направлена **внутрь**", //замена
				Description:     "Вирусное заражение",          //замена
				Status:          "None",
				Date:            "2024-02-06T15:37:52+03:00",
				Type:            "snort_alert",
				ObjectType:      "",
				Source:          "zsiеmSystems",
				SourceRef:       "TSK-8MSK-6-ZPM-240206-1215999",
				Case:            "alert",
				CaseTemplate:    "Alert_Snort", //замена
				//замена
				Tags: map[string][]string{
					"sensor:id":             {"8030105"},
					"misp:Payload delivery": {"email-src"},
				},
				TagsAll: []string{
					"'Webhook:send=ES'",
					"'Sensor:id=\"8030105\"'",
					"misp:Payload delivery=\"email-src\"",
				},
				CustomFields: datamodels.CustomFields{
					"first-time": &datamodels.CustomFieldDateType{
						Order: 0,
						Date:  "2024-02-06T15:20:30+03:00",
					},
					//замена
					"last-time": &datamodels.CustomFieldDateType{
						Order: 0,
						Date:  "2024-02-07T22:48:13+03:00",
					},
				},
			},
		}

		It("Ряд полей в EventMessageForEsAlert должны быть успешно заменены", func() {
			num, err := oldStruct.ReplacingOldValues(newStruct)

			Expect(err).ShouldNot(HaveOccurred())
			//кол-во замененных полей
			Expect(num).Should(Equal(23))

			//меняется
			Expect(oldStruct.GetBase()).Should(BeTrue())
			//меняется
			Expect(oldStruct.GetStartDate()).Should(Equal("2024-02-13T5:12:24+03:00"))
			//НЕ меняется
			Expect(oldStruct.GetObjectId()).Should(Equal("~4192"))
			//меняется
			Expect(oldStruct.GetObjectType()).Should(Equal("ALERT"))
			//меняется
			Expect(oldStruct.GetOrganisation()).Should(Equal("GCM-MSK"))
			//меняется
			Expect(oldStruct.GetOperation()).Should(Equal("update"))
			//меняется
			Expect(oldStruct.GetOrganisationId()).Should(Equal("~419211"))
			//НЕ меняется
			Expect(oldStruct.GetRequestId()).Should(Equal("55994d44f3b276c1:6483e5ec:18d786c2f83:-8000:138497"))

			//--- Details ---
			//НЕ меняется
			Expect(oldStruct.Details.GetSourceRef()).Should(Equal("TSK-8MSK-6-ZPM-240206-1215999"))
			//меняется
			Expect(oldStruct.Details.GetTitle()).Should(Equal("протоколы: **smtp/tcp**"))
			//меняется
			Expect(oldStruct.Details.GetDescription()).Should(Equal("использует протоколы: **smtp/tcp**"))
			//меняется
			Expect(len(oldStruct.Details.GetTags())).Should(Equal(4))
			//меняется
			Expect(len(oldStruct.Details.GetTagsAll())).Should(Equal(6))

			//--- Object ---
			//меняется
			Expect(oldStruct.Object.GetCreatedBy()).Should(Equal("d.zablotsky@cloud.gcm"))
			//меняется
			Expect(oldStruct.Object.GetCreatedAt()).Should(Equal("2024-02-10T23:25:14+03:00"))
			//меняется
			Expect(oldStruct.Object.GetUpdatedAt()).Should(Equal("2024-02-06T15:15:14+03:00"))
			//меняется
			Expect(oldStruct.Object.GetUnderliningType()).Should(Equal("ALERT"))
			//меняется
			Expect(oldStruct.Object.GetTitle()).Should(Equal("Атака направлена **внутрь**"))
			//меняется
			Expect(oldStruct.Object.GetDescription()).Should(Equal("Вирусное заражение"))
			//НЕ меняется
			Expect(oldStruct.Object.GetSource()).Should(Equal("zsiеmSystems"))
			//меняется
			Expect(oldStruct.Object.GetCaseTemplate()).Should(Equal("Alert_Snort"))
			//меняется
			Expect(len(oldStruct.Object.GetTags())).Should(Equal(2))
			//меняется
			Expect(len(oldStruct.Object.GetTagsAll())).Should(Equal(4))

			customFields := oldStruct.Object.GetCustomFields()
			_, _, _, firstTime := customFields["first-time"].Get()
			_, _, _, lastTime := customFields["last-time"].Get()
			//НЕ меняется
			Expect(firstTime).Should(Equal("2024-02-06T15:20:30+03:00"))
			//меняется
			Expect(lastTime).Should(Equal("2024-02-07T22:48:13+03:00"))

			fmt.Println("---=== EventMessageForEsAlert ===---")
			fmt.Println(oldStruct.ToStringBeautiful(0))
		})
	})

	Context("Тест 3. Проверка замены старых значений AlertMessageForEsAlert объекта, новыми значениями, если они отличаются", func() {
		oldStruct := datamodels.AlertMessageForEsAlert{
			Tlp:       2,
			Date:      "1970-01-01T03:00:00+03:00",
			CreatedAt: "2024-02-07T11:11:11+03:00",
			// UpdatedAt: ,
			UpdatedBy:       "webhook@cloud.gcm",
			UnderliningId:   "~88026357960",
			Status:          "New",
			Type:            "snort",
			UnderliningType: "__Snort",
			Description:     "free alerts",
			CaseTemplate:    "sonr",
			SourceRef:       "TSK-8MSK-6-ZPM-240206-1216137",
			Tags: map[string][]string{
				"sensor:id":  {"8030105"},
				"ats:reason": {"Редко встречающиеся признаки ВПО"},
			},
			TagsAll: []string{
				"Sensor:id=\"8030105\"",
				"ATs:reason=\"Редко встречающиеся признаки ВПО\"",
				"'Webhook:send=ES'",
			},
			CustomFields: datamodels.CustomFields{
				"first-time": &datamodels.CustomFieldDateType{
					Order: 0,
					Date:  "2024-01-01T05:22:30+03:00",
				},
				"last-time": &datamodels.CustomFieldDateType{
					Order: 0,
					Date:  "2024-01-17T00:18:13+03:00",
				},
			},
			Artifacts: map[string][]datamodels.ArtifactForEsAlert{
				"coordinates": {
					{
						Ioc:           false,
						Tlp:           1,
						UnderliningId: "~84302220012",
						Id:            "~84302220012",
						CreatedAt:     "2024-01-26T13:02:01+03:00",
						//UpdatedAt: ,
						StartDate: "2024-01-26T13:02:01+03:00",
						CreatedBy: "friman@email.net",
						Data:      "63.5656 89.12",
						DataType:  "coordinates",
						Message:   "Any message",
						Tags: map[string][]string{
							"sensor:id":     {"1111111"},
							"geoip:country": {"CH"},
						},
						TagsAll: []string{
							"Sensor:id=\"1111111\"",
							"geoip:country=CH",
							"'Webhook:send=ES'",
						},
					},
				},
				"ipaddr": {
					{
						Ioc:           true,
						Tlp:           2,
						UnderliningId: "~306522241",
						Id:            "~306522241",
						CreatedAt:     "2024-01-16T03:32:01+03:00",
						StartDate:     "2024-01-04T19:32:01+03:00",
						CreatedBy:     "example@email.net",
						Data:          "5.63.123.99",
						DataType:      "ipaddr",
						Message:       "ffdffd fdg",
						Tags: map[string][]string{
							"sensor:id":     {"3411"},
							"geoip:country": {"RU"},
						},
						TagsAll: []string{
							"Sensor:id=\"3411\"",
							"geoip:country=RU",
							"'Webhook:send=ES'",
						},
					},
				},
			},
		}

		newStruct := datamodels.AlertMessageForEsAlert{
			Tlp:             3, //замена
			Date:            "1970-01-01T03:00:00+03:00",
			CreatedAt:       "2024-02-10T10:00:41+03:00", //замена
			UpdatedAt:       "2024-02-11T12:34:48+03:00", //замена
			UpdatedBy:       "webexample@cloud.gcm",      //замена
			UnderliningId:   "~88026357960",
			Status:          "Update",       //замена
			Type:            "snort_alert",  //замена
			UnderliningType: "snort_alert",  //замена
			Description:     "free alerts!", //замена
			CaseTemplate:    "snort",        //замена
			SourceRef:       "TSK-8MSK-6-ZPM-240206-1216137",
			//замена
			Tags: map[string][]string{
				"ats:reason": {"Редко встречающиеся признаки ВПО"},
			},
			//замена
			TagsAll: []string{
				"ATs:reason=\"Редко встречающиеся признаки ВПО\"",
				"'Webhook:send=ES'",
				"APPA:Direction=\"inbound\"",
			},
			CustomFields: datamodels.CustomFields{
				//замена
				"first-time": &datamodels.CustomFieldDateType{
					Order: 0,
					Date:  "2024-01-22T15:13:10+03:00",
				},
				"last-time": &datamodels.CustomFieldDateType{
					Order: 0,
					Date:  "2024-01-17T00:18:13+03:00",
				},
			},
			Artifacts: map[string][]datamodels.ArtifactForEsAlert{
				"coordinates": {
					{
						Ioc:           true, //замена
						Tlp:           3,    //замена
						UnderliningId: "",   //НЕ замена
						Id:            "~84302220012",
						CreatedAt:     "2024-01-27T22:17:17+03:00", //замена
						StartDate:     "2024-01-26T13:02:01+03:00",
						CreatedBy:     "friman@email.net",
						Data:          "63.5656 89.1211", //замена
						DataType:      "coordinates",
						Message:       "Any message",
						//замена
						Tags: map[string][]string{
							"sensor:id":     {"12345667"},
							"geoip:country": {"CH", "US"},
						},
						//замена
						TagsAll: []string{
							"Sensor:id=\"12345667\"",
							"geoip:country=CH",
							"geoip:country=US",
							"'Webhook:send=ES'",
						},
					},
					//добавление
					{
						Ioc:           false,
						Tlp:           1,
						UnderliningId: "",
						Id:            "~8430120011",
						CreatedAt:     "2024-01-27T22:17:17+03:00",
						StartDate:     "2024-01-26T13:02:01+03:00",
						CreatedBy:     "friman@email.net",
						Data:          "55.89.12.11",
						DataType:      "ip address",
						Message:       "funy description",
						Tags: map[string][]string{
							"sensor:id":     {"43522"},
							"geoip:country": {"RU"},
						},
						TagsAll: []string{
							"Sensor:id=\"12345667\"",
							"geoip:country=RU",
							"'Webhook:send=ES'",
						},
					},
				},
				"ip_home": {
					{
						Ioc:           false,
						Tlp:           1,
						UnderliningId: "~7344456683",
						Id:            "~7344456683",
						CreatedAt:     "2024-01-17T13:12:01+03:00",
						StartDate:     "2024-01-04T19:32:01+03:00",
						CreatedBy:     "example@email.net",
						Data:          "5.63.123.99",
						DataType:      "ipaddr",
						Message:       "ffdffd fdg",
						Tags: map[string][]string{
							"geoip:country": {"RU"},
						},
						TagsAll: []string{
							"geoip:country=RU",
							"'Webhook:send=ES'",
						},
					},
				},
			},
		}
		It("Ряд полей в AlertMessageForEsAlert должны быть успешно заменены", func() {
			num, err := oldStruct.ReplacingOldValues(newStruct)

			/*

				Этот тест не отрабатывает нормально, надо смотреть подробнее
				в чем причина


			*/

			Expect(err).ShouldNot(HaveOccurred())
			//кол-во замененных полей
			Expect(num).Should(Equal(18))

			//Alert
			//меняется
			Expect(oldStruct.GetTlp()).Should(Equal(uint64(3)))
			//меняется
			//НЕ меняется
			Expect(oldStruct.GetDate()).Should(Equal("1970-01-01T03:00:00+03:00"))
			//меняется
			Expect(oldStruct.GetCreatedAt()).Should(Equal("2024-02-10T10:00:41+03:00"))
			//меняется
			Expect(oldStruct.GetUpdatedAt()).Should(Equal("2024-02-11T12:34:48+03:00"))
			//меняется
			Expect(oldStruct.GetStatus()).Should(Equal("Update"))
			//меняется
			Expect(oldStruct.GetType()).Should(Equal("snort_alert"))
			//меняется
			Expect(oldStruct.GetUnderliningType()).Should(Equal("snort_alert"))
			//меняется
			Expect(oldStruct.GetDescription()).Should(Equal("free alerts!"))
			//меняется
			Expect(oldStruct.GetCaseTemplate()).Should(Equal("snort"))
			//меняется
			Expect(len(oldStruct.GetTags())).Should(Equal(2))

			customFields := oldStruct.GetCustomFields()
			_, _, _, firstTime := customFields["first-time"].Get()
			_, _, _, lastTime := customFields["last-time"].Get()
			//меняется
			Expect(firstTime).Should(Equal("2024-01-22T15:13:10+03:00"))
			//НЕ меняется
			Expect(lastTime).Should(Equal("2024-01-17T00:18:13+03:00"))

			Expect(len(oldStruct.GetArtifacts())).Should(Equal(3))
			coordinates, ok := oldStruct.GetArtifacts()["coordinates"]
			Expect(ok).Should(BeTrue())
			Expect(len(coordinates)).Should(Equal(3))
			Expect(coordinates[0].GetIoc()).Should(BeTrue())
			Expect(coordinates[0].GetTlp()).Should(Equal(3))
			Expect(coordinates[0].GetCreatedAt()).Should(Equal("2024-01-27T22:17:17+03:00"))
			Expect(coordinates[0].GetData()).Should(Equal("63.5656 89.1211"))
			Expect(coordinates[0].GetMessage()).Should(Equal("Any message"))
			Expect(len(coordinates[0].GetTags())).Should(Equal(2))
			Expect(len(coordinates[0].GetTagsAll())).Should(Equal(4))
			Expect(coordinates[1].GetIoc()).ShouldNot(BeTrue())
			Expect(coordinates[1].GetTlp()).Should(Equal(1))
			Expect(coordinates[1].GetCreatedAt()).Should(Equal("2024-01-27T22:17:17+03:00"))
			Expect(coordinates[1].GetData()).Should(Equal("55.89.12.11"))
			Expect(coordinates[1].GetMessage()).Should(Equal("funy description"))
			Expect(len(coordinates[1].GetTags())).Should(Equal(2))
			Expect(len(coordinates[1].GetTagsAll())).Should(Equal(3))

			ipaddr, ok := oldStruct.GetArtifacts()["ipaddr"]
			Expect(ok).Should(BeTrue())
			Expect(len(ipaddr[0].GetTags())).Should(Equal(2))
			Expect(len(ipaddr[0].GetTagsAll())).Should(Equal(3))

			ipHome, ok := oldStruct.GetArtifacts()["ip_home"]
			Expect(ok).Should(BeTrue())
			Expect(ipHome[0].GetData()).Should(Equal("5.63.123.99"))
			Expect(len(ipHome[0].GetTags())).Should(Equal(2))
			Expect(len(ipHome[0].GetTagsAll())).Should(Equal(3))

			fmt.Println("---=== AlertMessageForEsAlert ===---")
			fmt.Println(oldStruct.ToStringBeautiful(0))
		})
	})
})
