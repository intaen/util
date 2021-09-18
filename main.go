package main

import (
	"fmt"
	"strings"

	inamonth "github.com/intaen/go-month"
	"github.com/intaen/go-str"
)

func main() {
	data := "PPKM diperpanjang hingga 2021-10-07"
	date := str.GetStr(data, -1, 10)     // Get last string
	month := inamonth.ConvertDate(date) // Convert date to date with name of month

	fmt.Println(strings.Replace(data, date, month, -1)) // Replace data
}
