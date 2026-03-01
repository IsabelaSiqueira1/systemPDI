package http

import (
	"encoding/json"
	"net/http"
)

// CreateService godoc
// @Summary Criar serviço
// @Description Cria um novo serviço no catálogo.
// @Tags Serviços
// @Accept json
// @Produce json
// @Param body body createServiceRequest true "Payload"
// @Success 201 {object} serviceDTO
// @Failure 400 {string} string "JSON inválido"
// @Router /v1/services [post]
func (h *Handlers) CreateService(w http.ResponseWriter, r *http.Request) {
	var req createServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	svc := h.Services.Create(req.Name)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(serviceDTO{
		ID:   svc.ID,
		Name: svc.Name,
	})
}