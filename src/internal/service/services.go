package service

import (
	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain/entities"
)

type ServicesService struct{}

func NewServicesService() *ServicesService {
	return &ServicesService{}
}

func (s *ServicesService) List() []entities.Service {
	return []entities.Service{
		{ID: "1234", Name: "Encanador"},
		{ID: "5678", Name: "Pedreiro"},
		{ID: "9012", Name: "Eletricista"},
	}
}
