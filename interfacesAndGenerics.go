package main

import (
	"fmt"
	"strconv"
)

type AnyNumber interface {
	int | int8 | int16 | int32 | float32 | float64
}

func Summation[T AnyNumber](a T, b T) T {
	return a + b
}



type NumCollector interface {
	NumRet() int
}

type ElementOfArray struct {
	value string
	index int
}

type Element struct {
	value string
}

func (e *ElementOfArray) NumRet() int {
	val, err := strconv.Atoi(e.value)
	if err != nil {
		panic(err)
	}
	val += e.index

	return val
}

func (e *Element) NumRet() int {
	val, err := strconv.Atoi(e.value)
	if err != nil {
		panic(err)
	}

	return val
}

func AnyFunction(n NumCollector) {
	fmt.Println(n.NumRet())
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



	var el1 = Element{value: "16"}
	var el2 = ElementOfArray{value: "16", index: 11}

	AnyFunction(&el1)
	AnyFunction(&el2)
}