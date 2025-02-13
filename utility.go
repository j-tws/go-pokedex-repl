package main

import (
	"encoding/json"
	"fmt"
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

func makeGetRequest(url string) (locationArea, error) {
	var locationArea locationArea
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

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationArea); err != nil {
		return locationArea, fmt.Errorf("Error decoding response body: %w", err)
	}

	return locationArea, nil
}

func setConfig(c *config, l locationArea){
	c.next = l.Next
	c.previous = l.Previous
}