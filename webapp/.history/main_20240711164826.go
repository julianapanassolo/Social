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
	
	
	
	
	fmt.Println("Rodando WebApp!")
	log.Fatal(http.ListenAndServe(":2000", r))
}