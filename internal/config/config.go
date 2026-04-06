package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port    int               `json:"port"`
	Domains map[string]string `json:"domains"`
}

func Load() (*Config, error) {
	file, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}