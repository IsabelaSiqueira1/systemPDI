package domain

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrServiceNameRequired  = errors.New("nome do serviço é obrigatróio")
	ErrServiceAlreadyExists = errors.New("serviço já existe")
	ErrNotImplemented       = errors.New("not implemented")
)

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func WriteError(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(APIError{Code: code, Message: message})
}
