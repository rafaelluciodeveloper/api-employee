package configs

import (
	"os"
)

var cfg *config

type config struct {
	DB DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Load() error {
	cfg = new(config)

	cfg.DB = DBConfig{
		Host:     os.Getenv("HOST_DB"),
		Port:     os.Getenv("PORT_DB"),
		User:     os.Getenv("USER_DB"),
		Password: os.Getenv("PASSWORD_DB"),
		Database: os.Getenv("DATABASE_DB"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}
