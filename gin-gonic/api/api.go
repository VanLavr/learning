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

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, v := range albums {
		if v.ID == id {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"not found this album..."})
}

func PostAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println(err.Error())
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}