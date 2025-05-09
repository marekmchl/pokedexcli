package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// cleanInput splits the users string input into "words" based on whitespace, lowercases the input, and trims
// any leading or trailing whitespace.
func cleanInput(text string) []string {
	result := []string{}
	for _, word := range strings.Fields(text) {
		result = append(result, strings.ToLower(word))
	}
	return result
}

func repl() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			panic(fmt.Errorf("scanner failed"))
		}
		fmt.Printf("\nYour command was: %v\n", cleanInput(scanner.Text())[0])
	}
}
