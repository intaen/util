# util

Compilation of util and can be use to simplify your code and make it easier.

## String

Simple library to cut-off string. With this, you can get first, sub (middle), or last string as simple as that.
To get first string, use GetStr with 1 for parameter pos; last string, use GetStr with -1 for parameter pos; sub
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
	return str.GetStr(phn, 3, 1)
}

func GetStudentMiddleName(name string, firstlen, lastlen int) string {
	return str.SubStr(name, firstlen, lastlen)
}

func GetLastPhoneNumber(phn string) string {
	return str.GetStr(phn, 4, -1)
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

# Time
## Month

Simple library to get name of month (full or short form) or number of month in Indonesia.

## Installation

Download and install library:
```go
go get github.com/intaen/go-time
```

## Import
Import in your code
```go
import "github.com/intaen/go-month"
```

## Example
```go
package main

import (
	"fmt"

	"github.com/intaen/go-month"
)

func main() {
    // Example to get date with name of month
    data := "Andi akan pergi ke Amerika pada 2021-09-01"
	conv := strings.Split(data, " ")
	month := inamonth.ConvertDate(conv[6])

	var dt []string
	dt = append(dt, conv[0], conv[1], conv[2], conv[3], conv[4], conv[5], month)
	fmt.Println(strings.Join(dt, " "))
	// Result: Andi akan pergi ke Amerika pada 1 September 2021

	// Example to number of month
    data := inamonth.April
	fmt.Println(int(data))
	// Result: 4
}
```

# go-main

Simple library to check if email format is valid. You can obfuscate email with asterisk, so it'll be hidden.

## Installation

Download and install library:
```go
go get github.com/intaen/go-mail
```

## Import
Import in your code
```go
import "github.com/intaen/go-mail"
```

## Example
```go
package main

import (
	"fmt"

	"github.com/intaen/go-mail"
)

func main() {
	email := "testing.gomail@gmail.com"
	isValid, localpart, domain := IsEmailValid(email)
	if isValid {
		email = ObfuscateEmail(localpart, domain)
		fmt.Println(email)
	}
	// Result: true, testing.gomail, gmail
    // Result: te***********l@gmail.com
}
```