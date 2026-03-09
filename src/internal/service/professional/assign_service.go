package professional

import (
	"strings"

	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain"
	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain/entities"
)

func (s *Service) AssignService(professionalID, serviceID string) (entities.Professional, error) {
	serviceID = strings.TrimSpace(serviceID)

	professional, err := s.repo.FindByID(professionalID)
	if err != nil {
		return entities.Professional{}, err
	}

	_, err = s.servicesRepo.FindByID(serviceID)
	if err != nil {
		return entities.Professional{}, err
	}

	if professional.Status != entities.ProfessionalOffDuty {
		return entities.Professional{}, domain.ErrProfessionalMustBeOffDuty
	}

	if professional.ServiceID != nil {
		return entities.Professional{}, domain.ErrProfessionalAlreadyAssigned
	}

	return s.repo.UpdateServiceID(professionalID, &serviceID)
}
