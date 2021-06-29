package assert

import "testing"

func TestGeneralPurpose(t *testing.T) {

	str, err := Is("test").MinLen(4).String()

	if err != nil {
		t.FailNow()
	}

	if str != "test" {
		t.FailNow()
	}

}
