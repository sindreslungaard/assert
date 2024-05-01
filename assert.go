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

func First(e ...error) error {

	for _, err := range e {
		if err != nil {
			return err
		}
	}

	return nil
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

// MinNum asserts that the value of the source
// (formatted as int) is less than or equal to the value
func (a *Assertion) MinNum(val int) *Assertion {

	a.add(func(i interface{}) error {

		num, err := strconv.Atoi(fmt.Sprintf("%v", i))

		if err != nil {
			return fmt.Errorf("assert.minnum")
		}

		if num >= val {
			return nil
		}

		return fmt.Errorf("assert.minnum")
	})

	return a

}

// MaxNum asserts that the value of the source
// (formatted as int) is greater than or equal to the value
func (a *Assertion) MaxNum(val int) *Assertion {

	a.add(func(i interface{}) error {

		num, err := strconv.Atoi(fmt.Sprintf("%v", i))

		if err != nil {
			return fmt.Errorf("assert.maxnum")
		}

		if num <= val {
			return nil
		}

		return fmt.Errorf("assert.maxnum")
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

// Email asserts that the source (formatted as string) matches an email pattern
func (a *Assertion) Email() *Assertion {
	return a.Regex("^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$")
}

// Alpha asserts that the source (formatted as string) only containes letters a-z and A-Z
func (a *Assertion) Alpha() *Assertion {
	return a.Regex("^[a-zA-Z]+$")
}

// AlphaNumeric asserts that the source (formatted as string) only containes characters a-z, A-Z and 0-9
func (a *Assertion) AlphaNumeric() *Assertion {
	return a.Regex("^[a-zA-Z0-9]+$")
}
