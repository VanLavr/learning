package main

import (
	"fmt"
)

func main() {
	var a [120]string
	for i := 0; i < 120; i++ {
		a[i] = "@"
	}
	fmt.Print(a)
}
