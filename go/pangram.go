package pangram

import (
	"regexp"
	"strings"
)

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	re := regexp.MustCompile("[^a-z]")
	input = re.ReplaceAllString(input, "")
	letters := map[rune]bool{}

	for _, letter := range input {
		letters[letter] = true
	}

	return len(letters) == 26
}
