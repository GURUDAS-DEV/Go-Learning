package goroutines

import (
	"fmt"
	"time"
)

func Main1() {
	go LoopRoutine()

	fmt.Println("Hello! From Gurudas")
	fmt.Println("Hello! From Gurudas")
	fmt.Println("Hello! From Gurudas")
	time.Sleep(1 * time.Second)
}

func LoopRoutine() {
	for i := 0; i <= 1000; i++ {
		fmt.Printf("%d ", i)
	}
}
