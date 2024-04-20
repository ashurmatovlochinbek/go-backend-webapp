package config

import (
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
	err := godotenv.Load()

	if err != nil {
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

//psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
