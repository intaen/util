package str

import "testing"

func TestGetStrFirst(t *testing.T) {
	data := "Test Getstring"
	pos := -1
	strlen := 9
	expected := "Getstring"
	res := GetStr(data, strlen, pos)
	if res != expected {
		t.Errorf("This function doesn't work, if we find last position from '%s', the result must be '%s' not '%s'", data, expected, res)
	}
}

func TestGetStrLast(t *testing.T) {
	data := "Test Getstring"
	pos := 1
	strlen := 4
	expected := "Test"
	res := GetStr(data, strlen, pos)
	if res != expected {
		t.Errorf("This function doesn't work, if we find last position from '%s', the result must be '%s' not '%s'", data, expected, res)
	}
}