package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	utils.CarregarTemplates()
	r := router.Gerar()	
	
	fmt.Println("Escutando na porta!")
	log.Fatal(http.ListenAndServe(":5000", r))
}