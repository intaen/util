package inamonth

import "testing"

// Positive case
func TestINAMonthP(t *testing.T) {
	num := 10
	expected := "Oktober"
	res := MonthNumber(num)
	if res.String() != expected {
		t.Errorf("This function doesn't work, result must be '%s' not '%s'", expected, res)
	}
}

func TestINAShortMonthP(t *testing.T) {
	num := 9
	expected := "Sep"
	res := MonthNumber.ShortForm(MonthNumber(num))
	if res != expected {
		t.Errorf("This function doesn't work, result must be '%s' not '%s'", expected, res)
	}
}

func TestINANumMonth(t *testing.T) {
	data := Agustus
	expected := 8
	res := int(data)
	if res != expected {
		t.Errorf("This function doesn't work, result must be '%d' not '%d'", expected, res)
	}
}

func TestConvertDate(t *testing.T) {
	data := "2021-03-17"
	expected := "17 Maret 2021"
	res := ConvertDate(data)
	if res != expected {
		t.Errorf("This function doesn't work, result must be '%s' not '%s'", expected, res)
	}
}

// Negative case
func TestINAMonthN(t *testing.T) {
	num := 20
	expected := "Bulan tidak ditemukan"
	res := MonthNumber(num)
	if res.String() != expected {
		t.Errorf("This function doesn't work, result must be '%s' not '%s'", expected, res)
	}
}

func TestINAShortMonthN(t *testing.T) {
	num := -1
	expected := "Bulan tidak ditemukan"
	res := MonthNumber.ShortForm(MonthNumber(num))
	if res != expected {
		t.Errorf("This function doesn't work, result must be '%s' not '%s'", expected, res)
	}
}

func TestConvertDateFailed(t *testing.T) {
	data := "2021-20-10"
	expected := "Bulan tidak ditemukan"
	res := ConvertDate(data)
	if res != expected {
		t.Errorf("This function doesn't work, result must be '%s' not '%s'", expected, res)
	}
}
