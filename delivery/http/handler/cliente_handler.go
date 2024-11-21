package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gomesmatheus/tc-cliente/domain/entity"
	"github.com/gomesmatheus/tc-cliente/usecase"
)

type ClienteHandler struct {
	clienteUseCases usecase.ClienteUseCases
}

func NewClienteHandler(clienteUseCases usecase.ClienteUseCases) *ClienteHandler {
	return &ClienteHandler{
		clienteUseCases: clienteUseCases,
	}
}

func (c *ClienteHandler) CriacaoRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		var cliente entity.Cliente
		err = json.Unmarshal(body, &cliente)
		if err != nil {
			fmt.Println("Error parsing request body")
			w.WriteHeader(400)
			w.Write([]byte("400 bad request"))
			return
		}

		fmt.Println(cliente)

		cliente, err = c.clienteUseCases.Cadastrar(cliente)
		if err != nil {
			fmt.Println("Erro ao cadastrar o cliente", err)
			w.WriteHeader(500)
			w.Write([]byte("Erro ao cadastrar o cliente"))
			return
		}

		w.WriteHeader(201)
		w.Write([]byte("Cliente inserido"))
	}

	return
}

func (c *ClienteHandler) IdentificacaoRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		cpf, err := strconv.ParseInt(strings.Split(r.URL.Path, "/")[2], 10, 64)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(400)
			w.Write([]byte("Formato de CPF inválido"))

		}

		cliente, err := c.clienteUseCases.Recuperar(cpf)
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte(fmt.Sprintf("Cliente com CPF %d não identificado", cpf)))
			return
		}
		response, _ := json.Marshal(cliente)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(response)
	}

	return
}
