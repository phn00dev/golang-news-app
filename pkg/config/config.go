package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HttpServer     httpServer
	FolderConfig   folderConfig
	DatabaseConfig databaseConfig
	RedisConfig    redisConfig
}

type httpServer struct {
	Host    string
	Port    string
	Header  string
	AppName string
}

type folderConfig struct {
	PublicFolder string
}

type databaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SslMode  string
	TimeZone string
}

type redisConfig struct {
	Host string
	Port string
}

func GetConfig() (*Config, error) {
	if err := godotenv.Load("../.env"); err != nil {
		return nil, err
	}
	cfg := &Config{
		HttpServer: httpServer{
			Host:    os.Getenv("HTTP_SERVER_HOST"),
			Port:    os.Getenv("HTTP_SERVER_PORT"),
			Header:  os.Getenv("HTTP_SERVER_HEADER"),
			AppName: os.Getenv("HTTP_SERVER_APP_NAME"),
		},
		FolderConfig: folderConfig{
			PublicFolder: os.Getenv("PUBLIC_FOLDER"),
		},
		DatabaseConfig: databaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			SslMode:  os.Getenv("DB_SLL_MODE"),
			TimeZone: os.Getenv("DB_TIME_ZONE"),
		},
		RedisConfig: redisConfig{
			Host: os.Getenv("REDIS_HOST"),
			Port: os.Getenv("REDIS_PORT"),
		},
	}

	return cfg, nil
}
