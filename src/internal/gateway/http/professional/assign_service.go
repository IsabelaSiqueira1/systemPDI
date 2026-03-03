package professional

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain"
)

type assignServiceRequest struct {
	ServiceID string `json:"service_id"`
}

// AssignService godoc
// @Summary Vincular profissional a um serviço
// @Description Vincula um profissional a um serviço.
// @Tags Profissionais
// @Accept json
// @Produce json
// @Param professionalId path string true "ID do profissional"
// @Param body body assignServiceRequest true "Payload"
// @Success 200 {object} professionalDTO
// @Failure 400 {object} domain.APIError "Requisição inválida"
// @Failure 404 {object} domain.APIError "Profissional ou serviço não encontrado"
// @Failure 409 {object} domain.APIError "Profissional não está INDISPONIVEL"
// @Router /v1/professionals/{professionalId}/service [put]
func (h *Handler) AssignService(w http.ResponseWriter, r *http.Request) {
	professionalID := chi.URLParam(r, "professionalId")
	if professionalID == "" {
		domain.WriteError(w, http.StatusBadRequest, "missing_professional_id", "ID do profissional é obrigatório")
		return
	}

	var req assignServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		domain.WriteError(w, http.StatusBadRequest, "invalid_body", "JSON inválido")
		return
	}

	if req.ServiceID == "" {
		domain.WriteError(w, http.StatusBadRequest, "missing_service_id", "service_id é obrigatório")
		return
	}

	updated, err := h.professionals.AssignService(professionalID, req.ServiceID)
	if err != nil {
		respondProfessionalError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(professionalDTO{
		ID:        updated.ID,
		Name:      updated.Name,
		Email:     updated.Email,
		Status:    string(updated.Status),
		ServiceID: updated.ServiceID,
	})
}
