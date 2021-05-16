// pass_fail reports whether a grade is passing or failing
package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	var status string

	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()

	if err != nil {
		log.Fatal(err)
	}

	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
