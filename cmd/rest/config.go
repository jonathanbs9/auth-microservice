package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	PostgresHost              string
	PostgresUser              string
	PostgresPassword          string
	PostgresDBName            string
	PostgresPort              string
	PostgresSSLMode           string
	PostgresTimezone          string
	DefaultAdminFirstName     string
	DefaultAdminLastName      string
	DefaultAdminEmail         string
	DefaultCredentialPassword string
}

func loadConfig() *config {
	err := godotenv.Load()
	fmt.Println("DB_HOST:")
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}
	return &config{
		PostgresHost:              os.Getenv("DB_HOST"),
		PostgresUser:              os.Getenv("DB_USER"),
		PostgresPassword:          os.Getenv("DB_PASS"),
		PostgresDBName:            os.Getenv("DB_NAME"),
		PostgresPort:              os.Getenv("DB_PORT"),
		PostgresSSLMode:           os.Getenv("DB_SSL_MODE"),
		PostgresTimezone:          os.Getenv("DB_TIMEZONE"),
		DefaultAdminFirstName:     os.Getenv("ADMIN_NAME"),
		DefaultAdminLastName:      os.Getenv("ADMIN_LASTANAME"),
		DefaultAdminEmail:         os.Getenv("ADMIN_EMAIL"),
		DefaultCredentialPassword: os.Getenv("CREDENTIAL_PASSWORD"),
	}
}
