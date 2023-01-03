package main

import (
	"fmt"
	"time"
	"sync"
)

var (
	Lock = sync.Mutex{}
	money int = 100
)

func main() {
	fmt.Printf("money count start: %d\n", money)

	go spend()
	go earn()

	time.Sleep(time.Millisecond * 3000)
	fmt.Printf("money count finish: %d\n", money)
//	fmt.Scanln()
}

func spend() {
	for i := 0; i < 1001; i++ {
		time.Sleep(time.Millisecond * 1)
		Lock.Lock()
		money -= 10
		Lock.Unlock()
	}
	fmt.Println("spending done...")
//  fmt.Println(*money)
}

func earn() {
	for i := 0; i < 1001; i++ {
		time.Sleep(time.Millisecond * 1)
		Lock.Lock()
		money += 10
		Lock.Unlock()
	}
	fmt.Println("earning done...")
//	fmt.Println(*money)
}