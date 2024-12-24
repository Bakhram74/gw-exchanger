package config

import (
	"log"

	"github.com/Bakhram74/gw-exchanger/pkg"
	"github.com/joho/godotenv"
)

type Config struct {
	Env     string
	Storage StorageConfig

	HttpPort string
}

type StorageConfig struct {
	PostgresHost     string
	PostgresPort     string
	PostgresDatabase string
	PostgresUsername string
	PostgresPassword string
	PostgresSslMode  string
}

func NewConfig() Config {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalf("Error loading config.env file")
	}
	storage := StorageConfig{
		PostgresHost:     pkg.GetEnv("HOST_DB", "localhost"),
		PostgresPort:     pkg.GetEnv("PORT_DB", "5432"),
		PostgresDatabase: pkg.GetEnv("DATABASE", "wallet"),
		PostgresUsername: pkg.GetEnv("USERNAME_DB", "postgres"),
		PostgresPassword: pkg.GetEnv("PASSWORD_DB", "secret"),
		PostgresSslMode:  pkg.GetEnv("SSL_MODE", "disable"),
	}

	config := Config{

		// HttpPort: pkg.GetEnv("HTTP_PORT", "9090"),
		Env:      pkg.GetEnv("ENVIRONMENT", "local"),
		Storage:  storage,
	}
	return config
}
