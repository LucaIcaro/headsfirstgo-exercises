package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// first, open the file
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	// second, scan the file
	// create the scanner
	scanner := bufio.NewScanner(file)
	// loop the scan (so, read lines) until it returns "false"
	for scanner.Scan() {
		// print the line found by the scanner
		fmt.Println(scanner.Text())
	}
	// close the file
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// check for errors in the scanner
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
