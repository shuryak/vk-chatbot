package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Server `yaml:"server"`
		VK     `yaml:"vk"`
	}

	Server struct {
		Port string `yaml:"port" env:"SERVER_PORT" env-default:":8080"`
	}

	VK struct {
		GroupID         int    `yaml:"group_id" env:"VK_GROUP_ID"`
		Token           string `yaml:"token" env:"VK_TOKEN"`
		ConfirmationKey string `yaml:"confirm_key" env:"VK_CONFIRM_KEY"`
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
