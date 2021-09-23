package mail

import (
	"strings"
)

// ObfuscateEmail is a function to modify email show in public with asterisk
func ObfuscateEmail(email string) string {
	conv := strings.Split(email, "@")

	// Prepare data
	llocpart := len(conv[0])       // Get length of local-part
	obf := conv[0][2 : llocpart-1] // Get local-part that we will modify into asterisk (obfuscated)
	lobf := len(obf)               // Get length of local-part that will be obfuscated

	// If length local-part less than 6 chars, it's needed because in microsoft it's allowed
	if llocpart < 6 {
		obf = conv[0][1 : llocpart-1]
		lobf = len(obf)
	}

	// Create asterisk as much as local-part that will be obfuscated
	ast := strings.Repeat("*", lobf)

	// Replace it
	localpart := strings.Replace(conv[0], obf, ast, -1)

	return localpart + "@" + conv[1]
}
