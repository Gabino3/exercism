// Package bob contains Bob's responses to various phrases
package bob

import (
	"strings"
)

// Hey returns Bob's responses to remarks
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)
	isYelling := remark == strings.ToUpper(remark) && remark != strings.ToLower(remark)
	isQuestion := strings.HasSuffix(remark, "?")

	if remark == "" {
		return "Fine. Be that way!"
	} else if isYelling && isQuestion {
		return "Calm down, I know what I'm doing!"
	} else if isQuestion {
		return "Sure."
	} else if isYelling {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
