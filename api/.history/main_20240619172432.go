package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    // Definindo o timezone diretamente no código
    os.Setenv("TZ", "America/Sao_Paulo")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Esc!")
    })

    fmt.Println("Escutando na porta 5000")
    log.Fatal(http.ListenAndServe(":5000", nil))
}
