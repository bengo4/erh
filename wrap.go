package erh

import (
	"bytes"
	"fmt"
)

// Wrap returns an error that formats as err with the point Wrap is called, and the supplied message.
// Wrap also records the cause of err.
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
	buf.WriteString(err.Error())
	addFileLine(buf)

	return &wrap{
		cause: Cause(err),
		msg:   buf.String(),
	}
}

// Cause returns the underlying cause of the error, if possible.
//
// This is same thing as github.com/pkg/errors.Cause() (https://github.com/pkg/errors)
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	if err == nil {
		return nil
	}

	c, ok := err.(causer)
	if !ok {
		return err
	}

	return c.Cause()
}

type wrap struct {
	cause error
	msg   string
}

// Error
func (w *wrap) Error() string {
	return w.msg
}

// Cause
func (w *wrap) Cause() error {
	return w.cause
}
