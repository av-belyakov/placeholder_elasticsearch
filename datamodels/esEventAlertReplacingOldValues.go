package datamodels

import (
	"reflect"
)

// ReplacingOldValues заменяет старые значения структуры EventMessageForEsAlert
// новыми значениями. Изменяемые поля:
// CommonEventType
// Details - детальная информация о событии
// Object - объект события
func (e *EventMessageForEsAlert) ReplacingOldValues(element EventMessageForEsAlert) (int, error) {
	var (
		err                  error
		countReplacingFields int
	)

	currentStruct := reflect.ValueOf(e).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(element)
	typeOfNewStruct := newStruct.Type()

DONE:
	for i := 0; i < currentStruct.NumField(); i++ {
		for j := 0; j < newStruct.NumField(); j++ {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			// для обработки общих полей
			//**************************
			if typeOfCurrentStruct.Field(i).Name == "CommonEventType" {
				num, errC := e.CommonEventType.Get().ReplacingOldValues(*element.CommonEventType.Get())
				if errC != nil {
					err = errC

					break DONE
				}

				countReplacingFields += num
			}

			//для обработки поля "Details"
			if typeOfCurrentStruct.Field(i).Name == "Details" {
				countReplacingFields += e.Details.ReplacingOldValues(element.GetDetails())

				continue
			}

			//для обработки поля "Object"
			if typeOfCurrentStruct.Field(i).Name == "Object" {
				countReplacingFields += e.Object.ReplacingOldValues(element.GetObject())

				continue
			}
		}
	}

	return countReplacingFields, err
}

// ReplacingOldValues заменяет старые значения структуры EventMessageForEsAlertDetails
// новыми значениями. Изменяемые поля:
// SourceRef - ссылка
// Title - заголовок
// Description - описание
// GeoIp - сюда помещаются теги относящиеся к географическому позиционированию
// SensorId - сюда помещаются теги относящиеся к номерам сигнатур
// Reasons - сюда помещаются теги относящиеся к причине возникновения события
func (d *EventMessageForEsAlertDetails) ReplacingOldValues(element EventMessageForEsAlertDetails) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(d).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(element)
	typeOfNewStruct := newStruct.Type()

	for i := 0; i < currentStruct.NumField(); i++ {
		for j := 0; j < newStruct.NumField(); j++ {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			// для обработки поля "Tags"
			//**************************
			if typeOfCurrentStruct.Field(i).Name == "Tags" {
				newTags, okNew := newStruct.Field(j).Interface().(map[string][]string)
				if !okNew {
					continue
				}

				for key, value := range newTags {
					for _, v := range value {
						if d.SetValueTags(key, v) {
							countReplacingFields++
						}
					}
				}

				continue
			}

			// для обработки поля "TagsAll"
			//**************************
			if typeOfCurrentStruct.Field(i).Name == "TagsAll" {
				if list, ok := replacingSlice(currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			if !currentStruct.Field(i).Equal(newStruct.Field(j)) {
				if !currentStruct.Field(i).CanSet() {
					continue
				}

				if str, ok := newStruct.Field(j).Interface().(string); ok {
					//не обновлять текущие значения новыми пустыми значениями
					if str == "" {
						continue
					}
				}

				currentStruct.Field(i).Set(newStruct.Field(j))
				countReplacingFields++
			}
		}
	}

	return countReplacingFields
}

// ReplacingOldValues заменяет старые значения структуры EventMessageForEsAlertObject
// новыми значениями. Изменяемые поля:
// Tags - список тегов
// TagsAll - список всех тегов
// CustomFields - настраиваемые поля
func (o *EventMessageForEsAlertObject) ReplacingOldValues(element EventMessageForEsAlertObject) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(o).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(element)
	typeOfNewStruct := newStruct.Type()

	for i := 0; i < currentStruct.NumField(); i++ {
		for j := 0; j < newStruct.NumField(); j++ {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			// для обработки общих полей
			//**************************
			if typeOfCurrentStruct.Field(i).Name == "CommonEventAlertObject" {
				countReplacingFields += o.CommonEventAlertObject.Get().ReplacingOldValues(*element.CommonEventAlertObject.Get())
			}

			// для обработки поля "Tags"
			//**************************
			if typeOfCurrentStruct.Field(i).Name == "Tags" {
				newTags, okNew := newStruct.Field(j).Interface().(map[string][]string)
				if !okNew {
					continue
				}

				for key, value := range newTags {
					for _, v := range value {
						if o.SetValueTags(key, v) {
							countReplacingFields++
						}
					}
				}

				continue
			}

			// для обработки поля "TagsAll"
			//*****************************
			if typeOfCurrentStruct.Field(i).Name == "TagsAll" {
				if list, ok := replacingSlice(currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			//для обработки поля "CustomFields"
			//*********************************
			if typeOfCurrentStruct.Field(i).Name == "CustomFields" {
				currentCustomFields, okCurr := currentStruct.Field(i).Interface().(CustomFields)
				newCustomFields, okNew := newStruct.Field(j).Interface().(CustomFields)
				if !okCurr || !okNew {
					continue
				}

				for k, v := range newCustomFields {
					currentCustomFields[k] = v
				}

				o.SetValueCustomFields(currentCustomFields)
				countReplacingFields++

				continue
			}
		}
	}

	return countReplacingFields
}
