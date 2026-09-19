package Interface

import "fmt"

func stringerLearn() {
	a := user{"Gurudas", 22}
	z := user{"Naruto", 18}
	fmt.Println(a, z)
}

type user struct {
	Name string
	Age  int
}

func (u user) String() string {
	return fmt.Sprintf("Teri Maaki Chut %s", u.Name)
}
