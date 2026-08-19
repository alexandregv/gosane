//nolint:unused // test file
package noinits

// init() function should be flagged.
func init() {} // want "init functions are not allowed"

// An init function with a receiver should NOT be flagged.
type Dummy struct {}
func (d *Dummy) init() {}
