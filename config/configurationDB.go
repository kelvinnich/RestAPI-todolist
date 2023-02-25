package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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

	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Printf("failed to load env file %v", errEnv)
	}
	c.DBConfig = DBConfig{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		User: os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		DBName: os.Getenv("DBNAME"),
	}

	c.APIConfig = APIConfig{
		ApiPort: os.Getenv("APIPORT"),
		ApiHost: os.Getenv("APIHOST"),
	}

	return c
}

func NewConfig() Config{
	cfg := Config{}

	return cfg.ReadConfigFile()
}

