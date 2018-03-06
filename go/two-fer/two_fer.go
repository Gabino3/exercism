package twofer

import (
	"fmt"
)

// ShareWith returns the string "One for {s}, one for me.", where {s} is the passed in value
// or "you" if an empty string is passed in
func ShareWith(s string) string {
	if s == "" {
		s = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", s)
}
