package main

import (
	g "Hello/Greeting"
	"errors"
	"fmt"
	"strings"
)

func main() {
	var name string

	fmt.Print("Enter your name : ")
	fmt.Scan(&name)

	returnedGreeting, err := g.Greet(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(returnedGreeting)
}

func truncuateName(name string) (string, error) {
	if name == "" {
		return "", errors.New("Please don't Provide a empty string")
	}

	return strings.TrimSpace(name), nil
}
