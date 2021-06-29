package assert

import (
	"fmt"
	"regexp"
	"strconv"
)

type AssertFunc func(interface{}) error

func assert(src interface{}, tests []AssertFunc) error {
	for _, test := range tests {
		err := test(src)

		if err != nil {
			return err
		}
	}

	return nil
}

type Assertion struct {
	src    interface{}
	ok     bool
	errors error
	tests  []AssertFunc
}

func Is(src interface{}) *Assertion {
	return &Assertion{
		src:   src,
		ok:    true,
		tests: []AssertFunc{},
	}
}

// add adds an AssertFunc to the list of tests to run later
func (a *Assertion) add(f AssertFunc) {
	a.tests = append(a.tests, f)
}

// String runs all assertions and returns the source as a string and errors if any
func (a *Assertion) String() (string, error) {
	err := assert(a.src, a.tests)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", a.src), nil
}

// Int runs all assertions and returns the source as an int and errors if any
func (a *Assertion) Int() (int, error) {
	err := assert(a.src, a.tests)

	if err != nil {
		return 0, err
	}

	num, err := strconv.Atoi(fmt.Sprintf("%v", a.src))

	if err != nil {
		return 0, fmt.Errorf("assert.int")
	}

	return num, nil
}

// Float runs all assertions and returns the source as a float64 and errors if any
func (a *Assertion) Float64() (float64, error) {
	err := assert(a.src, a.tests)

	if err != nil {
		return 0, err
	}

	num, err := strconv.ParseFloat(fmt.Sprintf("%v", a.src), 64)

	if err != nil {
		return 0, fmt.Errorf("assert.float")
	}

	return num, nil
}

// NotEmpty asserts that the string representation of the
// source is not equal to ""
func (a *Assertion) NotEmpty() *Assertion {

	a.add(func(i interface{}) error {
		if fmt.Sprintf("%v", i) != "" {
			return nil
		}

		return fmt.Errorf("assert.required")
	})

	return a

}

// MinLen asserts that the character length of the source
// (formatted as string) is greater than or equal to the value
func (a *Assertion) MinLen(val int) *Assertion {

	a.add(func(i interface{}) error {
		if len([]rune(fmt.Sprintf("%v", i))) >= val {
			return nil
		}

		return fmt.Errorf("assert.minlength")
	})

	return a

}

// MaxLen asserts that the character length of the source
// (formatted as string) is less than or equal to the value
func (a *Assertion) MaxLen(val int) *Assertion {

	a.add(func(i interface{}) error {
		if len([]rune(fmt.Sprintf("%v", i))) <= val {
			return nil
		}

		return fmt.Errorf("assert.maxlength")
	})

	return a

}

// Regex asserts that the source (formatted as string) matches
// the regex pattern specified
func (a *Assertion) Regex(val string) *Assertion {

	a.add(func(i interface{}) error {

		r, err := regexp.Compile(val)

		if err != nil {
			return fmt.Errorf("assert.regex.compile")
		}

		if r.MatchString(fmt.Sprintf("%v", i)) {
			return nil
		}

		return fmt.Errorf("assert.regex")
	})

	return a

}
