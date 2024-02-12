package supportingfunctions

// SliceJoinUniq создает новый срез состоящий только из уникальных
// значений срезов sliceOne и sliceTwo
func SliceJoinUniq[T comparable](sliceOne, sliceTwo []T) []T {
	numOne, numTwo := len(sliceOne), len(sliceTwo)

	if numOne == 0 && numTwo == 0 {
		return []T{}
	}

	if numOne == 0 {
		return sliceTwo
	}

	if numTwo == 0 {
		return sliceOne
	}

	newList := make([]T, 0, numOne+numTwo)

	for _, v := range append(sliceOne, sliceTwo...) {
		if SliceElemIsExist[T](v, newList) {
			continue
		}

		newList = append(newList, v)
	}

	return newList
}

// SliceElemIsExist проверяет наличие элемента в срезе
func SliceElemIsExist[T comparable](item T, list []T) bool {
	for _, v := range list {
		if item == v {
			return true
		}
	}

	return false
}
