package tokenizer

import (
	"regexp"
	"strings"
)

func Tokenize(token string){
	token = normalize(strings.ToLower(token))
}

func normalize(value string) string{
	value = regexp.MustCompile("(?i)[èéêë]").ReplaceAllString(value, "e")
	value = regexp.MustCompile("(?i)[àáâäãå]").ReplaceAllString(value, "e")
	value = regexp.MustCompile("(?i)[ç]").ReplaceAllString(value, "c")
	value = regexp.MustCompile("(?i)[îï]").ReplaceAllString(value, "i")
	return value
}