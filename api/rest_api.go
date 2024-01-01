package main

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/lib/pq"
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

const (
  host      = "localhost"
  port      = 5432
  user      = "github"
  password  = "123456"
  dbname    = "Uber"
)

func InitPg() {
  var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  db, err := pq.Open(psqlInfo)

  if err != nil {
    fmt.Println("error: ",err)
  } else {
    fmt.Println("successfully init!")
  }
  defer db.Close()
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
  //r.Run("localhost:8080")

  InitPg()
}

