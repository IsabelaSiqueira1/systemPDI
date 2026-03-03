package service

import (
	"errors"
	"net/http"

	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain"
)

func respondServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrServiceNameRequired):
		domain.WriteError(w, http.StatusBadRequest, "name_required", "Service name is required.")
	case errors.Is(err, domain.ErrServiceAlreadyExists):
		domain.WriteError(w, http.StatusConflict, "service_already_exists", "Service already exists.")
	default:
		domain.WriteError(w, http.StatusInternalServerError, "internal_error", "Internal server error.")
	}
}
