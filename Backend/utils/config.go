package utils

import "github.com/caarlos0/env/v11"

type database struct {
	DbString  string `env:"BK_DB_STRING" envDefault:"postgres://postgres:postgres@localhost:5432/pingpong"`
	DbMaxIdle int    `env:"BK_DB_MAX_IDLE" envDefault:"10"`
	DbMaxOpen int    `env:"BK_DB_MAX_OPEN" envDefault:"100"`
}

type config struct {
	Dgn            string `env:"BK_DGN" envDefault:"local"`
	ServerIdentity string `env:"BK_SERVER_IDENTITY" envDefault:"webserver"`
	Database       database
}

func GetConfig() (*config, error) {
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
