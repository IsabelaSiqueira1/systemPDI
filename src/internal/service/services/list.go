package services

import "github.com/IsabelaSiqueira1/systemPDI/src/internal/domain/entities"

func (s *Service) List() []entities.Service {
	return s.repo.List()
}
