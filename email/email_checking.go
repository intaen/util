package mail

import (
	"errors"
	"log"
	"strings"
)

type EmailProvider struct {
	Google    string
	Microsoft string
}

var SpecialChars = []string{"`", "~", "!", "#", "$", "%", "^", "&", "*", "(", ")", "+", "=", ":", ";", "\"", "'", "\\", "?", "<", ">", "{", "}", "[", "]"}
var Service = EmailProvider{
	Google:    "Google",
	Microsoft: "Microsoft",
}

// IsEmailValid is a function to check if email valid based on length and chars
func IsEmailValid(email string) (bool, string, string) {
	// Check if email has '@' sign more than one
	if strings.Count(email, "@") > 1 {
		log.Println("Email Invalid! It's contain '@' sign more than one")
		return false, "", ""
	}

	// Split email into local-part and domain
	conv := strings.Split(email, "@")
	if conv[0] == "" || conv[1] == "" {
		log.Println("Email Invalid! Local-part or domain is empty")
		return false, "", ""
	}
	provdr := GetEmailProvider(conv[1]) // Get email service provider

	// Check domain
	err := CheckDomain(conv[1])
	if err != nil {
		log.Println("Email Invalid!", err)
		return false, "", ""
	}

	// Check local-part
	err = CheckLocalPart(conv[0], provdr)
	if err != nil {
		log.Println("Email Invalid!", err)
		return false, "", ""
	}

	return true, conv[0], conv[1]
}

// GetService.EmailProvider is a function to give provider from its domain
func GetEmailProvider(domain string) string {
	if strings.Contains(domain, "gmail") {
		return Service.Google
	} else if strings.Contains(domain, "outlook") || strings.Contains(domain, "hotmail") {
		return Service.Microsoft
	}

	return ""
}

// CheckDomain is a function to check if domain contains special chars
func CheckDomain(domain string) error {
	// Check if data has '.' more than one
	if strings.Count(domain, ".") > 1 {
		return errors.New("domain contains '.' (period) more than one")
	}

	// Check if data has special character in it
	spchars := strings.Join(SpecialChars, "")
	if strings.ContainsAny(domain, spchars+"-_") {
		return errors.New("domain contains special characters " + spchars + "-_")
	}

	return nil
}

// CheckLocalPart is a function to check if local-part or username from each provider already qualified (based on length or special chars on it)
func CheckLocalPart(localpart, provdr string) error {
	if provdr == Service.Google {
		// Append more special character
		SpecialChars = append(SpecialChars, "-_")

		// Check length of local part
		if len(localpart) < 5 {
			return errors.New("local-part is less than 5 chars")
		}

		// Check if data has '.' more than one
		if strings.Count(localpart, ".") > 1 {
			return errors.New("local-part contains '.' (period) more than one")
		}
	}

	// Check if data has special character in it
	spchars := strings.Join(SpecialChars, "")
	if strings.ContainsAny(localpart, spchars) {
		return errors.New("local-part contains special characters " + spchars)
	}

	return nil

}
