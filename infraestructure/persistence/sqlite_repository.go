package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gomesmatheus/tc-cliente/domain/entity"
	_ "github.com/mattn/go-sqlite3"
)

type DbConnectionsMock struct {
	Db *sql.DB
}

func (conns DbConnectionsMock) RegistrarCliente(cliente entity.Cliente) (entity.Cliente, error) {
	query := `INSERT INTO clientes (cpf, nome, email) VALUES (?, ?, ?)`
	_, err := conns.Db.ExecContext(context.Background(), query, cliente.Cpf, cliente.Nome, cliente.Email)
	if err != nil {
		fmt.Println("Erro ao inserir cliente na base de dados:", err)
		return cliente, err
	}

	fmt.Println("Cliente inserido com sucesso:", cliente)

	return cliente, nil
}

func (conns DbConnectionsMock) BuscarCliente(cpf int64) (entity.Cliente, error) {
	var cliente entity.Cliente
	query := `SELECT cpf, nome, email FROM clientes WHERE cpf = ?`

	err := conns.Db.QueryRowContext(context.Background(), query, cpf).Scan(&cliente.Cpf, &cliente.Nome, &cliente.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Cliente com CPF %d não encontrado.\n", cpf)
		} else {
			fmt.Println("Erro buscando por CPF:", err)
		}
		return cliente, err
	}

	fmt.Println("Cliente encontrado:", cliente)
	return cliente, nil
}
