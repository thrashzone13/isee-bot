package main

import (
	"log"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func CheckIfError(err error, params ...string) {
	if err != nil {
		if len(params) > 0 {
			log.Fatal(params[0])
			return
		}
		log.Fatal(err)
	}
}

func ThousandSeparator(num int) string {
	p := message.NewPrinter(language.English)
	r := p.Sprintf("%d", num)

	return r
}
