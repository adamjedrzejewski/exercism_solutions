package isogram

import (
	"strings"
	"unicode/utf8"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	word = strings.ReplaceAll(word, " ", "")
	word = strings.ReplaceAll(word, "-", "")

	uniqueLetters := map[rune]bool{}
	for _, c := range word {
		uniqueLetters[c] = true
	}

	return utf8.RuneCountInString(word) == len(uniqueLetters)
}
