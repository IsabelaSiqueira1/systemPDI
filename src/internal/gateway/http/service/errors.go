package service

import (
	"errors"
	"net/http"

	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain"
)

func respondServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrServiceNameRequired):
		domain.WriteError(w, http.StatusBadRequest, "Nome do serviço é obrigatório")
	case errors.Is(err, domain.ErrServiceAlreadyExists):
		domain.WriteError(w, http.StatusConflict, "Serviço já existe")
	default:
		domain.WriteError(w, http.StatusInternalServerError, "Erro interno")
	}
}
