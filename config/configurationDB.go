package config

import (
	"github.com/spf13/viper"
)

type APIConfig struct {
	ApiPort string
	ApiHost string
}

type DBConfig struct {
	Host string
	Port string
	User string
	Password string
	DBName string
}

type Config struct {
	APIConfig
	DBConfig
}

func (c Config) ReadConfigFile() Config {

	c.DBConfig = DBConfig{
		Host: viper.GetString("HOST"),
		Port: viper.GetString("PORT"),
		User: viper.GetString("USER"),
		Password: viper.GetString("PASSWORD"),
		DBName: viper.GetString("DBNAME"),
	}

	c.APIConfig = APIConfig{
		ApiPort: viper.GetString("APIPORT"),
		ApiHost: viper.GetString("APIHOST"),
	}

	return c
}

func NewConfig() Config{
	cfg := Config{}

	return cfg.ReadConfigFile()
}

