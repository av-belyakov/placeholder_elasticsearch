package testhandlereventforalert_test

import (
	"fmt"
	"reflect"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/datamodels/commonalert"
	"placeholder_elasticsearch/datamodels/commonalertartifact"
	commonevent "placeholder_elasticsearch/datamodels/commonevent"
	commonobjectevent "placeholder_elasticsearch/datamodels/commonobjectevent"
	"placeholder_elasticsearch/listhandlercommon"
	"placeholder_elasticsearch/listhandlerthehivejson"
	testing "placeholder_elasticsearch/test"
)

var _ = Describe("Testeventforalert", Ordered, func() {
	var (
		event        *datamodels.EventMessageTheHiveAlert = datamodels.NewEventMessageTheHiveAlert()
		eventObject  *datamodels.EventAlertObject         = datamodels.NewEventAlertObject()
		eventDetails *datamodels.EventAlertDetails        = datamodels.NewEventAlertDetails()

		alert *datamodels.AlertMessageTheHiveAlert = datamodels.NewAlertMessageTheHiveAlert()

		sa listhandlerthehivejson.SupportiveAlertArtifacts = *listhandlerthehivejson.NewSupportiveAlertArtifacts()

		eventObjectCustomFields datamodels.CustomFields = datamodels.CustomFields{}
		alertObjectCustomFields datamodels.CustomFields = datamodels.CustomFields{}

		listHandlerEvent        map[string][]func(interface{})
		listHandlerEventDetails map[string][]func(interface{})

		listHandlerEventObject             map[string][]func(interface{})
		listHandlerEventObjectCustomFields map[string][]func(interface{})

		listHandlerAlert             map[string][]func(interface{})
		listHandlerAlertCustomFields map[string][]func(interface{})
		listHandlerAlertArtifacts    map[string][]func(interface{})

		verifiedTheHiveAlert *datamodels.VerifiedTheHiveAlert = datamodels.NewVerifiedTheHiveAlert()
	)

	BeforeAll(func() {
		// ------ EVENT ------
		listHandlerEvent = listhandlerthehivejson.NewListHandlerEventAlertElement(event)

		// ------ EVENT OBJECT ------
		listHandlerEventObject = listhandlerthehivejson.NewListHandlerEventAlertObjectElement(eventObject)

		// ------ EVENT OBJECT CUSTOMFIELDS ------
		listHandlerEventObjectCustomFields = listhandlercommon.NewListHandlerAlertCustomFieldsElement(eventObjectCustomFields)

		// ------ EVENT DETAILS ------
		listHandlerEventDetails = listhandlerthehivejson.NewListHandlerEventAlertDetailsElement(eventDetails)

		// ------ ALERT ------
		listHandlerAlert = listhandlerthehivejson.NewListHandlerAlertElement(alert)

		// ------ ALERT CUSTOMFIELDS ------
		listHandlerAlertCustomFields = listhandlercommon.NewListHandlerAlertCustomFieldsElement(alertObjectCustomFields)

		// ------ ALERT ARTIFACTS ------
		listHandlerAlertArtifacts = listhandlerthehivejson.NewListHandlerAlertArtifactsElement(&sa)
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
			Expect(len(anyEventDetails.GetTags())).Should(Equal(5))
		})

		It("Объект Event.object должен быть успешно заполнен", func() {
			//event.object
			anyEventObject := eventObject.Get()

			Expect(anyEventObject.GetFollow()).Should(BeTrue())
			Expect(anyEventObject.GetSeverity()).Should(Equal(uint64(2)))
			Expect(anyEventObject.GetTlp()).Should(Equal(uint64(3)))
			Expect(anyEventObject.GetPap()).Should(Equal(uint64(5)))
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
			Expect(len(anyEventObject.GetTags())).Should(Equal(3))
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

			Expect(anyAlert.GetFollow()).Should(BeTrue())
			Expect(anyAlert.GetTlp()).Should(Equal(uint64(2)))
			Expect(anyAlert.GetSeverity()).Should(Equal(uint64(3)))
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
			Expect(len(anyAlert.GetTags())).Should(Equal(4))
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

			Expect(anyArtifacts[0].GetIoc()).Should(BeTrue())
			Expect(anyArtifacts[0].GetSighted()).Should(BeTrue())
			Expect(anyArtifacts[0].GetIgnoreSimilarity()).Should(BeTrue())
			Expect(anyArtifacts[0].GetTlp()).Should(Equal(uint64(2)))
			Expect(anyArtifacts[0].GetUnderliningId()).Should(Equal("~84302291040"))
			Expect(anyArtifacts[0].GetId()).Should(Equal("~84302291040"))
			Expect(anyArtifacts[0].GetUnderliningType()).Should(Equal("case_artifact"))
			Expect(anyArtifacts[0].GetCreatedAt()).Should(Equal("2024-02-06T15:05:15+03:00"))
			Expect(anyArtifacts[0].GetUpdatedAt()).Should(Equal("1970-01-01T03:00:00+03:00"))
			Expect(anyArtifacts[0].GetStartDate()).Should(Equal("2024-02-06T15:05:15+03:00"))
			Expect(anyArtifacts[0].GetCreatedBy()).Should(Equal("e.anisimova@cloud.gcm"))
			Expect(anyArtifacts[0].GetUpdatedBy()).Should(Equal(""))
			Expect(anyArtifacts[0].GetData()).Should(Equal("http://anisimova.cloud.gcm:8005/sessions?expression=file%20%3D%3D%20%2Fopt%2Farkime%2Fraw%2F2024_02_06_15_03_01_522219____1707221109_2024_02_06____15_05_09_364449.pcap&date=-1"))
			Expect(anyArtifacts[0].GetDataType()).Should(Equal("url_arkime"))
			Expect(anyArtifacts[0].GetMessage()).Should(Equal("Ссылка на Arkime"))
			Expect(len(anyArtifacts[0].GetTags())).Should(Equal(1))

			Expect(anyArtifacts[1].GetIoc()).Should(BeTrue())
			Expect(anyArtifacts[1].GetSighted()).Should(BeTrue())
			Expect(anyArtifacts[1].GetIgnoreSimilarity()).Should(BeTrue())
			Expect(anyArtifacts[1].GetTlp()).Should(Equal(uint64(1)))
			Expect(anyArtifacts[1].GetUnderliningId()).Should(Equal(""))
			Expect(anyArtifacts[1].GetId()).Should(Equal("~84998705312"))
			Expect(anyArtifacts[1].GetUnderliningType()).Should(Equal("case_artifact"))
			Expect(anyArtifacts[1].GetCreatedAt()).Should(Equal("2024-02-06T15:05:14+03:00"))
			Expect(anyArtifacts[1].GetUpdatedAt()).Should(Equal("2024-02-06T15:05:14+03:00"))
			Expect(anyArtifacts[1].GetStartDate()).Should(Equal("2024-02-06T15:05:14+03:00"))
			Expect(anyArtifacts[1].GetCreatedBy()).Should(Equal(""))
			Expect(anyArtifacts[1].GetUpdatedBy()).Should(Equal("zsiem@cloud.gcm"))
			Expect(anyArtifacts[1].GetData()).Should(Equal("ftp://ftp.cloud.gcm/traffic/8030030/1707221109_2024_02_06____15_05_09_364449.pcap"))
			Expect(anyArtifacts[1].GetDataType()).Should(Equal("url_pcap"))
			Expect(anyArtifacts[1].GetMessage()).Should(Equal("Download a piece of traffic"))
			Expect(len(anyArtifacts[1].GetTags())).Should(Equal(5))
		})

		It("Должен быть полностью заполнен объект verifiedTheHiveAlert", func() {
			verified := verifiedTheHiveAlert.Get()

			//fmt.Println("---=== VERIFEDTHEHIVEALERT ===---")
			//fmt.Println(verified.ToStringBeautiful(0))

			Expect(verified.GetID()).Should(Equal("fhe78f838f88fg488398f8e3"))
			Expect(verified.GetElasticsearchID()).Should(Equal("3883f8f9-839r983hf848g8h-f84"))
			Expect(verified.GetSource()).Should(Equal("GCM"))
			Expect(verified.GetCreateTimestatmp()).Should(Equal("2024-02-06T15:37:52+03:00"))
		})
	})

	Context("Тест 2. Проверка замены старых значений EventMessageTheHiveAlert объекта, новыми значениями, если они отличаются", func() {
		oldStruct := datamodels.EventMessageTheHiveAlert{
			CommonEventType: commonevent.CommonEventType{
				Base:           false,
				StartDate:      "2024-02-06T15:20:30+03:00",
				RootId:         "~84625227848",
				ObjectId:       "~4192",
				ObjectType:     "alert",
				Organisation:   "GCM",
				OrganisationId: "~4192",
				Operation:      "new",
				RequestId:      "55994d44f3b276c1:6483e5ec:18d786c2f83:-8000:138497",
			},
			Details: datamodels.EventAlertDetails{
				SourceRef:   "TSK-8MSK-6-ZPM-240206-1215999",
				Title:       "222",
				Description: "111",
				Tags: []string{
					"ATs:reason=\"INFO Controlled FGS\"",
					"Sensor:id=\"8030066\"",
				},
			},
			Object: datamodels.EventAlertObject{
				Severity: 1,
				Pap:      1,
				CommonEventAlertObject: commonobjectevent.CommonEventAlertObject{
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
				},
				Tags: []string{
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
				/*map[string]datamodels.CustomerFields{
					"first-time": &datamodels.CustomFieldDateType{
						Order: 0,
						Date:  "2024-02-06T15:20:30+03:00",
					},
					"last-time": &datamodels.CustomFieldDateType{
						Order: 0,
						Date:  "2024-02-06T15:20:30+03:00",
					},
				},*/
			},
		}

		newStruct := datamodels.EventMessageTheHiveAlert{
			CommonEventType: commonevent.CommonEventType{
				Base:           true,                        //замена
				StartDate:      "2024-02-13T05:12:24+03:00", //замена
				RootId:         "~84625227848",
				ObjectId:       "~4192",
				ObjectType:     "ALERT",   //замена
				Organisation:   "GCM-MSK", //замена
				OrganisationId: "~419211", //замена
				Operation:      "update",  //замена
				RequestId:      "55994d44f3b276c1:6483e5ec:18d786c2f83:-8000:138497",
			},
			Details: datamodels.EventAlertDetails{
				SourceRef:   "TSK-8MSK-6-ZPM-240206-1215999",
				Title:       "протоколы: **smtp/tcp**",            //замена
				Description: "использует протоколы: **smtp/tcp**", //замена
				//замена
				Tags: []string{
					"ATs:reason=\"INFO Controlled FGS\"",
					"Sensor:id=\"8030066\"",
					"'APPA:Direction=\"inbound\"'",
					"ATs:geoip=\"Китай\"",
					"misp:Network activity=\"snort\"",
				},
			},
			Object: datamodels.EventAlertObject{
				Severity: 1,
				Pap:      1,
				CommonEventAlertObject: commonobjectevent.CommonEventAlertObject{
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
				},
				//замена
				Tags: []string{
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
				/*map[string]datamodels.CustomerFields{
					"first-time": &datamodels.CustomFieldDateType{
						Order: 0,
						Date:  "2024-02-06T15:20:30+03:00",
					},
					//замена
					"last-time": &datamodels.CustomFieldDateType{
						Order: 0,
						Date:  "2024-02-07T22:48:13+03:00",
					},
				},*/
			},
		}

		It("Ряд полей в EventMessageTheHiveAlert должны быть успешно заменены", func() {
			num, err := oldStruct.ReplacingOldValues(newStruct)

			Expect(err).ShouldNot(HaveOccurred())
			//кол-во замененных полей
			Expect(num).Should(Equal(18))

			//fmt.Println("---=== VERIFEDTHEHIVEALERT ===---")
			//fmt.Println(oldStruct.ToStringBeautiful(0))

			//меняется
			Expect(oldStruct.GetBase()).Should(BeTrue())
			//меняется
			Expect(oldStruct.GetStartDate()).Should(Equal("2024-02-13T05:12:24+03:00"))
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
			Expect(len(oldStruct.Details.GetTags())).Should(Equal(5))

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
			Expect(len(oldStruct.Object.GetTags())).Should(Equal(4))

			customFields := oldStruct.Object.GetCustomFields()
			_, _, _, firstTime := customFields["first-time"].Get()
			_, _, _, lastTime := customFields["last-time"].Get()
			//НЕ меняется
			Expect(firstTime).Should(Equal("2024-02-06T15:20:30+03:00"))
			//меняется
			Expect(lastTime).Should(Equal("2024-02-07T22:48:13+03:00"))

			Expect(true).Should(BeTrue())
		})
	})

	Context("Тест 3. Проверка замены старых значений AlertMessageTheHiveAlert объекта, новыми значениями, если они отличаются", func() {
		oldStruct := datamodels.AlertMessageTheHiveAlert{
			Follow: false,
			CommonAlertType: commonalert.CommonAlertType{
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
			},
			Severity: 2,
			Tags: []string{
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
			/*map[string]datamodels.CustomerFields{
				"first-time": &datamodels.CustomFieldDateType{
					Order: 0,
					Date:  "2024-01-01T05:22:30+03:00",
				},
				"last-time": &datamodels.CustomFieldDateType{
					Order: 0,
					Date:  "2024-01-17T00:18:13+03:00",
				},
			},*/
			Artifacts: []datamodels.AlertArtifact{
				{
					CommonArtifactType: commonalertartifact.CommonArtifactType{
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
					},
					Sighted:          false,
					IgnoreSimilarity: true,
					UpdatedBy:        "friman@email.net",
					Tags: []string{
						"Sensor:id=\"1111111\"",
						"geoip:country=CH",
						"'Webhook:send=ES'",
					},
				},
				{
					CommonArtifactType: commonalertartifact.CommonArtifactType{
						Ioc:           true,
						Tlp:           2,
						UnderliningId: "~306522241",
						Id:            "~306522241",
						CreatedAt:     "2024-01-16T03:32:01+03:00",
						//UpdatedAt: ,
						StartDate: "2024-01-04T19:32:01+03:00",
						CreatedBy: "example@email.net",
						Data:      "5.63.123.99",
						DataType:  "ipaddr",
						Message:   "ffdffd fdg",
					},
					Sighted:          false,
					IgnoreSimilarity: true,
					UpdatedBy:        "example@email.net",
					Tags: []string{
						"Sensor:id=\"3411\"",
						"geoip:country=RU",
						"'Webhook:send=ES'",
					},
				},
			},
		}

		newStruct := datamodels.AlertMessageTheHiveAlert{
			Follow: true, //замена
			CommonAlertType: commonalert.CommonAlertType{
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
			},
			Severity: 4, //замена
			//замена
			Tags: []string{
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
			/*map[string]datamodels.CustomerFields{
				//замена
				"first-time": &datamodels.CustomFieldDateType{
					Order: 0,
					Date:  "2024-01-22T15:13:10+03:00",
				},
				"last-time": &datamodels.CustomFieldDateType{
					Order: 0,
					Date:  "2024-01-17T00:18:13+03:00",
				},
			},*/
			Artifacts: []datamodels.AlertArtifact{
				{
					CommonArtifactType: commonalertartifact.CommonArtifactType{
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
					},
					Sighted:          true,                        //замена
					IgnoreSimilarity: false,                       //замена
					UpdatedAt:        "2024-02-01T02:00:47+03:00", //замена
					UpdatedBy:        "grintman@email.net",        //замена
					//замена
					Tags: []string{
						"Sensor:id=\"12345667\"",
						"geoip:country=CH",
						"geoip:country=US",
						"'Webhook:send=ES'",
					},
				},
				{
					CommonArtifactType: commonalertartifact.CommonArtifactType{
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
					},
					Sighted:          false,
					IgnoreSimilarity: true,
					UpdatedAt:        "2024-02-14T29:13:11+03:00",
					UpdatedBy:        "vbbbba@email.net",
					Tags: []string{
						"geoip:country=RU",
						"'Webhook:send=ES'",
					},
				},
			},
		}
		It("Ряд полей в AlertMessageTheHiveAlert должны быть успешно заменены", func() {
			num, err := oldStruct.ReplacingOldValues(newStruct)

			Expect(err).ShouldNot(HaveOccurred())
			//кол-во замененных полей
			Expect(num).Should(Equal(22))

			fmt.Println("---=== AlertMessageTheHiveAlert ===---")
			fmt.Println(oldStruct.ToStringBeautiful(0))

			//Alert
			//меняется
			Expect(oldStruct.GetFollow()).Should(BeTrue())
			//меняется
			Expect(oldStruct.GetTlp()).Should(Equal(uint64(3)))
			//меняется
			Expect(oldStruct.GetSeverity()).Should(Equal(uint64(4)))
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
			Expect(len(oldStruct.GetTags())).Should(Equal(4))

			customFields := oldStruct.GetCustomFields()
			_, _, _, firstTime := customFields["first-time"].Get()
			_, _, _, lastTime := customFields["last-time"].Get()
			//меняется
			Expect(firstTime).Should(Equal("2024-01-22T15:13:10+03:00"))
			//НЕ меняется
			Expect(lastTime).Should(Equal("2024-01-17T00:18:13+03:00"))

			for _, v := range oldStruct.GetArtifacts() {
				if v.GetId() != "~84302220012" {
					continue
				}
				//fmt.Println(v.ToStringBeautiful(1))

				//не должно быть изменено на пустое значение
				Expect(v.GetUnderliningId()).ShouldNot(Equal(""))
				//меняется
				Expect(v.GetIoc()).Should(BeTrue())
				//меняется
				Expect(v.GetSighted()).Should(BeTrue())
				//меняется
				Expect(v.GetIgnoreSimilarity()).ShouldNot(BeTrue())
				//меняется
				Expect(v.GetTlp()).Should(Equal(uint64(3)))
				//меняется
				Expect(v.GetIoc()).Should(BeTrue())
				//меняется
				Expect(v.GetCreatedAt()).Should(Equal("2024-01-27T22:17:17+03:00"))
				//меняется
				Expect(v.GetUpdatedAt()).Should(Equal("2024-02-01T02:00:47+03:00"))
				//НЕ меняется
				Expect(v.GetCreatedBy()).Should(Equal("friman@email.net"))
				//меняется
				Expect(v.GetUpdatedBy()).Should(Equal("grintman@email.net"))
				//меняется
				Expect(v.GetData()).Should(Equal("63.5656 89.1211"))
				//меняется
				Expect(len(v.GetTags())).Should(Equal(5))

				customFields := oldStruct.GetCustomFields()
				_, _, _, firstTime := customFields["first-time"].Get()
				_, _, _, lastTime := customFields["last-time"].Get()
				//меняется
				Expect(firstTime).Should(Equal("2024-01-22T15:13:10+03:00"))
				//НЕ меняется
				Expect(lastTime).Should(Equal("2024-01-17T00:18:13+03:00"))
			}

			Expect(len(oldStruct.GetArtifacts())).Should(Equal(3))
		})
	})
})
