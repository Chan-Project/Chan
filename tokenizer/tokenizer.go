package tokenizer

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func Tokenize(token string) {
	token = normalize(strings.ToLower(token))
	token2 := removeStopWords(token)
	fmt.Println(token2)
}

// Text preprocessing
func normalize(value string) string {
	for i, v := range map[string]string{"(?i)[èéêë]": "e", "(?i)[àáâäãå]": "a", "(?i)[ç]": "c", "(?i)[îï]": "i", "[',?;.:/!§)(]": ""} {
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
