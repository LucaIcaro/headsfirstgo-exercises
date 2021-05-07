package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "GX rXcks!"
	replacer := strings.NewReplacer("X", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
