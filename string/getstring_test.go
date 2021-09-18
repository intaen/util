package str

import "testing"

func TestGetStr(t *testing.T) {
	data := "Test Getstring"
	pos := -1
	strlen := 9
	expected := "Getstring"
	res := GetStr(data, strlen, pos)
	if res != expected {
		t.Errorf("This function doesn't work, if we find last position from '%s', the result must be '%s' not '%s'", data, expected, res)
	}
}
