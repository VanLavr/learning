package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"MyWebService/RequestHandlers"
)

func main() {
	router := gin.Default()
	router.GET("/albums", RequestsHandlers.GetAlbums)
	
	router.Run("localhost:8080")
}