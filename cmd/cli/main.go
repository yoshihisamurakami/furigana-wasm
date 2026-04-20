package main

import (
	"encoding/json"
	"fmt"
	"os"

	textparse "furigana-wasm/textparser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: cli <text>")
		os.Exit(1)
	}

	result, err := textparse.Parse(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(result)
}
