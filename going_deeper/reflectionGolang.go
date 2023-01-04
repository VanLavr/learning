package main

import (
	"fmt"
)

type A int

func (a A) String() string {
	return "aaaa"
}

func main() {
	var a A = 12
	fmt.Println(a.String())

	b := &a
	fmt.Println(b.String())

	fmt.Printf("a: %d, b(=&a): %p", a, b)
}