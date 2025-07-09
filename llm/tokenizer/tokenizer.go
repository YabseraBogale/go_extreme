package tokenizer

import (
	"fmt"
	"regexp"
	"strings"
)

type Tokenizer struct {
	StrToInt    map[string]int
	IntToString map[int]string
}

func (t *Tokenizer) Encode(text string) ([]int, error) {
	re := regexp.MustCompile(`([,.:;?_!"()\']|--|\s)`)
	preprocessed := re.Split(text, -1)

	var cleanedPreprocessed []string
	for _, item := range preprocessed {
		trimmedItem := strings.TrimSpace(item)
		if trimmedItem != "" {
			cleanedPreprocessed = append(cleanedPreprocessed, trimmedItem)
		}
	}

	ids := make([]int, len(cleanedPreprocessed))
	for i, s := range cleanedPreprocessed {
		id, ok := t.StrToInt[s]
		if !ok {
			return nil, fmt.Errorf("unknown token: '%s'", s)
		}
		ids[i] = id
	}
	return ids, nil
}
func (t *Tokenizer) Decode(ids []int) (string, error) {
	words := make([]string, len(ids))
	for i, id := range ids {
		s, ok := t.IntToString[id]
		if !ok {
			return "", fmt.Errorf("unknown token ID: %d", id)
		}
		words[i] = s
	}
	text := strings.Join(words, " ")

	// Regular expression to replace spaces before specified punctuations
	re := regexp.MustCompile(`\s+([,.?!"()\'])`)
	text = re.ReplaceAllString(text, `$1`)
	return text, nil
}
