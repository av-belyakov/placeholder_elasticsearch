package commonobservable

import "reflect"

// ReplacingOldValues заменяет старые значения структуры CommonEventType
// новыми значениями. Изменяемые поля:
// Ioc - индикатор компрометации
// Tlp - tlp
// UnderliningId - уникальный идентификатор
// Id - уникальный идентификатор
// UnderliningType - тип
// CreatedAt - время создания
// CreatedBy - кем создан
// StartDate - дата начала
// Data - данные
// DataType - тип данных
// Message - сообщение
func (o *CommonObservableType) ReplacingOldValues(element CommonObservableType) int {
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
