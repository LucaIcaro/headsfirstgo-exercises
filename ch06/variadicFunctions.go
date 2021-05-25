package main

import "fmt"

func exampleSum(numbers ...int) int {
	// a variadic function can have ONLY THE LAST PARAMETER variadic argument
  var sum int
  for _, number := range(numbers) {
	  sum += number
  }
  return sum
}

func main() {
	fmt.Println(exampleSum(1,2,3,4,5))
}