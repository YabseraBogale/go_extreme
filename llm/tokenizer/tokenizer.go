package tokenizer

import (
	"regexp"
	"strings"
)

type Tokenizer struct {
	StrToInt    map[string]int
	IntToString map[int]string
}

func (t *Tokenizer) Encode(text string) []int {
	re := regexp.MustCompile(`([,.?_!"()\']|--\\s)`)
	preprocessed := re.Split(text, -1)
	var ids []int
	for _, item := range preprocessed {
		ids = append(ids, t.StrToInt[strings.TrimSpace(item)])
	}
	return ids
}

func (t *Tokenizer) Decode(ids []int) string {
	var keys []string
	for k := range t.IntToString {
		keys = append(keys, t.IntToString[k])
	}
	re := regexp.MustCompile(`\s+([,.?!"()\\]')`)
	text := strings.Join(keys, " ")
	text = re.ReplaceAllLiteralString(text, `\\1`)
	return text
}
