package entities

type ProfessionalStatus string

const (
	ProfessionalAvailable ProfessionalStatus = "AVAILABLE"
	ProfessionalBusy      ProfessionalStatus = "BUSY"
	ProfessionalOffDuty   ProfessionalStatus = "UNAVAILABLE"
)

type Professional struct {
	ID           string
	Name         string
	Email        string
	PasswordHash string

	Status    ProfessionalStatus
	ServiceID *string
}
