package config

import (
	"github.com/caarlos0/env/v7"
)

type Config struct {
	User     string `env:"USER" envDefault:"root"`
	Password string `env:"PASSWORD" envDefault:"root"`
	URL      string `env:"URL" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:":8080"`
}

func NewConfig() *Config {
	cfg := Config{}
	env.Parse(&cfg)
	return &cfg
}
