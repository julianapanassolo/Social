package respostas

import "net/http"


// JSON = Retorna um resposta 
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
}