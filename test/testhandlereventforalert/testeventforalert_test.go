package testhandlereventforalert_test

import (
	"reflect"

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

		eventObjectCustomFields map[string]datamodels.CustomerFields = make(map[string]datamodels.CustomerFields)

		listHandlerEvent        map[string][]func(interface{})
		listHandlerEventDetails map[string][]func(interface{})

		listHandlerEventObject             map[string][]func(interface{})
		listHandlerEventObjectCustomFields map[string][]func(interface{})
	)

	BeforeAll("", func() {
		// ------ EVENT ------
		listHandlerEvent = listhandlerthehivejson.NewListHandlerEventAlertElement(&event)

		// ------ EVENT OBJECT ------
		listHandlerEventObject = listhandlerthehivejson.NewListHandlerEventAlertObjectElement(&eventObject)

		// ------ EVENT OBJECT CUSTOMFIELDS ------
		listHandlerEventObjectCustomFields = listhandlerthehivejson.NewListHandlerEventObjectCustomFieldsElement(eventObjectCustomFields)

		// ------ EVENT DETAILS ------
		listHandlerEventDetails = listhandlerthehivejson.NewListHandlerEventAlertDetailsElement(&eventDetails)
	})

	Context("Тест 1. Проверка заполнения объекта для хранения events", func() {
		var (
			anyEvent        *datamodels.EventMessageTheHiveAlert
			anyEventObject  *datamodels.EventAlertObject
			anyEventDetails *datamodels.EventAlertDetails
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
				}
			}()

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

			close(sendData)

			eventObject.SetValueCustomFields(eventObjectCustomFields)

			event.SetValueObject(eventObject)
			event.SetValueDetails(eventDetails)
		})

		It("Объект Event должен быть успешно заполнен", func() {
			//event
			anyEvent = event.Get()

			Expect(anyEvent.GetBase()).Should(BeTrue())
			Expect(anyEvent.GetStartDate()).Should(Equal("2024-01-12T15:07:47+03:00"))
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
			anyEventDetails = eventDetails.Get()

			Expect(anyEventDetails.GetSourceRef()).Should(Equal("TSK-8MSK-6-ZPM-240206-1215999"))
			Expect(anyEventDetails.GetTitle()).Should(Equal("Зафиксированна КА"))
			Expect(anyEventDetails.GetDescription()).Should(Equal("**Задача переданная из смежной системы"))
			Expect(len(anyEventDetails.GetTags())).Should(Equal(5))
		})

		It("Объект Event.object должен быть успешно заполнен", func() {
			//event.object
			anyEventObject = eventObject.Get()

			Expect(anyEventObject.GetFollow()).Should(BeTrue())
			Expect(anyEventObject.GetSeverity()).Should(Equal(uint64(2)))
			Expect(anyEventObject.GetTlp()).Should(Equal(uint64(3)))
			Expect(anyEventObject.GetPap()).Should(Equal(uint64(5)))
			Expect(anyEventObject.GetUnderliningId()).Should(Equal("~85455464790"))
			Expect(anyEventObject.GetId()).Should(Equal("~85771464712"))
			Expect(anyEventObject.GetCreatedBy()).Should(Equal("v.kovanko@cloud.gcm"))
			Expect(anyEventObject.GetUpdatedBy()).Should(Equal("y.kovalenko@cloud.gcm"))
			Expect(anyEventObject.GetCreatedAt()).Should(Equal("2024-01-11T16:37:55+03:00"))
			Expect(anyEventObject.GetUpdatedAt()).Should(Equal("2024-01-12T15:27:06+03:00"))
			/*
				!!!!!!!!!!!!
				Тут надо доделать
				!!!!!!!!!!!!!!
			*/

			/*
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
				Expect(anyEventObject.GetUnderliningType()).Should(Equal("case"))
				Expect(anyEventObject.GetStartDate()).Should(Equal("2024-01-11T16:37:55+03:00"))
				Expect(anyEventObject.GetImpactStatus()).Should(Equal("With Impact"))
				Expect(anyEventObject.GetStatus()).Should(Equal("Open"))
				Expect(anyEventObject.GetResolutionStatus()).Should(Equal("True Positive"))
			*/
		})

		It("Объект Event.object.customFields должен быть успешно заполнен", func() {
			//event.object.customFields
			//Expect(len(anyEventObject.GetCustomFields())).Should(Equal(9))
		})
	})
})
