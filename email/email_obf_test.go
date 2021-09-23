package mail

import "testing"

func TestObfuscateEmail1(t *testing.T) {
	data1 := "testing.email"
	data2 := "gmail.com"
	expected := "te**********l@gmail.com"
	res := ObfuscateEmail(data1, data2)
	if res != expected {
		t.Errorf("This function doesn't work, result must be '%s' not '%s'", expected, res)
	}
}


func TestObfuscateEmail2(t *testing.T) {
	data1 := "test"
	data2 := "outlook.com"
	expected := "t**t@outlook.com"
	res := ObfuscateEmail(data1, data2)
	if res != expected {
		t.Errorf("This function doesn't work, result must be '%s' not '%s'", expected, res)
	}
}