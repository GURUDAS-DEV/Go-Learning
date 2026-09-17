package deferLearn

import "fmt"

func LearnDefer() {
	num := 20
	defer fmt.Println("Number is : ", num)
	num = 30
}

func LearnDeferFunc() {
	num := 20

	defer func() {
		fmt.Println(num)
	}()

	num = 1
}

func MultipleDefer() {
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	defer fmt.Println("Defer 3")
}
