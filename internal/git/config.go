package git

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed config.json
var assets embed.FS

func Assets() embed.FS {
	return assets
}

type Config struct {
	GitCommand   string
	Repositories []Repository
}

type Repository struct {
	Path   string
	Remote string
	Name   string
}

func DefaultConfig() *Config {
	b, err := assets.ReadFile("config.json")
	if err != nil {
		panic(fmt.Sprintf("1 there was a problem reading the config file: %s", err))
	}

	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		panic(fmt.Sprintf("3 there was a problem parsing the config file: %s", err))
	}

	return &cfg
}
