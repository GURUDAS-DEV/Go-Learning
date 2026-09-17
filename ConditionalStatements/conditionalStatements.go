package conditionalstatements

import (
	"fmt"
	"runtime"
)

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

func IfElse() {
	var age int
	fmt.Print("Enter the Age : ")
	fmt.Scan(&age)

	if age < 18 {
		fmt.Println("You are underage")
	} else {
		fmt.Println("You are Adult")
	}
}

func ShortIfElse() {
	var val int
	fmt.Print("ENTER : ")
	fmt.Scan(&val)
	if age := 12; age+val < 18 {
		fmt.Println("ADULT")
	} else {
		fmt.Println("NOT DEAD")
	}
}

func Swtich() {
	switch os := runtime.GOOS; os {
	case "Darwin":
		fmt.Println("DARWIN")
		break
	case "windows":
		fmt.Println("Windows")
		break
	default:
		fmt.Println("ERROR")
	}
}
