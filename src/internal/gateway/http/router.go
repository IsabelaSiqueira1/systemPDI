package http

import (
	"net/http"

	"github.com/IsabelaSiqueira1/systemPDI/src/internal/service"
	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	servicesSvc := service.NewServicesService()
	h := NewHandlers(servicesSvc)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/services", h.ListServices)
	})

	return r
}
