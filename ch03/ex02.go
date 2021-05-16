package main

import "fmt"

func perimeter(width float64, lenght float64) float64 {
	return (width * 2) + (lenght * 2)
}

func main() {
	total := perimeter(8.2, 10) + perimeter(5, 5.4) + perimeter(6.2, 4.5)
	fmt.Println("You'll need", total, "meters of fencing")
}
