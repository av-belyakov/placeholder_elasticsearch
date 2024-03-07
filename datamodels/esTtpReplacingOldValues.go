package datamodels

// ReplacingOldValues заменяет старые значения структуры TtpsMessageEs
// новыми значениями. Изменяемые поля:
// Ttp
func (ttp *TtpsMessageEs) ReplacingOldValues(element TtpsMessageEs) int {
	var countReplacingFields int

	for key, value := range element.Ttp {
		currentTtp, ok := ttp.GetKeyTtp(key)
		if !ok {
			ttp.SetKeyTtp(key, value)

			continue
		}

		modifiedTtp, num := comparisonListsTtp(currentTtp, value)
		countReplacingFields += num
		ttp.SetKeyTtp(key, modifiedTtp)
	}

	return countReplacingFields
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
