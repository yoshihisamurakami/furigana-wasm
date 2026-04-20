package textparser

import (
	"strings"
	"unicode"
)

func Separate(text string, ruby string) []TextAndRuby {
	ret := []TextAndRuby{}

	if !containsHiragana(text) {
		ret = append(ret, TextAndRuby{
			Text: text,
			Ruby: ruby,
		})
		return ret
	}

	sep := extractFirstHiragana(text)

	textSplited := splitBeforeAfter(text, sep)
	rubySplited := splitBeforeAfter(ruby, sep)
	for i, r := range textSplited {
		if r == "" {
			continue
		}
		var rubyPart string
		if i < len(rubySplited) {
			rubyPart = rubySplited[i]
		}
		ret = append(ret, TextAndRuby{
			Text: r,
			Ruby: rubyPart,
		})
	}
	return ret
}

func containsHiragana(s string) bool {
	for _, r := range s {
		if unicode.In(r, unicode.Hiragana) {
			return true
		}
	}
	return false
}

func extractFirstHiragana(s string) string {
	var result []rune
	inHiragana := false

	for _, r := range s {
		if unicode.In(r, unicode.Hiragana) {
			result = append(result, r)
			inHiragana = true
		} else if inHiragana {
			break
		}
	}

	return string(result)
}

// splitBeforeAfter は指定された文字列 s を、部分文字列 sep の前後で分割します。
func splitBeforeAfter(s, sep string) []string {
	index := strings.Index(s, sep)
	if index == -1 {
		return []string{s}
	}
	before := s[:index]
	after := s[index+len(sep):]

	return []string{before, after}
}
