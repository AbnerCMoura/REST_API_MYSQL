package config

import (
	"fmt"
	"os"
)

type Config struct {
	PublicHost string
	Port       string
	DbUser     string
	DbPassword string
	DbAddress  string
	DbName     string
}

var Envs = InitConfig()

func InitConfig() Config {
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "3306"),
		DbUser:     getEnv("DB_USER", "root"),
		DbPassword: getEnv("DB_PASSWORD", "pass"),
		DbAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost"), getEnv("PORT", "3306")),
		DbName:     getEnv("DB_NAME", "ecom"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
