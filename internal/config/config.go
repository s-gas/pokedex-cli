package config

type Config struct {
	Url  string
	Next string
	Prev string
}

func New() Config {
	return Config{
		Prev: "",
		Next:  "https://pokeapi.co/api/v2/location-area",
	}
}
