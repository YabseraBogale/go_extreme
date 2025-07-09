package tokenizer

import (
	"regexp"
	"strings"
)

type Tokenizer struct {
	StrToInt    map[string]int
	IntToString map[int]string
}

func (t *Tokenizer) Init(b []byte) {
	strToInt := make(map[string]int)
	intToString := make(map[int]string)
	for i, j := range strings.Split(string(b), " ") {
		intToString[i] = j
		strToInt[j] = i
	}
	t.IntToString = intToString
	t.StrToInt = strToInt
}

func (t *Tokenizer) Encode(text string) []int {
	re := regexp.MustCompile(`([,.:;?_!"()\']|--|\s)`)
	preprocessed := re.Split(text, -1)
	var ids []int
	for _, item := range preprocessed {
		i, ok := t.StrToInt[strings.TrimSpace(item)]
		if !ok {
			ids = append(ids, -1)
		} else {
			ids = append(ids, i)
		}

	}
	return ids
}

func (t *Tokenizer) Decode(ids []int) string {
	var keys []string
	for k := range ids {
		if k == -1 {
			keys = append(keys, t.IntToString[k])
		} else {
			keys = append(keys, "<|unk|>")
		}

	}
	text := strings.Join(keys, " ")
	re := regexp.MustCompile(`\s+([,.?!"()\'])`)
	text = re.ReplaceAllString(text, `$1`)
	return text
}
