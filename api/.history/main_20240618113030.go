package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando API!")
	fmt.Println("Alteração git")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))


}