package arrays

// Contains проверяет, есть ли искомый элемент в массиве
func Contains(a []interface{}, e interface{}) bool {
	for _, v := range a {
		if v == e {
			return true
		}
	}

	return false
}

// Index возвращает индекс искомого элемента или -1
func Index(a []interface{}, e interface{}) int {
	for k, v := range a {
		if v == e {
			return k
		}
	}

	return -1
}

// Remove удаляет элементы в количестве length, начиная с start
func Remove(a []interface{}, start, length int) []interface{} {
	if length == -1 {
		return a[:start]
	}

	return append(a[:start], a[start+length:]...)
}
