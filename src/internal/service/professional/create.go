package professional

import (
	"net/mail"
	"strings"

	"github.com/google/uuid"

	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain"
	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain/entities"
)

func (s *Service) Create(name, email, password string) (entities.Professional, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" {
		return entities.Professional{}, domain.ErrProfessionalNameRequired
	}
	if email == "" {
		return entities.Professional{}, domain.ErrProfessionalEmailRequired
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return entities.Professional{}, domain.ErrProfessionalEmailInvalid
	}

	professional := entities.Professional{
		ID:        uuid.NewString(),
		Name:      name,
		Email:     email,
		Status:    entities.ProfessionalOffDuty,
		ServiceID: nil,
	}

	return s.repo.Create(professional)
}
