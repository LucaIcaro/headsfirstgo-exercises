package main

import "fmt"

func main() {
	amount := 6
	double(&amount)
	fmt.Println(amount)
	// fmt.Println(&amount)

	// var myInt int
	// fmt.Println(reflect.TypeOf(&myInt))
	// fmt.Println(reflect.TypeOf(myInt))

	// var myIntPointer *int
	// fmt.Println(reflect.TypeOf(myIntPointer))
	// myIntPointer = &myInt
	// fmt.Println(*myIntPointer)
	// fmt.Println(myIntPointer)
	// myInt = 4
	// fmt.Println(*myIntPointer)
	// fmt.Println(myIntPointer)
	// *myIntPointer = 8
	// fmt.Println(*myIntPointer)
	// fmt.Println(myInt)

	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)

}

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func double(number *int) {
	*number *= 2
}
