package main

import (
	"fmt"
	"my.api.my/api"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(api.Greet())

	router := gin.Default()
	router.GET("/albums/:id", api.GetAlbumById)
	router.GET("/albums", api.GetAlbums)
	router.POST("/albums", api.PostAlbums)

	router.Run("localhost:8080")
}