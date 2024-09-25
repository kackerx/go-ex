package main

import (
	"fmt"
	"sync"
)

func main() {
	p := sync.Pool{New: func() any {
		return "hehe"
	}}

	fmt.Println(p.Get())
	fmt.Println(p.Get())
	fmt.Println(p.Get())
}
