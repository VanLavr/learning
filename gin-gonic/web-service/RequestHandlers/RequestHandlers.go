package RequestHandlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album {
	{ID: "1", Title: "Bricks", Artist: "Arctic Monkeys", Price: 13.92},
	{ID: "2", Title: "Do I wanna know", Artist: "Arctic Monkeys", Price: 20.99},
	{ID: "3", Title: "Knee socks", Artist: "Arctic Monkeys", Price: 9.32},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
