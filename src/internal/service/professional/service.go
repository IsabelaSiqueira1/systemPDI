package professional

import "github.com/IsabelaSiqueira1/systemPDI/src/internal/repository"

type Service struct {
	repo           *repository.ProfessionalsRepository
	servicesRepo   *repository.ServicesRepository
	attendanceRepo *repository.AttendancesRepository
}

func NewService(repo *repository.ProfessionalsRepository, servicesRepo *repository.ServicesRepository, attendanceRepo *repository.AttendancesRepository) *Service {
	return &Service{repo: repo, servicesRepo: servicesRepo, attendanceRepo: attendanceRepo}
}
