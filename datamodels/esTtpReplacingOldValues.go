package datamodels

import "reflect"

// ReplacingOldValues заменяет старые значения структуры TtpsMessageEs
// новыми значениями. Изменяемые поля:
// Ttp
func (ttp *TtpsMessageEs) ReplacingOldValues(element TtpsMessageEs) (int, error) {
	var (
		err                  error
		countReplacingFields int
	)

	currentStruct := reflect.ValueOf(ttp).Elem()
	typeOfCurrentStruct := currentStruct.Type()

	newStruct := reflect.ValueOf(element)
	typeOfNewStruct := newStruct.Type()

	for i := 0; i < currentStruct.NumField(); i++ {
		for j := 0; j < newStruct.NumField(); j++ {
			if typeOfCurrentStruct.Field(i).Name != typeOfNewStruct.Field(j).Name {
				continue
			}

			// для обработки поля "Ttp"
			if typeOfCurrentStruct.Field(i).Name == "Ttp" {
				newTtp, okNew := newStruct.Field(j).Interface().(map[string][]TtpMessage)
				if !okNew {
					continue
				}

				for key, value := range newTtp {
					currentTtp, ok := ttp.GetKeyTtp(key)
					if !ok {
						ttp.SetKeyTtp(key, value)

						continue
					}

					modifiedTtp, num := comparisonListsTtp(currentTtp, value)
					countReplacingFields += num
					ttp.SetKeyTtp(key, modifiedTtp)
				}

				continue
			}
		}
	}

	return countReplacingFields, err
}

// comparisonListsTtp объединяет два списка
func comparisonListsTtp(currentTtp, newTtp []TtpMessage) ([]TtpMessage, int) {
	var countReplacingFields int

	for _, value := range newTtp {
		var isExist bool

		for k, v := range currentTtp {
			if value.GetUnderliningId() == v.GetUnderliningId() {
				isExist = true
				countReplacingFields += currentTtp[k].ReplacingOldValues(value)

				break
			}
		}

		if !isExist {
			currentTtp = append(currentTtp, value)
		}
	}

	return currentTtp, countReplacingFields
}
