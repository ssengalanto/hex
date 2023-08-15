package hex

import "reflect"

// Deref takes a pointer to a value of type T and returns the de-referenced value.
// If the pointer is nil, it returns a zeroed value of type T.
func Deref[T any](ptr *T) T {
	field := reflect.ValueOf(ptr)

	if field.IsNil() {
		return reflect.Zero(field.Type().Elem()).Interface().(T)
	}

	return field.Elem().Interface().(T)
}
