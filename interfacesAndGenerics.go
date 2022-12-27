package main

import "fmt"

type Greeter interface {
	greet(string)
}

type GoodGreet struct {}
type BadGreet struct {}

func (gg *GoodGreet) greet(s string) {
	fmt.Printf("Nice to meet you, %s!\n", s)
}

func (bg *BadGreet) greet(s string) {
	fmt.Printf("F*** you, %s!\n", s)
}

func SayHi(g Greeter, name string) {
	g.greet(name)
}

func main() {
	var (
		GoodBoy = "Ivan"
		BadBoy = "Kalim"
		G = GoodGreet{}
		B = BadGreet{}
	)

	SayHi(&B, BadBoy)
	SayHi(&G, GoodBoy)
}