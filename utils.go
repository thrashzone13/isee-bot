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

func GetProperMessage(status int) string {
	switch status {
	case 0:
		return "Enter salary"
	case 1:
		return "Do you own a house?"
	case 2:
		return "Enter house area"
	case 3:
		return "Enter number of family members"
	case 4:
		return ""
	default:
		log.Fatal("Invalid user status")
	}
	return ""
}
