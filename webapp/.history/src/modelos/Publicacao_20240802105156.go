package pu

import (
    "html/template"
    "net/http"
    "time"
)

// Publicacao representa uma publicação feita por um usuário
type Publicacao struct {
    ID        uint64    `json:"id,omitempty"`
    Titulo    string    `json:"titulo,omitempty"`
    Conteudo  string    `json:"conteudo,omitempty"`
    AutorID   uint64    `json:"autorId,omitempty"`
    AutorNick string    `json:"autorNick,omitempty"`
    Curtidas  uint64    `json:"curtidas"`
    CriadaEm  time.Time `json:"criadaEm,omitempty"`
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Dados de exemplo
        publicacoes := []Publicacao{
            {ID: 1, Titulo: "Primeira Publicação", Conteudo: "Este é o conteúdo da primeira publicação.", AutorID: 1, AutorNick: "Autor1", Curtidas: 10, CriadaEm: time.Now()},
            {ID: 2, Titulo: "Segunda Publicação", Conteudo: "Este é o conteúdo da segunda publicação.", AutorID: 2, AutorNick: "Autor2", Curtidas: 20, CriadaEm: time.Now()},
        }

        tmpl := template.Must(template.ParseFiles("views/home.html"))
        err := tmpl.Execute(w, publicacoes)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

    // Servir arquivos estáticos (CSS, JS, etc.)
    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

    http.ListenAndServe(":5000", nil)
}
