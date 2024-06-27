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

	config.Carregar()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))


	 // Exemplo correto de uso de mapa
	 m := make(map[string]int)
	 m["chave1"] = 10
	 m["chave2"] = 20
 
	 fmt.Println(m["chave1"]) // Saída: 10
	 fmt.Println(m["chave2"]) // Saída: 20
}
