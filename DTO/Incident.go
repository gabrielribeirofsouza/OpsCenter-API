package dto

type IncidentCreateResponse struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}
type IncidentAssumeResponse struct {
	Message         string `json:"message"`
	NameResponsable string `json:"name_responsable"`
}
type IncidentUpdateResponse struct{
	Message string `json:"message"`
	ID string `json:"id_history"`
}
