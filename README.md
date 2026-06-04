*This project has been created as part of the Boot.dev course.*

# pokedex-cli

Command-line REPL that fetches data from [PokéAPI](https://pokeapi.co/).

## How to install

Clone the repository and change to the project's directory:

```bash
git clone https://github.com/s-gas/pokedex-cli.git
cd pokedex-cli
```

Install the binary:

```bash
go install
```

## How to run

```bash
pokedex-cli
```

## Commands

- `help`: List all commands
- `exit`: Exit the program
- `map`: List the next 20 location areas
- `mapb`: List the previous 20 location areas
- `explore <location-area>`: List the Pokémon that you can encounter in an area
- `catch <pokemon>`: Tries to catch a Pokémon
- `inspect <pokemon>`: Displays information about the Pokémon
- `pokedex`: List all the Pokémons in the Pokédex
