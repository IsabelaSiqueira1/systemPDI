package http

import (
	"encoding/json"
	"net/http"
)

type serviceDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ListServices godoc
// @Summary List services
// @Description Returns all registered services (in-memory).
// @Tags services
// @Produce json
// @Success 200 {array} serviceDTO
// @Router /v1/services [get]
func (h *Handlers) ListServices(w http.ResponseWriter, r *http.Request) {
	services := h.Services.List()

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
