package main

import (
  "fmt"
  "math"
)

// function basic
func swap(x string, y string) (string, string) {
  return y, x
} 
// named return values
func split_bills(sum int) (x,y int) {
  x = sum * 4/9
  y = sum - x
  return
}
// if with a short statement
func pow_limit(x,n,lim float64) float64 {
  if v:= math.Pow(x,n); v < lim {
    return v
  }
  return lim
}
func main() {
  a, b := swap("hello","world")
  fmt.Println(a,b)

  fmt.Println(split_bills(17))
  fmt.Println(
    pow_limit(3,2,10),
    pow_limit(3,3,20),
  )
}
