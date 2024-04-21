package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  string
	PgDriver           string
}

func GetPostgresConfig() (*PostgresConfig, error) {
	log.Println("Loading PostgresConfig")
	err := godotenv.Load()

	if err != nil {
		log.Println("Error occured with loading PostgresConfig")
		return nil, err
	}

	return &PostgresConfig{
		PostgresqlHost:     os.Getenv("host"),
		PostgresqlPort:     os.Getenv("port"),
		PostgresqlUser:     os.Getenv("user"),
		PostgresqlPassword: os.Getenv("password"),
		PostgresqlDbname:   os.Getenv("dbname"),
		PostgresqlSSLMode:  os.Getenv("sslmode"),
		PgDriver:           os.Getenv("driver"),
	}, nil

}
