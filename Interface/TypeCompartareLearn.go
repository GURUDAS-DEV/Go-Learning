package Interface

import "fmt"

func TypeCompLearn() {
	s1 := []int{1, 2, 3, 5, 6, 7, 234, 5, 2, 4, 63, 3}
	fmt.Println(Index(s1, 2))

	s2 := []string{"Hello", "Moo", "Cow", "Dog"}
	fmt.Println(Index(s2, "Hello"))
}

func Index[T comparable](s []T, val T) int {
	for idx, ele := range s {
		if val == ele {
			return idx
		}
	}
	return -1
}
