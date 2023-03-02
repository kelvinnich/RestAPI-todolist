package config

import (
	"log"

	"github.com/spf13/viper"
)

type APIConfig struct {
	ApiPort string
	ApiHost string
}

type DBConfig struct {
	Host string
	Port int
	User string
	Password string
	DBName string
}

type Config struct {
	APIConfig
	DBConfig
}

func (c Config) ReadConfigFile() Config {

	vi := viper.New()
	vi.SetConfigFile(".env")
	err := vi.ReadInConfig()
	if err != nil {
		log.Printf("failed to read env file")
	}
	c.DBConfig = DBConfig{
		Host: vi.GetString("HOST"),
		Port: vi.GetInt("PORT"),
		User: vi.GetString("USER"),
		Password: vi.GetString("PASSWORD"),
		DBName: vi.GetString("DBNAME"),
	}

	c.APIConfig = APIConfig{
		ApiPort: vi.GetString("APIPORT"),
		ApiHost: vi.GetString("APIHOST"),
	}

	return c
}

func NewConfig() Config{
	cfg := Config{}

	return cfg.ReadConfigFile()
}

