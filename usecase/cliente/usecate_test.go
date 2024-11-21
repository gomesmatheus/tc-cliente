package usecase

import (
	"testing"

	"github.com/gomesmatheus/tc-cliente/domain/entity"
	"github.com/gomesmatheus/tc-cliente/infraestructure/database"
)

func TestCadastrar_ValidCliente(t *testing.T) {
	repo, _ := database.NewClienteRepositoryLocal()
	useCases := NewClienteUseCases(repo)

	cliente := entity.Cliente{
		Cpf:   12345678901,
		Nome:  "Test User",
		Email: "test@example.com",
	}

	result, err := useCases.Cadastrar(cliente)
	if err != nil {
		t.Fatalf("Cadastrar failed: %v", err)
	}

	if result != cliente {
		t.Errorf("Expected %v, got %v", cliente, result)
	}
}

func TestCadastrar_InvalidCliente(t *testing.T) {
	repo, _ := database.NewClienteRepositoryLocal()
	useCases := NewClienteUseCases(repo)

	invalidClientes := []entity.Cliente{
		{Cpf: 0, Nome: "Test User", Email: "test@example.com"},  // Invalid CPF
		{Cpf: 12345678901, Nome: "", Email: "test@example.com"}, // Invalid Nome
		{Cpf: 12345678901, Nome: "Test User", Email: ""},        // Invalid Email
	}

	for _, cliente := range invalidClientes {
		_, err := useCases.Cadastrar(cliente)
		if err == nil {
			t.Errorf("Expected error for invalid cliente: %v", cliente)
		}
	}
}

func TestRecuperar_ClienteExists(t *testing.T) {
	repo, _ := database.NewClienteRepositoryLocal()
	useCases := NewClienteUseCases(repo)

	cliente := entity.Cliente{
		Cpf:   12345678901,
		Nome:  "Test User",
		Email: "test@example.com",
	}

	repo.RegistrarCliente(cliente)

	result, err := useCases.Recuperar(cliente.Cpf)
	if err != nil {
		t.Fatalf("Recuperar failed: %v", err)
	}

	if result != cliente {
		t.Errorf("Expected %v, got %v", cliente, result)
	}
}

func TestRecuperar_ClienteDoesNotExist(t *testing.T) {
	repo, _ := database.NewClienteRepositoryLocal()
	useCases := NewClienteUseCases(repo)

	_, err := useCases.Recuperar(12345678901)
	if err == nil {
		t.Fatalf("Expected error when cliente does not exist")
	}

	expectedErr := "sql: no rows in result set"
	if err.Error() != expectedErr {
		t.Errorf("Expected error: %s, got: %s", expectedErr, err.Error())
	}
}
