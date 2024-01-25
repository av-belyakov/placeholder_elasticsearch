package testhandlerevent_test

import (
	"fmt"
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/listhandlerthehivejson"
	testing "placeholder_elasticsearch/test"
)

var _ = Describe("Testevent", Ordered, func() {
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
	)

	BeforeAll(func() {
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

	Context("Тест 1. Проверка заполнения объекта для хранения events", func() {
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
			Expect(anyEvent.GetStartDate()).Should(Equal("2024-01-12T15:07:47+03:00"))
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
			Expect(anyEventObject.GetEndDate()).Should(Equal("1970-01-01T03:00:00+03:00"))
			Expect(anyEventObject.GetCaseId()).Should(Equal(uint64(34411)))
			Expect(anyEventObject.GetDescription()).Should(Equal("Атака направлена **наружу**"))
			Expect(anyEventObject.GetFlag()).Should(BeTrue())
			Expect(anyEventObject.GetTlp()).Should(Equal(uint64(3)))
			Expect(anyEventObject.GetPap()).Should(Equal(uint64(5)))
			Expect(anyEventObject.GetUnderliningType()).Should(Equal("case"))
			Expect(anyEventObject.GetUnderliningId()).Should(Equal("~85771464712"))
			Expect(anyEventObject.GetStartDate()).Should(Equal("2024-01-11T16:37:55+03:00"))
			Expect(anyEventObject.GetImpactStatus()).Should(Equal("With Impact"))
			Expect(anyEventObject.GetStatus()).Should(Equal("Open"))
			Expect(anyEventObject.GetCreatedBy()).Should(Equal("b.polyakov@cloud.gcm"))
			Expect(anyEventObject.GetCreatedAt()).Should(Equal("2024-01-11T16:37:55+03:00"))
			Expect(anyEventObject.GetUpdatedAt()).Should(Equal("2024-01-12T15:27:06+03:00"))
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
			Expect(anyEventDetails.GetEndDate()).Should(Equal("2024-01-12T15:44:48+03:00"))
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
