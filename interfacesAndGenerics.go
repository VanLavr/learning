package main

import "fmt"

type AnyNumber interface {
	int | int8 | int16 | int32 | float32 | float64
}

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

func Summation[T AnyNumber](a T, b T) T {
	return a + b
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

	res1 := Summation(1, 3)
	res2 := Summation(1.1, 3.3)
	fmt.Println(res1, res2)
}