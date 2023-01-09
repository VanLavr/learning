package main

import (
	"fmt"
	"my.api.my/api"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hi!")
	fmt.Println(api.Greet())

	router := gin.Default()
	router.GET("/albums", api.GetAlbums)
	router.Run("localhost:8080")
}