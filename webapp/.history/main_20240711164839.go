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
	
	fmt.Println("Escutando!")
	log.Fatal(http.ListenAndServe(":2000", r))
}