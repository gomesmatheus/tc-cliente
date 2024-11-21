package database

import (
	"context"
	"fmt"

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

func NewPostgresDb(url string) (*pgx.Conn, error) {
	config, err := pgx.ParseConfig(url)
	if err != nil {
		fmt.Println("Error parsing config", err)
		return nil, err
	}

	db, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		fmt.Println("Error creating database connection", err)
		return nil, err
	}
	// setup create table
	if _, err := db.Exec(context.Background(), createTables); err != nil {
		fmt.Println("Error creating table Clientes", err)
		return nil, err
	}

	return db, err
}
