package main

import (
	"./tokenizer"
	"fmt"
)

func main()  {
	fmt.Println(tokenizer.Tokenize("Vous avez fait vos preuves en gestion de projet pendant 3 ans minimum"))
}
