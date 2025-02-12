package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name					string
	description 	string
	callback			func() error
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
	}
}

func main(){
	commandsMap := cliCommandMap()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		command, ok := commandsMap[input[0]]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		command.callback()
	}
}

func cleanInput(text string) []string{
	output := []string{}
	slice := strings.Split(strings.Trim(text, " "), " ")

	for _, word := range slice {
		output = append(output, strings.ToLower(word))
	}

	return output
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for key, value := range cliCommandMap() {
		fmt.Printf("%s: %s\n", key, value.description)
	}

	return nil
}
