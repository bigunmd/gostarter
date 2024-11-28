package main

import (
	"fmt"

	"github.com/bigunmd/gostarter/internal/heroes"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
)

func main() {
	filepath := pflag.StringP("config", "c", "", "Configuration file in json/yaml/env format (default:\"\")")
	pflag.Parse()

	cfg := &heroes.Config{}
	if err := readConfig(*filepath, cfg); err != nil {
		log.Fatal().Err(err).Msg("failed to read config")
	}

	if err := heroes.Run(cfg); err != nil {
		log.Fatal().Err(err).Msg("failed to run heroes")
	}
}

func readConfig(filepath string, cfg *heroes.Config) error {
	if filepath != "" {
		if err := cleanenv.ReadConfig(filepath, cfg); err != nil {
			return fmt.Errorf("cannot read configuration file: %w", err)
		}
	} else {
		if err := cleanenv.ReadEnv(cfg); err != nil {
			return fmt.Errorf("cannot read environment: %w", err)
		}
	}

	return nil
}
