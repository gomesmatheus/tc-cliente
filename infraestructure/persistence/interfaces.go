package persistence

import "github.com/gomesmatheus/tc-cliente/domain/entity"

type ClienteRepository interface {
	RegistrarCliente(entity.Cliente) (entity.Cliente, error)
	BuscarCliente(cpf int64) (entity.Cliente, error)
}
