package services

import (
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain"
	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain/entities"
)

func (s *Service) Create(name string) (entities.Service, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return entities.Service{}, domain.ErrServiceNameRequired
	}

	createdService := entities.Service{
		ID:        uuid.NewString(),
		Name:      name,
		CreatedAt: time.Now().UTC(),
	}

	return s.repo.Create(createdService)
}
