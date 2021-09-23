package mail

import (
	"testing"
)

// Positive case
func TestIsEmailValidP1(t *testing.T) {
	data := "testing.email-positive@outlook.com"
	expectedstr := "te**********l"
	expected := true
	res, locpart, _ := IsEmailValid(data)
	if res != expected {
		t.Errorf("This function doesn't work, result must be '%s' with status '%t' not '%s' with status '%t'", expectedstr, expected, locpart, res)
	}
}

func TestIsEmailValidP2(t *testing.T) {
	data := "testing.email.positive@outlook.com"
	expectedstr := "te*******************e"
	expected := true
	res, locpart, _ := IsEmailValid(data)
	if res != expected {
		t.Errorf("This function doesn't work, result must be '%s' with status '%t' not '%s' with status '%t'", expectedstr, expected, locpart, res)
	}
}

// Negative case
func TestIsEmailValidN1(t *testing.T) {
	data := "testing@emailnegative@gmail.com"
	expectedstr := ""
	expected := false
	res, locpart, _ := IsEmailValid(data)
	if res != expected {
		t.Errorf("This function doesn't work, result must be '%s' with status '%t' not '%s' with status '%t'", expectedstr, expected, locpart, res)
	}
}

func TestIsEmailValidN2(t *testing.T) {
	data := "testing.email.negative@gmail.com"
	expectedstr := ""
	expected := false
	res, locpart, _ := IsEmailValid(data)
	if res != expected {
		t.Errorf("This function doesn't work, result must be '%s' with status '%t' not '%s' with status '%t'", expectedstr, expected, locpart, res)
	}
}

func TestIsEmailValidN3(t *testing.T) {
	data := "testing.email-negative@gmail.com"
	expectedstr := ""
	expected := false
	res, locpart, _ := IsEmailValid(data)
	if res != expected {
		t.Errorf("This function doesn't work, result must be '%s' with status '%t' not '%s' with status '%t'", expectedstr, expected, locpart, res)
	}
}

func TestIsEmailValidN4(t *testing.T) {
	data := "testing.email.negative@outlook-.com"
	expectedstr := ""
	expected := false
	res, locpart, _ := IsEmailValid(data)
	if res != expected {
		t.Errorf("This function doesn't work, result must be '%s' with status '%t' not '%s' with status '%t'", expectedstr, expected, locpart, res)
	}
}

func TestIsEmailValidN5(t *testing.T) {
	data := "test@gmail.com"
	expectedstr := ""
	expected := false
	res, locpart, _ := IsEmailValid(data)
	if res != expected {
		t.Errorf("This function doesn't work, result must be '%s' with status '%t' not '%s' with status '%t'", expectedstr, expected, locpart, res)
	}
}

func TestIsEmailValidN6(t *testing.T) {
	data := "testing.emailnegative@gmail..com"
	expectedstr := ""
	expected := false
	res, locpart, _ := IsEmailValid(data)
	if res != expected {
		t.Errorf("This function doesn't work, result must be '%s' with status '%t' not '%s' with status '%t'", expectedstr, expected, locpart, res)
	}
}

func TestIsEmailValidN7(t *testing.T) {
	data := "testing.emailnegative@"
	expectedstr := ""
	expected := false
	res, locpart, _ := IsEmailValid(data)
	if res != expected {
		t.Errorf("This function doesn't work, result must be '%s' with status '%t' not '%s' with status '%t'", expectedstr, expected, locpart, res)
	}
}
