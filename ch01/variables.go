package main

import (
	"fmt"
	//	"reflect"
)

func main() {
	/*	var quantity int
		var lenght, width float64
		var customerName string

		fmt.Println(customerName, quantity, lenght)
	*/

	quantity := 2
	customerName := "pippo"
	lenght, width := 1.2, 2.4

	fmt.Println(customerName, "has ordered", float64(quantity), " of things ")
	fmt.Println("each one of ", int(lenght*width), "square meters")
}
