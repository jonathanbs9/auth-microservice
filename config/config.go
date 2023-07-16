package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresHost          string
	PostgresUser          string
	PostgresPassword      string
	PostgresDBName        string
	PostgresPort          string
	PostgresSSLMode       string
	PostgresTimezone      string
	DefaultAdminFirstName string
	DefaultAdminLastName  string
	DefaultAdminEmail     string
	DefaultAdminPassword  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}
	return &Config{
		PostgresHost:          os.Getenv("DB_HOST"),
		PostgresUser:          os.Getenv("DB_USER"),
		PostgresPassword:      os.Getenv("DB_PASS"),
		PostgresDBName:        os.Getenv("DB_NAME"),
		PostgresPort:          os.Getenv("DB_PORT"),
		PostgresSSLMode:       os.Getenv("DB_SSL_MODE"),
		PostgresTimezone:      os.Getenv("DB_TIMEZONE"),
		DefaultAdminFirstName: os.Getenv("ADMIN_NAME"),
		DefaultAdminLastName:  os.Getenv("ADMIN_LASTNAME"),
		DefaultAdminEmail:     os.Getenv("ADMIN_EMAIL"),
		DefaultAdminPassword:  os.Getenv("ADMIN_PASSWORD"),
	}
}
