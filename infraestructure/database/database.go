package database

import (
	"github.com/gomesmatheus/tc-cliente/infraestructure/persistence"
)

func NewClienteRepository() (persistence.ClienteRepository, error) {
	pgDb, _ := NewPostgresDb("postgres://postgres:123@cliente-db:5432/postgres")

	return persistence.DbConnections{
		Db:    pgDb,
		Redis: NewRedisDb(),
	}, nil
}

func NewClienteRepositoryLocal() (persistence.ClienteRepository, error) {
	db := NewSqliteDB()

	return persistence.DbConnectionsMock{
		Db: db,
	}, nil
}
