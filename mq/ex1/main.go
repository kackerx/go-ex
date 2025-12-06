package main

import "fmt"

func main() {

}

type Person struct {
	Name string
	age  int
}

func (p *Person) SayHello() {
	a := 1
	fmt.Println(a)
	fmt.Println("Hello, my name is", p.Name)
}
