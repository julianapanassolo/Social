package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	fmt.Println("Rodando WebApp!")
	
	
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":2000", r))
}