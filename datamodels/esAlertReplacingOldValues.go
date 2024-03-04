package datamodels

import (
	"reflect"
)

// ReplacingOldValues заменяет старые значения структуры AlertMessageForEsAlert
// новыми значениями. Изменяемые поля:
// Tags - теги
// TagsAll - все теги
// CustomFields - настраиваемые поля
// Artifacts - артефакты
func (am *AlertMessageForEsAlert) ReplacingOldValues(element AlertMessageForEsAlert) (int, error) {
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
			//**************************
			if typeOfCurrentStruct.Field(i).Name == "Tags" {
				newTags, okNew := newStruct.Field(j).Interface().(map[string][]string)
				if !okNew {
					continue
				}

				for key, value := range newTags {
					for _, v := range value {
						if am.SetValueTags(key, v) {
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
			if typeOfCurrentStruct.Field(i).Name == "Artifacts" {
				newArtifact, okNew := newStruct.Field(j).Interface().(map[string][]ArtifactForEsAlert)
				if !okNew {
					continue
				}

				for key, value := range newArtifact {
					currentArtifacts, ok := am.GetKeyArtifacts(key)
					if !ok {
						am.SetKeyArtifacts(key, value)

						continue
					}

					modifiedArtifacts, num := comparisonListsArtifacts(currentArtifacts, value)
					countReplacingFields += num
					am.SetKeyArtifacts(key, modifiedArtifacts)
				}

				continue
			}
		}
	}

	return countReplacingFields, err
}

// ReplacingOldValues заменяет старые значения структуры ArtifactForEsAlert
// новыми значениями. Изменяемые поля:
// CommonArtifactType
// Tags - список тегов
// TahsAll - все теги
func (a *ArtifactForEsAlert) ReplacingOldValues(element ArtifactForEsAlert) int {
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
						if a.SetValueTags(key, v) {
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
		}
	}

	return countReplacingFields
}

// comparisonListsArtifacts объединяет два списка
func comparisonListsArtifacts(currentArtifacts, newArtifacts []ArtifactForEsAlert) ([]ArtifactForEsAlert, int) {
	var countReplacingFields int

	for _, value := range newArtifacts {
		var isExist bool

		for k, v := range currentArtifacts {
			if value.GetId() == v.GetId() || value.GetUnderliningId() == v.GetUnderliningId() {
				isExist = true
				countReplacingFields += currentArtifacts[k].ReplacingOldValues(value)

				break
			}
		}

		if !isExist {
			currentArtifacts = append(currentArtifacts, value)
		}
	}

	return currentArtifacts, countReplacingFields
}
