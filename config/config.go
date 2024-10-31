package config

import (
	"simple_rest_crud/pkg/logging"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	IsDebug  *bool    `yaml:"is_debug" env-required:"true"`
	Database PgConfig `yaml:"database"`
	Auth     struct {
		SaltLenth int `yaml:"salt_length"`
	} `yaml:"auth"`
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

func GetConfig(logger *logging.Logger) (*Config, error) {

	var (
		once sync.Once
		cfg  *Config
		err  error
	)

	once.Do(func() {
		logger.WithFields(logrus.Fields{}).Info("reading app conf")

		cfg = &Config{}
		err = cleanenv.ReadConfig("config.yaml", cfg)

	})
	return cfg, err
}
