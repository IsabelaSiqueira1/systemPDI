package professional

import (
	"time"

	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain"
	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain/entities"
	"github.com/google/uuid"
)

func (s *Service) StartShift(serviceID, professionalID string) (entities.Professional, error) {
	_, err := s.servicesRepo.FindByID(serviceID)
	if err != nil {
		return entities.Professional{}, err
	}

	professional, err := s.repo.FindByID(professionalID)
	if err != nil {
		return entities.Professional{}, err
	}

	if professional.ServiceID == nil || *professional.ServiceID != serviceID {
		return entities.Professional{}, domain.ErrProfessionalServiceMismatch
	}

	if professional.Status == entities.ProfessionalBusy {
		return entities.Professional{}, domain.ErrProfessionalBusy
	}

	if professional.Status == entities.ProfessionalAvailable {
		return entities.Professional{}, domain.ErrProfessionalAlreadyOnDuty
	}

	_, err = s.attendanceRepo.Create(entities.Attendance{
		ID:             uuid.NewString(),
		ProfessionalID: professionalID,
		ServiceID:      serviceID,
		StartedAt:      time.Now().UTC(),
		EndedAt:        nil,
	})
	if err != nil {
		return entities.Professional{}, err
	}

	return s.repo.UpdateStatus(professionalID, entities.ProfessionalAvailable)
}
