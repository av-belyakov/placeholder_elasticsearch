package supportingfunctions

// ConversionAnyToInt преобразовывает любое числовое значение в int или
// возвращает 0 если было получено не число
func ConversionAnyToInt(i interface{}) int {
	if result, ok := i.(int); ok {
		return result
	}

	if result, ok := i.(int16); ok {
		return int(result)
	}

	if result, ok := i.(int32); ok {
		return int(result)
	}

	if result, ok := i.(int64); ok {
		return int(result)
	}

	if result, ok := i.(uint); ok {
		return int(result)
	}

	if result, ok := i.(uint16); ok {
		return int(result)
	}

	if result, ok := i.(uint32); ok {
		return int(result)
	}

	if result, ok := i.(uint64); ok {
		return int(result)
	}

	if result, ok := i.(float32); ok {
		return int(result)
	}

	if result, ok := i.(float64); ok {
		return int(result)
	}

	return 0
}
