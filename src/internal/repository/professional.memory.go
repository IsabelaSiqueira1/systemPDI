package repository

import (
	"strings"
	"sync"

	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain"
	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain/entities"
)

type ProfessionalsRepository struct {
	mu            sync.RWMutex
	professionals []entities.Professional
}

func NewProfessionalsRepository() *ProfessionalsRepository {
	return &ProfessionalsRepository{
		professionals: []entities.Professional{},
	}
}

func (r *ProfessionalsRepository) ExistsByEmail(email string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	normalizedEmail := strings.TrimSpace(email)
	for _, p := range r.professionals {
		if strings.EqualFold(p.Email, normalizedEmail) {
			return true
		}
	}

	return false
}

func (r *ProfessionalsRepository) Save(professional entities.Professional) (entities.Professional, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, p := range r.professionals {
		if strings.EqualFold(p.Email, strings.TrimSpace(professional.Email)) {
			return entities.Professional{}, domain.ErrProfessionalEmailAlreadyUsed
		}
	}

	r.professionals = append(r.professionals, professional)
	return professional, nil
}

func (r *ProfessionalsRepository) Create(professional entities.Professional) (entities.Professional, error) {
	return r.Save(professional)
}

func (r *ProfessionalsRepository) FindByID(id string) (entities.Professional, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, p := range r.professionals {
		if p.ID == id {
			return p, nil
		}
	}
	return entities.Professional{}, domain.ErrProfessionalNotFound
}

func (r *ProfessionalsRepository) UpdateServiceID(id string, serviceID *string) (entities.Professional, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, p := range r.professionals {
		if p.ID == id {
			r.professionals[i].ServiceID = serviceID
			return r.professionals[i], nil
		}
	}
	return entities.Professional{}, domain.ErrProfessionalNotFound
}
