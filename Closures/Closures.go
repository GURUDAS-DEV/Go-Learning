package closures

import "fmt"

func ClousersPractice() {
	main()
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

func adder() func(int) int {
	sum := 10

	return func(i int) int {
		sum += i
		return sum
	}
}
