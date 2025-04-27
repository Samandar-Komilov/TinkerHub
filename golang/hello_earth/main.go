package main

import (
	"flag"
	"fmt"
)

type language string

var phrasebook = map[language]string{
	"en": "Hello World!",
	"ru": "Привет, мир!",
	"uz": "Assalom dunyo!",
}

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language, e.g. en")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
