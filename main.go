package main

import (
	"fmt"
	"strings"
)

// cleanInput splits the users string input into "words" based on whitespace, lowercases the input, and trims
// any leading or trailing whitespace.
func cleanInput(text string) []string {
	return strings.Fields(text)
}

func main() {
	fmt.Println("Hello, World!")
}
