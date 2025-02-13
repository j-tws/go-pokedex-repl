package main

import (
	"fmt"
	"os"
)

const locationAreaURL = "https://pokeapi.co/api/v2/location-area"

type cliCommand struct {
	name					string
	description 	string
	callback			func(*config) error
}

type locationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func cliCommandMap() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Display the next 20 location areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Display the previous 20 location areas",
			callback: commandMapB,
		},
	}
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for key, value := range cliCommandMap() {
		fmt.Printf("%s: %s\n", key, value.description)
	}

	return nil
}

func commandMap(c *config) error {
	var url string

	if c.next == "" {
		url = locationAreaURL
	} else {
		url = c.next
	}

	locationArea, err := makeGetRequest(url)

	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	for _, area := range locationArea.Results {
		fmt.Println(area.Name)
	}
	
	setConfig(c, locationArea)

	return nil
}

func commandMapB(c *config) error {
	if c.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url := *c.previous

	locationArea, err := makeGetRequest(url)

	for _, area := range locationArea.Results {
		fmt.Println(area.Name)
	}

	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	setConfig(c, locationArea)

	return nil
}