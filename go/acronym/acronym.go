// Package acronym should have a package comment that summarizes what it's about.
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	acronym := ""

	lookingForLetter := true
	for _, c := range s {
		if lookingForLetter && unicode.IsLetter(c) {
			lookingForLetter = false
			acronym += string(c)
		}
		if lookingForLetter == false && unicode.IsLetter(c) == false {
			lookingForLetter = true
		}
	}

	return strings.ToUpper(acronym)
}
