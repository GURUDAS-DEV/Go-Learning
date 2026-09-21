package goroutines

import "fmt"  

func Channels() {
	ch := make(chan int)

	go func() {
		ch <- 42s
	}()

	value := <-ch

	fmt.Println(value)
}
