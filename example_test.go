package erh_test

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/bengo4/erh"
)

func ExampleCause() {
	err0 := fmt.Errorf("something wrong")
	err1 := erh.Wrap(err0)
	err2 := erh.Wrap(err1)
	fmt.Println(erh.Cause(err0) == err0)
	fmt.Println(erh.Cause(err1) == err0)
	fmt.Println(erh.Cause(err2) == err0)
	// Output:
	// true
	// true
	// true
}

func ExampleErrorf() {
	err := erh.Errorf("something wrong")
	fmt.Println(err)
	// Output: something wrong[example_test.go:26]
}

func ExampleWrap() {
	err0 := fmt.Errorf("something wrong")
	err1 := erh.Wrap(err0)
	err2 := erh.Wrap(err1)
	fmt.Println(err1)
	fmt.Println(err2)
	// Output:
	// something wrong[example_test.go:33]
	// something wrong[example_test.go:33][example_test.go:34]
}

func ExampleWrap_nil() {
	err := erh.Wrap(nil)
	fmt.Println(err)
	// Output: <nil>
}

func ExampleWrap_message() {
	err0 := fmt.Errorf("something wrong")
	err1 := erh.Wrap(err0, "wrapped")
	err2 := erh.Wrap(err1, "wrapped again")
	fmt.Println(err1)
	fmt.Println(err2)
	// Output:
	// wrapped; something wrong[example_test.go:50]
	// wrapped again; wrapped; something wrong[example_test.go:50][example_test.go:51]
}

func ExampleWrap_messageFormatted() {
	err0 := fmt.Errorf("something wrong")
	err1 := erh.Wrap(err0, "additional message, p:%d", 123)
	fmt.Println(err1)
	// Output: additional message, p:123; something wrong[example_test.go:61]
}

func ExampleWrap_unwrap() {
	err0 := fmt.Errorf("something wrong")
	err1 := erh.Wrap(err0)
	err2 := erh.Wrap(err1)
	fmt.Println(errors.Unwrap(err0) == nil)
	fmt.Println(errors.Unwrap(err1) == err0)
	fmt.Println(errors.Unwrap(err2) == err1)
	// Output:
	// true
	// true
	// true
}

func ExampleWrap_is() {
	err0 := fmt.Errorf("something wrong")
	err1 := erh.Wrap(err0)
	err2 := erh.Wrap(err1)
	fmt.Println(errors.Is(err1, err0))
	fmt.Println(errors.Is(err2, err1))
	fmt.Println(errors.Is(err2, err0))
	// Output:
	// true
	// true
	// true
}

func ExampleAs() {
	_, err0 := os.Open("non-existing")
	err1 := erh.Wrap(err0)
	// target1 will be unwrapped to err0 which is *fs.PathError.
	target1, ok := erh.As[*fs.PathError](err1)
	fmt.Println("target1:", ok)
	fmt.Println("target1:", target1)
	// target2 will not be unwrapped because err1 matches to error interface.
	target2, ok := erh.As[error](err1)
	fmt.Println("target2:", ok)
	fmt.Println("target2:", target2)
	// target3 will be nil because both err1 and err0 do not match to *os.LinkError.
	target3, ok := erh.As[*os.LinkError](err1)
	fmt.Println("target3:", ok)
	fmt.Println("target3:", target3)
	// Output:
	// target1: true
	// target1: open non-existing: no such file or directory
	// target2: true
	// target2: open non-existing: no such file or directory[example_test.go:94]
	// target3: false
	// target3: <nil>
}
