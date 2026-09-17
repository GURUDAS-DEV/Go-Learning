package pointers

import "fmt"

func PointerLearning() {
	var point *int
	val := 23
	point = &val

	fmt.Println("Num : ", val)
	*point = *point + 1

	fmt.Println("Num after editing: ", *point)
}
