package domain

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrServiceNameRequired  = errors.New("nome do serviço é obrigatróio")
	ErrServiceAlreadyExists = errors.New("serviço já existe")
	ErrServiceNotFound      = errors.New("serviço não encontrado")

	ErrProfessionalNotFound         = errors.New("profissional não encontrado")
	ErrProfessionalNameRequired     = errors.New("nome do profissional é obrigatório")
	ErrProfessionalEmailRequired    = errors.New("email é obrigatório")
	ErrProfessionalPasswordRequired = errors.New("senha é obrigatória")
	ErrProfessionalEmailInvalid     = errors.New("email inválido")
	ErrProfessionalEmailAlreadyUsed = errors.New("email já cadastrado")
	ErrProfessionalMustBeOffDuty    = errors.New("profissional deve estar INDISPONIVEL para ser vinculado a um serviço")
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
