package main

import (
	"encoding/json"
	"fmt"
	"internal/pokecache"
	"io"
	"net/http"
	"strings"
)

func cleanInput(text string) []string{
	output := []string{}
	slice := strings.Split(strings.Trim(text, " "), " ")

	for _, word := range slice {
		output = append(output, strings.ToLower(word))
	}

	return output
}

func makeGetRequest(url string, cache *pokecache.Cache) (locationArea, error) {
	var locationArea locationArea
	// go hit the cache
	cachedData, exist := cache.Get(url)

	if exist {
		if err := json.Unmarshal(cachedData, &locationArea); err != nil {
			return locationArea, fmt.Errorf("Error decoding response body: %w", err)
		}
	
		return locationArea, nil
	}

	// if cache miss
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return locationArea, fmt.Errorf("Error forming request: %w", err)
	}

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return locationArea, fmt.Errorf("Error making request: %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return locationArea, fmt.Errorf("Error reading response: %w", err)
	}

	cache.Add(url, data)

	if err := json.Unmarshal(data, &locationArea); err != nil {
		return locationArea, fmt.Errorf("Error decoding response body: %w", err)
	}

	return locationArea, nil
}

func setConfig(c *config, l locationArea){
	c.next = l.Next
	c.previous = l.Previous
}