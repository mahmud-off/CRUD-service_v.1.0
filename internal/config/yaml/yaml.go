package yaml

import (
	"os"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	db        string
	db_source string
}

func NewConfig() (*ServerConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(os.Getenv("CFG_PATH"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := ServerConfig{
		db:        viper.GetString("db"),
		db_source: os.Getenv("DB_SOURCE"),
	}
	return &cfg, nil
}

func (s *ServerConfig) GetDB() string {
	return s.db
}

func (s *ServerConfig) GetDBSource() string {
	return s.db_source
}
