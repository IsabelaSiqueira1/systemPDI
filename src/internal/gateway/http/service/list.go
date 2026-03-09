package service

import (
	"encoding/json"
	"net/http"

	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain"
)

// ListServices godoc
// @Summary Listar serviços
// @Description Retorna todos os serviços cadastrados.
// @Tags serviços
// @Produce json
// @Success 200 {array} serviceDTO
// @Failure 500 {object} domain.APIError
// @Router /v1/services [get]
func (h *Handler) ListServices(w http.ResponseWriter, r *http.Request) {
	services := h.services.List()

	out := make([]serviceDTO, 0, len(services))
	for _, s := range services {
		out = append(out, serviceDTO{
			ID:   s.ID,
			Name: s.Name,
		})
	}

	payload, err := json.Marshal(out)
	if err != nil {
		domain.WriteError(w, http.StatusInternalServerError, "Erro interno")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(payload)
}
