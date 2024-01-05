package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

var Config *config

type config struct {
	Wallet []*wallet `yaml:"wallet"`
}

type wallet struct {
	Name string `yaml:"name"`
	Addr string `yaml:"address"`
}

func LoadConfig(path string) {
	var cfg *config
	if content, e := os.ReadFile(path); e != nil {
		panic("no valid config file")
	} else {
		if e := yaml.Unmarshal(content, &cfg); e != nil {
			panic("config file unmarshal error")
		}
	}

	Config = cfg
}
