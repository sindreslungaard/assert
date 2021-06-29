package assert

import "testing"

func TestGeneralPurpose(t *testing.T) {

	src := "test"

	str, err := Is(src).NotEmpty().MinLen(4).String()

	if err != nil {
		t.Fatal(err)
	}

	if str != src {
		t.FailNow()
	}

}
