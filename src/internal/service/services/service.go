package services

import "github.com/IsabelaSiqueira1/systemPDI/src/internal/repository"

type Service struct {
	repo *repository.ServicesRepository
}

func NewService(repo *repository.ServicesRepository) *Service {
	return &Service{repo: repo}
}
