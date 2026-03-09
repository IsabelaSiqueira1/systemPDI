package entities

import "time"

type Attendance struct {
	ID             string
	ProfessionalID string
	ServiceID      string
	StartedAt      time.Time
	EndedAt        *time.Time
}
