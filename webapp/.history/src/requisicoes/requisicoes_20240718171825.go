package requisicoes

import (
	"io"
	"net/http"
)


// FazerRequisicaoComAutenticacao = É 
func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	
}