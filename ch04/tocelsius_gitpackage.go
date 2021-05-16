// tocelsius converts a temperature from
// fahrenheit to celsius

package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/keyboard"
)

func main() {
	fmt.Print("Enter a temperature in Fahreneheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit -32) *5 /9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}