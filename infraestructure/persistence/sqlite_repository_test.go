package persistence

import (
	"context"
	"database/sql"
	"testing"

	"github.com/gomesmatheus/tc-cliente/domain/entity"
	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open in-memory database: %v", err)
	}

	createTableQuery := `
	CREATE TABLE clientes (
		cpf INTEGER PRIMARY KEY,
		nome TEXT NOT NULL,
		email TEXT NOT NULL
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}

	return db
}

func TestRegistrarCliente(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	conns := DbConnectionsMock{Db: db}

	cliente := entity.Cliente{
		Cpf:   12345678901,
		Nome:  "Test User",
		Email: "test@example.com",
	}

	registeredCliente, err := conns.RegistrarCliente(cliente)
	if err != nil {
		t.Fatalf("RegistrarCliente failed: %v", err)
	}

	var count int
	query := `SELECT COUNT(*) FROM clientes WHERE cpf = ? AND nome = ? AND email = ?`
	err = db.QueryRowContext(context.Background(), query, cliente.Cpf, cliente.Nome, cliente.Email).Scan(&count)
	if err != nil {
		t.Fatalf("Failed to validate inserted data: %v", err)
	}

	if count != 1 {
		t.Errorf("Expected 1 row, got %d", count)
	}

	if registeredCliente != cliente {
		t.Errorf("Expected %v, got %v", cliente, registeredCliente)
	}
}

func TestBuscarCliente(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	conns := DbConnectionsMock{Db: db}

	cliente := entity.Cliente{
		Cpf:   12345678901,
		Nome:  "Test User",
		Email: "test@example.com",
	}
	insertQuery := `INSERT INTO clientes (cpf, nome, email) VALUES (?, ?, ?)`
	_, err := db.ExecContext(context.Background(), insertQuery, cliente.Cpf, cliente.Nome, cliente.Email)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	foundCliente, err := conns.BuscarCliente(cliente.Cpf)
	if err != nil {
		t.Fatalf("BuscarCliente failed: %v", err)
	}

	if foundCliente != cliente {
		t.Errorf("Expected %v, got %v", cliente, foundCliente)
	}
}

func TestBuscarClienteNotFound(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	conns := DbConnectionsMock{Db: db}

	_, err := conns.BuscarCliente(99999999999)
	if err == nil {
		t.Fatal("Expected error for non-existent CPF, got nil")
	}

	if err.Error() != sql.ErrNoRows.Error() {
		t.Errorf("Expected error %v, got %v", sql.ErrNoRows, err)
	}
}

func TestRegistrarClienteDuplicateCPF(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	conns := DbConnectionsMock{Db: db}

	cliente := entity.Cliente{
		Cpf:   12345678901,
		Nome:  "Test User",
		Email: "test@example.com",
	}

	_, err := conns.RegistrarCliente(cliente)
	if err != nil {
		t.Fatalf("Failed to insert the first cliente: %v", err)
	}

	_, err = conns.RegistrarCliente(cliente)
	if err == nil {
		t.Fatal("Expected error for duplicate CPF, got nil")
	}
}

func TestBuscarClienteInvalidCPFFormat(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	conns := DbConnectionsMock{Db: db}

	invalidCpf := int64(-1)
	_, err := conns.BuscarCliente(invalidCpf)
	if err == nil {
		t.Fatal("Expected error for invalid CPF, got nil")
	}
}

func TestBuscarClienteNotFoundError(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	conns := DbConnectionsMock{Db: db}

	_, err := conns.BuscarCliente(99999999999)
	if err == nil {
		t.Fatal("Expected error for non-existent CPF, got nil")
	}
	if err.Error() != sql.ErrNoRows.Error() {
		t.Errorf("Expected error %v, got %v", sql.ErrNoRows, err)
	}
}

func TestBuscarClienteSuccess(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	conns := DbConnectionsMock{Db: db}

	cliente := entity.Cliente{
		Cpf:   12345678901,
		Nome:  "Test User",
		Email: "test@example.com",
	}

	insertQuery := `INSERT INTO clientes (cpf, nome, email) VALUES (?, ?, ?)`

	_, err := db.ExecContext(context.Background(), insertQuery, cliente.Cpf, cliente.Nome, cliente.Email)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	retrievedCliente, err := conns.BuscarCliente(cliente.Cpf)
	if err != nil {
		t.Fatalf("Failed to retrieve cliente: %v", err)
	}

	if retrievedCliente != cliente {
		t.Errorf("Expected %v, got %v", cliente, retrievedCliente)
	}
}
