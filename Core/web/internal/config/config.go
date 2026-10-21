package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"development" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-default:"./storage" env-required:"true"`
	HTTPServer  HTTPServer
}

type HTTPServer struct {
	Port        string        `yaml:"port" env-default:"8000"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"5s"`
}

func MustLoad() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("CONFIG_PATH does not exist %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig
}
