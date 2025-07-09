package main

import (
	"fmt"
	"llm/tokenizer"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("the verdict.txt")
	if err != nil {
		log.Println(err)
	}
	var token tokenizer.Tokenizer
	token.Init(file)

	result := token.Encode("Hello World")
	fmt.Println(result)

	fmt.Println(token.Decode(result))
}
