//https://golangbyexample.com/interface-in-golang/

package main

import "fmt"

type animal interface {
    breathe()
    walk()
}

type lion struct {
     age int
}

func (l lion) breathe() {
    fmt.Println("Lion breathes")
}

func (l lion) walk() {
    fmt.Println("Lion walk")
}

type dog struct {
     age int
}

func (l dog) breathe() {
    fmt.Println("Dog breathes")
}

func (l dog) walk() {
    fmt.Println("Dog walk")
}

func main() {
    var a animal
    
    a = lion{age: 10}
    a.breathe()
    a.walk()
  
    a = dog{age: 5}
    a.breathe()
    a.walk()
}