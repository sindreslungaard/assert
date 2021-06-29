# assert <img src="https://storage.googleapis.com/gopherizeme.appspot.com/gophers/1dc4851c11e2ec6e05533c2e7d87df1687cf97fc.png" alt="gopher" width="130" align="right">

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

    // str == "some-value"
}
```

Every assertion formats the input to a type it prefers to test on before executing, so you can do number assertions on string values and so on

```go
num, err := assert.Is("55").MinNum(50).Int()
num, err := assert.Is(55).MinNum(50).Int()
// Both of these are OK
```

Example validation for a http form POST request

```go
func mySignupHandler(c *gin.Context) {

    username, err := assert.Is(c.PostForm("username")).MinLen(3).MaxLen(20).AlphaNumeric().String()

    if err != nil {
        // let the user know
    }

    password, err := assert.Is(c.PostForm("password")).MinLen(8)

    if err != nil {
        // let the user know
    }

}
```

We can recude the number of error checks with `assert.First`

```go
func mySignupHandler(c *gin.Context) {

    username, err1 := assert.Is(c.PostForm("username")).MinLen(3).MaxLen(20).AlphaNumeric().String()
    email,    err2 := assert.Is(c.PostForm("email")).Email().String()
    password, err3 := assert.Is(c.PostForm("password")).MinLen(8).String()

    err := assert.First(err1, err2, err3)

    if err != nil {
        // let the user know
    }

}
```

### Available validations

| Validation   | Argument type | Description                                                                                  |
| ------------ | ------------- | -------------------------------------------------------------------------------------------- |
| NotEmpty     |               | Checks if the input (formatted to string) is equal to ""                                     |
| MinLen       | int           | Checks if the length of input (formatted to string) is greater than or equal to the argument |
| MaxLen       | int           | Checks if the length of input (formatted to string) is less than or equal to the argument    |
| MinNum       | int           | Checks if the value of input (formatted to int) is greater than or equal to the argument     |
| MaxNum       | int           | Checks if the value of input (formatted to int) is less than or equal to the argument        |
| Regex        | string        | Checks if the input (formatted to string) matches the provided regex pattern                 |
| Email        |               | Checks if the input (formatted to string) is a valid email                                   |
| Alpha        |               | Checks if the input (formatted to string) only contains a-z and A-Z                          |
| AlphaNumeric |               | Checks if the input (formatted to string) only contains a-z, A-Z and 0-9                     |

### Available type cast assertions

Type casts should always be last in your assertion chain, they are what executes all assertions in the chain and returns any errors or the input converted to the new type.

| Type    | Description                                                                                    |
| ------- | ---------------------------------------------------------------------------------------------- |
| String  | Formats the input as a string using `fmt.Sprintf`                                              |
| Int     | Formats the input as an int using `fmt.Sprintf -> strconv.Atoi`                                |
| Float64 | Formats the input as a float64 using `fmt.Sprintf -> strconv.ParseFloat` with 64 bit precision |
