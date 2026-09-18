package ArrayLearning

import "fmt"

func Main() {
	LearnArray()
}

func LearnArray() {
	var arr [10]int

	for i := 0; i < 10; i++ {
		fmt.Print("Enter the Number : ")
		fmt.Scan(&arr[i])
	}

	for _, ele := range arr {
		fmt.Printf("Value : %d\n", ele)
	}

	var sli []int = arr[1:5]

	fmt.Println(sli)
}

func experiment() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("The Length of Array : %d, the Cap : %d\n", len(s), cap(s))
	fmt.Printf("Array : %v\n\n", s)
}

func main() {
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func appending() {
	var arr [3]int
	fmt.Print(arr[1])
}
