package main

import "fmt"

func main() {
  var a [2]string // list of 2 string
  a[0] = "Hello"
  a[1] = "World"
  fmt.Println(a[0], a[1]) // Hello World

  primes := [6]int {2,3,5,7,11,13}
  fmt.Println(primes)

  var sl []int = primes[1:4] // from element 1 - 3
  fmt.Println(sl)
  var sl2 = primes[1:3]
  sl2[0] = 0
  fmt.Println(
    sl,
    sl2,
  )
}
