package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database   Database
	ServerPort int `envconfig:"SERVER_PORT"`
}

type Database struct {
	Host     string `envconfig:"DATABASE_HOST" required:"true"`
	Port     int    `envconfig:"DATABASE_PORT" required:"true"`
	User     string `envconfig:"DATABASE_USER" required:"true"`
	Password string `envconfig:"DATABASE_PASSWORD" required:"true"`
	Name     string `envconfig:"DATABASE_NAME" required:"true"`
}

func NewParseConfig() (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("[INIT] Unable to load configuration. %+v\n", err)
	}

	cnf := Config{}
	err = envconfig.Process("", &cnf)
	return cnf, err
}
