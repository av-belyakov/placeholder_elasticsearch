package testhandlerobservables_test

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

var _ = Describe("Testobservablesreports", Ordered, func() {
	var (
		event        datamodels.EventMessageTheHive = datamodels.EventMessageTheHive{}
		eventObject  datamodels.EventObject         = datamodels.EventObject{}
		eventDetails datamodels.EventDetails        = datamodels.EventDetails{}

		eventObjectCustomFields  map[string]datamodels.CustomerFields = make(map[string]datamodels.CustomerFields)
		eventDetailsCustomFields map[string]datamodels.CustomerFields = make(map[string]datamodels.CustomerFields)

		listHandlerEvent                    map[string][]func(interface{})
		listHandlerEventDetails             map[string][]func(interface{})
		listHandlerEventDetailsCustomFields map[string][]func(interface{})

		listHandlerEventObject             map[string][]func(interface{})
		listHandlerEventObjectCustomFields map[string][]func(interface{})

		so                     *listhandlerthehivejson.SupportiveObservables
		sor                    *listhandlerthehivejson.SupportiveObservablesReports
		listHandlerObservables map[string][]func(interface{})
	)

	BeforeAll(func() {
		so = listhandlerthehivejson.NewSupportiveObservables()
		sor = listhandlerthehivejson.NewSupportiveObservablesReports()
		listHandlerObservables = listhandlerthehivejson.NewListHandlerObservablesElement(so)

		// ------ EVENT ------
		listHandlerEvent = listhandlerthehivejson.NewListHandlerEventElement(&event)

		// ------ EVENT OBJECT ------
		listHandlerEventObject = listhandlerthehivejson.NewListHandlerEventObjectElement(&eventObject)

		// ------ EVENT OBJECT CUSTOMFIELDS ------
		listHandlerEventObjectCustomFields = listhandlerthehivejson.NewListHandlerEventObjectCustomFieldsElement(eventObjectCustomFields)

		// ------ EVENT DETAILS ------
		listHandlerEventDetails = listhandlerthehivejson.NewListHandlerEventDetailsElement(&eventDetails)

		// ------ EVENT DETAILS CUSTOMFIELDS ------
		listHandlerEventDetailsCustomFields = listhandlerthehivejson.NewListHandlerEventDetailsCustomFieldsElement(eventDetailsCustomFields)
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
			Expect(observables[0].GetUnderliningCreatedAt()).Should(Equal(uint64(1690968664227)))
			Expect(observables[0].GetUnderliningCreatedBy()).Should(Equal("uds@crimea-rcm"))
			Expect(observables[0].GetUnderliningUpdatedAt()).Should(Equal(uint64(1704977151860)))
			Expect(observables[0].GetSighted()).ShouldNot(BeTrue())
			Expect(observables[0].GetStartDate()).Should(Equal(uint64(1690968664227)))
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
			Expect(observables[1].GetUnderliningCreatedAt()).Should(Equal(uint64(1705049449272)))
			Expect(observables[1].GetUnderliningCreatedBy()).Should(Equal("zhmurchuk@mail.rcm"))
			Expect(observables[1].GetUnderliningUpdatedAt()).Should(Equal(uint64(1705049448855)))
			Expect(observables[1].GetSighted()).ShouldNot(BeTrue())
			Expect(observables[1].GetStartDate()).Should(Equal(uint64(1705049449272)))
			Expect(len(observables[1].GetTags())).Should(Equal(2))

			for k, v := range observables[1].GetTags() {
				if k == 0 {
					Expect(v).Should(Equal("misp:Network activity=\"ip-src\""))
				}

				if k == 1 {
					Expect(v).Should(Equal("b2m:ip_ext=206.189.15.25"))
				}
			}

			//Test observables.reports

		})
	})

	Context("Тест 3. Проверка заполнения объекта для хранения events", func() {
		var (
			anyEvent        *datamodels.EventMessageTheHive
			anyEventObject  *datamodels.EventObject
			anyEventDetails *datamodels.EventDetails
		)

		BeforeAll(func() {
			sendData := make(chan struct {
				num   int
				name  string
				value interface{}
			})

			go func() {
				for data := range sendData {
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

					//event.details.customFields element
					if lf, ok := listHandlerEventDetailsCustomFields[data.name]; ok {
						//fmt.Printf("%d. (event.details.customFields) v.ElemName: %s, v.ElemValue: '%v'\n", data.num, data.name, data.value)
						for _, f := range lf {
							f(data.value)
						}
					}
				}
			}()

			for k, v := range testing.GetEventOne() {
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
			eventDetails.SetValueCustomFields(eventDetailsCustomFields)

			event.SetValueObject(eventObject)
			event.SetValueDetails(eventDetails)
		})

		It("Объект Event должен быть успешно заполнен", func() {
			//event
			anyEvent = event.Get()

			Expect(anyEvent.GetObjectId()).Should(Equal("~419385432"))
			Expect(anyEvent.GetObjectType()).Should(Equal("case"))
			Expect(anyEvent.GetBase()).Should(BeTrue())
			Expect(anyEvent.GetStartDate()).Should(Equal(uint64(1705061267325)))
			Expect(anyEvent.GetOrganisation()).Should(Equal("RCM"))
			Expect(anyEvent.GetOperation()).Should(Equal("update"))
			Expect(anyEvent.GetRootId()).Should(Equal("~419385432"))
			Expect(anyEvent.GetRequestId()).Should(Equal("019f0dbc0ab90bbe:-58339429:18b66b86afa:-8000:780802"))
			Expect(anyEvent.GetOrganisationId()).Should(Equal("~20488"))
		})

		It("Объект Event.object должен быть успешно заполнен", func() {
			//event.object
			anyEventObject = eventObject.Get()

			Expect(anyEventObject.GetId()).Should(Equal("~85771464712"))
			Expect(len(anyEventObject.GetTags())).Should(Equal(8))
			Expect(anyEventObject.GetSummary()).Should(Equal("trigger"))
			Expect(anyEventObject.GetOwner()).Should(Equal("b.polyakov@cloud.gcm"))
			Expect(anyEventObject.GetUpdatedBy()).Should(Equal("d.sergeev@cloud.gcm"))
			Expect(anyEventObject.GetTitle()).Should(Equal("ПНПО \"Ammyy Admin\""))
			Expect(anyEventObject.GetSeverity()).Should(Equal(uint64(2)))
			Expect(anyEventObject.GetEndDate()).Should(Equal(uint64(0)))
			Expect(anyEventObject.GetCaseId()).Should(Equal(uint64(34411)))
			Expect(anyEventObject.GetDescription()).Should(Equal("Атака направлена **наружу**"))
			Expect(anyEventObject.GetFlag()).Should(BeTrue())
			Expect(anyEventObject.GetTlp()).Should(Equal(uint64(3)))
			Expect(anyEventObject.GetPap()).Should(Equal(uint64(5)))
			Expect(anyEventObject.GetUnderliningType()).Should(Equal("case"))
			Expect(anyEventObject.GetUnderliningId()).Should(Equal("~85771464712"))
			Expect(anyEventObject.GetStartDate()).Should(Equal(uint64(1704980275686)))
			Expect(anyEventObject.GetImpactStatus()).Should(Equal("With Impact"))
			Expect(anyEventObject.GetStatus()).Should(Equal("Open"))
			Expect(anyEventObject.GetCreatedBy()).Should(Equal("b.polyakov@cloud.gcm"))
			Expect(anyEventObject.GetCreatedAt()).Should(Equal(uint64(1704980275725)))
			Expect(anyEventObject.GetUpdatedAt()).Should(Equal(uint64(1705062426568)))
			Expect(anyEventObject.GetResolutionStatus()).Should(Equal("True Positive"))
		})

		It("Объект Event.object.customFields должен быть успешно заполнен", func() {
			//event.object.customFields
			Expect(len(anyEventObject.GetCustomFields())).Should(Equal(8))
		})

		It("Объект Event.details.customFields должен быть успешно заполнен", func() {
			//event.details
			anyEventDetails = eventDetails.Get()

			Expect(anyEventDetails.GetSummary()).Should(Equal("FP (Обращение на getz-club.ru) с 185.4.65.151"))
			Expect(anyEventDetails.GetStatus()).Should(Equal("Resolved"))
			Expect(anyEventDetails.GetImpactStatus()).Should(Equal("NotApplicable"))
			Expect(anyEventDetails.GetEndDate()).Should(Equal(uint64(1705063488183)))
			Expect(anyEventDetails.GetResolutionStatus()).Should(Equal("FalsePositive"))
			//Expect(anyEventDetails.).Should(Equal(""))
		})

		It("Объект Event.details.customFields должен быть успешно заполнен", func() {
			//event.details.customFields
			Expect(len(anyEventDetails.GetCustomFields())).Should(Equal(6))
		})

		It("Должен быть выведен полный объект Event", func() {
			fmt.Printf("---=== EVENT OBJECT ===---\n%v\v", anyEvent.Get())

			Expect(true).Should(BeTrue())
		})
	})
})
