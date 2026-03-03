package professional

import (
	"errors"
	"net/http"

	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain"
)

func respondProfessionalError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrProfessionalNotFound):
		domain.WriteError(w, http.StatusNotFound, "professional_not_found", "Profissional não encontrado")
	case errors.Is(err, domain.ErrServiceNotFound):
		domain.WriteError(w, http.StatusNotFound, "service_not_found", "Serviço não encontrado")
	case errors.Is(err, domain.ErrProfessionalMustBeOffDuty):
		domain.WriteError(w, http.StatusConflict, "professional_not_off_duty", "Profissional deve estar INDISPONIVEL para ser vinculado")
	case errors.Is(err, domain.ErrProfessionalNameRequired):
		domain.WriteError(w, http.StatusBadRequest, "name_required", "Nome é obrigatório")
	case errors.Is(err, domain.ErrProfessionalEmailRequired):
		domain.WriteError(w, http.StatusBadRequest, "email_required", "Email é obrigatório")
	case errors.Is(err, domain.ErrProfessionalPasswordRequired):
		domain.WriteError(w, http.StatusBadRequest, "password_required", "Senha é obrigatória")
	case errors.Is(err, domain.ErrProfessionalEmailInvalid):
		domain.WriteError(w, http.StatusBadRequest, "email_invalid", "Email inválido")
	case errors.Is(err, domain.ErrProfessionalEmailAlreadyUsed):
		domain.WriteError(w, http.StatusConflict, "email_already_used", "Email já cadastrado")
	default:
		domain.WriteError(w, http.StatusInternalServerError, "internal_error", "Erro interno")
	}
}
