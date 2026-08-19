//nolint:unused // test file
package noshortif

func ij() (int, int) {
	return 1, 2
}

func main() {
	// Short-if should be flagged.
	if i := 1; i == 2 { // want "short-if syntax is not allowed"
		_ = i
	}

	if i, j := ij(); i == 2 { // want "short-if syntax is not allowed"
		_ = j
	}

	// Regular if syntax should NOT be flagged.
	i := 1
	if i == 2 {
		_ = i
	}
}
