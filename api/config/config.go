package config

import "github.com/caarlos0/env/v6"

type ServiceConfig struct {
	Host string `env:"HOST" envDefault:""`
	Port string `env:"PORT" envDefault:"3000"`
}

func FromEnv() (*ServiceConfig, error) {
	c := &ServiceConfig{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}
	return c, nil
}
