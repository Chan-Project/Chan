package tokenizer

import (
	"fmt"
	"strings"
)

func Tokenize(token string){
	fmt.Printf("%q\n", strings.Split(token, " "))
}