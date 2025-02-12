package main

import (
	"bufio"
	"fmt"
	"os"
)

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
