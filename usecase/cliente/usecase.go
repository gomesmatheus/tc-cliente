package usecase

import (
	"errors"

	"github.com/gomesmatheus/tc-cliente/domain/entity"
	"github.com/gomesmatheus/tc-cliente/infraestructure/persistence"
)

type clienteUseCases struct {
	database persistence.ClienteRepository
}

func NewClienteUseCases(clienteRepository persistence.ClienteRepository) *clienteUseCases {
	return &clienteUseCases{
		database: clienteRepository,
	}
}

func (usecase *clienteUseCases) Cadastrar(cliente entity.Cliente) (entity.Cliente, error) {
	if !isClienteValido(cliente) {
		return cliente, errors.New("Cliente inválido")
	}

	return usecase.database.RegistrarCliente(cliente)
}

func (usecase *clienteUseCases) Recuperar(cpf int64) (entity.Cliente, error) {
	// adicionar validação de cpf
	return usecase.database.BuscarCliente(cpf)
}

func isClienteValido(cliente entity.Cliente) bool {
	return cliente.Cpf != 0 && cliente.Nome != "" && cliente.Email != ""
}
