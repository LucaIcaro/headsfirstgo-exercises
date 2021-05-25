package main

import "fmt"

func main() {
	var mySlice []int        // this defines a slice with nil value
	mySlice = make([]int, 7) // the slice is now a 7-elements slice
	fmt.Println(mySlice)

	notes := []string{"do", "re", "mi", "fa", "sol", "la", "si"} // this slice is prepopulated

	sliceOperator := notes[1:3]
	fmt.Println(sliceOperator)

	// every slice has an underlying array

	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	fmt.Println(s1, s2, s3, s4)
	s4[0] = "XX"
	fmt.Println(s1, s2, s3, s4)
}
