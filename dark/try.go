package dark

import (
	"fmt"
)

/*
Must wraps a function returning a value and an error, and panics if the error is not nil.

Usage: Must(f()) where f() returns (T, error)

Example:

	func success() (string, error) {
		return "hello", nil
	}

	func failure() (string, error) {
		return "", errors.New("some error")
	}

	func main() {
		r := dark.Must(success())
		fmt.Println(r) // prints "hello"

		r = dark.Must(failure()) // panics
		fmt.Println(r)           // never reached
	}
*/
func Must[T any](r T, err error) T {
	if err != nil {
		panic(err)
	}
	return r
}

/*
MustNil panics if the error is not nil.
*/
func MustNil(err error) {
	Must[any](nil, err)
}

/*
Try wraps a function and a catch function, and calls the catch function if the wrapped function panics.
*/
func Try(tryFc func(), catchFc func(error)) {
	err := AbortOnErr(tryFc)
	if err != nil {
		catchFc(err)
	}
}

/*
AbortOnErr wraps a function and aborts on panic.
The function aborts on the first panic, and returns the panic error as a return value.
*/
func AbortOnErr(fc func()) (returnError error) {
	defer func() {
		if err := recover(); err != nil {
			// cast err to error or create a new error from err
			if _, ok := err.(error); ok {
				returnError = err.(error)
			}

			// convert err to string

			returnError = fmt.Errorf("Try failed: %v", err)
		}
	}()
	fc()
	return nil
}
