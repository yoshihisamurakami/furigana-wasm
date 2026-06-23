package textparser

import (
	ipaneologd "github.com/ikawaha/kagome-dict-ipa-neologd"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

type TextAndRuby struct {
	Text string `json:"text"`
	Ruby string `json:"ruby"`
}

func Parse(text string) (map[string]interface{}, error) {
	t, err := tokenizer.New(ipaneologd.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return nil, err
	}

	var furiganaDetails []interface{}
	for _, v := range t.Tokenize(text) {
		reading, _ := v.Reading()
		readingHiragana := KatakanaToHiragana(reading)
		if ContainsKanji(v.Surface) {
			textAndRuby := TextAndRuby{
				Text: v.Surface,
				Ruby: readingHiragana,
			}
			cutExtra := CutExtraString(textAndRuby.Text, textAndRuby.Ruby)
			separate := Separate(cutExtra.Text, cutExtra.Ruby)
			for _, tr := range separate {
				furiganaDetails = append(furiganaDetails, map[string]interface{}{
					"text": tr.Text,
					"ruby": tr.Ruby,
				})
			}
		}
	}

	return map[string]interface{}{
		"originalText":    text,
		"furiganaDetails": furiganaDetails,
	}, nil
}
