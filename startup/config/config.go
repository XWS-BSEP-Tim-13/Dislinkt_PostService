package config

import "os"

type Config struct {
	Port            string
	PostDBHost      string
	PostDBPort      string
	AccessKey       string
	SecretAccessKey string
}

func NewConfig() *Config {
	return &Config{
		Port:            os.Getenv("POST_SERVICE_PORT"),
		PostDBHost:      os.Getenv("POST_DB_HOST"),
		PostDBPort:      os.Getenv("POST_DB_PORT"),
		AccessKey:       os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}
}
