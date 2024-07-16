package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Carregar()
	utils.CarregarTemplates()
	r := router.Gerar()	

		
	fmt.Printf("Escutando na porta %d\n")
	log.Fatal(http.ListenAndServe(":3000", r))
}
