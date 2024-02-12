package datamodels

import (
	"fmt"
	"reflect"

	"placeholder_elasticsearch/supportingfunctions"
)

// ReplacingOldValues заменяет старые значения структуры EventMessageTheHiveAlert
// новыми значениями. Изменяемые поля:
//
//	Base           bool              `json:"base" bson:"base"`
//	StartDate      string            `json:"startDate" bson:"startDate"` //в формате RFC3339
//	RootId         string            `json:"rootId" bson:"rootId"`
//	ObjectId       string            `json:"objectId" bson:"objectId"`
//	ObjectType     string            `json:"objectType" bson:"objectType"`
//	Organisation   string            `json:"organisation" bson:"organisation"`
//	OrganisationId string            `json:"organisationId" bson:"organisationId"`
//	Operation      string            `json:"operation" bson:"operation"`
//	RequestId      string            `json:"requestId" bson:"requestId"`
func (e *EventMessageTheHiveAlert) ReplacingOldValues(element EventMessageTheHiveAlert) (int, error) {
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
				if !currentStruct.Field(i).Equal(newStruct.Field(j)) {
					curRootId := currentStruct.Field(i).String()
					newRootId := newStruct.Field(i).String()
					err = fmt.Errorf("the values of the 'rootId' field in the compared objects do not match, current rootId = '%s', new rootId = '%s'", curRootId, newRootId)

					break DONE
				}
			}

			//для обработки поля "Details"
			if typeOfCurrentStruct.Field(i).Name == "Details" {
				countReplacingFields += e.Details.ReplacingOldValues(element.GetDetails())

				continue
			}

			//для обработки поля "Object"
			if typeOfCurrentStruct.Field(i).Name == "Object" {
				countReplacingFields += e.Object.ReplacingOldValues(element.GetObject())

				continue
			}

			if !currentStruct.Field(i).Equal(newStruct.Field(j)) {
				if !currentStruct.Field(i).CanSet() {
					continue
				}

				currentStruct.Field(i).Set(newStruct.Field(j))
				countReplacingFields++
			}
		}
	}

	return countReplacingFields, err
}

// ReplacingOldValues заменяет старые значения структуры EventAlertDetails
// новыми значениями. Изменяемые поля:
// SourceRef - ссылка
// Title - заголовок
// Description - описание
// Tags - список тегов
func (d *EventAlertDetails) ReplacingOldValues(element EventAlertDetails) int {
	var countReplacingFields int

	currentStruct := reflect.ValueOf(d).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(element)
	typeOfNewStruct := newStruct.Type()

	for i := 0; i < currentStruct.NumField(); i++ {
		for j := 0; j < newStruct.NumField(); j++ {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			// для обработки поля "Tags"
			if typeOfCurrentStruct.Field(i).Name == "Tags" {
				if reflect.DeepEqual(currentStruct.Field(i), newStruct.Field(j)) {
					continue
				}

				currentTags, okCurr := currentStruct.Field(i).Interface().([]string)
				newTags, okNew := newStruct.Field(j).Interface().([]string)
				if !okCurr || !okNew {
					continue
				}

				if !currentStruct.Field(i).CanSet() {
					continue
				}

				list := reflect.ValueOf(supportingfunctions.SliceJoinUniq[string](currentTags, newTags))
				currentStruct.Field(i).Set(list)
				countReplacingFields++

				continue
			}

			if !currentStruct.Field(i).Equal(newStruct.Field(j)) {
				if !currentStruct.Field(i).CanSet() {
					continue
				}

				currentStruct.Field(i).Set(newStruct.Field(j))
				countReplacingFields++
			}
		}
	}

	fmt.Println("COUNT Details:", countReplacingFields)

	return countReplacingFields
}

// ReplacingOldValues заменяет старые значения структуры EventAlertObject
// новыми значениями. Изменяемые поля:
// Follow - следовать
// Tlp - tlp
// Pap - pap
// Severity - строгость
// UnderliningId - уникальный идентификатор
// Id - уникальный идентификатор
// CreatedBy - кем создан
// UpdatedBy - кем обновлен
// CreatedAt - дата создания (формат RFC3339)
// UpdatedAt - дата обновления (формат RFC3339)
// UnderliningType - тип
// Title - заголовок
// Description - описание
// Tags - список тегов
// Status - статус
// CustomFields - настраиваемые поля
// Date - дата (формат RFC3339)
// Type - тип
// Source - источник
// SourceRef - ссылка на источник
// Case - кейс
// CaseTemplate - шаблон обращения
// ObjectType - тип объекта
// Tags - теги
// CustomFields - настраиваемые поля
func (o *EventAlertObject) ReplacingOldValues(element EventAlertObject) int {
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

			// для обработки поля "Tags"
			if typeOfCurrentStruct.Field(i).Name == "Tags" {
				if reflect.DeepEqual(currentStruct.Field(i), newStruct.Field(j)) {
					continue
				}

				currentTags, okCurr := currentStruct.Field(i).Interface().([]string)
				newTags, okNew := newStruct.Field(j).Interface().([]string)
				if !okCurr || !okNew {
					continue
				}

				list := reflect.ValueOf(supportingfunctions.SliceJoinUniq[string](currentTags, newTags))
				currentStruct.Field(i).Set(list)
				countReplacingFields++

				continue
			}

			//для обработки поля "CustomFields"
			if typeOfCurrentStruct.Field(i).Name == "CustomFields" {
				currentCustomFields, okCurr := currentStruct.Field(i).Interface().(map[string]CustomerFields)
				newCustomFields, okNew := newStruct.Field(j).Interface().(map[string]CustomerFields)
				if !okCurr || !okNew {
					continue
				}

				for k, v := range newCustomFields {
					currentCustomFields[k] = v
				}

				o.SetValueCustomFields(currentCustomFields)
				countReplacingFields++

				continue
			}

			if !currentStruct.Field(i).Equal(newStruct.Field(j)) {
				if !currentStruct.Field(i).CanSet() {
					continue
				}

				currentStruct.Field(i).Set(newStruct.Field(j))
				countReplacingFields++
			}
		}
	}

	return countReplacingFields
}
