package main

import "fmt"

func logger(prf string) (func (messege string)) {
	return func(messege string) {
		fmt.Printf("[%s] %s\n", prf, messege)
	}
}

func main() {
	warning := logger("WARNING")
	warning("Incorrect sexual otientation...")
}