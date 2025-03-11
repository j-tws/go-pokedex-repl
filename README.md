# Lil Pokemon CLI

This is an exercise by boot.dev to create a cli app with Go and its HTTP package.

## Usage

List of commands:

- help
- exit
- map
- mapb
- explore
- catch
- inspect
- pokedex

### `help`

Displays all available commands along with it's description

```
Pokedex > help
Welcome to the Pokedex!
Usage:

pokedex: List all pokemons in pokedex
exit: Exit the pokedex
help: Displays a help message
map: Display the next 20 location areas
mapb: Display the previous 20 location areas
explore: Get pokemon of the area
catch: Catch a pokemon!
inspect: Get the stats of a caught pokemon
```

### `exit`

Exits the REPL

### `map` and `mapb`

Fetch the first 20 location areas. Subsequent `map` calls the next 20 location areas.

`mapb` will print out the previous 20 location areas.

`map` will cache the results thus making the same request will just end up fetching the data from the cache instead

```
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
...

Pokedex > map
mt-coronet-1f-route-216
mt-coronet-1f-route-211
mt-coronet-b1f
great-marsh-area-1
great-marsh-area-2
...

Pokedex > mapb
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
...
```

### `explore`

The `explore` commands list out all pokemons available in an area

```
Pokedex > explore trophy-garden-area
pikachu
pichu
roselia
staravia
kricketune
```

### `catch`

Catch a pokemon based on it's base experience. The higher the experience, the harder it is to catch the pokemon

```
Pokedex > catch mew
Throwing a Pokeball at mew...
mew escaped
Pokedex > catch mew
Throwing a Pokeball at mew...
mew was caught!
You may now inspect it with the inspect command.
```

### `inspect`

Print out the stats of a caught pokemon. If the pokemon is not caught, it will return an error message saying the pokemon isn't caught.

```
Pokedex > inspect mew
Name: mew
Height: 4
Weight: 40
Stats:
  -hp: 100
  -attack: 100
  -defense: 100
  -special-attack: 100
  -special-defense: 100
  -speed: 100
Types:
  -psychic
```

### pokedex

List out all caught pokemons in the pokedex.

```
Pokedex > pokedex
Your Pokedex:
  - mew
  - pikachu
```
