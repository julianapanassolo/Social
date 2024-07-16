package controllers

import (
	"encoding/json"
	"net/http"
)

func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal()
}