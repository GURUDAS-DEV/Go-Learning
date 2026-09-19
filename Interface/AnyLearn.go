package Interface

import "fmt"

func AnyLearn() {
	var val any
	val = "Hello"
	fmt.Println(val)

	val = 2343
	fmt.Println(val)
}
