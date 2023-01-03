package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	fmt.Println("main started...")
	time.Sleep(time.Millisecond * 500)
	var wg sync.WaitGroup

	wg.Add(1)
	go WG(&wg)
	wg.Wait()

	time.Sleep(time.Millisecond * 500)
	fmt.Println("main finished...")
}

func WG(wg *sync.WaitGroup) {
	defer wg.Done()

	
	fmt.Println("executing WG...")
	
	time.Sleep(time.Millisecond * 1500)
	
	fmt.Println("WG executed...")
}