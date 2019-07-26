package main

import (
	"os"
	"path"

	"github.com/naoina/toml"
)

type Config struct {
	Scheme string `toml:"scheme"`
	Endpoint string `toml:"endpoint"` }

func FindConfig() (string) {
	home, _ := os.UserHomeDir()
	locations := [2]string{
		"config.toml",
		path.Join(home, ".config.toml") }

	for _, l := range locations {
		if FileExists(l) { return l }
	}
	return ""
}

func FileExists(fpath string) (bool) {
	info, err := os.Stat(fpath)
	if os.IsNotExist(err) {
		return false }
	return !info.IsDir()
}

func ReadConfig(fpath string) (config *Config, err error) {
	config = new(Config)
	f, err := os.Open(fpath)
	if err != nil { return }
	defer f.Close()

	err = toml.NewDecoder(f).Decode(config)
	if err != nil { return }

	return
}
