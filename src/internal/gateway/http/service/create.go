package service

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain"
)

// CreateService godoc
// @Summary Create service
// @Description Creates a new service in the catalog.
// @Tags Services
// @Accept json
// @Produce json
// @Param body body createServiceRequest true "Payload"
// @Success 201 {object} serviceDTO
// @Failure 400 {object} domain.APIError "Invalid request"
// @Failure 409 {object} domain.APIError "Service already exists"
// @Router /v1/services [post]
func (h *Handler) CreateService(w http.ResponseWriter, r *http.Request) {
	var req createServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		domain.WriteError(w, http.StatusBadRequest, "invalid_body", "Invalid JSON")
		return
	}

	if strings.TrimSpace(req.Name) == "" {
		domain.WriteError(w, http.StatusBadRequest, "name_required", "Service name is required.")
		return
	}

	createdService, err := h.services.Create(req.Name)
	if err != nil {
		respondServiceError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(serviceDTO{
		ID:   createdService.ID,
		Name: createdService.Name,
	})
}
