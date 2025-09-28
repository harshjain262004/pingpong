package utils

import "github.com/caarlos0/env/v11"

type config struct {
	ServerIdentity string `env:"BK_SERVER_IDENTITY" envDefault:"webserver"`
}

func GetConfig() (*config, error) {
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
