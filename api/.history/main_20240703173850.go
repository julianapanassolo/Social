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
	// Definir a variável de ambiente TZ
	os.Setenv("TZ", "America/Sao_Paulo")

	db, erro := banco.Conectar()
	if erro != nil {
		return
	}
	defer db.Close()

	config.Carregar()
	r := router.Gerar(db)

	fmt.Println(config.SECRET_KEY)

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))


}
