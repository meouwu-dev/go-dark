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
	defer func() {
		if err := recover(); err != nil {
			// cast err to error or create a new error from err
			if _, ok := err.(error); ok {
				catchFc(err.(error))
				return
			}

			// convert err to string
			catchFc(fmt.Errorf("Try failed: %v", err))
		}
	}()
	tryFc()
}
