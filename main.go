package main

import (
	"syscall/js"

	textparse "furigana-wasm/textparser"

	"github.com/ikawaha/kagome-dict/ipa"
)

func tokenize(_ js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	result, err := textparse.Parse(args[0].String())
	if err != nil {
		return nil
	}
	return result
}

func registerCallbacks() {
	_ = ipa.Dict()
	js.Global().Set("kagome_tokenize", js.FuncOf(tokenize))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
