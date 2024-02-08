package testhandlereventforalert_test

import (
	"fmt"
	"reflect"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/listhandlerthehivejson"
	testing "placeholder_elasticsearch/test"
)

var _ = Describe("Testeventforalert", Ordered, func() {
	var (
		event        datamodels.EventMessageTheHiveAlert = datamodels.EventMessageTheHiveAlert{}
		eventObject  datamodels.EventAlertObject         = datamodels.EventAlertObject{}
		eventDetails datamodels.EventAlertDetails        = datamodels.EventAlertDetails{}

		alert datamodels.AlertMessageTheHiveAlert = datamodels.AlertMessageTheHiveAlert{}

		sa listhandlerthehivejson.SupportiveAlertArtifacts = *listhandlerthehivejson.NewSupportiveAlertArtifacts()

		eventObjectCustomFields map[string]datamodels.CustomerFields = make(map[string]datamodels.CustomerFields)
		alertObjectCustomFields map[string]datamodels.CustomerFields = make(map[string]datamodels.CustomerFields)

		listHandlerEvent        map[string][]func(interface{})
		listHandlerEventDetails map[string][]func(interface{})

		listHandlerEventObject             map[string][]func(interface{})
		listHandlerEventObjectCustomFields map[string][]func(interface{})

		listHandlerAlert             map[string][]func(interface{})
		listHandlerAlertCustomFields map[string][]func(interface{})
		listHandlerAlertArtifacts    map[string][]func(interface{})

		verifiedTheHiveAlert datamodels.VerifiedTheHiveAlert = datamodels.VerifiedTheHiveAlert{}
	)

	BeforeAll(func() {
		// ------ EVENT ------
		listHandlerEvent = listhandlerthehivejson.NewListHandlerEventAlertElement(&event)

		// ------ EVENT OBJECT ------
		listHandlerEventObject = listhandlerthehivejson.NewListHandlerEventAlertObjectElement(&eventObject)

		// ------ EVENT OBJECT CUSTOMFIELDS ------
		listHandlerEventObjectCustomFields = listhandlerthehivejson.NewListHandlerAlertCustomFieldsElement(eventObjectCustomFields)

		// ------ EVENT DETAILS ------
		listHandlerEventDetails = listhandlerthehivejson.NewListHandlerEventAlertDetailsElement(&eventDetails)

		// ------ ALERT ------
		listHandlerAlert = listhandlerthehivejson.NewListHandlerAlertElement(&alert)

		// ------ ALERT CUSTOMFIELDS ------
		listHandlerAlertCustomFields = listhandlerthehivejson.NewListHandlerAlertCustomFieldsElement(alertObjectCustomFields)

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

			event.SetValueObject(eventObject)
			event.SetValueDetails(eventDetails)

			alert.SetValueCustomFields(alertObjectCustomFields)
			alert.SetValueArtifacts(sa.GetArtifacts())

			verifiedTheHiveAlert.SetID("fhe78f838f88fg488398f8e3")
			verifiedTheHiveAlert.SetElasticsearchID("3883f8f9-839r983hf848g8h-f84")
			verifiedTheHiveAlert.SetSource("GCM")
			verifiedTheHiveAlert.SetCreateTimestatmp("2024-02-06T15:37:52+03:00")
			verifiedTheHiveAlert.SetEvent(event)
			verifiedTheHiveAlert.SetAlert(alert)
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

			fmt.Println("---=== VERIFEDTHEHIVEALERT ===---")
			fmt.Println(verified.ToStringBeautiful(0))

			Expect(verified.GetID()).Should(Equal("fhe78f838f88fg488398f8e3"))
			Expect(verified.GetElasticsearchID()).Should(Equal("3883f8f9-839r983hf848g8h-f84"))
			Expect(verified.GetSource()).Should(Equal("GCM"))
			Expect(verified.GetCreateTimestatmp()).Should(Equal("2024-02-06T15:37:52+03:00"))
		})
	})
})
