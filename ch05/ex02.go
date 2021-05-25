package main

import "fmt"

func main() {
	numbers := [9]int{1, 0, 2, 0, 1, 0, 0, 1, 2}
	var occurences [3]int

	for _, number := range numbers {
		occurences[number]++
	}

	for index, value := range occurences {
		fmt.Printf("%d occured %d times\n", value, index)
	}
}
