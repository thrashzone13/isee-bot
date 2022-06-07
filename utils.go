package main

import (
	"log"
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
