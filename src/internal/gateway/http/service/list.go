package service

import (
	"encoding/json"
	"net/http"
)

// ListServices godoc
// @Summary Listar serviços
// @Description Retorna todos os serviços cadastrados.
// @Tags serviços
// @Produce json
// @Success 200 {array} serviceDTO
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

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(out)
}
