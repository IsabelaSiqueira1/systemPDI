package professional

import (
	"errors"
	"net/http"

	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain"
)

func respondProfessionalError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrProfessionalNotFound):
		domain.WriteError(w, http.StatusNotFound, "Profissional não encontrado")
	case errors.Is(err, domain.ErrServiceNotFound):
		domain.WriteError(w, http.StatusNotFound, "Serviço não encontrado")
	case errors.Is(err, domain.ErrProfessionalMustBeOffDuty):
		domain.WriteError(w, http.StatusConflict, "Profissional deve estar UNAVAILABLE para ser vinculado")
	case errors.Is(err, domain.ErrProfessionalAlreadyAssigned):
		domain.WriteError(w, http.StatusConflict, "Profissional já possui serviço vinculado")
	case errors.Is(err, domain.ErrProfessionalServiceMismatch):
		domain.WriteError(w, http.StatusForbidden, "Profissional não pertence ao serviço")
	case errors.Is(err, domain.ErrProfessionalBusy):
		domain.WriteError(w, http.StatusConflict, "Profissional está BUSY")
	case errors.Is(err, domain.ErrProfessionalAlreadyOnDuty):
		domain.WriteError(w, http.StatusConflict, "Profissional já está em expediente")
	case errors.Is(err, domain.ErrProfessionalNameRequired):
		domain.WriteError(w, http.StatusBadRequest, "Nome é obrigatório")
	case errors.Is(err, domain.ErrProfessionalEmailRequired):
		domain.WriteError(w, http.StatusBadRequest, "Email é obrigatório")
	case errors.Is(err, domain.ErrProfessionalPasswordRequired):
		domain.WriteError(w, http.StatusBadRequest, "Senha é obrigatória")
	case errors.Is(err, domain.ErrProfessionalEmailInvalid):
		domain.WriteError(w, http.StatusBadRequest, "Email inválido")
	case errors.Is(err, domain.ErrProfessionalEmailAlreadyUsed):
		domain.WriteError(w, http.StatusConflict, "Email já cadastrado")
	default:
		domain.WriteError(w, http.StatusInternalServerError, "Erro interno")
	}
}
