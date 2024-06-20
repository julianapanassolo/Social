package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Definindo o timezone diretamente no código
	os.Setenv("TZ", "America/Sao_Paulo")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	   fmt.Fprintf(w, "Hello, world!")
	   })
   
	config.Carregar()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}