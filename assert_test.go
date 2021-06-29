package assert

import "testing"

func TestGeneralPurposeString(t *testing.T) {

	src := "test"

	str, err := Is(src).NotEmpty().MinLen(4).String()

	if err != nil {
		t.Fatal(err)
	}

	if str != src {
		t.FailNow()
	}

}

func TestGeneralPurposeInt(t *testing.T) {

	src := 55

	num, err := Is(src).NotEmpty().Int()

	if err != nil {
		t.Fatal(err)
	}

	if num != src {
		t.Fatal("not equal")
	}

}

func TestGeneralPurposeFloat64(t *testing.T) {

	src := 55.55555555555555555555555555

	num, err := Is(src).NotEmpty().Float64()

	if err != nil {
		t.Fatal(err)
	}

	if num != src {
		t.Fatal("not equal")
	}

}
