package erh

import (
	"errors"
	"fmt"
	"strings"
)

// Wrap returns a new error which wraps err.
//
// A message acording to the format specifier can be added to the returned error.
//
// The error message of the new error also includes the short filename and the line number of the place where Wrap is called.
//
// If err is nil, Wrap returns nil.
//
// The err wrapped by Wrap can be retrieved by errors.Unwrap. So [errors.Is] can be used with Wrap.
//
// The very first error of repeatedly wrapped errors can be retrieved by Cause.
func Wrap(err error, a ...interface{}) error {
	if err == nil {
		return nil
	}

	var b strings.Builder
	if len(a) > 0 {
		format := fmt.Sprintf("%v; ", a[0])
		fmt.Fprintf(&b, format, a[1:]...)
	}
	b.WriteString("%w")
	addFileLine(&b)
	return fmt.Errorf(b.String(), err)
}

// Cause returns the very first error of repeatedly wrapped errors.
//
// This works like https://pkg.go.dev/github.com/pkg/errors#Cause.
func Cause(err error) error {
	for {
		err2 := errors.Unwrap(err)
		if err2 == nil {
			return err
		}

		err = err2
	}
}
