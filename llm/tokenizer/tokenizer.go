package tokenizer

import (
	"regexp"
	"strings"
)

type Tokenizer struct {
	StrToInt    map[string]int
	IntToString string
}

func (t *Tokenizer) Encode(text string) []int {
	re := regexp.MustCompile(`([,.?_!"()\']|--\\s)`)
	preprocessed := re.Split(text, -1)
	var result []int
	for _, item := range preprocessed {
		result = append(result, t.StrToInt[strings.TrimSpace(item)])
	}
	return result
}
