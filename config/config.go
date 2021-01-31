package config

import (
	"github.com/caarlos0/env"
	"time"
)

type App struct {
	Name  string `env:"APP_NAME" envDefault:"news-aggregator"`
	Debug bool   `env:"APP_DEBUG"`
}

type HTTPConfig struct {
	Port uint `env:"HTTP_SERVER_PORT"`
}

type Database struct {
	User        string        `env:"DB_USER" envDefault:"test"`
	Pass        string        `env:"DB_PASS" envDefault:"test"`
	Host        string        `env:"DB_HOST" envDefault:"localhost"`
	Port        uint          `env:"DB_PORT" envDefault:"8086"`
	Name        string        `env:"DB_NAME" envDefault:"news-db"`
	MaxConns    int           `env:"DB_MAX_CONNS"`
	MaxLifetime time.Duration `env:"DB_MAX_LIFETIME"`

	Timeout time.Duration `env:"DB_TIMEOUT" envDefault:"30s"`
	Retries int           `env:"DB_RETRIES" envDefault:"5"`
}

type Logger struct {
	PodName      string `env:"LOGGER_POD_NAME"  envDefault:"kube01"`
	PodNode      string `env:"LOGGER_POD_NODE" envDefault:"node01"`
	PodNamespace string `env:"LOGGER_POD_NAMESPACE" envDefault:"namespace01"`

	Address string `env:"LOGGER_ADDRESS" envDefault:"10.110.0.70:12201"`
	Level   string `env:"LOGGER_LEVEL" envDefault:"DEBUG"`
}

type Config struct {
	App
	Logger
	HTTPConfig
	Database
}

func (c *Config) Parse() (err error) {
	if err = env.Parse(&c.HTTPConfig); err != nil {
		return err
	}

	if err = env.Parse(&c.Logger); err != nil {
		return err
	}

	if err = env.Parse(&c.App); err != nil {
		return err
	}

	if err = env.Parse(&c.Database); err != nil {
		return err
	}

	return
}
