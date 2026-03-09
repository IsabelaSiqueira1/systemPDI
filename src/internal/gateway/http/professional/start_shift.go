package professional

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain"
)

// StartShift godoc
// @Summary Iniciar expediente
// @Description Inicia o expediente do profissional e altera seu status para AVAILABLE.
// @Tags Profissionais
// @Produce json
// @Param serviceId path string true "ID do serviço"
// @Param professionalId path string true "ID do profissional"
// @Success 200 {object} professionalDTO
// @Failure 400 {object} domain.APIError "Requisição inválida"
// @Failure 403 {object} domain.APIError "Profissional não pertence ao serviço"
// @Failure 404 {object} domain.APIError "Profissional ou serviço não encontrado"
// @Failure 409 {object} domain.APIError "Profissional BUSY ou já em expediente"
// @Failure 500 {object} domain.APIError "Erro interno"
// @Router /v1/services/{serviceId}/professionals/{professionalId}/start-shift [put]
func (h *Handler) StartShift(w http.ResponseWriter, r *http.Request) {
	serviceID := chi.URLParam(r, "serviceId")
	if serviceID == "" {
		domain.WriteError(w, http.StatusBadRequest, "ID do serviço é obrigatório")
		return
	}

	professionalID := chi.URLParam(r, "professionalId")
	if professionalID == "" {
		domain.WriteError(w, http.StatusBadRequest, "ID do profissional é obrigatório")
		return
	}

	updated, err := h.professionals.StartShift(serviceID, professionalID)
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
