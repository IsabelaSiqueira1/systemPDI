package http

import (
	"net/http"

	professionalgateway "github.com/IsabelaSiqueira1/systemPDI/src/internal/gateway/http/professional"
	servicegateway "github.com/IsabelaSiqueira1/systemPDI/src/internal/gateway/http/service"
	"github.com/IsabelaSiqueira1/systemPDI/src/internal/repository"
	svcprofessional "github.com/IsabelaSiqueira1/systemPDI/src/internal/service/professional"
	svcservices "github.com/IsabelaSiqueira1/systemPDI/src/internal/service/services"
	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	servicesRepo := repository.NewServicesRepository()
	servicesSvc := svcservices.NewService(servicesRepo)
	servicesHandler := servicegateway.NewHandler(servicesSvc)

	professionalsRepo := repository.NewProfessionalsRepository()
	professionalsSvc := svcprofessional.NewService(professionalsRepo, servicesRepo)
	professionalsHandler := professionalgateway.NewHandler(professionalsSvc)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/services", servicesHandler.ListServices)
		r.Post("/services", servicesHandler.CreateService)

		r.Post("/professionals", professionalsHandler.CreateProfessional)
		r.Put("/professionals/{professionalId}/service", professionalsHandler.AssignService)
	})

	return r
}
