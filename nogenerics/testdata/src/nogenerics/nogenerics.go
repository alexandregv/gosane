//nolint:unused // test file
package nogenerics

// Function with type parameters should be flagged.
func Max[T int | float64](a, b T) T { // want "generics are not allowed"
	if a > b {
		return a
	}
	return b
}

// Type with type parameters should be flagged.
type Stack[T any] struct { // want "generics are not allowed"
	items []T
}

// Method on generic type (type params come from Stack, not from this method).
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Regular function without type parameters should NOT be flagged.
func Add(a, b int) int {
	return a + b
}
