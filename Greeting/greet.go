package greeting

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

func Greet(name string) (string, error) {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return "", errors.New("Please Provide a name")
	}

	message := fmt.Sprintf(randomString(), trimmedName)

	return message, nil
}

func randomString() string {
	format := []string{
		"Welcome! %v How you are Doing!",
		"Hey, %v Welcome to our site",
		"How are You %v",
	}

	return format[rand.Intn(len(format))]
}

func testing() {
	const num string = "dag"

	println(num)

}
