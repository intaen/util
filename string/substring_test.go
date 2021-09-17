package str

import "testing"

func TestSubStr(t *testing.T) {
	data := "Test Substring"
	firstlen := 5
	lastlen := 6
	expected := "Sub"
	res := SubStr(data, firstlen, lastlen)
	if res != expected {
		t.Errorf("This function doesn't work, result must be '%s' not '%s'", expected, res)
	}
}
