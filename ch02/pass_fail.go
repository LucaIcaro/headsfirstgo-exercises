// pass_fail reports whether a grade is passing or failing
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var status string

	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)

	//	input, _ := reader.ReadString('\n')
	/*	the "_" is a "blank identifier", a special type of variable that is
		used to gather data that will be discarded
	*/

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	grade, err := strconv.ParseFloat(input, 64)

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
