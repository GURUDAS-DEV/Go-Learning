package method

import "fmt"

type BigInt int32

func (a BigInt) give(b BigInt) BigInt {
	return max(b, a)
}

func MethodToType() {
	var val BigInt = 123123434
	fmt.Println(val.give(23))
	fmt.Println(val)
}
