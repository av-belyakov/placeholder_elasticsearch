package datamodels

import (
	"reflect"
)

// ReplacingOldValues заменяет старые значения структуры AlertMessageTheHiveAlert
// новыми значениями. Изменяемые поля:
// Follow - следовать
// Severity - строгость
// Tags - теги
// CustomFields - настраиваемые поля
// Artifacts - артифакты
func (am *AlertMessageTheHiveAlert) ReplacingOldValues(element AlertMessageTheHiveAlert) (int, error) {
	var (
		err                  error
		countReplacingFields int
	)

	currentStruct := reflect.ValueOf(am).Elem()
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
			if typeOfCurrentStruct.Field(i).Name == "CommonAlertType" {
				countReplacingFields += am.CommonAlertType.ReplacingOldValues(*element.CommonAlertType.Get())

				continue
			}

			// для обработки поля "Tags"
			if typeOfCurrentStruct.Field(i).Name == "Tags" {
				if list, ok := replacingSliceString(currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			// для обработки поля "CustomFields"
			if typeOfCurrentStruct.Field(i).Name == "CustomFields" {
				currentCustomFields, okCurr := currentStruct.Field(i).Interface().(CustomFields)
				newCustomFields, okNew := newStruct.Field(j).Interface().(CustomFields)
				if !okCurr || !okNew {
					continue
				}

				for k, v := range newCustomFields {
					currentCustomFields[k] = v
				}

				am.SetValueCustomFields(currentCustomFields)
				countReplacingFields++

				continue
			}

			// для обработки поля "Artifacts"
			//Значения данного поля обновляются если есть совпадение в
			//полях 'id' или '_id' между новым артефактом и тем который
			//уже имеется.
			if typeOfCurrentStruct.Field(i).Name == "Artifacts" {
				currentCustomFields, okCurr := currentStruct.Field(i).Interface().([]AlertArtifact)
				newCustomFields, okNew := newStruct.Field(j).Interface().([]AlertArtifact)
				if !okCurr || !okNew {
					continue
				}

				for _, v := range newCustomFields {
					var isExist bool

					for key, value := range currentCustomFields {
						if v.GetId() == value.GetId() || v.GetUnderliningId() == value.GetUnderliningId() {
							countReplacingFields += currentCustomFields[key].ReplacingOldValues(*v.Get())
							isExist = true

							break
						}
					}

					if !isExist {
						currentCustomFields = append(currentCustomFields, v)
					}

					am.SetValueArtifacts(currentCustomFields)
				}

				continue
			}

			//обработка полей содержащихся в AlertMessageTheHiveAlert
			//и не относящихся к вышеперечисленым значениям
			//*******************************************************
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

	return countReplacingFields, err
}

// ReplacingOldValues заменяет старые значения структуры AlertArtifact
// новыми значениями. Изменяемые поля:
// Sighted - видящий
// IgnoreSimilarity - игнорировать похожие
// CommonArtifactType
// UpdatedAt - время обновления
// UpdatedBy - кем обновлен
// Tags - список тегов
func (a *AlertArtifact) ReplacingOldValues(element AlertArtifact) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(a).Elem()
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
			if typeOfCurrentStruct.Field(i).Name == "CommonArtifactType" {
				countReplacingFields += a.CommonArtifactType.ReplacingOldValues(*element.CommonArtifactType.Get())

				continue
			}

			// для обработки поля "Tags"
			if typeOfCurrentStruct.Field(i).Name == "Tags" {
				if list, ok := replacingSliceString(currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

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
