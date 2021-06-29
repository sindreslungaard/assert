# assert <img src="https://storage.googleapis.com/gopherizeme.appspot.com/gophers/1dc4851c11e2ec6e05533c2e7d87df1687cf97fc.png" alt="gopher" width="40" align="right">

Dead simple 0 dependency "typed" golang validator

### Motivation

Most go validators use either struct tags or a string based representation of the validation rules, and while there's nothing wrong with that, sometimes you just want to quickly validate an input and perhaps cast it to a different type from which you received it.

With assert you may also rest assured you did not typo any of the validation rules as all rules are functions, hence a bit more "typed" than something like `validate:"required,min=4,max=5"`

### Installing

```
go get github.com/sindreslungaard/assert
```

### Examples

Simple example of validating the character length of an interface and returning it as a string

```go
package main

import "github.com/sindreslungaard/assert"

func main() {
    // This can be any value, any type.
    src := "some-value"

    str, err := assert.Is(src).MinLen(5).MaxLen(10).String()

    if err != nil {
        // handle error
    }

    // str = some-value
}
```

### Available validations

| Validation | Argument type | Description                                                                                  |
| ---------- | ------------- | -------------------------------------------------------------------------------------------- |
| NotEmpty   |               | Checks if the input (formatted to string) is equal to ""                                     |
| MinLen     | int           | Checks if the length of input (formatted to string) is greater than or equal to the argument |
| MaxLen     | int           | Checks if the length of input (formatted to string) is less than or equal to the argument    |

### Available type cast assertions

Type casts should always be last in your assertion chain, they are what executes all assertions in the chain and returns any errors or the input converted to the new type.

| Type    | Description                                                                                    |
| ------- | ---------------------------------------------------------------------------------------------- |
| String  | Formats the input as a string using `fmt.Sprintf`                                              |
| Int     | Formats the input as an int using `fmt.Sprintf -> strconv.Atoi`                                |
| Float64 | Formats the input as a float64 using `fmt.Sprintf -> strconv.ParseFloat` with 64 bit precision |
