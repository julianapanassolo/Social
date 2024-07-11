package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
)

func main() {
	fmt.Println("Rodando WebApp!")
	r := router
	log.Fatal(http.ListenAndServe(":3000", nil))
}