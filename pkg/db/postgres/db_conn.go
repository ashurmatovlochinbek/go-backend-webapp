package postgres

import (
	"database/sql"
	"fmt"
	"log"
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
		log.Println("Could not connect to database: db_conn.go:33")
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)

	if err = db.Ping(); err != nil {
		log.Println("Could not ping to database: db_conn.go:43")
		return nil, err
	}

	log.Println("Succesfully connected to database: db_conn.go:47")
	return db, nil
}
