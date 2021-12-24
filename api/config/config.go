package config

import "github.com/caarlos0/env/v6"

type ServiceConfig struct {
	Host           string `env:"HOST" envDefault:""`
	Port           string `env:"PORT" envDefault:"8080"`
	FrontURLOrigin string `env:"FRONT_URL_ORIGIN,required"`
}

func FromEnv() (*ServiceConfig, error) {
	c := &ServiceConfig{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}
	return c, nil
}
