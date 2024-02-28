package commonobjectevent

import (
	"fmt"
	"reflect"
)

// ReplacingOldValues заменяет старые значения структуры EventMessageForEsAlertObject
// новыми значениями. Изменяемые поля:
// Flag - флаг
// CaseId - уникальный идентификатор дела
// Severity - строгость
// Tlp - tlp
// Pap - pap
// StartDate - начальная дата
// EndDate - конечная дата
// CreatedAt - дата создания
// UpdatedAt - дата обновления
// UnderliningId - уникальный идентификатор
// Id - уникальный идентификатор
// CreatedBy - кем создан
// UpdatedBy - кем обновлен
// UnderliningType - тип
// Title - заголовок
// Description - описание
// ImpactStatus - краткое описание воздействия
// ResolutionStatus - статус разрешения
// Status - статус
// Summary - резюме
// Owner - владелец
func (e *CommonEventCaseObject) ReplacingOldValues(element CommonEventCaseObject) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(e).Elem()
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

				fmt.Println("func 'CommonEventCaseObject' Field Name:", typeOfCurrentStruct.Field(i).Name, " Value:", newStruct.Field(j))

			}
		}
	}

	fmt.Println("func 'CommonEventCaseObject' countReplacingFields:", countReplacingFields)

	return countReplacingFields
}
