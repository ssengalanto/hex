package hex

// FindIndex returns the index of the first element in slice s that satisfies fn(T).
// If no values satisfy fn(t) it will return -1.
func FindIndex[T any](s []T, fn func(el T) bool) int {
	for i, el := range s {
		if fn(el) {
			return i
		}
	}

	return -1
}
