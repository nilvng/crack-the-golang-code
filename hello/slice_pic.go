package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
  
  matrix := make([][]uint8, dy)

  for i, _ := range matrix {
    matrix[i] = make([]uint8,dx)
    for j, _ := range matrix[i] {
      matrix[i][j] = uint8 (i*j)
    }
  }
  return matrix
}

func main() {
  pic.Show(Pic)
}
