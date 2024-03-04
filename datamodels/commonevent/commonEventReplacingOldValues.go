package commonevent

import (
	"fmt"
	"reflect"
)

// ReplacingOldValues заменяет старые значения структуры CommonEventType
// новыми значениями. Изменяемые поля:
// Base - основа
// StartDate - начальная дата
// RootId - главный уникальный идентификатор
// ObjectId - уникальный идентификатор объекта
// ObjectType - тип объекта
// Organisation - наименование организации
// OrganisationId - уникальный идентификатор организации
// Operation - операция
// RequestId - уникальный идентификатор запроса
func (e *CommonEventType) ReplacingOldValues(element CommonEventType) (int, error) {
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

			if typeOfCurrentStruct.Field(i).Name == "RootId" {
				if str, ok := currentStruct.Field(i).Interface().(string); ok {
					if str == "" {
						currentStruct.Field(i).Set(newStruct.Field(j))
						countReplacingFields++

						continue
					}
				}

				if !currentStruct.Field(i).Equal(newStruct.Field(j)) {
					curRootId := currentStruct.Field(i).String()
					newRootId := newStruct.Field(i).String()
					err = fmt.Errorf("the values of the 'rootId' field in the compared objects do not match, current rootId = '%s', new rootId = '%s'", curRootId, newRootId)

					break DONE
				}
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
