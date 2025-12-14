package main

import "fmt"

// 双写
type DoubleWrite string

const (
	DoubleWriteSRC_ONLY  DoubleWrite = "SRC_ONLY"
	DoubleWriteSRC_FIRST DoubleWrite = "SRC_FIRST"
	DoubleWriteDST_FIRST DoubleWrite = "DST_FIRST"
	DoubleWriteDST_ONLY  DoubleWrite = "DST_ONLY"
)

func main() {
	fmt.Println("Hello, World!")
}
