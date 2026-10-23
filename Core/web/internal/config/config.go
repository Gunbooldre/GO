package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/caarlos0/env/v11"
)

const (
	Envlocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

var cfg Config

type Config struct {
	Env         string `env:"ENV" envDefault:"local"`
	StoragePath string `env:"STORAGE_PATH" envDefault:"./storage"`
	HTTPServer  HTTPServer
}

type HTTPServer struct {
	Port        string        `env:"PORT" envDefault:"8000"`
	Timeout     time.Duration `env:"TIMEOUT" envDefault:"5s"`
	IdleTimeout time.Duration `env:"IDLE_TIMEOUT" envDefault:"60s"`
}

func New() (*Config, error) {
	c := Config{}
	if err := env.Parse(&c); err != nil {
		return nil, fmt.Errorf("config: parse env: %w", err)
	}
	if err := c.validate(); err != nil {
		return nil, err
	}

	cfg = c

	return &cfg, nil
}

func Get() *Config {
	return &cfg
}

func (c *Config) validate() error {
	rules := []struct {
		failed bool
		err    error
	}{
		{strings.TrimSpace(c.Env) == "", fmt.Errorf("config: Env is required")},
	}
	for _, rule := range rules {
		if rule.failed {
			return rule.err
		}
	}
	return nil
}
