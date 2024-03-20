package datamodels

import "reflect"

// ReplacingOldValues заменяет старые значения структуры ObservablesMessageEs
// новыми значениями. Изменяемые поля:
// Observables
func (o *ObservablesMessageEs) ReplacingOldValues(element ObservablesMessageEs) int {
	var countReplacingFields int

	for key, value := range element.Observables {
		currentObservables, ok := o.GetKeyObservables(key)
		if !ok {
			o.SetKeyObservables(key, value)

			continue
		}

		modifiedObservables, num := comparisonListsObservables(currentObservables, value)
		countReplacingFields += num
		o.SetKeyObservables(key, modifiedObservables)
	}

	return countReplacingFields
}

// ReplacingOldValues заменяет старые значения структуры ObservableMessageEs
// новыми значениями. Изменяемые поля:
// SensorId - идентификатор сенсора
// SnortSid - список идентификаторов сигнатур
// Tags - список тегов
// TagsAll - список всех тегов
// Attachment - приложенные данные
// Reports - список отчетов
func (o *ObservableMessageEs) ReplacingOldValues(element ObservableMessageEs) int {
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
			if typeOfCurrentStruct.Field(i).Name == "CommonObservableType" {
				countReplacingFields += o.CommonObservableType.Get().ReplacingOldValues(*element.CommonObservableType.Get())
			}

			// для обработки поля "SnortSid"
			//******************************
			if typeOfCurrentStruct.Field(i).Name == "SnortSid" {
				if list, ok := replacingSliceString(currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

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
				if list, ok := replacingSliceString(currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			// для обработки поля "Attachment"
			//********************************
			if typeOfCurrentStruct.Field(i).Name == "Attachment" {
				countReplacingFields += o.Attachment.ReplacingOldValues(*element.GetAttachment())
				continue
			}

			// для обработки поля "Reports"
			//*****************************
			if typeOfCurrentStruct.Field(i).Name == "Reports" {
				/*newRT, okNew := newStruct.Field(j).Interface().(map[string]ReportTaxonomies)
				if !okNew {
					continue
				}

				for key, value := range newRT {
					rt, ok := o.GetTaxonomies(key)
					if !ok {
						o.AddValueReports(key, value)

						continue
					}

					num, _ := rt.ReplacingOldValues(value)
					countReplacingFields += num
				}*/

				continue
			}

			//обработка полей содержащихся в ObservableMessageEs
			//и не относящихся к вышеперечисленым значениям
			//***************************************************
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

// comparisonListsObservables объединяет два списка
func comparisonListsObservables(currentObs, newObs []ObservableMessageEs) ([]ObservableMessageEs, int) {
	var countReplacingFields int

	for _, value := range newObs {
		var isExist bool

		for k, v := range currentObs {
			if value.GetUnderliningId() == v.GetUnderliningId() {
				isExist = true
				countReplacingFields += currentObs[k].ReplacingOldValues(value)

				break
			}
		}

		if !isExist {
			currentObs = append(currentObs, value)
		}
	}

	return currentObs, countReplacingFields
}
