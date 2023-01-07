package main

import (
	"fmt"
	"net/http"
)

func main() {
	resourse, err := http.Get("https://example.com/")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("StatusCode:", resourse.StatusCode)
}