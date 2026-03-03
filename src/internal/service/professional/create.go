package professional

import (
	"net/mail"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

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
	if strings.TrimSpace(password) == "" {
		return entities.Professional{}, domain.ErrProfessionalPasswordRequired
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return entities.Professional{}, domain.ErrProfessionalEmailInvalid
	}

	passwordHashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return entities.Professional{}, err
	}

	professional := entities.Professional{
		ID:           uuid.NewString(),
		Name:         name,
		Email:        email,
		PasswordHash: string(passwordHashBytes),
		Status:       entities.ProfessionalOffDuty,
		ServiceID:    nil,
	}

	return s.repo.Create(professional)
}
