package erh

import (
	"fmt"
	"strings"
)

// Errorf returns a new error with a formatted message accordings to a format specifier.
//
// Additionally, the message of the returned error includes the short filename and the line number of the place where Errorf is called.
func Errorf(format string, a ...interface{}) error {
	var b strings.Builder
	fmt.Fprintf(&b, format, a...)
	addFileLine(&b)
	return fmt.Errorf("%s", b.String())
}
