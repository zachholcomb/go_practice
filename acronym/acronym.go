package acronym

import (
	"strings"
)

func Abbreviate(input string) (acronym string){
	words := strings.ReplaceAll(input, "-", " ")
	words = strings.ReplaceAll(words, "_", " ")
	wordSet := strings.Fields(string(words))
	for _, word := range wordSet {
		letter := strings.Split(word, "")
		acronym += strings.ToUpper(string(letter[0]))
	}
	return acronym
}