package requisicoes

import (
	"io"
	"net/http"

	
)

// FazerRequisicaoComAutenticacao = É utilizada para colocar o token na requisição
func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	request, erro := http.NewRequest(metodo, url, dados)
	if erro != nil{
		return nil, erro
	}

	cookie, _
}