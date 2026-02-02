package config

import "github.com/mahmud-off/crudUsers/internal/config/yaml"

type Config interface {
	GetDB() string
	GetDBSource() string
}

func NewConfig() (Config, error) {
	return yaml.NewConfig()
}
