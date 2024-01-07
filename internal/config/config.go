package config

import (
	"os"
)

type Config struct {
	HttpServer HttpServer
	DB         Database
	Admin      Admin
}

type HttpServer struct {
	Host string
	Port string
}
type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Admin struct {
	Email    string
	Password string
}

var config *Config

func GetConfig() *Config {
	if config != nil {
		return config
	}

	return &Config{
		HttpServer: HttpServer{
			Host: os.Getenv("HTTP_HOST"),
			Port: os.Getenv("HTTP_PORT"),
		},
		DB: Database{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Name:     os.Getenv("POSTGRES_NAME"),
		},
		Admin: Admin{
			Email:    os.Getenv("ADMIN_EMAIL"),
			Password: os.Getenv("ADMIN_PASSWORD"),
		},
	}
}
