package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 100)

	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	for item := range ch {
		fmt.Println(item)
	}

	fmt.Println("main end")
	select {}
}
