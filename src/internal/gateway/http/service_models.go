package http

type serviceDTO struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

type createServiceRequest struct {
	Name string `json:"name"`
}