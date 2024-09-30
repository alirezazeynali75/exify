package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Configs struct {
	App   App       `envPrefix:"APP_"`
	Mysql Mysql     `envPrefix:"MYSQL_"`
	A     AProvider `envPrefix:"A_"`
	B     BProvider `envPrefix:"B_"`
	Http  Http      `envPrefix:"HTTP_"`
}

func Configure() (*Configs, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("reading .env file error: %w", err)
	}

	cfg := &Configs{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("parsing configuration error: %w", err)
	}

	return cfg, nil
}
