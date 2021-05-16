package main

import (
	"fmt"
	"log"
)

func perimeter(width float64, lenght float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("A width of %f is invalid", width)
	} else if lenght < 0 {
		return 0, fmt.Errorf("A lenght of %f is invalid", lenght)
	}
	return (width * 2) + (lenght * 2), nil
}

func main() {
	total, err := perimeter(-10, 5)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("You'll need", total, "meters of fencing")
	}
}
