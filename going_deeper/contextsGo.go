package main

import (
	"fmt"
	"context"
	"time"
)

const (
	user = "uid"
)

func main() {
	ctx := context.WithValue(context.Background(), user, 12345)
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		fmt.Scanln()
		cancel()
	}()

	someFunc(ctx)
}

func someFunc(ctx context.Context) {
	id := ctx.Value(user)
	
	select {
	case <- time.After(time.Millisecond * 5000):
		fmt.Printf("Process for id %d done...\n", id)
	case <- ctx.Done():
		fmt.Printf("Canceled...\n")
	}
}