package repository

import (
	"strings"
	"sync"

	"github.com/IsabelaSiqueira1/ProjetinhoPDI/ProjetinhoPDI/manipulacao-dados/collections/linkedlist"
	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain"
	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain/entities"
)

type ServicesRepository struct {
	mu       sync.RWMutex
	services *linkedlist.LinkedList[entities.Service]
}

func NewServicesRepository() *ServicesRepository {
	return &ServicesRepository{
		services: linkedlist.New[entities.Service](),
	}
}

func (r *ServicesRepository) List() []entities.Service {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]entities.Service, 0, r.services.Len())
	r.services.ForEach(func(service entities.Service) {
		out = append(out, service)
	})
	return out
}

func (r *ServicesRepository) Create(createdService entities.Service) (entities.Service, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	name := strings.TrimSpace(createdService.Name)

	_, found := r.services.Find(func(existing entities.Service) bool {
		return strings.EqualFold(existing.Name, name)
	})
	if found {
		return entities.Service{}, domain.ErrServiceAlreadyExists
	}

	r.services.Insert(createdService)
	return createdService, nil
}

func (r *ServicesRepository) FindByID(id string) (entities.Service, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	service, found := r.services.Find(func(service entities.Service) bool {
		return service.ID == id
	})
	if found {
		return service, nil
	}

	return entities.Service{}, domain.ErrServiceNotFound
}
