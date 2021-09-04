package main

import "fmt"

func printLines(array []string) {
	for _, line := range array {
		fmt.Println(line)
	}
}

func main() {
	daysOfWeek := [7]string{"Sunday", "Monday", "Tuesday",
		"Wednesday", "Thursday", "Friday", "Saturday"}
	workingDays := daysOfWeek[1:6]
	printLines(workingDays)
}
