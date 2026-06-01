package config

type Config struct {
	Url  string
	Next string
	Prev string
}

func Init() Config {
	return Config{
		Url:  "https://pokeapi.co/api/v2/location-area",
		Next: "",
		Prev: "",
	}
}
