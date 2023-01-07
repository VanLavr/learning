package main

import "fmt"

func logger(prf string) (func (messege string)) {
	return func(messege string) {
		fmt.Printf("[%s] %s\n", prf, messege)
	}
}

func main() {
	//warning := logger("WARNING")
	//warning("Incorrect sexual otientation...")
	/*
	y := func(a int) int {
		return a * -1
	}

	func(b int, retfunc func(a int) int) {
		fmt.Println(retfunc(b))
	}(5, y)
	*/

	/*
	x := func(b int, retfunc func(a int) int) {
		fmt.Println(retfunc(b))
	}

	y := func(a int) int {
		return a * -1
	}
	x(5, y)
	*/
}