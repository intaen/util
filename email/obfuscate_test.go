package mail

import "testing"

func TestObfuscateEmail1(t *testing.T) {
	data := "testing.email@gmail.com"
	expected := "te**********l@gmail.com"
	res := ObfuscateEmail(data)
	if res != expected {
		t.Errorf("This function doesn't work, result must be '%s' not '%s'", expected, res)
	}
}


func TestObfuscateEmail2(t *testing.T) {
	data := "test@outlook.com"
	expected := "t**t@outlook.com"
	res := ObfuscateEmail(data)
	if res != expected {
		t.Errorf("This function doesn't work, result must be '%s' not '%s'", expected, res)
	}
}