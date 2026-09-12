package conditionalstatements

import "fmt"

func ForLoop() {
	var val int
	fmt.Print("Enter a number : ")
	fmt.Scan(&val)

	for i := 1; i < 11; i++ {
		fmt.Printf("%d X %d = %d \n", val, i, val*i)
	}
}

func WhileLoop() {
	var val, i int
	fmt.Print("Enter a number : ")
	fmt.Scan(&val)

	for i <= 10 {
		fmt.Printf("%d X %d = %d \n", val, i, val*i)
		i++
	}
}
