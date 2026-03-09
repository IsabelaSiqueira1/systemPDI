package service

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain"
)

// CreateService godoc
// @Summary Cadastrar serviço
// @Description Cadastra um novo serviço no catálogo.
// @Tags Serviços
// @Accept json
// @Produce json
// @Param body body createServiceRequest true "Payload"
// @Success 201 {object} serviceDTO
// @Failure 400 {object} domain.APIError "Requisição inválida"
// @Failure 409 {object} domain.APIError "Serviço já existe"
// @Failure 500 {object} domain.APIError "Erro interno"
// @Router /v1/services [post]
func (h *Handler) CreateService(w http.ResponseWriter, r *http.Request) {
	var req createServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		domain.WriteError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	if strings.TrimSpace(req.Name) == "" {
		domain.WriteError(w, http.StatusBadRequest, "Nome do serviço é obrigatório")
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
