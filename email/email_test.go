package mail

import (
	"testing"
)

// Positive case
// Local-part
// Special characters
func TestIsEmailValidP1(t *testing.T) {
	data := "testing.email-positive@outlook.com"
	expected := true
	result := IsEmailValid(data)
	if result != expected {
		t.Errorf("This function doesn't work, result must be '%t' not '%t'", expected, result)
	}
}

// Local-part
// Two '.' (period)
func TestIsEmailValidP2(t *testing.T) {
	data := "testing.email.positive@outlook.com"
	expected := true
	result := IsEmailValid(data)
	if result != expected {
		t.Errorf("This function doesn't work, result must be '%t' not '%t'", expected, result)
	}
}

// Domain
// Others
func TestIsEmailValidP3(t *testing.T) {
	data := "testing.email@baf.id" // domain not covered
	expected := true
	result := IsEmailValid(data)
	if result != expected {
		t.Errorf("This function doesn't work, result must be '%t' not '%t'", expected, result)
	}
}

// Negative case
// Two '@' sign
func TestIsEmailValidN1(t *testing.T) {
	data := "testing@emailnegative@gmail.com"
	expected := false
	result := IsEmailValid(data)
	if result != expected {
		t.Errorf("This function doesn't work, result must be '%t' not '%t'", expected, result)
	}
}

// Empty data after/before '@' sign
func TestIsEmailValidN2(t *testing.T) {
	data := "testing.emailnegative@"
	expected := false
	result := IsEmailValid(data)
	if result != expected {
		t.Errorf("This function doesn't work, result must be '%t' not '%t'", expected, result)
	}
}

// Local-part
// Two '.' (period)
func TestIsEmailValidN3(t *testing.T) {
	data := "testing.email.negative@gmail.com"
	expected := false
	result := IsEmailValid(data)
	if result != expected {
		t.Errorf("This function doesn't work, result must be '%t' not '%t'", expected, result)
	}
}

// Local-part
// Length less than 6
func TestIsEmailValidN4(t *testing.T) {
	data := "test@gmail.com"
	expected := false
	result := IsEmailValid(data)
	if result != expected {
		t.Errorf("This function doesn't work, result must be '%t' not '%t'", expected, result)
	}
}

// Local-part
// Length more than 30 chars
func TestIsEmailValidN5(t *testing.T) {
	data := "useremailuseremailuserpokoknyapunya.userdeh@gmail.com"
	expected := false
	result := IsEmailValid(data)
	if result != expected {
		t.Errorf("This function doesn't work, result must be '%t' not '%t'", expected, result)
	}
}

// Local-part
// Length more than 64 chars
func TestIsEmailValidN6(t *testing.T) {
	data := "useremailuseremailuserpokoknyapunya.userdehaduhharuspanjangbanget123@hotmail.com"
	expected := false
	result := IsEmailValid(data)
	if result != expected {
		t.Errorf("This function doesn't work, result must be '%t' not '%t'", expected, result)
	}
}

// Local-part
// Special characters
func TestIsEmailValidN7(t *testing.T) {
	data := "testing.email-negative@gmail.com"
	expected := false
	result := IsEmailValid(data)
	if result != expected {
		t.Errorf("This function doesn't work, result must be '%t' not '%t'", expected, result)
	}
}

// Domain
// Length more than 255 chars
func TestIsEmailValidN8(t *testing.T) {
	data := "testaja@testingemailyangpanjangnyaduaratuslimapuluhlimauntukdomaintapigataubisasampesegitugakyajiriniajabarusembilanpuluhtigasekarangbaruseratusduapuluhlimayaallahpanjangbeneryagamungkinsihyaadayangdomainnyasepanjanginiyaallahmasihbanyakkurangnyagimanainiplisplisyeheeet.com"
	expected := false
	result := IsEmailValid(data)
	if result != expected {
		t.Errorf("This function doesn't work, result must be '%t' not '%t'", expected, result)
	}
}

// Domain
// Two '.' (period)
func TestIsEmailValidN9(t *testing.T) {
	data := "testing.emailnegative@gmail..com"
	expected := false
	result := IsEmailValid(data)
	if result != expected {
		t.Errorf("This function doesn't work, result must be '%t' not '%t'", expected, result)
	}
}

// Domain
// Special characters
func TestIsEmailValidN10(t *testing.T) {
	data := "testing.email.negative@outlook-.com"
	expected := false
	result := IsEmailValid(data)
	if result != expected {
		t.Errorf("This function doesn't work, result must be '%t' not '%t'", expected, result)
	}
}
