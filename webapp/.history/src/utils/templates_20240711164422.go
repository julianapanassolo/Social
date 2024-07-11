package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// CarregarTemplates = Insere os templates html na variável templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func ExecutarTemplate(w http.ResponseWriter, template string)