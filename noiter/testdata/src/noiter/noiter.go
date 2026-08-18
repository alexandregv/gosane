//nolint:unused // test file
package noiter

import (
	"iter" // want "iter package is not allowed"
	"maps"
)

func Seq(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range n {
			if !yield(i) {
				break
			}
		}
	}
}

func main() {
	// Range over function should be flagged.
	for i := range Seq(10) { // want "range-over-func is not allowed"
		_ = i
	}

	m := map[string]string{"foo": "bar"}
	for key := range maps.Keys(m) { // want "range-over-func is not allowed"
		_ = key
	}

	// Regular range over slice should NOT be flagged.
	s := []int{1, 2, 3}
	for i := range s {
		_ = i
	}
}
