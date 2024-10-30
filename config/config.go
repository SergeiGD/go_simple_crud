package config

import (
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug  *bool    `yaml:"is_debug" env-required:"true"`
	Database PgConfig `yaml:"database"`
}

type PgConfig struct {
	Username   string `env:"DB_USERNAME" env-required:"true"`
	Password   string `env:"DB_PASSWORD" env-required:"true"`
	Host       string `env:"DB_HOST" env-required:"true"`
	Port       string `env:"DB_PORT" env-required:"true"`
	Database   string `env:"DB_DATABASE" env-required:"true"`
	MaxAttemps int    `yaml:"max_attemps"`
	Timeount   int    `yaml:"timeout"`
	ConnDelay  int    `yaml:"conn_delay"`
}

func GetConfig() (*Config, error) {

	var (
		once sync.Once
		cfg  *Config
		err  error
	)

	once.Do(func() {
		// TODO: logger
		fmt.Println("read application configuration")

		cfg = &Config{}
		err = cleanenv.ReadConfig("config.yaml", cfg)

	})
	return cfg, err
}
