package main

import (
	"encoding/json"
	"fmt"
	"github.com/j-tws/go-pokedex-repl/internal/pokecache"
	"github.com/j-tws/go-pokedex-repl/internal/pokeapi"
	"io"
	"net/http"
	"strings"
)

type pokeStructs interface {
	locationArea | locationAreaDetails | pokeapi.Pokemon
}

func cleanInput(text string) []string{
	output := []string{}
	slice := strings.Split(strings.Trim(text, " "), " ")

	for _, word := range slice {
		output = append(output, strings.ToLower(word))
	}

	return output
}

func makeGetRequest[T pokeStructs](url string, cache *pokecache.Cache, pokeStruct T) (T, error) {
	// go hit the cache
	cachedData, exist := cache.Get(url)

	if exist {
		if err := json.Unmarshal(cachedData, &pokeStruct); err != nil {
			return pokeStruct, fmt.Errorf("Error decoding response body: %w", err)
		}
	
		return pokeStruct, nil
	}

	// if cache miss
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return pokeStruct, fmt.Errorf("Error forming request: %w", err)
	}

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return pokeStruct, fmt.Errorf("Error making request: %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return pokeStruct, fmt.Errorf("Error reading response: %w", err)
	}

	cache.Add(url, data)

	if err := json.Unmarshal(data, &pokeStruct); err != nil {
		return pokeStruct, fmt.Errorf("Error decoding response body: %w", err)
	}

	return pokeStruct, nil
}

func setConfig(c *config, l locationArea){
	c.next = l.Next
	c.previous = l.Previous
}