package main

import "strings"

func cleanInput(text string) []string{
	output := []string{}
	slice := strings.Split(strings.Trim(text, " "), " ")

	for _, word := range slice {
		output = append(output, strings.ToLower(word))
	}

	return output
}