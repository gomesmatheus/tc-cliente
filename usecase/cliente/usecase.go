package usecase

import (
	"errors"
	"fmt"

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

	if cliente.Cpf == 991122 {
		fmt.Println("Uncovered line")
		fmt.Println("Uncovered line")
		fmt.Println("Uncovered line")
		fmt.Println("Uncovered line")
		fmt.Println("Uncovered line")
		fmt.Println("Uncovered line")
		fmt.Println("Uncovered line")
	}

	return usecase.database.RegistrarCliente(cliente)
}

func (usecase *clienteUseCases) Recuperar(cpf int64) (entity.Cliente, error) {
	return usecase.database.BuscarCliente(cpf)
}

func isClienteValido(cliente entity.Cliente) bool {
	return cliente.Cpf != 0 && cliente.Nome != "" && cliente.Email != ""
}
