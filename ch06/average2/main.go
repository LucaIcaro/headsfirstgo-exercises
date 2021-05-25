// average calculates the average of several numbers
package main

import (
	"datafile"
	"fmt"
	"log"
	"os"
)

func average(numbers ...float64) float64 {
        var sum float64 = 0
        for _, number := range numbers {
                sum += number
        }
        return sum / float64(len(numbers))
}

func main () {
        // this function will require the file as argument
        arguments := os.Args[1:]
        numbers, err := datafile.GetFloats(arguments[0])
        if err != nil {
                log.Fatal(err)
        }
        // the ellipsis will transform the slice in variadic arguments
        fmt.Printf("Average: %0.2f\n", average(numbers...))
}
