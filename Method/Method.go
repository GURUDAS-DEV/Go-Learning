package method

import "fmt"

func Main() {
	PointerWithReciever()
}

type num struct {
	x, y int
}

func (v num) maxFun() int {
	return max(v.x, v.y)
}

func main() {
	val := num{1, 2}
	fmt.Println(val.maxFun())
}
