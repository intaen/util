package main

import (
	"fmt"
	"strings"
	"github.com/intaen/go-str"

	inamonth "github.com/intaen/go-month"
)

func main() {
	data := "PPKM diperpanjang hingga 2021-10-07"
	date := str.GetStr(data, 10, -1)    // Get last string
	month := inamonth.ConvertDate(date) // Convert date to date with name of month

	fmt.Println(strings.Replace(data, date, month, -1)) // Replace data
}
