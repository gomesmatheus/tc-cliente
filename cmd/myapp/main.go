package main

import (
	"fmt"
	"log"
	"net/http"

	handlers "github.com/gomesmatheus/tc-cliente/delivery/http/handler"
	"github.com/gomesmatheus/tc-cliente/infraestructure/database"
	usecase "github.com/gomesmatheus/tc-cliente/usecase/cliente"
)

func main() {
	clienteRepository, err := database.NewClienteRepository()

	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	clienteUseCases := usecase.NewClienteUseCases(clienteRepository)
	clienteHandler := handlers.NewClienteHandler(clienteUseCases)
	http.HandleFunc("/cliente", clienteHandler.CriacaoRoute)
	http.HandleFunc("/cliente/", clienteHandler.IdentificacaoRoute)

	fmt.Println("Cliente ms running!")
	log.Fatal(http.ListenAndServe(":3334", nil))
}
