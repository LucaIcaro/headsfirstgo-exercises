package main

import "fmt"

func main() {
	weekdays := [5]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	for index, day := range weekdays {
		fmt.Println(index, day)
	}
}
