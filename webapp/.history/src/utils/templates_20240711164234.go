package utils

import "html/template"

var templates *template.Template

// CarregarTemplates = Insere os templates html
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}