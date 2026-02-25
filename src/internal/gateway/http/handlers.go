package http

import "github.com/IsabelaSiqueira1/systemPDI/src/internal/service"

type Handlers struct {
	Services *service.ServicesService
}

func NewHandlers(services *service.ServicesService) *Handlers {
	return &Handlers{Services: services}
}
