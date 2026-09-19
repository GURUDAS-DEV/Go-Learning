package method

import "fmt"

func PointerWithReciever() {
	master()
}

type val struct {
	x, y string
}

func change(v *val) {
	v.x = "Changed x from outer function"
	v.y = "Changed y explict outer function"
}

func (v *val) change() {
	v.x = "Changed x from explict inner function"
	v.y = "Changed y from explict inner function"
}

func (v val) changeCopy() {
	v.x = "fucking change"
	v.y = "fucking Changed"

}

func master() {
	values := val{"Unchanged", "Unchanged"}
	fmt.Println(values)

	// change(&values)
	// fmt.Println(values)

	v := &values

	v.change()
	fmt.Println(values)
	fmt.Println(v)

	// values.changeCopy()
	// fmt.Println(values)
}
