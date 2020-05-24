package tokenizer

import (
	"fmt"
	"github.com/kljensen/snowball"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"regexp"
	"strings"
)

func Tokenize(originalString string) {
	originalString = normalize(strings.ToLower(originalString))
	originalString = stemmer(originalString, "french")
	token := removeStopWords(originalString)
	fmt.Println(token)
}

// Text preprocessing
func normalize(value string) string {
	for i, v := range map[string]string{
		"(?i)[èéêë]":    "e",
		"(?i)[àáâäãå]":  "a",
		"(?i)[ç]":       "c",
		"(?i)[îï]":      "i",
		"[',?;.:/!§)(]": " ",
		"[\\d]":         ""} {
		value = regexp.MustCompile(i).ReplaceAllString(value, v)
	}
	return value
}

func removeStopWords(value string) (nl []string) {
	fmt.Println(value)
	data, err := ioutil.ReadFile("./tokenizer/stopwords/french.txt")
	if err != nil {
		return nil
	}
	for _, v := range strings.Split(value, " ") {
		if !contains(v, strings.Split(string(data), "\n")) {
			nl = append(nl, v)
		}
	}
	return
}

func contains(value string, array []string) bool {
	for _, v := range array {
		if strings.TrimSpace(v) == value {
			return true
		}
	}
	return false
}

func displayArray(array []string) {
	for _, v := range array {
		fmt.Printf("* %s\n", v)
	}
}

// Stemmer function handler kljensen
func stemmer(str, lang string) (data string) {
	data, err := snowball.Stem(str, lang, false)
	if err != nil {
		log.Fatal("Error during the stemming", err)
		return
	}
	return
}
