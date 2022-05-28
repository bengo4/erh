package erh

import (
	"bytes"
	"errors"
	"fmt"
)

// Wrap returns a new error based on err.
//
// Wrap formats an error message with wrapping err.
//
// Wrap also add the short filename and the line number of the place where Wrap is called to the error message.
//
// The err wrapped by Wrap can be retrieved by errors.Unwrap.
//
// The very first error of repeatedly wrapped errors can retrieved by Cause.
//
// If err is nil, Wrap returns nil.
func Wrap(err error, a ...interface{}) error {
	if err == nil {
		return nil
	}

	buf := new(bytes.Buffer)
	if len(a) > 0 {
		format := fmt.Sprintf("%v; ", a[0])
		fmt.Fprintf(buf, format, a[1:]...)
	}
	buf.WriteString("%w")
	addFileLine(buf)
	return fmt.Errorf(buf.String(), err)
}

// Cause returns the very first error of repeatedly wrapped errors.
//
// This works like github.com/pkg/errors.Cause (https://github.com/pkg/errors)
func Cause(err error) error {
	for {
		err2 := errors.Unwrap(err)
		if err2 == nil {
			return err
		}

		err = err2
	}
}
