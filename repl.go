package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommandRegistry() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func makeUsage() string {
	result := ""
	for _, command := range getCommandRegistry() {
		result += command.name + ": " + command.description + "\n"
	}
	return result
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println(makeUsage())
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
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

func repl(scanner *bufio.Scanner) {
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
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
