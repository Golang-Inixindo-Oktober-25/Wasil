package helper

import (
	"fmt"
	"regexp"
)

func Info(msg string) {
	fmt.Println("[INFO]: ", msg)
}

func Error(msg string) {
	fmt.Println("[ERROR]: ", msg)
}

func IsEmailValid(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	match, err := regexp.MatchString(pattern, email)
	if err != nil {
		Error(err.Error())
		return false
	}

	if match {
		return true
	} else {
		return false
	}
}
