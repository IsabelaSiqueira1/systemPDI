package repository

import (
	"strings"
	"sync"

	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain"
	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain/entities"
)

type ServicesRepository struct {
	mu       sync.RWMutex
	services []entities.Service
}

func NewServicesRepository() *ServicesRepository {
	return &ServicesRepository{
		services: []entities.Service{},
	}
}

func (r *ServicesRepository) List() []entities.Service {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]entities.Service, len(r.services))
	copy(out, r.services)
	return out
}

func (r *ServicesRepository) Create(createdService entities.Service) (entities.Service, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	name := strings.TrimSpace(createdService.Name)

	for _, existing := range r.services {
		if strings.EqualFold(existing.Name, name) {
			return entities.Service{}, domain.ErrServiceAlreadyExists
		}
	}

	r.services = append(r.services, createdService)
	return createdService, nil
}

func (r *ServicesRepository) FindByID(id string) (entities.Service, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, s := range r.services {
		if s.ID == id {
			return s, nil
		}
	}
	return entities.Service{}, domain.ErrServiceNotFound
}