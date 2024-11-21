package usecase

import "github.com/gomesmatheus/tc-cliente/domain/entity"

type ClienteUseCases interface {
	Cadastrar(entity.Cliente) (entity.Cliente, error)
	Recuperar(cpf int64) (entity.Cliente, error)
}
