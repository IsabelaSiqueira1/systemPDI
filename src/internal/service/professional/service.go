package professional

import "github.com/IsabelaSiqueira1/systemPDI/src/internal/repository"

type Service struct {
	repo         *repository.ProfessionalsRepository
	servicesRepo *repository.ServicesRepository
}

func NewService(repo *repository.ProfessionalsRepository, servicesRepo *repository.ServicesRepository) *Service {
	return &Service{repo: repo, servicesRepo: servicesRepo}
}
