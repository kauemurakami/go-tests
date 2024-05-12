package addresses

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Verify if address contains a valid type in first word
func AddressType(address string) string {
	validTypes := []string{
		"street", "avenue", "road", "highway",
	}
	// address in lowercase
	lowercaseAddress := strings.ToLower(address)
	// Split text in array separing peer empty spaces
	// ex split with empty space result 0-RUA 1-ABC 2-DEF
	// and set in firstWordAddress recovering position 0
	// of the created array
	firstWordAddress := strings.Split(lowercaseAddress, " ")[0]
	isValid := false

	for _, t := range validTypes {

		if t == firstWordAddress {

			isValid = true
		}
	}

	if isValid {
		caser := cases.Title(language.BrazilianPortuguese)
		return caser.String(firstWordAddress) // To uppercase first letter
	}
	return "Invalid type"
}
