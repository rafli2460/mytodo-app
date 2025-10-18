package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func Init() map[string]string {
	config, err := godotenv.Read()
	if err != nil {
		log.Panic().Str("Context", "Load Env").Err(err)
	}

	return config
}
