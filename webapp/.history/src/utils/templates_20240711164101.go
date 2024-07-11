package utils

import "html/template"

var templates *template.Template
func CarregarTemplates() {
	templates = template.Must(template.Par)
}