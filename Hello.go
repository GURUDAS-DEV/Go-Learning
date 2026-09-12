package main

import (
	g "Hello/Greeting"
	p "Hello/NumericConstant"
	"errors"
	"fmt"
	m "math"
	"strings"
)

func main() {
	p.Main()
}

func truncuateName(name string) (string, error) {
	if name == "" {
		return "", errors.New("Please don't Provide a empty string")
	}

	return strings.TrimSpace(name), nil
}

func greetFromOutside() {
	var name string

	fmt.Print("Enter your name : ")
	fmt.Scan(&name)

	returnedGreeting, err := g.Greet(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(returnedGreeting)
}

func omitParams(num1, num2, num3 int, name string) {

}

func helper() {
	var a, b int = 3, 4
	var f float64 = m.Sqrt(float64(a*a + b*b))
	var z uint = uint(f)
	fmt.Println(a, b, z)
	var x int = 'w'
	println(x)
	const abs int = 23
}
