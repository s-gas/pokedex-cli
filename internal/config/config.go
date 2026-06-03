package config

import (
	"net/url"
	"fmt"
)

type Config struct {
	UrlLocation	*url.URL
	UrlPokemon	*url.URL
	Next 				string
	Prev 				string
}

func New() (Config, error) {
	uLocation, err := url.Parse("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		return Config{}, fmt.Errorf("New: %w", err)
	}
	uPokemon, err := url.Parse("https://pokeapi.co/api/v2/pokemon")
	if err != nil {
		return Config{}, fmt.Errorf("New: %w", err)
	}
	return Config{
		UrlLocation:	uLocation,
		UrlPokemon:		uPokemon,
		Prev: 				"",
		Next: 				uLocation.String(),
	}, nil
}
