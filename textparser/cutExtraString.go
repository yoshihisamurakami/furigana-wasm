package textparser

func CutExtraString(text string, ruby string) TextAndRuby {
	if len(ruby) == 0 {
		return TextAndRuby{
			Text: text,
			Ruby: ruby,
		}
	}

	reversedText := reverseString(text)
	reversedRuby := reverseString(ruby)

	// reversedTextに対して1文字づつ繰り返し
	reversedTextRunes := []rune(reversedText)
	reversedRubyRunes := []rune(reversedRuby)

	// 文字数が何文字目まで同じか
	someCharsCount := -1
	for rtrIndex, reversedTextRune := range reversedTextRunes {
		if rtrIndex >= len(reversedRubyRunes) {
			break
		}
		if reversedRubyRunes[rtrIndex] == reversedTextRune {
			someCharsCount = rtrIndex
		} else {
			break
		}
	}
	if someCharsCount == -1 {
		return TextAndRuby{
			Text: text,
			Ruby: ruby,
		}
	} else {
		someCharsCount = someCharsCount + 1
		return TextAndRuby{
			Text: removeLastRune(text, someCharsCount),
			Ruby: removeLastRune(ruby, someCharsCount),
		}
	}
}

func reverseString(s string) string {
	// 文字列をruneのスライスに変換
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// 先頭と末尾のruneを交換
		runes[i], runes[j] = runes[j], runes[i]
	}
	// runeスライスを文字列に戻す
	return string(runes)
}

func removeLastRune(s string, num int) string {
	runes := []rune(s)
	if num <= 0 || len(runes) < num {
		return s
	}
	return string(runes[:len(runes)-num])
}
