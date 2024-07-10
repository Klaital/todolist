package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"log/slog"
	"os"
)

type Config struct {
	DbHost     string `env:"DB_HOST" envDefault:"localhost"`
	DbUser     string `env:"POSTGRES_USER" envDefault:"todo_tester"`
	DbPassword string `env:"POSTGRES_PASSWORD" envDefault:"todotester123"`
	DbName     string `env:"POSTGRES_DB" envDefault:"todolist_test"`
	DbPort     int    `env:"DB_PORT" envDefault:"5431"`
}

func Load() *Config {
	var c Config
	err := env.Parse(&c)
	if err != nil {
		slog.Error("Failed to parse environment", "error", err)
		os.Exit(1)
	}
	return &c
}

func (c *Config) DbDsn(dbPrefix string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/todolist_test?sslmode=disable", c.DbUser, c.DbPassword, c.DbHost, c.DbPort)
}
