package erh

import (
	"bytes"
	"fmt"
)

// Errorf returns a new error with a formatted message accordings to a format specifier.
//
// Additionally, the message of the returned error includes the short filename and the line number of the place where Errorf is called.
func Errorf(format string, a ...interface{}) error {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, format, a...)
	addFileLine(buf)
	return fmt.Errorf(buf.String())
}
