package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/j-tws/go-pokedex-repl/internal/pokeapi"
	"github.com/j-tws/go-pokedex-repl/internal/pokecache"
)

const (
	locationAreaURL = "https://pokeapi.co/api/v2/location-area/"
	pokemonURL = "https://pokeapi.co/api/v2/pokemon/"
)

type cliCommand struct {
	name					string
	description 	string
	callback			func(*config, *pokecache.Cache, string, map[string]pokeapi.Pokemon) error
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

type locationAreaDetails struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
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
		"explore": {
			name: "explore",
			description: "Get pokemon of the area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Catch a pokemon!",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Get the stats of a caught pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "List all pokemons in pokedex",
			callback: commandPokedex,
		},
	}
}

func commandExit(_ *config, _ *pokecache.Cache, _ string, _ map[string]pokeapi.Pokemon) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *config, _ *pokecache.Cache, _ string, _ map[string]pokeapi.Pokemon) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for key, value := range cliCommandMap() {
		fmt.Printf("%s: %s\n", key, value.description)
	}

	return nil
}

func commandMap(c *config, cache *pokecache.Cache, _ string, _ map[string]pokeapi.Pokemon) error {
	var url string
	var locationArea locationArea

	if c.next == "" {
		url = locationAreaURL
	} else {
		url = c.next
	}

	locationAreaData, err := makeGetRequest(url, cache, locationArea)

	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	for _, area := range locationAreaData.Results {
		fmt.Println(area.Name)
	}

	setConfig(c, locationAreaData)

	return nil
}

func commandMapB(c *config, cache *pokecache.Cache, _ string, _ map[string]pokeapi.Pokemon) error {
	var locationArea locationArea

	if c.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url := *c.previous
	locationAreaData, err := makeGetRequest(url, cache, locationArea)

	for _, area := range locationAreaData.Results {
		fmt.Println(area.Name)
	}

	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	setConfig(c, locationAreaData)

	return nil
}

func commandExplore(c *config, cache *pokecache.Cache, areaName string, _ map[string]pokeapi.Pokemon) error {
	url := locationAreaURL + areaName

	var locationAreaDetails locationAreaDetails

	locationAreaDetailsData, err := makeGetRequest(url, cache, locationAreaDetails)

	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	for _, pokemonEncounter := range locationAreaDetailsData.PokemonEncounters {
		fmt.Println(pokemonEncounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(
	c *config,
	cache *pokecache.Cache,
	pokemonName string,
	pokedex map[string]pokeapi.Pokemon,
) error {
	url := pokemonURL + pokemonName
	var pokemonStruct pokeapi.Pokemon

	pokemonData, err := makeGetRequest(url, cache, pokemonStruct)

	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	num := rand.Intn(pokemonData.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)

	if num < 100 {
		pokedex[pokemonName] = pokemonData
		fmt.Printf("%v was caught!\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%v escaped\n", pokemonName)
	}

	return nil
}

func commandInspect(
	_ *config,
	_ *pokecache.Cache,
	pokemon string,
	pokedex map[string]pokeapi.Pokemon,
) error {
	p, exist := pokedex[pokemon]

	if !exist {
		return fmt.Errorf("you have not caught that pokemon")
	}

	base := fmt.Sprintf("Name: %v\nHeight: %v\nWeight: %v\n", p.Name, p.Height, p.Weight)
	stats := "Stats:\n"
	types := "Types:\n"

	for _, stat := range p.Stats {
		stats = stats + fmt.Sprintf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	for _, stat := range p.Types {
		types = types + fmt.Sprintf("  -%v\n", stat.Type.Name)
	}

	fmt.Println(base + stats + types)

	return nil
}

func commandPokedex(
	_ *config,
	_ *pokecache.Cache,
	_ string,
	pokedex map[string]pokeapi.Pokemon,
) error {
	fmt.Println("Your Pokedex:")

	for pokemonName := range pokedex {
		fmt.Printf("  - %v\n", pokemonName)
	}

	return nil
}