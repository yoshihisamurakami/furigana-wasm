package textparser

import (
	"unicode"
)

// 漢字が含まれているかどうかを判定する関数
func ContainsKanji(s string) bool {
	for _, r := range s {
		// Unicode範囲をチェック
		if unicode.In(r, unicode.Han) {
			return true
		}
	}
	return false
}
