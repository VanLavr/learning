package main

import (
	"fmt"
	"time"
)

func main() {
	out1 := make(chan string)
	out2 := make(chan string)
	out3 := make(chan string)
	out4 := make(chan string)

	go ProcessManager(out1, "A")
	go ProcessManager(out2, "B")
	go ProcessManager(out3, "C")
	go ProcessManager(out4, "D")
	
	for {
		message1,  opened1 := <- out1
		fmt.Println(message1)
		time.Sleep(time.Millisecond * 500)

		for {
			message2, opened2 := <- out2
			fmt.Println(message2)
			time.Sleep(time.Millisecond * 500)
			if !opened2 {
				break
			}
		}

		for {
			message3, opened3 := <- out3
			fmt.Println(message3)
			time.Sleep(time.Millisecond * 500)
			if !opened3 {
				break
			}
		}

		for {
			message4, opened4 := <- out4
			fmt.Println(message4)
			time.Sleep(time.Millisecond * 500)
			if !opened4 {
				break
			}
		}

		if !opened1 {
			break
		}
	}

	fmt.Scanln()
}

func ProcessManager(out chan string, processName string) {
	defer close(out)
	out <- fmt.Sprintf("process %s started...\n", processName)
	out <- fmt.Sprintf("process %s finished...\n", processName)
}