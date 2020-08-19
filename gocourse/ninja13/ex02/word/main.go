package word

import "strings"

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

func Count(s string) int {
	inWhitespace := true
	isWhitespaceRune := true

	if s == "" {
		return 0
	}

	wordCounter := 0
	for _, runeValue := range s {
		if runeValue == ' ' || runeValue == '\t' {
			isWhitespaceRune = true
		} else {
			isWhitespaceRune = false
		}

		if inWhitespace && !isWhitespaceRune {
			wordCounter++
			inWhitespace = false
		} else if !inWhitespace && isWhitespaceRune {
			inWhitespace = true
		}
	}

	return wordCounter
}
