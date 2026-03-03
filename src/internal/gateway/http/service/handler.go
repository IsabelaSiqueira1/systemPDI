package service

import (
	servicessvc "github.com/IsabelaSiqueira1/systemPDI/src/internal/service/services"
)

type Handler struct {
	services *servicessvc.Service
}

func NewHandler(services *servicessvc.Service) *Handler {
	return &Handler{services: services}
}
