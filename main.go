package main

import (
	"bufio"
	"fmt"
	"os"
)

type config struct {
	next			string
	previous 	*string
}

func main(){
	commandsMap := cliCommandMap()
	scanner := bufio.NewScanner(os.Stdin)
	config := config{}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		command, ok := commandsMap[input[0]]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		command.callback(&config)
	}
}
