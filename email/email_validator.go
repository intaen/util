package mail

import (
	"errors"
	"strings"
)

type EmailProvider struct {
	Google    string
	Microsoft string
	Others    string
}

var SpecialChars = []string{"`", "~", "!", "#", "$", "%", "^", "&", "*", "(", ")", "+", "=", ":", ";", "\"", "'", "\\", "?", "<", ">", "{", "}", "[", "]"}
var Service = EmailProvider{
	Google:    "Google",
	Microsoft: "Microsoft",
	Others:    "Others",
}

// GetService.EmailProvider is a function to give provider from its domain
func GetEmailProvider(domain string) string {
	if strings.Contains(domain, "gmail") {
		return Service.Google
	} else if strings.Contains(domain, "outlook") || strings.Contains(domain, "hotmail") {
		return Service.Microsoft
	}

	return Service.Others
}

// CheckDomain is a function to check if domain contains special chars
func CheckDomain(domain string) error {
	// Check length of domain
	if len(domain) > 255 {
		return errors.New("local-part is more than 255 chars")
	}

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
		if len(localpart) < 6 || len(localpart) > 30 {
			return errors.New("local-part is less than 6 chars or more than 30 chars")
		}

		// Check if data has '.' more than one
		if strings.Count(localpart, ".") > 1 {
			return errors.New("local-part contains '.' (period) more than one")
		}
	} else if provdr == Service.Microsoft || provdr == Service.Others {
		// Check length of local part
		if len(localpart) > 64 {
			return errors.New("local-part is more than 64 chars")
		}
	}

	// Check if data has special character in it
	spchars := strings.Join(SpecialChars, "")
	if strings.ContainsAny(localpart, spchars) {
		return errors.New("local-part contains special characters " + spchars)
	}

	return nil

}
