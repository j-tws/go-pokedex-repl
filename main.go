package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/j-tws/go-pokedex-repl/internal/pokeapi"
	"github.com/j-tws/go-pokedex-repl/internal/pokecache"
)

type config struct {
	next			string
	previous 	*string
}

func main(){
	commandsMap := cliCommandMap()
	scanner := bufio.NewScanner(os.Stdin)
	config := config{}
	newCache := pokecache.NewCache(1 * time.Hour)
	pokedex := map[string]pokeapi.Pokemon{}
	var param string

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		command, ok := commandsMap[input[0]]
		if len(input) > 1 {
			param = input[1]
		}

		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(&config, &newCache, param, pokedex)

		if err != nil {
			fmt.Println(err)
		}
	}
}
