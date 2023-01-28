package hex

// Find returns the first element in slice `s` that satisfies `fn(T)` and true.
// If no values satisfy `fn(T)` it will return zero value of the element and false.
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
// If no values satisfy `fn(t)` it will return -1.
func FindIndex[T any](s []T, fn func(el T) bool) int {
	for i, el := range s {
		if fn(el) {
			return i
		}
	}

	return -1
}

// Map creates a new slice from calling `fn(T, int)` for every slice element.
func Map[T any, R any](s []T, fn func(el T, i int) R) []R {
	res := make([]R, len(s))

	for i, el := range s {
		res[i] = fn(el, i)
	}

	return res
}

// Filter filters down to just the elements from slice `s` that satisfies `fn(T, int)`.
func Filter[T any](s []T, fn func(el T, index int) bool) []T {
	var result []T

	for i, el := range s {
		if fn(el, i) {
			result = append(result, el)
		}
	}

	if len(result) == 0 {
		return []T{}
	}

	return result
}

// Chunk creates a slice of elements splits into groups the length of size.
// If slice can't be split evenly, the final chunk will be the remaining elements.
func Chunk[T any](s []T, sz int) [][]T {
	if sz <= 0 {
		panic("`sz` must be greater than zero")
	}

	n := len(s) / sz
	if len(s)%sz != 0 {
		n += 1
	}

	res := make([][]T, 0, n)

	for i := 0; i < n; i++ {
		l := (i + 1) * sz
		if l > len(s) {
			l = len(s)
		}
		res = append(res, s[i*sz:l])
	}

	return res
}

// Prepend prepends an element in slice `s`.
func Prepend[T any](s []T, el T) []T {
	s = append([]T{el}, s...)
	return s
}
