package mail

import (
	"log"
	"strings"
)

// IsEmailValid is a function to check if email valid based on length and chars
func IsEmailValid(email string) bool {
	// Check if email has '@' sign more than one
	if strings.Count(email, "@") > 1 {
		log.Println("Email Invalid! It's contain '@' sign more than one")
		return false
	}

	// Split email into local-part and domain
	conv := strings.Split(email, "@")
	if conv[0] == "" || conv[1] == "" {
		log.Println("Email Invalid! Local-part or domain is empty")
		return false
	}

	provdr := GetEmailProvider(conv[1]) // Get email service provider

	// Check domain
	err := CheckDomain(conv[1])
	if err != nil {
		log.Println("Email Invalid!", err)
		return false
	}

	// Check local-part
	err = CheckLocalPart(conv[0], provdr)
	if err != nil {
		log.Println("Email Invalid!", err)
		return false
	}

	return true
}
