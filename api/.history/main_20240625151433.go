package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
	"os"
)

func init() {
	
}

func main() {
	// Definir a variável de ambiente TZ
	os.Setenv("TZ", "America/Sao_Paulo")

	config.Carregar()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))


}
