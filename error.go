package erh

import (
	"bytes"
	"fmt"
)

// Errorf returns an error that formats according to a format specifier and adds the point it was called.
func Errorf(f string, a ...interface{}) error {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, f, a...)
	addFileLine(buf)
	return fmt.Errorf(buf.String())
}
