package bob

import (
	"strings"
	"regexp"
)

func Hey(remark string) string {
	if IsUpperCase(remark) && IsQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	} else if IsUpperCase(remark) {
		return "Whoa, chill out!"
	} else if IsQuestion(remark) {
		return "Sure."
	} else if !HasWord(remark) {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}

func HasWord(remark string) bool {
	hasWord, _ := regexp.MatchString(`\w+`, remark)
	return hasWord
}

func IsUpperCase(remark string) bool {
	hasLetter, _ := regexp.MatchString(`[A-Z]+[a-z]{0}`, remark)
	return hasLetter && remark == strings.ToUpper(remark)
}

func IsQuestion(remark string) bool {
	isQuestion, _ := regexp.MatchString(`[\?]+[\s]*$`, remark)
	return isQuestion
}