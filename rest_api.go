package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

var uberPools = [] struct{
  Id int `json:"ID"`
  Destination string `json:"destination"`
  Slots int `json:"slots"`
  Taken int `json:"taken"`
  Total float64 `json:"total"` 
} {
  {1, "Melbourne Central", 4, 1, 15.5},
  {2, "Flinder Station", 4, 2, 12.5},
}

func main() {
  r := gin.Default()
  r.GET("/ping", func (c *gin.Context){
    c.JSON(http.StatusOK, gin.H {
      "message": "pong",
    })
  })

  r.GET("/pools", func (c *gin.Context) {
    c.JSON(http.StatusOK, uberPools)
  })
  r.Run("localhost:8080")
}

