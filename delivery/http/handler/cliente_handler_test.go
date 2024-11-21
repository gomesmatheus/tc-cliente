package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gomesmatheus/tc-cliente/domain/entity"
)

type MockClienteUseCases struct {
	MockCadastrar func(cliente entity.Cliente) (entity.Cliente, error)
	MockRecuperar func(cpf int64) (entity.Cliente, error)
}

func (m *MockClienteUseCases) Cadastrar(cliente entity.Cliente) (entity.Cliente, error) {
	return m.MockCadastrar(cliente)
}

func (m *MockClienteUseCases) Recuperar(cpf int64) (entity.Cliente, error) {
	return m.MockRecuperar(cpf)
}

func TestCriacaoRoute_Success(t *testing.T) {
	mockUseCases := &MockClienteUseCases{
		MockCadastrar: func(cliente entity.Cliente) (entity.Cliente, error) {
			return cliente, nil
		},
	}
	handler := NewClienteHandler(mockUseCases)

	cliente := entity.Cliente{Cpf: 12345678901, Nome: "Test User", Email: "test@example.com"}
	body, _ := json.Marshal(cliente)

	req := httptest.NewRequest("POST", "/clientes", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.CriacaoRoute(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}
	if rec.Body.String() != "Cliente inserido" {
		t.Errorf("expected response body 'Cliente inserido', got %s", rec.Body.String())
	}
}

func TestCriacaoRoute_InvalidBody(t *testing.T) {
	mockUseCases := &MockClienteUseCases{
		MockCadastrar: func(cliente entity.Cliente) (entity.Cliente, error) {
			return cliente, nil
		},
	}
	handler := NewClienteHandler(mockUseCases)

	req := httptest.NewRequest("POST", "/clientes", bytes.NewReader([]byte("invalid-json")))
	rec := httptest.NewRecorder()

	handler.CriacaoRoute(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
	if rec.Body.String() != "400 bad request" {
		t.Errorf("expected response body '400 bad request', got %s", rec.Body.String())
	}
}

func TestCriacaoRoute_FailureOnCadastrar(t *testing.T) {
	mockUseCases := &MockClienteUseCases{
		MockCadastrar: func(cliente entity.Cliente) (entity.Cliente, error) {
			return cliente, errors.New("mock failure")
		},
	}
	handler := NewClienteHandler(mockUseCases)

	cliente := entity.Cliente{Cpf: 12345678901, Nome: "Test User", Email: "test@example.com"}
	body, _ := json.Marshal(cliente)

	req := httptest.NewRequest("POST", "/clientes", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.CriacaoRoute(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, rec.Code)
	}
	if rec.Body.String() != "Erro ao cadastrar o cliente" {
		t.Errorf("expected response body 'Erro ao cadastrar o cliente', got %s", rec.Body.String())
	}
}

func TestIdentificacaoRoute_Success(t *testing.T) {
	mockUseCases := &MockClienteUseCases{
		MockRecuperar: func(cpf int64) (entity.Cliente, error) {
			return entity.Cliente{Cpf: cpf, Nome: "Test User", Email: "test@example.com"}, nil
		},
	}
	handler := NewClienteHandler(mockUseCases)

	cpf := int64(12345678901)
	req := httptest.NewRequest("GET", "/clientes/"+strconv.FormatInt(cpf, 10), nil)
	rec := httptest.NewRecorder()

	handler.IdentificacaoRoute(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var cliente entity.Cliente
	err := json.Unmarshal(rec.Body.Bytes(), &cliente)
	if err != nil {
		t.Fatalf("unexpected error unmarshalling response: %v", err)
	}

	if cliente.Cpf != cpf {
		t.Errorf("expected CPF %d, got %d", cpf, cliente.Cpf)
	}
}

func TestIdentificacaoRoute_InvalidCPF(t *testing.T) {
	mockUseCases := &MockClienteUseCases{
		MockRecuperar: func(cpf int64) (entity.Cliente, error) {
			return entity.Cliente{Cpf: cpf, Nome: "Test User", Email: "test@example.com"}, nil
		},
	}
	handler := NewClienteHandler(mockUseCases)

	req := httptest.NewRequest("GET", "/clientes/invalid-cpf", nil)
	rec := httptest.NewRecorder()

	handler.IdentificacaoRoute(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestIdentificacaoRoute_NotFound(t *testing.T) {
	mockUseCases := &MockClienteUseCases{
		MockRecuperar: func(cpf int64) (entity.Cliente, error) {
			return entity.Cliente{}, errors.New("not found")
		},
	}
	handler := NewClienteHandler(mockUseCases)

	cpf := int64(12345678901)
	req := httptest.NewRequest("GET", "/clientes/"+strconv.FormatInt(cpf, 10), nil)
	rec := httptest.NewRecorder()

	handler.IdentificacaoRoute(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
	expectedBody := fmt.Sprintf("Cliente com CPF %d não identificado", cpf)
	if rec.Body.String() != expectedBody {
		t.Errorf("expected response body '%s', got '%s'", expectedBody, rec.Body.String())
	}
}
