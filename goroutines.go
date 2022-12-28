package main

import (
	"fmt"
	"time"
)

func main() {
	go count("FISH")
	count("sheep")
}

func count(word string) {
	for i := 0; true; i++ {
		fmt.Println(i, word)
		time.Sleep(time.Millisecond * 500)
	}
}