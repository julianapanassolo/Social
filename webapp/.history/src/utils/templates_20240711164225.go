package utils

import "html/template"

var templates *template.Template

// CarregarTemplates = 
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}