package textparser

import (
	"strings"
)

// カタカナをひらがなに変換する関数
func KatakanaToHiragana(katakana string) string {
	var result strings.Builder

	for _, r := range katakana {
		switch {
		// 長音記号はそのまま
		case r == 'ー':
			result.WriteRune(r)

		// カタカナ（ァ〜ヶ）をひらがなへ変換
		case r >= 'ァ' && r <= 'ヶ':
			result.WriteRune(r - 0x60)

		default:
			result.WriteRune(r)
		}
	}
	return result.String()
}
