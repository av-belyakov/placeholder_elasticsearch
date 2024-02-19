package datamodels

import "reflect"

// ReplacingOldValues заменяет старые значения структуры AlertMessageForEsAlert
// новыми значениями. Изменяемые поля:
// Tlp - номер группы разделяющие общие цели
// Date - дата (формат RFC3339)
// CreatedAt - дата создания (формат RFC3339)
// UpdatedAt - дата обновления (формат RFC3339)
// UpdatedBy - кем обновлен
// UnderliningId - уникальный идентификатор
// Status - статус
// Type - тип
// UnderliningType - тип
// Description - описание
// CaseTemplate - шаблон обращения
// SourceRef - ссылка на источник
// Tags - теги
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

			// для обработки поля "Tags"
			//**************************
			if typeOfCurrentStruct.Field(i).Name == "Tags" {

				/* !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
				*
				* 	Эта часть еще не проверялась, надо выполнить тесты
				*
				 */

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
				newCustomFields, okNew := newStruct.Field(j).Interface().(map[string][]ArtifactForEsAlert)
				if !okNew {
					continue
				}

				for key, value := range newCustomFields {
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

// ReplacingOldValues заменяет старые значения структуры ArtifactForEsAlert
// новыми значениями. Изменяемые поля:
// Ioc - индикатор компрометации
// Sighted - видящий
// IgnoreSimilarity - игнорировать похожие
// Tlp - tlp
// UnderliningId - уникальный идентификатор
// Id - уникальный идентификатор
// UnderliningType - тип
// CreatedAt - время создания
// CreatedBy - кем создан
// StartDate - дата начала
// UpdatedAt - время обновления
// UpdatedBy - кем обновлен
// Data - данные
// DataType - тип данных
// Message - сообщение
// Tags - список тегов
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

			// для обработки поля "Tags"
			//**************************
			if typeOfCurrentStruct.Field(i).Name == "Tags" {

				/* !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
				*
				* 	Эта часть еще не проверялась, надо выполнить тесты
				*
				 */

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

// comparisonListsArtifacts объединяет два списка
func comparisonListsArtifacts(currentArtifacts, newArtifacts []ArtifactForEsAlert) (modifay []ArtifactForEsAlert, countReplacingFields int) {
	modifay = make([]ArtifactForEsAlert, len(currentArtifacts))

	for _, value := range newArtifacts {
		var isExist bool

		for k, v := range currentArtifacts {
			if value.GetId() == v.GetId() || value.GetUnderliningId() == v.GetUnderliningId() {
				isExist = true
				countReplacingFields += currentArtifacts[k].ReplacingOldValues(value)
				modifay = append(modifay, currentArtifacts[k])

				break
			}
		}

		if !isExist {
			modifay = append(modifay, value)
		}
	}

	return modifay, countReplacingFields
}
