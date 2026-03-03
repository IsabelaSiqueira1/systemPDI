package professional

import (
	professionalsvc "github.com/IsabelaSiqueira1/systemPDI/src/internal/service/professional"
)

type Handler struct {
	professionals *professionalsvc.Service
}

func NewHandler(professionals *professionalsvc.Service) *Handler {
	return &Handler{professionals: professionals}
}
