package pointers

import "fmt"

func PointerLearning() {
	var point *int
	val := 23
	point = &val

	fmt.Println("Pointer value : ", *point)
	fmt.Println("real value : ", val)

}
