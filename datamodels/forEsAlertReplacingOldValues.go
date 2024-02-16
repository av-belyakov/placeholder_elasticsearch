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

				newCustomFields, okNew := newStruct.Field(j).Interface().(map[string][]string)
				if !okNew {
					continue
				}

				for key, value := range newCustomFields {
					for _, v := range value {
						if o.SetValueTags(key, v) {
							countReplacingFields++
						}
					}
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

			//********************************************************
			//
			// Ныжно выполнить следующие действия:
			// 1. Проверить как отрабатывает обработчик для
			//   if typeOfCurrentStruct.Field(i).Name == "Tags"
			// 2. Написать обработку для ArtifactForEsAlert
			// 3. Написать обработчики из группы listHandler... для
			//   обработки и формирования типов AlertMessageForEsAlert
			//   и EventMessageForEsAlert
			//
			//********************************************************

			/*if typeOfCurrentStruct.Field(i).Name == "Artifacts" {
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
			}*/

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

				newCustomFields, okNew := newStruct.Field(j).Interface().(map[string][]string)
				if !okNew {
					continue
				}

				for key, value := range newCustomFields {
					for _, v := range value {
						if o.SetValueTags(key, v) {
							countReplacingFields++
						}
					}
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
