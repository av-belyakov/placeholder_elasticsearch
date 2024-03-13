package datamodels

import "reflect"

//***************** TtpMessage ********************

// ReplacingOldValues заменяет старые значения структуры TtpMessage
// новыми значениями. Изменяемые поля:
// OccurDate - дата возникновения
// UnderliningCreatedAt - время создания
// UnderliningId - уникальный идентификатор
// UnderliningCreatedBy - кем создан
// PatternId - уникальный идентификатор шаблона
// Tactic - тактика
// ExtraData - дополнительные данные
func (a *TtpMessage) ReplacingOldValues(element TtpMessage) int {
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

			// для обработки поля "ExtraData"
			if typeOfCurrentStruct.Field(i).Name == "ExtraData" {
				countReplacingFields += a.ExtraData.ReplacingOldValues(element.GetExtraData())
				continue
			}

			//обработка полей содержащихся в TtpMessage
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

					//не обновлять значение если оно соответствует пустой дате
					if str == "1970-01-01T00:00:00+00:00" {
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

//***************** ExtraDataTtpMessage ********************

// ReplacingOldValues заменяет старые значения структуры ExtraDataTtpMessage
// новыми значениями. Изменяемые поля:
// Pattern - шаблон
// PatternParent - родительский шаблон
func (ed *ExtraDataTtpMessage) ReplacingOldValues(element ExtraDataTtpMessage) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(ed).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(element)
	typeOfNewStruct := newStruct.Type()

	for i := 0; i < currentStruct.NumField(); i++ {
		for j := 0; j < newStruct.NumField(); j++ {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			// для обработки поля "Pattern"
			if typeOfCurrentStruct.Field(i).Name == "Pattern" {
				countReplacingFields += ed.Pattern.ReplacingOldValues(element.GetPattern())
				continue
			}

			// для обработки поля "PatternParent"
			if typeOfCurrentStruct.Field(i).Name == "PatternParent" {
				countReplacingFields += ed.PatternParent.ReplacingOldValues(element.GetPatternParent())
				continue
			}
		}
	}

	return countReplacingFields
}

//***************** PatternExtraData ********************

// ReplacingOldValues заменяет старые значения структуры PatternExtraData
// новыми значениями. Изменяемые поля:
// RemoteSupport - удаленная поддержка
// Revoked - аннулированный
// UnderliningCreatedAt - время создания
// UnderliningCreatedBy - кем создан
// UnderliningId - уникальный идентификатор
// UnderliningType - тип
// DataSources - источники данных
// DefenseBypassed - чем выполнен обход защиты
// Description - описание
// ExtraData - дополнительные данные
// Name - наименование
// PatternId - уникальный идентификатор шаблона
// PatternType - тип шаблона
// PermissionsRequired - требуемые разрешения
// Platforms - список платформ
// SystemRequirements - системные требования
// Tactics - список тактик
// URL - URL
// Version - версия
func (ped *PatternExtraData) ReplacingOldValues(element PatternExtraData) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(ped).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(element)
	typeOfNewStruct := newStruct.Type()

	for i := 0; i < currentStruct.NumField(); i++ {
		for j := 0; j < newStruct.NumField(); j++ {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			// для обработки поля "Platforms"
			//*****************************
			if typeOfCurrentStruct.Field(i).Name == "Platforms" {
				if list, ok := replacingSliceString(currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			// для обработки поля "PermissionsRequired"
			//*****************************
			if typeOfCurrentStruct.Field(i).Name == "PermissionsRequired" {
				if list, ok := replacingSliceString(currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			// для обработки поля "DataSources"
			//*****************************
			if typeOfCurrentStruct.Field(i).Name == "DataSources" {
				if list, ok := replacingSliceString(currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			// для обработки поля "Tactics"
			//*****************************
			if typeOfCurrentStruct.Field(i).Name == "Tactics" {
				if list, ok := replacingSliceString(currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			//обработка полей содержащихся в PatternExtraData
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

					//не обновлять значение если оно соответствует пустой дате
					if str == "1970-01-01T00:00:00+00:00" {
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

//***************** AttachmentData ********************

// ReplacingOldValues заменяет старые значения структуры AttachmentData
// новыми значениями. Изменяемые поля:
// Size - размер
// Id - идентификатор
// Name - наименование
// ContentType - тип контента
// Hashes - список хешей
func (a *AttachmentData) ReplacingOldValues(element AttachmentData) int {
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

			// для обработки поля "Hashes"
			//**************************
			if typeOfCurrentStruct.Field(i).Name == "Hashes" {
				if list, ok := replacingSliceString(currentStruct.Field(i), newStruct.Field(j)); ok {
					currentStruct.Field(i).Set(list)
					countReplacingFields++
				}

				continue
			}

			//обработка полей содержащихся в ObservableMessageEs
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

	return countReplacingFields
}

//***************** ReportTaxonomies ********************

// ReplacingOldValues заменяет старые значения структуры ReportTaxonomies
// новыми значениями. Изменяемые поля:
// Taxonomies
func (tr *ReportTaxonomies) ReplacingOldValues(element ReportTaxonomies) (int, error) {
	var (
		err                  error
		countReplacingFields int
	)

	//На мой взгляд в данном случае taxonomy.ReplacingOldValues
	//здесь не нужна, так как не понятно по каким полям отслеживать
	//уникальность объекта
	tr.Taxonomies = append(tr.Taxonomies, element.Taxonomies...)

	//currentStruct := reflect.ValueOf(tr).Elem()
	//typeOfCurrentStruct := currentStruct.Type()

	//newStruct := reflect.ValueOf(element)
	//typeOfNewStruct := newStruct.Type()

	//for i := 0; i < currentStruct.NumField(); i++ {
	//	for j := 0; j < newStruct.NumField(); j++ {
	//		if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
	//			continue
	//		}
	//	}
	//}

	return countReplacingFields, err
}

//***************** Taxonomy ********************

// ReplacingOldValues заменяет старые значения структуры Taxonomy
// новыми значениями. Изменяемые поля:
// Level
// Namespace
// Predicate
// Value
func (t *Taxonomy) ReplacingOldValues(element Taxonomy) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(t).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(element)
	typeOfNewStruct := newStruct.Type()

	for i := 0; i < currentStruct.NumField(); i++ {
		for j := 0; j < newStruct.NumField(); j++ {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
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
