package respostas

import "net/http"


// JSON = Retorna um resposta em formato JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Wri
}