package hex

// Find returns the first element in slice `s` that satisfies `fn(T)` and true.
// If no values satisfy fn(T) it will return zero value of the element and false.
func Find[T any](s []T, fn func(el T) bool) (T, bool) {
	for _, el := range s {
		if fn(el) {
			return el, true
		}
	}

	var res T
	return res, false
}

// FindIndex returns the index of the first element in slice `s` that satisfies `fn(T)`.
// If no values satisfy fn(t) it will return -1.
func FindIndex[T any](s []T, fn func(el T) bool) int {
	for i, el := range s {
		if fn(el) {
			return i
		}
	}

	return -1
}