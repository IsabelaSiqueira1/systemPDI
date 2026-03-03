package entities

type ProfessionalStatus string

const (
	ProfessionalAvailable  ProfessionalStatus = "DISPONIVEL"
	ProfessionalBusy       ProfessionalStatus = "OCUPADO"
	ProfessionalOffDuty    ProfessionalStatus = "INDISPONIVEL"
)

type Professional struct {
	ID           string
	Name         string
	Email        string
	PasswordHash string

	Status           ProfessionalStatus
	ServiceID *string
}