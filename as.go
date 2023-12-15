package erh

import "errors"

// As is a wrapper of [errors.As].
//
// As finds the first error in err's chain that matches type T, and if so, returns the error of type T and true.
//
// If no error in the chain matches T, As returns zero value of T and false.
//
// Type T must conform to the error interface.
//
// An example of using As:
//
//	target, ok := erh.As[*fs.PathError](err)
//
// The code above is equivalent to the code below:
//
//	 var (
//		target *fs.PathError
//		_      error = target
//	 )
//	 ok := errors.As(err, &target)
func As[T error](err error) (T, bool) {
	var target T
	ok := errors.As(err, &target)
	return target, ok
}
