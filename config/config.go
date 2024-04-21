package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port           string `json:"port"`
	Endpoint       string `json:"endpoint"`
	Host           string `json:"host"`
	User           string `json:"user"`
	Password       string `json:"password"`
	Schema         string `json:"schema"`
	MaxConnections int    `json:"max_connections"`
}

func ReadConfig(filename string) (*Config, error) {
	var config Config

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
