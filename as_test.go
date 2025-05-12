package erh_test

import (
	"fmt"
	"io/fs"
	"os"
	"reflect"
	"testing"

	"github.com/bengo4/erh"
	"github.com/stretchr/testify/assert"
)

func TestAs(t *testing.T) {
	_, errPathError := os.Open("non-existing")
	errWrapped := &wrapError{msg: "wrapped", err: errPathError}
	errObject := objectError{msg: "object error"}
	errWrappedObject := &wrapError{msg: "wrapped", err: errObject}

	t.Run("erh.As[*fs.PathError](nil)", testAsNoMatch[*fs.PathError](nil))
	t.Run("erh.As[*fs.PathError](errPathError)", testAsMatchTo[*fs.PathError](errPathError, errPathError))
	t.Run("erh.As[*fs.PathError](errWrapped)", testAsMatchTo[*fs.PathError](errWrapped, errPathError))
	t.Run("erh.As[*fs.PathError](errObject)", testAsNoMatch[*fs.PathError](errObject))
	t.Run("erh.As[*fs.PathError](errWrappedObject)", testAsNoMatch[*fs.PathError](errWrappedObject))

	t.Run("erh.As[*os.LinkError](nil)", testAsNoMatch[*os.LinkError](nil))
	t.Run("erh.As[*os.LinkError](errPathError)", testAsNoMatch[*os.LinkError](errPathError))
	t.Run("erh.As[*os.LinkError](errWrapped)", testAsNoMatch[*os.LinkError](errWrapped))
	t.Run("erh.As[*os.LinkError](errObject)", testAsNoMatch[*os.LinkError](errObject))
	t.Run("erh.As[*os.LinkError](errWrappedObject)", testAsNoMatch[*os.LinkError](errWrappedObject))

	t.Run("erh.As[myE](nil)", testAsNoMatch[myE](nil))
	t.Run("erh.As[myE](errPathError)", testAsMatchTo[myE](errPathError, errPathError))
	t.Run("erh.As[myE](errWrapped)", testAsMatchTo[myE](errWrapped, errWrapped))
	t.Run("erh.As[myE](errObject)", testAsMatchTo[myE](errObject, errObject))
	t.Run("erh.As[myE](errWrappedObject)", testAsMatchTo[myE](errWrappedObject, errWrappedObject))

	t.Run("erh.As[*wrapError](nil)", testAsNoMatch[*wrapError](nil))
	t.Run("erh.As[*wrapError](errPathError)", testAsNoMatch[*wrapError](errPathError))
	t.Run("erh.As[*wrapError](errWrapped)", testAsMatchTo[*wrapError](errWrapped, errWrapped))
	t.Run("erh.As[*wrapError](errObject)", testAsNoMatch[*wrapError](errObject))
	t.Run("erh.As[*wrapError](errWrappedObject)", testAsMatchTo[*wrapError](errWrappedObject, errWrappedObject))

	t.Run("erh.As[objectError](nil)", testAsNoMatch[objectError](nil))
	t.Run("erh.As[objectError](errPathError)", testAsNoMatch[objectError](errPathError))
	t.Run("erh.As[objectError](errWrapped)", testAsNoMatch[objectError](errWrapped))
	t.Run("erh.As[objectError](errObject)", testAsMatchTo[objectError](errObject, errObject))
	t.Run("erh.As[objectError](errWrappedObject)", testAsMatchTo[objectError](errWrappedObject, errObject))

	t.Run("erh.As[*objectError](nil)", testAsNoMatch[*objectError](nil))
	t.Run("erh.As[*objectError](errPathError)", testAsNoMatch[*objectError](errPathError))
	t.Run("erh.As[*objectError](errWrapped)", testAsNoMatch[*objectError](errWrapped))
	t.Run("erh.As[*objectError](errObject)", testAsNoMatch[*objectError](errObject))
	t.Run("erh.As[*objectError](errWrappedObject)", testAsNoMatch[*objectError](errWrappedObject))
}

func testAsNoMatch[T error](err error) func(t *testing.T) {
	return testAs(err, assertErrorNoMatch[T]())
}

func testAsMatchTo[T error](err, matchTo error) func(t *testing.T) {
	return testAs(err, assertErrorMatchTo[T](matchTo))
}

func testAs[T error](err error, assertFunc func(*testing.T, T, bool)) func(t *testing.T) {
	return func(t *testing.T) {
		t.Helper()
		target, ok := erh.As[T](err)
		assertFunc(t, target, ok)
	}
}

func assertErrorNoMatch[T error]() func(t *testing.T, target T, ok bool) {
	return func(t *testing.T, target T, ok bool) {
		t.Helper()
		assert := assert.New(t)
		assert.False(ok)
		assert.Zero(target)
	}
}

func assertErrorMatchTo[T error](want error) func(t *testing.T, target T, ok bool) {
	return func(t *testing.T, target T, ok bool) {
		t.Helper()
		assert := assert.New(t)
		assert.True(ok)
		if assert.NotNil(target) {
			if reflect.ValueOf(target).Kind() == reflect.Ptr {
				assert.Same(want, target)
				t.Logf("want and target are same, target:%v", target)
			} else {
				assert.Equal(want, target)
				t.Logf("want and target are equal, target:%#+v", target)
			}
		}
	}
}

// myE is an interface compatible with error.
type myE interface {
	error
}

// wrapError is an wrapped error for testing.
type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return fmt.Sprintf("%s:%s", e.msg, e.err.Error())
}

func (e *wrapError) Unwrap() error {
	return e.err
}

// objectError is an error object for testing.
// Methods are implemeted to objectError, not to *objectError.
type objectError struct {
	msg string
}

func (e objectError) Error() string {
	return e.msg
}
