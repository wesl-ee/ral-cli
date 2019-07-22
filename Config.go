package main

import (
	"os"

	"github.com/naoina/toml"
)

type Config struct {
	Scheme string `toml:"scheme"`
	Site string `toml:"site"` }

func ReadConfig(path string) (config *Config, err error) {
	config = new(Config)
	f, err := os.Open(path)
	if err != nil { return }
	defer f.Close()

	err = toml.NewDecoder(f).Decode(config)
	if err != nil { return }

	return
}
