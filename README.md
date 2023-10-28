# go-dark

## Description

A go package containing functions that are hated by Gophers.

# Usage

## `Must`

The `Must` function is a function that takes a value and an error.
It returns the value if the error is nil, otherwise it panics.

### Example

```go
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
```

## `MustNil`

The `MustNil` is the same as `Must` but it only takes an error.

It panics if the error is not nil.

### Example

```go
func success() error {
    return nil
}

func failure() error {
    return errors.New("some error")
}

func main() {
    dark.MustNil(success())
    fmt.Println("success")

    dark.MustNil(failure()) // panics
    fmt.Println("failure")  // never reached
}

```

## `Try`

It is like the try/catch block in other languages.

### Example

```go
func failure() (string, error) {
    return "", errors.New("some error")
}

func main() {
    dark.Try(func() {
        dark.Must[string](failure())
        fmt.Printf("never reached")
    })(func(err any) {
        fmt.Printf("%v\n", err) // prints "some error"
    })
}
```



