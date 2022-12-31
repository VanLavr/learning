package main

import (
	"fmt"
	"time"
)

func main() {
	out1 := make(chan string)
	out2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			out1 <- "A: doing pullup"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Millisecond * 2500)
			out2 <- "BB: doing cluster"
		}
	}()

	for {
		select {
		case messege := <- out1:
			fmt.Println(messege)

		case messege := <- out2:
			fmt.Println(messege)
		}
	}
}