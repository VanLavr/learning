//U should NEVER close a channel where u r recieving data...

package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	outer1 := make(chan string)
	outer2 := make(chan string)

	go count("FISH", outer1)
	go count("sheep", outer2)

	for {
		messege, opened1 := <- outer1
		messege += " "
		msg, opened2 := <- outer2
		messege += msg

		if !(opened1 || opened2) {
			break
		}

		fmt.Println(messege)
	}
}

func count(word string, out chan string) {
	defer close(out)
	
	for i := 0; i < 5; i++ {
		NewI := strconv.Itoa(i)
		NewI += word
		time.Sleep(time.Millisecond * 500)
		out <- NewI
	}
}