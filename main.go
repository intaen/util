package main

import (
	"fmt"
	"strings"

	"github.com/intaen/go-mail"

	"github.com/intaen/go-str"

	inamonth "github.com/intaen/go-month"
)

func main() {
	data1 := "PPKM diperpanjang hingga 2021-10-07"
	date := str.GetStr(data1, 10, -1)   // Get last string
	month := inamonth.ConvertDate(date) // Convert date to date with name of month

	fmt.Println(strings.Replace(data1, date, month, -1)) // Replace data
	// Result: 2021-10-07
	// Result: 7 Oktober 2021
	// Result: PPKM diperpanjang hingga 7 Oktober 2021

	data2 := "Password akan dikirimkan ke email email.user@gmail.com"
	email := str.GetStr(data2, 20, -1) // Get last string
	isValid := mail.IsEmailValid(email)
	if isValid {
		obfemail := mail.ObfuscateEmail(email)
		fmt.Println(strings.Replace(data2, email, obfemail, -1)) // Replace data
	}
	// Result: true
	// Result: em*******r@gmail.com
	// Result: Password akan dikirimkan ke email em*******r@gmail.com
}
