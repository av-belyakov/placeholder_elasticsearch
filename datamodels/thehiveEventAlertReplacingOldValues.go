package datamodels

import (
	"reflect"
)

// ReplacingOldValues заменяет старые значения структуры EventMessageTheHiveAlert
// новыми значениями. Изменяемые поля:
// CommonEventType
// Details - детальная информация о событии
// Object - объект события
func (e *EventMessageTheHiveAlert) ReplacingOldValues(element EventMessageTheHiveAlert) (int, error) {
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
				num, errC := e.CommonEventType.ReplacingOldValues(*element.CommonEventType.Get())
				if errC != nil {
					err = errC

					break DONE
				}

				countReplacingFields += num
			}

			//для обработки поля "Details"
			//****************************
			if typeOfCurrentStruct.Field(i).Name == "Details" {
				countReplacingFields += e.Details.ReplacingOldValues(element.GetDetails())

				continue
			}

			//для обработки поля "Object"
			//***************************
			if typeOfCurrentStruct.Field(i).Name == "Object" {
				countReplacingFields += e.Object.ReplacingOldValues(element.GetObject())

				continue
			}
		}
	}

	return countReplacingFields, err
}

// ReplacingOldValues заменяет старые значения структуры EventAlertDetails
// новыми значениями. Изменяемые поля:
// SourceRef - ссылка
// Title - заголовок
// Description - описание
// Tags - список тегов
func (d *EventAlertDetails) ReplacingOldValues(element EventAlertDetails) int {
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
				if list, ok := replacingSlice[string](currentStruct.Field(i), newStruct.Field(j)); ok {
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

// ReplacingOldValues заменяет старые значения структуры EventAlertObject
// новыми значениями. Изменяемые поля:
// Follow - следовать
// Pap - pap
// Severity - строгость
// Tags - теги
// CustomFields - настраиваемые поля
func (o *EventAlertObject) ReplacingOldValues(element EventAlertObject) int {
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
				countReplacingFields += o.CommonEventAlertObject.ReplacingOldValues(*element.CommonEventAlertObject.Get())

				continue
			}

			// для обработки поля "Tags"
			//**************************
			if typeOfCurrentStruct.Field(i).Name == "Tags" {
				if list, ok := replacingSlice[string](currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			//для обработки поля "CustomFields"
			//**************************
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

			//обработка полей содержащихся в EventAlertObject
			//и не относящихся к вышеперечисленым значениям
			//***********************************************
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
