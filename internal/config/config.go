package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Server `yaml:"server"`
		PG     `yaml:"postgres"`
		Redis  `yaml:"redis"`
		VK     `yaml:"vk"`
		Logger `yaml:"logger"`
	}

	PG struct {
		PoolMax int    `yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env:"PG_URL"`
	}

	Redis struct {
		Host     string `yaml:"host" env:"REDIS_HOST"`
		Port     string `yaml:"port" env:"REDIS_PORT"`
		Name     string `yaml:"name" env:"REDIS_NAME"`
		Password string `yaml:"password" env:"REDIS_PASSWORD"`
	}

	Server struct {
		Port string `yaml:"port" env:"SERVER_PORT" env-default:":8080"`
	}

	VK struct {
		GroupID         int    `yaml:"group_id" env:"VK_GROUP_ID"`
		Token           string `yaml:"token" env:"VK_TOKEN"`
		ConfirmationKey string `yaml:"confirm_key" env:"VK_CONFIRM_KEY"`
	}

	Logger struct {
		Level string `yaml:"level" env:"LOG_LEVEL" env-default:"debug"`
	}
)

func ParseFileAndEnv(path string) (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %v", err)
	}

	// Env variables have more priority than YAML-declared
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func ParseEnv() (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
