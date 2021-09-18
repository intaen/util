package inamonth

import (
	"strconv"
	"strings"
)

type MonthNumber int

const (
	Januari MonthNumber = 1 + iota
	Februari
	Maret
	April
	Mei
	Juni
	Juli
	Agustus
	September
	Oktober
	November
	Desember
)

var months = [...]string{
	"Januari",
	"Februari",
	"Maret",
	"April",
	"Mei",
	"Juni",
	"Juli",
	"Agustus",
	"September",
	"Oktober",
	"November",
	"Desember",
}

var shortmonths = [...]string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"Mei",
	"Jun",
	"Jul",
	"Agt",
	"Sep",
	"Okt",
	"Nov",
	"Des",
}

var ErrMonthNotFound string = "Bulan tidak ditemukan"

func (n MonthNumber) String() string {
	if n > 12 || n < 1 {
		return ErrMonthNotFound
	}
	return months[n-1]
}

func (n MonthNumber) ShortForm() string {
	if n > 12 || n < 1 {
		return ErrMonthNotFound
	}
	return shortmonths[n-1]
}

// ConvertDate is a function to convert date with format RFC3339 without timezone to date with name of month
func ConvertDate(data string) string {
	// Split yyyy-mm-dd and convert it
	ndata := strings.Split(data, "-")
	num, _ := strconv.Atoi(ndata[1])
	month := MonthNumber(num).String()

	// Validate if number of month more than 12 or less than 1
	if month == ErrMonthNotFound {
		return ErrMonthNotFound
	}

	// Join all data
	var date []string
	date = append(date, ndata[2], month, ndata[0])

	return strings.Join(date, " ")
}
