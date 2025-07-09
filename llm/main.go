package main

import (
	"fmt"
	"llm/tokenizer"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("the verdict.txt")
	if err != nil {
		log.Println(err)
	}
	var token tokenizer.Tokenizer
	StrToInt := make(map[string]int)
	IntToString := make(map[int]string)
	for i, j := range strings.Split(string(file), " ") {
		IntToString[i] = j
		StrToInt[j] = i
	}
	token.IntToString = IntToString
	token.StrToInt = StrToInt
	result := token.Encode("I HAD always thought Jack Gisburn rather a cheap")
	fmt.Println(result)

	fmt.Println(token.Decode(result))
}
