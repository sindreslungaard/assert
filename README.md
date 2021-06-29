# assert

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
