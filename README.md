# util

Compilation of util i make and can be use to simplify your code and make it easier.

## String

Simple library to cut-off string. With this, you can get first, sub (middle), or last string as simple as that.
To get first string, use GetStr with -1 for parameter pos; last string, use GetStr with 1 for parameter pos; sub
string, use SubStr and send total string in first and last position you want to cut.

## Installation

Download and install library:
```go
go get github.com/intaen/go-str
```

## Import
Import in your code
```go
import "github.com/intaen/go-str"
```

## Example
```go
package main

import (
	"fmt"

	"github.com/intaen/go-str"
)

func GetCountryCode(phn string) string {
	return str.GetStr(phn, -1, 3)
}

func GetStudentMiddleName(name string, firstlen, lastlen int) string {
	return str.SubStr(name, firstlen, lastlen)
}

func GetLastPhoneNumber(phn string) string {
	return str.GetStr(phn, 1, 4)
}

func main() {
	// Example to get first string
	fmt.Println(GetCountryCode("+628593980XXXX"))
	// Result: +62

	// Example to get sub string
	fmt.Println(GetStudentMiddleName("John Doe Smith", 5, 6))
	// Result: Doe

	// Example to get last string
	fmt.Println(GetLastPhoneNumber("+628593980XXXX"))
	// Result: XXXX
}

```