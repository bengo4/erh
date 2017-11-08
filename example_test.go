package erh_test

import (
	"fmt"

	"github.com/bengo4/erh"
)

func ExampleCause() {
	err0 := fmt.Errorf("something wrong")
	err1 := erh.Wrap(err0)

	fmt.Println(erh.Cause(err0) == err0)
	fmt.Println(erh.Cause(err1) == err0)
	// Output:
	// true
	// true
}

func ExampleErrorf() {
	err := erh.Errorf("something wrong")
	fmt.Println(err)
	// Output: something wrong[example_test.go:21]
}

func ExampleWrap() {
	err0 := fmt.Errorf("something wrong")
	err1 := erh.Wrap(err0)
	fmt.Println(err1)
	// Output: something wrong[example_test.go:28]
}

func ExampleWrap_nil() {
	err := erh.Wrap(nil)
	fmt.Println(err)
	// Output: <nil>
}

func ExampleWrap_message() {
	cause := fmt.Errorf("something wrong")
	err := erh.Wrap(cause, "additional message")
	fmt.Println(err)
	// Output: additional message; something wrong[example_test.go:41]
}

func ExampleWrap_messageFormatted() {
	err0 := fmt.Errorf("something wrong")
	err1 := erh.Wrap(err0, "additional message, p:%d", 123)
	fmt.Println(err1)
	// Output: additional message, p:123; something wrong[example_test.go:48]
}
