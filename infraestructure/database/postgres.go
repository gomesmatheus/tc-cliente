package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

const (
	createTables = `
    CREATE TABLE IF NOT EXISTS clientes (
        cpf BIGINT PRIMARY KEY,
        nome VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE
    );
    `
)

const maxRetries = 10
const retryInterval = 2 * time.Second

func NewPostgresDb(url string) (*pgx.Conn, error) {
	var db *pgx.Conn
	var err error

	for i := 0; i < maxRetries; i++ {
		fmt.Printf("Attempt %d start", i+1)
		config, err := pgx.ParseConfig(url)
		if err != nil {
			fmt.Println("Error parsing config:", err)
			return nil, err
		}

		db, err = pgx.ConnectConfig(context.Background(), config)
		if err == nil {
			break
		}
		fmt.Printf("Attempt %d end: Error connecting to database: %v\n", i+1, err)
		time.Sleep(retryInterval)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL after %d attempts: %w", maxRetries, err)
	}

	if _, err := db.Exec(context.Background(), createTables); err != nil {
		fmt.Println("Error creating table clientes:", err)
		return nil, err
	}

	return db, nil
}
