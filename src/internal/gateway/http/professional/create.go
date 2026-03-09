package professional

import (
	"encoding/json"
	"net/http"

	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain"
)

// CreateProfessional godoc
// @Summary Cadastrar profissional
// @Description Cadastra um novo profissional.
// @Tags Profissionais
// @Accept json
// @Produce json
// @Param body body createProfessionalRequest true "Payload"
// @Success 201 {object} professionalDTO
// @Failure 400 {object} domain.APIError "Requisição inválida"
// @Failure 409 {object} domain.APIError "Email já cadastrado"
// @Failure 500 {object} domain.APIError "Erro interno"
// @Router /v1/professionals [post]
func (h *Handler) CreateProfessional(w http.ResponseWriter, r *http.Request) {
	var req createProfessionalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		domain.WriteError(w, http.StatusBadRequest, "invalid_body", "JSON inválido")
		return
	}

	createdProfessional, err := h.professionals.Create(req.Name, req.Email, req.Password)
	if err != nil {
		respondProfessionalError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(professionalDTO{
		ID:        createdProfessional.ID,
		Name:      createdProfessional.Name,
		Email:     createdProfessional.Email,
		Status:    string(createdProfessional.Status),
		ServiceID: createdProfessional.ServiceID,
	})
}
