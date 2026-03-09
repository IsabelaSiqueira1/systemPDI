package repository

import (
	"sync"

	"github.com/IsabelaSiqueira1/ProjetinhoPDI/ProjetinhoPDI/manipulacao-dados/collections/linkedlist"
	"github.com/IsabelaSiqueira1/systemPDI/src/internal/domain/entities"
)

type AttendancesRepository struct {
	mu          sync.RWMutex
	attendances *linkedlist.LinkedList[entities.Attendance]
}

func NewAttendancesRepository() *AttendancesRepository {
	return &AttendancesRepository{
		attendances: linkedlist.New[entities.Attendance](),
	}
}

func (r *AttendancesRepository) Create(attendance entities.Attendance) (entities.Attendance, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.attendances.Insert(attendance)
	return attendance, nil
}
