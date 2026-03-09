package repository

import (
	"strings"
	"sync"

	"github.com/IsabelaSiqueira1/ProjetinhoPDI/ProjetinhoPDI/manipulacao-dados/collections/linkedlist"
	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain"
	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain/entities"
)

type ProfessionalsRepository struct {
	mu            sync.RWMutex
	professionals *linkedlist.LinkedList[entities.Professional]
}

func NewProfessionalsRepository() *ProfessionalsRepository {
	return &ProfessionalsRepository{
		professionals: linkedlist.New[entities.Professional](),
	}
}

func (r *ProfessionalsRepository) ExistsByEmail(email string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	normalizedEmail := strings.TrimSpace(email)
	_, found := r.professionals.Find(func(p entities.Professional) bool {
		return strings.EqualFold(p.Email, normalizedEmail)
	})
	return found
}

func (r *ProfessionalsRepository) Save(professional entities.Professional) (entities.Professional, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.professionals.Insert(professional)
	return professional, nil
}

func (r *ProfessionalsRepository) Create(professional entities.Professional) (entities.Professional, error) {
	return r.Save(professional)
}

func (r *ProfessionalsRepository) FindByID(id string) (entities.Professional, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, found := r.professionals.Find(func(p entities.Professional) bool {
		return p.ID == id
	})
	if found {
		return p, nil
	}

	return entities.Professional{}, domain.ErrProfessionalNotFound
}

func (r *ProfessionalsRepository) UpdateServiceID(id string, serviceID *string) (entities.Professional, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	updated, found := r.professionals.Update(
		func(p entities.Professional) bool {
			return p.ID == id
		},
		func(p *entities.Professional) {
			p.ServiceID = serviceID
		},
	)

	if found {
		return updated, nil
	}

	return entities.Professional{}, domain.ErrProfessionalNotFound
}

func (r *ProfessionalsRepository) UpdateStatus(id string, status entities.ProfessionalStatus) (entities.Professional, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	updated, found := r.professionals.Update(
		func(p entities.Professional) bool {
			return p.ID == id
		},
		func(p *entities.Professional) {
			p.Status = status
		},
	)

	if found {
		return updated, nil
	}

	return entities.Professional{}, domain.ErrProfessionalNotFound
}
