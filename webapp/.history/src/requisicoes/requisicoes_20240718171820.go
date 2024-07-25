package requisicoes

import (
	"io"
	"net/http"
)


// 
func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	
}