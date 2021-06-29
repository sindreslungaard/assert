package assert

import "fmt"

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

// MinLen asserts that the character length of the source
// (formatted as string) is greater than or equal to the value
func (a *Assertion) MinLen(val int) *Assertion {

	a.add(func(i interface{}) error {
		if len([]rune(fmt.Sprintf("%v", a.src))) >= val {
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
		if len([]rune(fmt.Sprintf("%v", a.src))) <= val {
			return nil
		}

		return fmt.Errorf("assert.maxlength")
	})

	return a

}
