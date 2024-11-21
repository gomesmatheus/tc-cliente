package persistence

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gomesmatheus/tc-cliente/domain/entity"
	"github.com/jackc/pgx/v5"
)

type DbConnections struct {
	Db    *pgx.Conn
	Redis *redis.Client
}

func (conns DbConnections) RegistrarCliente(cliente entity.Cliente) (entity.Cliente, error) {
	_, err := conns.Db.Exec(context.Background(), "INSERT INTO clientes (cpf, nome, email) VALUES ($1, $2, $3)", cliente.Cpf, cliente.Nome, cliente.Email)
	if err != nil {
		fmt.Println("Erro ao inserir cliente na base de dados", err)
	}

	jsonData, _ := json.Marshal(cliente)

	var ctx = context.Background()
	err2 := conns.Redis.Set(ctx, strconv.FormatInt(cliente.Cpf, 10), jsonData, 72*time.Hour).Err()
	if err2 != nil {
		fmt.Printf("Could not set value: %v\n", err2)
	}

	return cliente, err
}

func (conns DbConnections) BuscarCliente(cpf int64) (entity.Cliente, error) {
	var cliente entity.Cliente
	var err error
	var ctx = context.Background()
	val, err2 := conns.Redis.Get(ctx, strconv.FormatInt(cpf, 10)).Result()
	if err2 != nil {
		fmt.Printf("Could not get value: %v\n", err2)
		err := conns.Db.QueryRow(context.Background(), "SELECT cpf, nome, email FROM clientes WHERE cpf = $1", cpf).Scan(&cliente.Cpf, &cliente.Nome, &cliente.Email)
		if err != nil {
			fmt.Println("Erro buscando por cpf", cpf, err)
		}
	}

	err = json.Unmarshal([]byte(val), &cliente)
	if err != nil {
		return cliente, err
	}

	return cliente, err
}
