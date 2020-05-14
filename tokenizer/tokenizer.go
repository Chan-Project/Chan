package tokenizer

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func Tokenize(token string) []string{
	token = normalize(strings.ToLower(token))
	token2 := removeStopWords(token)
	return token2
}

// Text preprocessing
func normalize(value string) string{
	value = regexp.MustCompile("(?i)[èéêë]").ReplaceAllString(value, "e")
	value = regexp.MustCompile("(?i)[àáâäãå]").ReplaceAllString(value, "e")
	value = regexp.MustCompile("(?i)[ç]").ReplaceAllString(value, "c")
	value = regexp.MustCompile("(?i)[îï]").ReplaceAllString(value, "i")
	value = regexp.MustCompile("[,?;.:/!§)(]").ReplaceAllString(value, "")
	return value
}

func removeStopWords(value string) []string{
	fmt.Println(value)
	data, err := ioutil.ReadFile("./tokenizer/stopwords/french.txt")
	if err != nil {
		return nil
	}
	var nl []string
	for _, v := range strings.Split(value, " "){
		if !contains(v, strings.Split(string(data), "\n")){
			nl = append(nl, v)
		}
	}
	return nl
}

func contains(value string, array []string) bool{
	for _, v := range array{
		if strings.TrimSpace(v) == value{
			return true
		}
	}
	return false
}

func displayArray(array []string){
	for _, v := range array {
		fmt.Printf("* %s\n", v)
	}
}


