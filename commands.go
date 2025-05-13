package main

import (
	"fmt"
	"os"
	"time"

	"github.com/marekmchl/pokedexcli/internal/pokeapi"
	"github.com/marekmchl/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
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
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas",
			callback:    commandMapBack,
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

func commandHelp(conf *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println(makeUsage())
	return nil
}

func commandExit(conf *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

var mapCache = pokecache.NewCache(5 * time.Second)

func commandMap(conf *Config) error {
	url := ""
	if conf.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

	} else {
		url = conf.Next
	}

	data, err := pokeapi.GetMap(url, mapCache)
	if err != nil {
		return err
	}

	conf.Next = data.Next
	conf.Previous = data.Previous

	for _, res := range data.Results {
		fmt.Println(res.Name)
	}

	return nil
}

func commandMapBack(conf *Config) error {
	url := ""
	if conf.Previous == "" {
		fmt.Println("you're on the first page")
		conf.Next = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
		return nil

	} else {
		url = conf.Previous
	}

	data, err := pokeapi.GetMap(url, mapCache)
	if err != nil {
		return err
	}

	conf.Next = data.Next
	conf.Previous = data.Previous

	for _, res := range data.Results {
		fmt.Println(res.Name)
	}

	return nil
}
