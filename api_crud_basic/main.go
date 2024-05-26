package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album {
	{ID: "1", Title: "Something to give each other", Artist: "Troye Sivan", Price: 12.22},
	{ID: "2", Title: "The best of me", Artist: "Bryan Adams", Price: 10.22},
}

func getAll(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

func add(c *gin.Context){
	var newAlbum album
	// bind the received JSON
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getById(c *gin.Context){
	var id = c.Query("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	// get size of albums
	if len(albums) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	} else {
		getAll(c)
	}
}


func main() {
	r := gin.Default()
	// r.GET("/albums", getAll)
	r.GET("/albums", getById)
	r.POST("/albums", add)
	r.Run("localhost:8080")
}
