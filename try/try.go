package dark

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
		r := Must(f())
		fmt.Println(r) // prints "hello"

		r = Must(failure()) // panics
		fmt.Println(r) // never reached
	}
*/
func Must[T any](r T, err error) T {
	if err != nil {
		panic(err)
	}
	return r
}

/*
Try wraps a function [tryFc] that may panic, and returns a function that wraps a function [catchFc] that handles the panic.
Example:

	func main() {
		Try(func() {
			panic(errors.New("some error"))
		})(func(err any) {
			fmt.Printf("%v\n", err) // prints "some error"
		})
	}
*/
func Try(tryFc func()) func(func(any)) {
	return func(catchFc func(any)) {
		defer func() {
			if err := recover(); err != nil {
				catchFc(err)
			}
		}()
		tryFc()
	}
}
