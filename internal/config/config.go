package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port   int    `yaml:"port"`
	Host   int    `yaml:"host"`
	DBpath string `yaml:"db_path"`
}

func Init() (*Config, error) {
	file, err := os.Open("config.yml")
	if err != nil {
		return nil, fmt.Errorf("cant get config: %v", err)
	}

	var cfg Config

	if err = yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
