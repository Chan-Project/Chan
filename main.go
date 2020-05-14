package main

import (
	"./tokenizer"
	"fmt"
)

func main()  {
	fmt.Println(tokenizer.Tokenize("Et salut toi, aimerais-tu une glace chocolatée ?"))
}
