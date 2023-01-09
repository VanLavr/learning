package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Greet() string {
	return fmt.Sprintf("Hello!")
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "rock", Artist: "Egor", Price: 10.99},
	{ID: "2", Title: "rep", Artist: "Ivan", Price: 11.99},
	{ID: "3", Title: "reggy", Artist: "Masha", Price: 12.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}