package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "sol", "la", "si"}
	fmt.Printf("%#v\n", notes)
	// standard loop on array
	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}
	// alternative loop (better and safer)
	for index, value := range notes {
		fmt.Println(index, value)
	}
	// it is possible to use the blank placeholder on the
	// loop definition
	for _, note := range notes {
		fmt.Println(note)
	}
}
