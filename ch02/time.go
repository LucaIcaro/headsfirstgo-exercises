package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(year)
}

// methods are functions associated wiht values of a given type

// functions belongs to a package ; methods belong to a value
