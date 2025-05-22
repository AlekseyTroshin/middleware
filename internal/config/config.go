package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Env         string `yaml:"env"`
	StoragePath string `yaml:"storage_path"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address string `yaml:"address"`
}

func MustConfig() *Config {
	root, err := os.Getwd()
	if err != nil {
		log.Fatalf("CONFIG is not set")
	}

	configPath := filepath.Join(root, "config", "config.yaml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
