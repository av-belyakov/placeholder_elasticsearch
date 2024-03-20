package datamodels

// ReplacingOldValues заменяет старые значения структуры TtpsMessageEs
// новыми значениями. Изменяемые поля:
// Ttp
func (ttp *TtpsMessageTheHive) ReplacingOldValues(element TtpsMessageTheHive) int {
	modifiedTtp, num := comparisonListsTtp(ttp.Ttp, element.Ttp)
	ttp.SetTtps(modifiedTtp)

	return num
}
