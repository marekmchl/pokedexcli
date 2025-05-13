package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Config struct {
	Next     string
	Previous string
}

// cleanInput splits the users string input into "words" based on whitespace, lowercases the input, and trims
// any leading or trailing whitespace.
func cleanInput(text string) []string {
	result := []string{}
	for _, word := range strings.Fields(text) {
		result = append(result, strings.ToLower(word))
	}
	return result
}

func repl(scanner *bufio.Scanner, conf *Config) {
	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			panic(fmt.Errorf("scanner failed"))
		}
		command, found := getCommandRegistry()[cleanInput(scanner.Text())[0]]
		if !found {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(conf)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
