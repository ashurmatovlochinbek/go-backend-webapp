package postgres

import (
	"database/sql"
	"fmt"
	"simple-go-app/config"
	"time"

	_ "github.com/lib/pq"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

func NewPsqlDB(c *config.PostgresConfig) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		c.PostgresqlHost,
		c.PostgresqlPort,
		c.PostgresqlUser,
		c.PostgresqlDbname,
		c.PostgresqlSSLMode,
		c.PostgresqlPassword,
	)

	db, err := sql.Open(c.PgDriver, dataSourceName)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
