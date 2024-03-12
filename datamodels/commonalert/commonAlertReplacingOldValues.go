package commonalert

import (
	"reflect"
)

// ReplacingOldValues заменяет старые значения структуры CommonEventType
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
func (a *CommonAlertType) ReplacingOldValues(element CommonAlertType) int {
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
