package acronym

import (
	"strings"
)

func Abbreviate(input string) (acronym string){
	wordSet := SeparateWords(input)
	for _, word := range wordSet {
		letter := strings.Split(word, "")
		acronym += strings.ToUpper(string(letter[0]))
	}
	return acronym
}

func SeparateWords(input string) []string {
	words := strings.ReplaceAll(input, "-", " ")
	words = strings.ReplaceAll(words, "_", " ")
	return strings.Fields(string(words))
}