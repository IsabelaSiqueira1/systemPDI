package professional

type createProfessionalRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type professionalDTO struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Status    string  `json:"status"`
	ServiceID *string `json:"service_id"`
}
