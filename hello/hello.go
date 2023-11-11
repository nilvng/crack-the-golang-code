package main

import "fmt"
import "rsc.io/quote"
func main(){
	var posInt uint64 = 12345
	//var negInt int32 = -12330

	fmt.Println(quote.Glass())

	var floatNum float32 = 10.1
	var sum = floatNum + float32(posInt)
	fmt.Println(sum)
}
