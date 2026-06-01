package config

import (
	"net/url"
	"fmt"
)

type Config struct {
	Url  *url.URL
	Next string
	Prev string
}

func New() (Config, error) {
	u, err := url.Parse("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		return Config{}, fmt.Errorf("New: %w", err)
	}
	return Config{
		Url:	u,
		Prev: "",
		Next: "https://pokeapi.co/api/v2/location-area",
	}, nil
}
