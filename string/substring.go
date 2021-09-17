package str

import "strings"

// SubStr is a function to trim data with fixed length and get data between it. Return result in string.
func SubStr(data string, firstlen, lastlen int) string {
	firstStr := data[:firstlen]
	lastStr := data[strings.LastIndex(data, "")-lastlen:]

	s1 := strings.TrimPrefix(data, firstStr)
	s2 := strings.TrimSuffix(s1, lastStr)

	return s2
}
