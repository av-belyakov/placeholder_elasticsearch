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
