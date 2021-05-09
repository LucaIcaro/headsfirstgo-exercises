package main

import "fmt"

func main() {
	var width, height, area float64
	width, height = 4.2, 3.0
	area = width * height

	fmt.Printf("%f\n", area)
	fmt.Println(area)

	resultString := fmt.Sprintf("%.2f\n", area)
	fmt.Println(resultString)

	fmt.Printf("The %s cost %d cents each.\n", "gumballs", 23)
	fmt.Printf("That will be $%0.2f please.\n", 0.23*5)

	fmt.Printf("%12s | %s\n", "Testing", "spaces with minimum width")
}
